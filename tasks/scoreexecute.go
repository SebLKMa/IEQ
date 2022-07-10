package tasks

import (
	"errors"
	"log"
	"time"

	db "github.com/seblkma/ieq/db/postgres"
	mdl "github.com/seblkma/ieq/models"
	rate "github.com/seblkma/ieq/ratings"
	awair "github.com/seblkma/ieq/sensors/awair"
	uhoo "github.com/seblkma/ieq/sensors/uhoo"
)

// Execute Implements interface Executable.Execute()
// This function executes an endless loop to get device, metrics, compute and
// store the scores at configured interval.
// Therefore it must be invoked using goroutine.
func (task *ScoringTask) Execute() error {
	if !task.Initialized {
		return errors.New("ScoringTask not properly initialized")
	}

	minutes := task.Cfg.TASK.Minutes
	duration := time.Duration(minutes * 60000)
	timenow := time.Now()

	// delay until the 5 minutes of the hour
	displayID := task.Cfg.VENDOR.DeviceDisplayID
	for {
		timenow = time.Now()
		_, minutes, _ := timenow.Clock()
		if minutes%5 == 0 { // only start task at 5 minute multiples of the hour
			break // lets do it
		}
		log.Printf("Waiting to start %s task...%s\n", displayID, timenow.Format("02/01/2006 15:04:05"))
		time.Sleep(1 * time.Second)
	}

	// this is the endless loop to get device, metrics, compute and store the scores
	for {
		timenow = time.Now()

		switch task.Cfg.VENDOR.Name {
		case "awair":
			sensor := awair.SensorInfo{}
			vendorDevID := task.Cfg.VENDOR.DeviceID
			sensor.Token = task.Cfg.VENDOR.Token
			sensor.Org = task.Cfg.VENDOR.Org
			devInfo, err := sensor.GetDeviceInfo(vendorDevID)
			log.Printf("Executing for %s at %s\n", devInfo.DeviceID, timenow.Format("2006/01/02 15:04:05"))
			if err != nil {
				log.Println(err)
				break
			}
			// write device error status to db
			err = db.CreateDeviceStatus(devInfo)
			if err != nil {
				log.Println(err)
				break
			}
			metrics, err := sensor.GetLatestMetrics(vendorDevID)
			if err != nil || metrics.Empty {
				log.Println(err)
				break
			}
			task.scoreThem(devInfo, metrics)
		case "uhoo":
			sensor := uhoo.SensorInfo{}
			vendorDevID := task.Cfg.VENDOR.DeviceID
			sensor.Token = task.Cfg.VENDOR.Token
			sensor.Org = task.Cfg.VENDOR.Org
			devInfo, err := sensor.GetDeviceInfo(vendorDevID)
			log.Printf("Executing for %s at %s\n", devInfo.DeviceID, timenow.Format("2006/01/02 15:04:05"))
			if err != nil {
				log.Println(err)
				break
			}
			// write device error status to db
			err = db.CreateDeviceStatus(devInfo)
			if err != nil {
				log.Println(err)
				break
			}
			metrics, err := sensor.GetLatestMetrics(vendorDevID)
			if err != nil || metrics.Empty {
				log.Println(err)
				break
			}
			task.scoreThem(devInfo, metrics)
		}

		// pause
		//time.Sleep(60000 * time.Millisecond) // 1 minute
		time.Sleep(duration * time.Millisecond)
	}
}

