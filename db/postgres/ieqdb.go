package db

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq" // the db driver supports database/sql
	mdl "github.com/seblkma/ieq/models"
)

// CreateMetric creates metrics record in database
func CreateMetric(data mdl.Metrics) error {
	db := connect()
	defer db.Close()

	sqlstmt := "INSERT INTO metrics(device_id, created_on, temperature, humidity, co2, voc, pm25, lighting, noise) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err := db.Exec(sqlstmt,
		data.DeviceID,
		time.Now(),
		data.Temperature,
		data.Humidity,
		data.CO2,
		data.VOC,
		data.PM25,
		data.Lighting,
		data.Noise)

	if err != nil {
		return err
	}

	fmt.Println("metrics created in database")
	return nil
}

// CreateMetricScore creates metricscores record in database
func CreateMetricScore(data mdl.MetricScore) error {
	db := connect()
	defer db.Close()

	sqlstmt := "INSERT INTO metricscores(device_id, created_on, temperature, humidity, co2, voc, pm25, lighting, noise) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err := db.Exec(sqlstmt,
		data.DeviceID,
		time.Now(),
		data.Temperature,
		data.Humidity,
		data.CO2,
		data.VOC,
		data.PM25,
		data.Lighting,
		data.Noise)

	if err != nil {
		return err
	}

	fmt.Println("metricscores created in database")
	return nil
}

// CreateIeqScore creates ieqscores record in database
func CreateIeqScore(data mdl.IeqScore) error {
	db := connect()
	defer db.Close()

	sqlstmt := "INSERT INTO ieqscores(device_id, created_on, scheme, thermal, iaq, lighting, noise, overall, thermal_weighting, iaq_weighting, lighting_weighting, noise_weighting) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
	_, err := db.Exec(sqlstmt,
		data.DeviceID,
		time.Now(),
		data.Scheme,
		data.Thermal,
		data.IAQ,
		data.Lighting,
		data.Noise,
		data.Overall,
		data.ThermalWeighting,
		data.IAQWeighting,
		data.LightingWeighting,
		data.NoiseWeighting)

	if err != nil {
		return err
	}

	fmt.Println("ieqscores created in database")
	return nil
}

// ReadLatestMetrics returns the latest metrics record
func ReadLatestMetrics(deviceID string) (data mdl.Metrics, err error) {
	data = mdl.Metrics{}
	db := connect()
	defer db.Close()

	stmt := "SELECT * FROM metrics WHERE device_id = $1 ORDER BY rowid DESC LIMIT 1"

	rows, err := db.Query(stmt, deviceID)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	var rowid int
	var results []mdl.Metrics
	for rows.Next() {
		item := mdl.Metrics{}
		err2 := rows.Scan(&rowid, &item.DeviceID, &item.CreatedOn, &item.Temperature, &item.Humidity, &item.CO2, &item.VOC, &item.PM25, &item.Lighting, &item.Noise)
		if err2 != nil {
			return data, err
		}
		results = append(results, item)
	}

	if len(results) == 0 {
		return data, errors.New("No record found")
	}

	// just the latest record
	data = results[0]
	return data, nil
}

// ReadLatestMetricScores returns the latest metricscores record
func ReadLatestMetricScores(deviceID string) (data mdl.MetricScore, err error) {
	data = mdl.MetricScore{}
	db := connect()
	defer db.Close()

	stmt := "SELECT * FROM metricscores WHERE device_id = $1 ORDER BY rowid DESC LIMIT 1"

	rows, err := db.Query(stmt, deviceID)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	var rowid int
	var results []mdl.MetricScore
	for rows.Next() {
		item := mdl.MetricScore{}
		err2 := rows.Scan(&rowid, &item.DeviceID, &item.CreatedOn, &item.Temperature, &item.Humidity, &item.CO2, &item.VOC, &item.PM25, &item.Lighting, &item.Noise)
		if err2 != nil {
			return data, err
		}
		results = append(results, item)
	}

	if len(results) == 0 {
		return data, errors.New("No record found")
	}

	// just the latest record
	data = results[0]
	return data, nil
}

// ReadLatestIeqScores returns the latest ieqscores record
func ReadLatestIeqScores(deviceID string) (data mdl.IeqScore, err error) {
	data = mdl.IeqScore{}
	db := connect()
	defer db.Close()

	stmt := "SELECT * FROM ieqscores WHERE device_id = $1 ORDER BY rowid DESC LIMIT 1"

	rows, err := db.Query(stmt, deviceID)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	var rowid int
	var results []mdl.IeqScore
	for rows.Next() {
		item := mdl.IeqScore{}
		err2 := rows.Scan(&rowid, &item.DeviceID, &item.CreatedOn, &item.Scheme, &item.Thermal, &item.IAQ, &item.Lighting, &item.Noise, &item.Overall, &item.ThermalWeighting, &item.IAQWeighting, &item.LightingWeighting, &item.NoiseWeighting)
		if err2 != nil {
			return data, err
		}
		results = append(results, item)
	}

	if len(results) == 0 {
		return data, errors.New("No record found")
	}

	// just the latest record
	data = results[0]
	return data, nil
}

// ReadMetrics returns just the metrics record
func ReadMetrics(deviceID string, count int) (results []mdl.Metrics, err error) {
	results = []mdl.Metrics{}
	db := connect()
	defer db.Close()

	stmt := "SELECT * FROM metrics WHERE device_id = $1 ORDER BY rowid DESC LIMIT " + strconv.Itoa(count)

	rows, err := db.Query(stmt, deviceID)
	if err != nil {
		return results, err
	}
	defer rows.Close()

	var rowid int
	for rows.Next() {
		item := mdl.Metrics{}
		err2 := rows.Scan(&rowid, &item.DeviceID, &item.CreatedOn, &item.Temperature, &item.Humidity, &item.CO2, &item.VOC, &item.PM25, &item.Lighting, &item.Noise)
		if err2 != nil {
			return results, err
		}
		results = append(results, item)
	}

	if len(results) == 0 {
		return results, errors.New("No record found")
	}

	return results, nil
}
