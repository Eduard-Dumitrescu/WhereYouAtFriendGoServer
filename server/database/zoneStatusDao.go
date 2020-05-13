package database

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"../models"
)

var (
	wg sync.WaitGroup // wg is used to wait for the program to finish.
)

// GetCitizensZoneStatus inserts new user to database and returns
func GetCitizensZoneStatus() ([]models.ZonesStatusPlaceAndCount, error) {
	conn, connectionError := sql.Open("mysql", DbURLDefault())
	defer conn.Close()

	if connectionError != nil {
		panic(connectionError.Error())
	}

	results, err := conn.Query("SELECT PostalCode, City, COUNT(if(IsInside = 1, 1, NULL)) AS InsideCount, COUNT(*) AS total FROM Citizens GROUP BY PostalCode, City")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	rezArr := make([]models.ZonesStatusPlaceAndCount, 1)

	for results.Next() {
		zoneStatusCount := models.ZonesStatusPlaceAndCount{}
		var err = results.Scan(&zoneStatusCount.PostalCode, &zoneStatusCount.City, &zoneStatusCount.IsInsideCount, &zoneStatusCount.TotalCount)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		rezArr = append(rezArr, zoneStatusCount)
	}

	return rezArr, nil
}

// GetZoneStatus data
func GetZoneStatus(postalCode string, city string) ([]models.ZonesStatusCount, error) {
	conn, connectionError := sql.Open("mysql", DbURLDefault())
	defer conn.Close()

	if connectionError != nil {
		panic(connectionError.Error())
	}

	results, err := conn.Query("SELECT InsideCount, TotalCount FROM ZoneStatus WHERE PostalCode=? AND City=?;", postalCode, city)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	rezArr := make([]models.ZonesStatusCount, 0)

	for results.Next() {
		zoneStatusCount := models.ZonesStatusCount{}
		var err = results.Scan(&zoneStatusCount.IsInsideCount, &zoneStatusCount.TotalCount)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		rezArr = append(rezArr, zoneStatusCount)
	}

	return rezArr, nil
}

// UpdateZoneStatus inserts new user to database and returns
func UpdateZoneStatus(zonesCount []models.ZonesStatusPlaceAndCount) error {
	conn, connectionError := sql.Open("mysql", DbURLDefault())
	defer conn.Close()

	if connectionError != nil {
		panic(connectionError.Error())
	}

	insertStatement, insertStatementError := conn.Prepare("INSERT INTO ZoneStatus (PostalCode, City, InsideCount, TotalCount) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE InsideCount = ?, TotalCount = ?;")
	//insertStatement, insertStatementError := conn.Prepare("UPDATE ZoneStatus SET InsideCount = ?, TotalCount = ? WHERE PostalCode = ? AND City = ?;")

	if insertStatementError != nil {
		errString := fmt.Sprintf("Error creating insert statement : %s ", insertStatementError)
		fmt.Println(errString)
		panic(insertStatementError.Error())
	}

	for _, el := range zonesCount {
		_, insertExecError := insertStatement.Exec(el.PostalCode, el.City, el.IsInsideCount, el.TotalCount, el.IsInsideCount, el.TotalCount)
		//_, insertExecError := insertStatement.Exec(el.IsInsideCount, el.TotalCount, el.PostalCode, el.City)
		if insertExecError != nil {
			errString := fmt.Sprintf("Error executing insert statement : %s ", insertExecError)
			fmt.Println(errString)
			panic(insertExecError.Error())
		}
	}

	return nil
}

// StartZoneStatusUpdate populates ZoneStatus every second
func StartZoneStatusUpdate() {
	isDone := true

	go func() {
		for range time.Tick(time.Second * 1) {
			if isDone {
				isDone = false
				go updateDb(&isDone)
			}
		}
	}()
}

func updateDb(done *bool) {
	zonesStatus, _ := GetCitizensZoneStatus()
	UpdateZoneStatus(zonesStatus)
	*done = true
	fmt.Println("Table updated")
}
