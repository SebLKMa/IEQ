package awair

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/christophwitzko/go-curl"
	mdl "github.com/seblkma/ieq/models"
)

// GetState implements device interface Device.GetState()
// Uses sensor cloud API to get device information.
func (sensor *SensorInfo) GetState(id string) (result string, err error) {

	apiurl := "https://developer-apis.awair.is/v1/orgs/" + sensor.Org + "/devices/awair-omni/" + id
	cb := func(st curl.IoCopyStat) error {
		fmt.Println(st.Stat)
		if st.Response != nil {
			fmt.Println(st.Response.Status)
			if st.Response.StatusCode != 200 {
				return errors.New(st.Response.Status)
			}
		}
		return nil
	}

	err, str, _ := curl.String(apiurl, cb, "method=", "GET",
		"disablecompression=", true,
		"header=", http.Header{
			"accept":       {"application/json"},
			"x-api-key":    {sensor.Token},
			"Content-Type": {"application/json"},
		},
	)
	if err != nil {
		return "", err
	}

	return str, nil
}

// GetRawMetrics implements device interface Device.GetLatestMetrics()
// Uses sensor cloud API to get metrics values.
func (sensor *SensorInfo) GetRawMetrics(id string) (result string, err error) {

	//var POST_DATA = "{ \"ip\": \"derrickyeo90.synology.me\", \"size\": 0, \"df_bit\": true, \"source\": \"\", \"vrf\": false}"

	//string1 := "{ \"ip\": \""
	//string2 := "\", \"size\": 0, \"df_bit\": true, \"source\": \"\", \"vrf\": false}"
	//POST_DATA := string1 + pingIp + string2
	//fmt.Println(POST_DATA)

	cb := func(st curl.IoCopyStat) error {
		fmt.Println(st.Stat)
		if st.Response != nil {
			fmt.Println(st.Response.Status)
			if st.Response.StatusCode != 200 {
				return errors.New(st.Response.Status)
			}
		}
		return nil
	}

	// curl "https://developer-apis.awair.is/v1/orgs/3332/devices/awair-omni/18453/air-data/latest" -H "x-api-key: <awair subscription key>"
	//err, str, resp := curl.String(

	apiurl := "https://developer-apis.awair.is/v1/orgs/" + sensor.Org + "/devices/awair-omni/" + id + "/air-data/latest"
	err, str, _ := curl.String(apiurl, cb, "method=", "GET",
		/*"data=", strings.NewReader(POST_DATA),*/
		"disablecompression=", true,
		"header=", http.Header{
			"accept":       {"application/json"},
			"x-api-key":    {sensor.Token},
			"Content-Type": {"application/json"},
		},
	)
	if err != nil {
		return "", err
	}
	//fmt.Println(resp.Header)
	//fmt.Println(str)
	return str, nil
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
	var g map[string]interface{}

	bytes := []byte(jsonData)
	err = json.Unmarshal(bytes, &g)
	if err != nil {
		result.Status = 0
		result.StatusDescription = err.Error()
		log.Println(err)
		return result, err
	}

	// iterate to get the data we want
	for k, v := range g {
		switch k {
		case "uuid":
			// Make sure is a string
			switch itemValue := v.(type) {
			case string:
				result.DeviceID = itemValue
			default:
				fmt.Println("Incorrect type for DeviceID: ", k)
			}
		case "mac_address":
			// Make sure is a string
			switch itemValue := v.(type) {
			case string:
				result.MacAddress = itemValue
			default:
				fmt.Println("Incorrect type for MacAddress: ", k)
			}
		case "display_name":
			// Make sure is a string
			switch itemValue := v.(type) {
			case string:
				result.DisplayName = itemValue
			default:
				fmt.Println("Incorrect type for DisplayName: ", k)
			}
		case "org_id":
			// Make sure is a string
			switch itemValue := v.(type) {
			case string:
				result.Org = itemValue
			default:
				fmt.Println("Incorrect type for Org: ", k)
			}
		case "connected":
			// Make sure is bool
			switch itemValue := v.(type) {
			case bool:
				if itemValue {
					result.Status = 1
					result.StatusDescription = "connected"
				} else {
					result.Status = 0
					result.StatusDescription = "disconnected"
				}
			default:
				fmt.Println("Incorrect type for Status: ", k)
			}
		}
	}

	fmt.Println(result)
	return result, nil
}
