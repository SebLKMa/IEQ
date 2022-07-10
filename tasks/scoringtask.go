package tasks

import (
	"fmt"
	"log"
	"os"

	conf "github.com/seblkma/ieq/configs"
	fml "github.com/seblkma/ieq/formulas"
	rate "github.com/seblkma/ieq/ratings"
	util "github.com/seblkma/ieq/utils"
	"gopkg.in/yaml.v3"
)

// ScoringTask properties
type ScoringTask struct {
	TemperatureFormula *fml.StandardFormula
	HumidityFormula    *fml.StandardFormula
	Co2Formula         *fml.MinIsGoodFormula
	VocFormula         *fml.MinIsGoodFormula
	Pm25Formula        *fml.MinIsGoodFormula
	NoiseFormula       *fml.MinIsGoodFormula
	LightingFormula    *fml.LightingFormula
	Cfg                *conf.AppConfig
	Initialized        bool
}

// NewScoringTask constructs a new ScoreTask instance
// ScoreTask properties will be initialized from configuration (file or database).
func NewScoringTask(configFile string) *ScoringTask {
	if !util.FileExists(configFile) {
		log.Fatalf("%s file not found in current directory.", configFile)
	}

	f, err := os.Open(configFile)
	if err != nil {
		log.Fatal(err)
	}
	//defer f.Close() // not deferring, close immediately after decode
	decoder := yaml.NewDecoder(f)

	task := ScoringTask{Initialized: false}

	err = decoder.Decode(&task.Cfg)
	if err != nil {
		log.Fatal("Failed to decode config yaml file. ", err)
	}
	f.Close()
	fmt.Printf("Config:\n%v\n", task.Cfg)

	task.TemperatureFormula = &fml.StandardFormula{}
	task.HumidityFormula = &fml.StandardFormula{}
	task.Co2Formula = &fml.MinIsGoodFormula{}
	task.VocFormula = &fml.MinIsGoodFormula{}
	task.Pm25Formula = &fml.MinIsGoodFormula{}
	task.NoiseFormula = &fml.MinIsGoodFormula{}
	task.LightingFormula = &fml.LightingFormula{}

	min := task.Cfg.Temperature.Min
	max := task.Cfg.Temperature.Max
	rate.Setup(task.TemperatureFormula, "Temperature", min, max)
	rate.PrintInfo(task.TemperatureFormula)

	min = task.Cfg.Humidity.Min
	max = task.Cfg.Humidity.Max
	rate.Setup(task.HumidityFormula, "Humidity", min, max)
	rate.PrintInfo(task.HumidityFormula)

	min = task.Cfg.CO2.Min
	max = task.Cfg.CO2.Max
	rate.Setup(task.Co2Formula, "CO2", min, max)
	rate.PrintInfo(task.Co2Formula)

	min = task.Cfg.VOC.Min
	max = task.Cfg.VOC.Max
	rate.Setup(task.VocFormula, "VOC", min, max)
	rate.PrintInfo(task.VocFormula)

	min = task.Cfg.PM25.Min
	max = task.Cfg.PM25.Max
	rate.Setup(task.Pm25Formula, "PM25", min, max)
	rate.PrintInfo(task.Pm25Formula)

	min = task.Cfg.NOISE.Min
	max = task.Cfg.NOISE.Max
	rate.Setup(task.NoiseFormula, "Noise", min, max)
	rate.PrintInfo(task.NoiseFormula)

	min = task.Cfg.LIGHTING.Min
	max = task.Cfg.LIGHTING.Max
	task.LightingFormula.SetScale(task.Cfg.LIGHTING.Scale) // Scale must be set first
	rate.Setup(task.LightingFormula, "Lighting", min, max)
	rate.PrintInfo(task.LightingFormula)

	task.Initialized = true

	return &task
}
