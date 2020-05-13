package controllers

import (
	"net/http"
	"strconv"

	"../database"
	"../models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CreateUser creates a new user
func CreateUser(context *gin.Context) {

	zoneData := new(models.AccountCreationData)
	jsonDataError := context.BindJSON(&zoneData)

	if jsonDataError != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": "Invalid json format"})
		return
	}

	if zoneData.PostalCode == "" || zoneData.City == "" {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": "Postal code and city cannot be empty"})
		return
	}

	guid, guidError := uuid.NewUUID()
	if guidError != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": "Error creating guid"})
		return
	}

	insertedID, dbError := database.InsertUser(guid.String(), zoneData.PostalCode, zoneData.City, zoneData.IsLocationFromAPI)

	if dbError != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": "Error db insertion"})
		return
	}

	context.JSON(http.StatusOK, models.CitizenIDAndGUID{ID: insertedID, UserGUID: guid.String()})
}

// UpdateIsInsideStatus updates status
func UpdateIsInsideStatus(context *gin.Context) {

	maplel := context.Request.Header
	userGUID := maplel["Userguid"]
	isInside := context.Request.URL.Query()["IsInside"]
	rez, _ := strconv.ParseBool(isInside[0])

	updatedRows, dbError := database.UpdateStatus(userGUID[0], rez)

	if dbError != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": "Error db insertion"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"Message": updatedRows})
}

//AuthMiddleWare checks for user guid
func AuthMiddleWare(context *gin.Context) {
	userGUID := context.Request.Header["Userguid"]
	if len(userGUID) == 0 {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Missing Token"})
		return
	}

	userID, dbError := database.GetUserIDByGUID(userGUID[0])

	if dbError != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Error db insertion"})
		return
	}

	if userID == -1 {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Error": "Invalid Token"})
		return
	}
}