// computes the metrics scores and store them in database for the device.
func (task *ScoringTask) scoreThem(devInfo mdl.DeviceInfo, metrics mdl.Metrics) {

	// write device error status to db
	err := db.CreateDeviceStatus(devInfo)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// no need to continue if device status is 0
	if devInfo.Status == 0 {
		return
	}

	// set up the weightings
	thermalRating := rate.Rating{}
	thermalRating.Setup("Thermal", task.Cfg.WEIGHTINGS.Thermal)

	iaqRating := rate.Rating{}
	iaqRating.Setup("IAQ", task.Cfg.WEIGHTINGS.IAQ)

	lightingRating := rate.Rating{}
	lightingRating.Setup("Lighting", task.Cfg.WEIGHTINGS.Lighting)

	noiseRating := rate.Rating{}
	noiseRating.Setup("Noise", task.Cfg.WEIGHTINGS.Noise)

	// start computing the metrics scores for the device
	dbmetrics := mdl.Metrics{}
	dbmetrics.DeviceID = devInfo.DeviceID
	dbmetricscore := mdl.MetricScore{}
	dbmetricscore.DeviceID = devInfo.DeviceID
	dbieqscore := mdl.IeqScore{}
	dbieqscore.DeviceID = devInfo.DeviceID

	score := rate.ComputeScore(task.TemperatureFormula, metrics.Temperature)
	thermalRating.AddIndex("Temperature", score)
	dbmetrics.Temperature = metrics.Temperature
	dbmetricscore.Temperature = score

	score = rate.ComputeScore(task.HumidityFormula, metrics.Humidity)
	thermalRating.AddIndex("Humidity", score)
	dbmetrics.Humidity = metrics.Humidity
	dbmetricscore.Humidity = score

	score = rate.ComputeScore(task.Co2Formula, metrics.CO2)
	iaqRating.AddIndex("CO2", score)
	dbmetrics.CO2 = metrics.CO2
	dbmetricscore.CO2 = score

	score = rate.ComputeScore(task.VocFormula, metrics.VOC)
	iaqRating.AddIndex("VOC", score)
	dbmetrics.VOC = metrics.VOC
	dbmetricscore.VOC = score

	score = rate.ComputeScore(task.Pm25Formula, metrics.PM25)
	iaqRating.AddIndex("PM25", score)
	dbmetrics.PM25 = metrics.PM25
	dbmetricscore.PM25 = score
	// include lighting score if required
	if lightingRating.Weighting() > 0 {
		score = rate.ComputeScore(task.LightingFormula, metrics.Lighting)
		lightingRating.AddIndex("Lighting", score)
		dbmetrics.Lighting = metrics.Lighting
		dbmetricscore.Lighting = score
	}
	// including noise score if required
	if noiseRating.Weighting() > 0 {
		score = rate.ComputeScore(task.NoiseFormula, metrics.Noise)
		noiseRating.AddIndex("Noise", score)
		dbmetrics.Noise = metrics.Noise
		dbmetricscore.Noise = score
	}

	// compute ratings for IEQ components
	thermalRating.SetRating()
	iaqRating.SetRating()
	// compute lighting rating if required
	if lightingRating.Weighting() > 0 {
		lightingRating.SetRating()
	}
	// compute noise rating if required
	if noiseRating.Weighting() > 0 {
		noiseRating.SetRating()
	}

	// compute IEQ overall rating
	ieqRating := rate.IEQRating{}
	ieqRating.Setup("Overall IEQ", 1.0)
	ieqRating.AddIndex(thermalRating.Name(), thermalRating.Rate())
	ieqRating.AddIndex(iaqRating.Name(), iaqRating.Rate())
	// include lighting rating if required
	if lightingRating.Weighting() > 0 {
		ieqRating.AddIndex(lightingRating.Name(), lightingRating.Rate())
	}
	// include noise rating if required
	if noiseRating.Weighting() > 0 {
		ieqRating.AddIndex(noiseRating.Name(), noiseRating.Rate())
	}
	ieqRating.SetRating()

	// start storing to database
	dbieqscore.Scheme = task.Cfg.WEIGHTINGS.Scheme
	dbieqscore.Thermal = thermalRating.Rate()
	dbieqscore.ThermalWeighting = thermalRating.Weighting()
	dbieqscore.IAQ = iaqRating.Rate()
	dbieqscore.IAQWeighting = iaqRating.Weighting()
	if lightingRating.Weighting() > 0 {
		dbieqscore.Lighting = lightingRating.Rate()
		dbieqscore.LightingWeighting = lightingRating.Weighting()
	}
	if noiseRating.Weighting() > 0 {
		dbieqscore.Noise = noiseRating.Rate()
		dbieqscore.NoiseWeighting = noiseRating.Weighting()
	}
	dbieqscore.Overall = ieqRating.Rate()

	for _, i := range ieqRating.Indices() {
		log.Printf("%v ", i)
	}
	log.Println()
	log.Printf("%s Rating: %g\n", ieqRating.Name(), ieqRating.Rate())

	// commit to database
	err = db.CreateMetric(dbmetrics)
	if err != nil {
		log.Println(err.Error())
	}
	err = db.CreateMetricScore(dbmetricscore)
	if err != nil {
		log.Println(err.Error())
	}
	err = db.CreateIeqScore(dbieqscore)
	if err != nil {
		log.Println(err.Error())
	}

}
