package database

import (
	"database/sql"
	"fmt"

	//use mysql driver instead of database/sql
	_ "github.com/go-sql-driver/mysql"
)

// InsertUser inserts new user to database and returns
func InsertUser(userGUID string, postalCode string, city string, IsLocationFromAPI bool) (int, error) {
	conn, connectionError := sql.Open("mysql", DbURLDefault())
	defer conn.Close()

	if connectionError != nil {
		errString := fmt.Sprintf("Error opening db connection : %s", connectionError)
		fmt.Println(errString)
		return -1, fmt.Errorf(errString)
	}

	insertStatement, insertStatementError := conn.Prepare("INSERT INTO Citizens (UserGuid, PostalCode, City, IsLocationFromApi) VALUES (?, ?, ?, ?);")

	if insertStatementError != nil {
		errString := fmt.Sprintf("Error creating insert statement : %s ", insertStatementError)
		fmt.Println(errString)
		return -1, fmt.Errorf(errString)
	}

	result, insertExecError := insertStatement.Exec(userGUID, postalCode, city, IsLocationFromAPI)
	if insertExecError != nil {
		errString := fmt.Sprintf("Error executing insert statement : %s ", insertExecError)
		fmt.Println(errString)
		return -1, fmt.Errorf(errString)
	}

	insertedID, idError := result.LastInsertId()
	if idError != nil {
		errString := fmt.Sprintf("Error getting inserted id but maybe insert worked : %s ", idError)
		fmt.Println(errString)
		return -1, fmt.Errorf(errString)
	}

	return int(insertedID), nil
}

// UpdateStatus update user status
func UpdateStatus(userGUID string, status bool) (int, error) {
	conn, connectionError := sql.Open("mysql", DbURLDefault())
	defer conn.Close()

	if connectionError != nil {
		errString := fmt.Sprintf("Error opening db connection : %s", connectionError)
		fmt.Println(errString)
		return -1, fmt.Errorf(errString)
	}

	insertStatement, insertStatementError := conn.Prepare("UPDATE Citizens SET IsInside=? WHERE UserGuid=?;")

	if insertStatementError != nil {
		errString := fmt.Sprintf("Error creating insert statement : %s ", insertStatementError)
		fmt.Println(errString)
		return -1, fmt.Errorf(errString)
	}

	result, insertExecError := insertStatement.Exec(status, userGUID)
	if insertExecError != nil {
		errString := fmt.Sprintf("Error executing update statement : %s ", insertExecError)
		fmt.Println(errString)
		return -1, fmt.Errorf(errString)
	}

	RowsAffected, idError := result.RowsAffected()
	if idError != nil {
		errString := fmt.Sprintf("Error getting inserted id but maybe insert worked : %s ", idError)
		fmt.Println(errString)
		return -1, fmt.Errorf(errString)
	}

	return int(RowsAffected), nil
}

// GetUserIDByGUID data
func GetUserIDByGUID(userGUID string) (int, error) {
	conn, connectionError := sql.Open("mysql", DbURLDefault())
	defer conn.Close()

	if connectionError != nil {
		panic(connectionError.Error())
	}

	results, err := conn.Query("SELECT Id FROM Citizens WHERE UserGUID=?;", userGUID)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	rez := -1

	for results.Next() {
		var err = results.Scan(&rez)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
	}

	return rez, nil
}
