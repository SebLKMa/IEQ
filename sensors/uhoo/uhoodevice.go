package uhoo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	mdl "github.com/seblkma/ieq/models"
)

// GetState implements device interface Device.GetState()
// Uses sensor cloud API to get device information.
func (sensor *SensorInfo) GetState(id string) (result string, err error) {
	apiurl := "https://api.uhooinc.com/v1/getdevicelist"
	token := sensor.Token
	org := sensor.Org
	apidata := "username=" + org + "&" + "password=" + token

	//reqBody := strings.NewReader(`username=sebmaspd@gmail.com&password=cbda313934cddaa638e13aca7c4dbf9a5a82260d2b899dda088c702c69a14eac`)
	reqBody := strings.NewReader(apidata)
	req, err := http.NewRequest("POST", apiurl, reqBody)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}

// GetRawMetrics implements device interface Device.GetLatestMetrics()
// Uses sensor cloud API to get metrics values.
func (sensor *SensorInfo) GetRawMetrics(id string) (result string, err error) {
	apiurl := "https://api.uhooinc.com/v1/getlatestdata"
	token := sensor.Token
	org := sensor.Org
	apidata := "username=" + org + "&" + "password=" + token + "&" + "serialNumber=" + id

	//reqBody = strings.NewReader(`username=sebmaspd@gmail.com&password=cbda313934cddaa638e13aca7c4dbf9a5a82260d2b899dda088c702c69a14eac&serialNumber=52ff6b066571525520131987`)
	reqBody := strings.NewReader(apidata)
	req, err := http.NewRequest("POST", apiurl, reqBody)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}

// GetLatestMetrics gets the latest raw metrics from device
func (sensor *SensorInfo) GetLatestMetrics(deviceID string) (result mdl.Metrics, err error) {
	result = mdl.Metrics{Empty: true, CreatedOn: time.Now()}

	jsonData, err := sensor.GetRawMetrics(deviceID)
	//fmt.Printf("%v\n", jsonData)
	if err != nil {
		log.Println(err)
		return result, err
	}

	// Unmarshal using a generic interface
	var g map[string]interface{}

	bytes := []byte(jsonData)
	err = json.Unmarshal(bytes, &g)
	if err != nil {
		log.Println(err)
		return result, err
	}

	// iterate to get the data we want
	for k, v := range g {
		switch k {
		case "Temperature":
			// Make sure is a string
			switch itemValue := v.(type) {
			case string:
				result.Empty = false
				result.Temperature, err = strconv.ParseFloat(itemValue, 64)
				if err != nil {
					return result, err
				}
			default:
				return result, errors.New("Incorrect type for Temperature")
			}
		case "Relative Humidity":
			// Make sure is a string
			switch itemValue := v.(type) {
			case string:
				result.Empty = false
				result.Humidity, err = strconv.ParseFloat(itemValue, 64)
				if err != nil {
					return result, err
				}
			default:
				return result, errors.New("Incorrect type for Humidity")
			}
		case "CO2":
			// Make sure is a string
			switch itemValue := v.(type) {
			case string:
				result.Empty = false
				result.CO2, err = strconv.ParseFloat(itemValue, 64)
				if err != nil {
					return result, err
				}
			default:
				return result, errors.New("Incorrect type for CO2")
			}
		case "TVOC":
			// Make sure is a string
			switch itemValue := v.(type) {
			case string:
				result.Empty = false
				result.VOC, err = strconv.ParseFloat(itemValue, 64)
				if err != nil {
					return result, err
				}
			default:
				return result, errors.New("Incorrect type for VOC")
			}
		case "PM2.5":
			// Make sure is a string
			switch itemValue := v.(type) {
			case string:
				result.Empty = false
				result.PM25, err = strconv.ParseFloat(itemValue, 64)
				if err != nil {
					return result, err
				}
			default:
				return result, errors.New("Incorrect type for PM25")
			}
		}
	}

	return result, nil
}

// GetDeviceInfo returns the current information of the device
func (sensor *SensorInfo) GetDeviceInfo(id string) (result mdl.DeviceInfo, err error) {
	result.VendorDeviceID = id
	result.CreatedOn = time.Now()
	jsonData, err := sensor.GetState(id)
	//fmt.Printf("%v\n", jsonData)
	if err != nil {
		result.Status = 0
		result.StatusDescription = err.Error()
		log.Println(err)
		return result, err
	}

	// Unmarshal using a generic interface
	var g interface{}

	bytes := []byte(jsonData)
	err = json.Unmarshal(bytes, &g)
	if err != nil {
		result.Status = 0
		result.StatusDescription = err.Error()
		log.Println(err)
		return result, err
	}

	// s is slice
	s, ok := g.([]interface{})
	if !ok {
		return result, errors.New("JSON type error")
	}

	// vendor API has no way to tell if device is online or offline
	result.Status = 1

	// iterate to get the data we want
	for k, v := range s {
		switch vt := v.(type) {
		case map[string]interface{}:
			fmt.Println(k, "is a map:")
			for kk, vv := range vt {
				//fmt.Println(kk, vv)
				switch kk {
				case "deviceName":
					switch itemValue := vv.(type) {
					case string:
						result.DeviceID = itemValue
						result.DisplayName = itemValue
					default:
						fmt.Println("Incorrect type for DeviceID: ", k)
					}
				case "serialNumber":
					switch itemValue := vv.(type) {
					case string:
						result.SerialNumber = itemValue
					default:
						fmt.Println("Incorrect type for SerialNumber: ", k)
					}
				case "company":
					switch itemValue := vv.(type) {
					case string:
						result.Org = itemValue
					default:
						fmt.Println("Incorrect type for Org: ", k)
					}
				case "macAddress":
					switch itemValue := vv.(type) {
					case string:
						result.MacAddress = itemValue
					default:
						fmt.Println("Incorrect type for MacAddress: ", k)
					}
				}

			}
		default:
			return result, errors.New("Unexpected JSON collection structure")
		}
	}

	return result, nil
}
