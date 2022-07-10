package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	db "github.com/seblkma/ieq/db/postgres"
	mdl "github.com/seblkma/ieq/models"
)

//=============================================================================
// HTTP Handlers
// defaultHandler handles / route, renders a simple hello
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	fmt.Fprintln(w, "guten tag! ich bin am Leben")
}

var (
	//templates = template.Must(template.ParseGlob("../gotemplates/*/*")) // parse all in dir and subdirs
	templates = template.Must(template.ParseFiles(
		"../gotemplates/common/customstyle.html",
		"../gotemplates/common/customscript.html",
		"../gotemplates/common/footer.html",
		"../gotemplates/common/metrics.html",
		"../gotemplates/common/metricscores.html",
		"../gotemplates/common/status.html",
		"../gotemplates/ieq/header.html",
		"../gotemplates/ieq/ieqscores.html",
		"../gotemplates/ieqcharts/chartcss.html",
		"../gotemplates/ieqcharts/heartscss.html",
		"../gotemplates/ieqcharts/scriptline.html",
		"../gotemplates/ieqcharts/scriptdonut.html",
		"../gotemplates/ieqcharts/scripthbarstacked.html",
		"../gotemplates/ieqcharts/scripthearts.html",
		"../gotemplates/ieqcharts/chartheader.html",
		"../gotemplates/ieqcharts/chart.html",
		"../gotemplates/ieqcharts/line.html",
		"../gotemplates/ieqcharts/hearts.html",
		"../gotemplates/ieqcharts/devicestatus.html",
		"../gotemplates/ieqcharts/chart_a.html",
		"../gotemplates/ieqcharts/scriptdonut_a.html",
		"../gotemplates/ieqcharts/scripthbarstacked_a.html",
		"../gotemplates/ieqcharts/line_a.html",
		"../gotemplates/ieqcharts/scriptline_a.html"))
)

// ieqNumbersHandler renders latest IEQ scores and metrics numbers
func ieqNumbersHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)

	// checks what templates are parsed successfully
	//fmt.Println(templates.DefinedTemplates())

	// https://golangbyexample.com/print-struct-variables-golang/

	deviceID, err := getDeviceIDFromURL(w, r)
	if err != nil {
		fmt.Fprintf(w, "Error encountered:\n%s\n", err.Error())
		return
	}

	dbmetrics, err := db.ReadLatestMetrics(deviceID)
	if err != nil {
		fmt.Fprintf(w, "Error reading IEQ database: %s\n", err.Error())
		return
	}
	//fmt.Fprintf(w, "%#v\n", dbmetrics)

	dbmetricscore, err := db.ReadLatestMetricScores(deviceID)
	if err != nil {
		fmt.Fprintf(w, "Error reading IEQ database: %s\n", err.Error())
		return
	}
	//fmt.Fprintf(w, "%#v\n", dbmetricscore)

	dbieqscore, err := db.ReadLatestIeqScores(deviceID)
	if err != nil {
		fmt.Fprintf(w, "Error reading IEQ database: %s\n", err.Error())
		return
	}
	//fmt.Fprintf(w, "%#v\n", dbieqscore)

	devInfo, err := db.ReadLastDeviceStatus(deviceID)
	if err != nil {
		fmt.Fprintf(w, "Error reading IEQ database: %s\n", err.Error())
		return
	}
	/* simple template test
	t := template.New("fieldname example")
	t, _ = t.Parse("Device {{.DeviceID}}")
	t.Execute(w, dbmetrics)
	*/

	if devInfo.Status == 0 {
		if err := templates.ExecuteTemplate(w, "status", devInfo); err != nil {
			log.Println(err)
		}
	}

	if err := templates.ExecuteTemplate(w, "ieqscores", dbieqscore); err != nil {
		log.Println(err)
	}
	if err := templates.ExecuteTemplate(w, "metricscores", dbmetricscore); err != nil {
		log.Println(err)
	}
	if err := templates.ExecuteTemplate(w, "metrics", dbmetrics); err != nil {
		log.Println(err)
	}

}

// getDeviceIdFromURL extracts and returns the Device ID from URL.
// If not found, error and empty string are returned.
func getDeviceIDFromURL(w http.ResponseWriter, r *http.Request) (string, error) {
	urlString := r.URL.String()
	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	qs := u.Query()
	result := qs.Get("device_id")
	if result == "" {
		return result, errors.New("device_id param missing in URL query string")
	}
	return result, nil
}

