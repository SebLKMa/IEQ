package db

import (
	"errors"
	"fmt"
	"time"

	_ "github.com/lib/pq" // the db driver supports database/sql
	mdl "github.com/seblkma/ieq/models"
)

// CreateDeviceStatus creates new record in database
func CreateDeviceStatus(data mdl.DeviceInfo) error {
	db := connect()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO devicestatus(device_id, created_on, status, status_desc) VALUES($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(data.DeviceID, time.Now(), data.Status, data.StatusDescription)
	if err != nil {
		return err
	}
	fmt.Println("device status created in database")
	return nil
}

// ReadLastDeviceStatus returns the last record
func ReadLastDeviceStatus(deviceID string) (data mdl.DeviceInfo, err error) {
	//data = DeviceStatus{}
	db := connect()
	defer db.Close()

	stmt := "SELECT * FROM devicestatus WHERE device_id = $1 ORDER BY rowid DESC LIMIT 1"

	rows, err := db.Query(stmt, deviceID)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	var rowid int
	var results []mdl.DeviceInfo
	for rows.Next() {
		item := mdl.DeviceInfo{}
		err2 := rows.Scan(&rowid, &item.DeviceID, &item.CreatedOn, &item.Status, &item.StatusDescription)
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
