package configs

// AppConfig will be populated from yaml file
type AppConfig struct {
	Temperature struct {
		Min float64 `yaml:"Min"`
		Max float64 `yaml:"Max"`
	} `yaml:"temperature"`
	Humidity struct {
		Min float64 `yaml:"Min"`
		Max float64 `yaml:"Max"`
	} `yaml:"humidity"`
	CO2 struct {
		Min float64 `yaml:"Min"`
		Max float64 `yaml:"Max"`
	} `yaml:"co2"`
	VOC struct {
		Min float64 `yaml:"Min"`
		Max float64 `yaml:"Max"`
	} `yaml:"voc"`
	PM25 struct {
		Min float64 `yaml:"Min"`
		Max float64 `yaml:"Max"`
	} `yaml:"pm25"`
	NOISE struct {
		Min float64 `yaml:"Min"`
		Max float64 `yaml:"Max"`
	} `yaml:"noise"`
	LIGHTING struct {
		Min   float64 `yaml:"Min"`
		Max   float64 `yaml:"Max"`
		Scale float64 `yaml:"Scale"`
	} `yaml:"lighting"`
	WEIGHTINGS struct {
		Scheme   string  `yaml:"scheme"`
		Thermal  float64 `yaml:"thermal"`
		IAQ      float64 `yaml:"iaq"`
		Lighting float64 `yaml:"lighting"`
		Noise    float64 `yaml:"noise"`
	} `yaml:"weightings"`
	TASK struct {
		Minutes int `yaml:"minutes"`
	} `yaml:"task"`
	VENDOR struct {
		Name            string `yaml:"Name"`
		DeviceDisplayID string `yaml:"DeviceDisplayID"`
		DeviceID        string `yaml:"DeviceID"`
		Org             string `yaml:"Org"`
		Token           string `yaml:"Token"`
	} `yaml:"vendor"`
}