// ieqchartsHandler renders latest IEQ donut and side-by-side chart
func ieqchartsHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)

	deviceID, err := getDeviceIDFromURL(w, r)
	if err != nil {
		fmt.Fprintf(w, "Error encountered:\n%s\n", err.Error())
		return
	}

	location, err := time.LoadLocation("Asia/Singapore")
	if err != nil {
		fmt.Fprintf(w, "Error time.LoadLocation %s\n", err.Error())
		return
	}

	dbmetricscores, err := db.ReadMetrics(deviceID, 10)
	if err != nil {
		fmt.Fprintf(w, "Error ReadMetrics IEQ database: %s\n", err.Error())
		return
	}

	dbmetricscore, err := db.ReadLatestMetricScores(deviceID)
	if err != nil {
		fmt.Fprintf(w, "Error ReadLatestMetricScores IEQ database: %s\n", err.Error())
		return
	}

	dbieqscore, err := db.ReadLatestIeqScores(deviceID)
	if err != nil {
		fmt.Fprintf(w, "Error ReadLatestIeqScores IEQ database: %s\n", err.Error())
		return
	}

	devInfo, err := db.ReadLastDeviceStatus(deviceID)
	if err != nil {
		fmt.Fprintf(w, "Error ReadLastDeviceStatus IEQ database: %s\n", err.Error())
		return
	}

	if devInfo.Status == 0 {
		if err := templates.ExecuteTemplate(w, "devicestatus", devInfo); err != nil {
			log.Println(err)
		}
	}

	// using anonymous struct to pass data into go template
	scores := struct {
		IeqScores    mdl.IeqScore
		MetricScores mdl.MetricScore
	}{dbieqscore, dbmetricscore}

	// show IEQ elements depending on weightings, otherwise just IAQ elements
	if dbieqscore.LightingWeighting > 0 && dbieqscore.NoiseWeighting > 0 {
		if err := templates.ExecuteTemplate(w, "chart", scores); err != nil {
			log.Println(err)
		}
	} else {
		if err := templates.ExecuteTemplate(w, "chart_a", scores); err != nil {
			log.Println(err)
		}
	}

	if err := templates.ExecuteTemplate(w, "hearts", scores); err != nil {
		log.Println(err)
	}

	// dbmetricscores are in time descending order
	// reversed will store metrics in time ascending order for line chart
	reversed := []mdl.Metrics{}
	utc := time.Now().UTC()
	local := utc
	for i, m := range dbmetricscores {
		utc = m.CreatedOn
		local = utc
		local = local.In(location)
		//t.Log("UTC", utc.Format("15:04"), local.Location(), local.Format("15:04"))
		//t.Logf("%v %v %g %g\n", m.CreatedOn, local, m.Temperature, m.Humidity)
		n := dbmetricscores[len(dbmetricscores)-1-i]
		reversed = append(reversed, n)
	}

	// just a local struct type to be passed to gotemplates
	metrics := struct {
		Times        []string
		Temperatures []float64
		Humidities   []float64
		CO2s         []float64
		VOCs         []float64
		PM25s        []float64
		Visuals      []float64
		Acoustics    []float64
	}{}

	for _, r := range reversed {
		utc = r.CreatedOn
		local = utc
		local = local.In(location)
		metrics.Times = append(metrics.Times, local.Format("15:04"))
		metrics.Temperatures = append(metrics.Temperatures, r.Temperature)
		metrics.Humidities = append(metrics.Humidities, r.Humidity)
		metrics.CO2s = append(metrics.CO2s, r.CO2)
		metrics.VOCs = append(metrics.VOCs, r.VOC)
		metrics.PM25s = append(metrics.PM25s, r.PM25)
		metrics.Visuals = append(metrics.Visuals, r.Lighting)
		metrics.Acoustics = append(metrics.Acoustics, r.Noise)
	}

	// show IEQ elements depending on weightings, otherwise just IAQ elements
	if dbieqscore.LightingWeighting > 0 && dbieqscore.NoiseWeighting > 0 {
		if err := templates.ExecuteTemplate(w, "line", metrics); err != nil {
			log.Println(err)
		}
	} else {
		if err := templates.ExecuteTemplate(w, "line_a", metrics); err != nil {
			log.Println(err)
		}
	}
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Println("arguments expected: listening Port")
		return
	}
	port := ":" + args[1]

	// http routes
	// http://127.0.0.1:<port>/
	http.HandleFunc("/ping", defaultHandler)

	// Example to get the IEQ numbers
	// http://localhost:<port>/ieq/numbers?device_id=awair-omni_18453
	http.HandleFunc("/ieq/numbers", ieqNumbersHandler)

	// Example to get the IEQ charts
	// http://localhost:<port>/ieq/device?device_id=awair-omni_18453
	http.HandleFunc("/ieq/device", ieqchartsHandler)

	// Temporary, this is for static mockup test only
	//http.Handle("/moqup", http.FileServer(http.Dir("./static"))) // serves index.html
	http.Handle("/moqup/", http.StripPrefix("/moqup/", http.FileServer(http.Dir("./static"))))

	log.Printf("IEQ up and running at port%s ...\n", port)
	http.ListenAndServe(port, nil)
}
