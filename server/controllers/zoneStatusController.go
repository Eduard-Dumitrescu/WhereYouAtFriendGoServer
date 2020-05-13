package controllers

import (
	"net/http"

	"../database"
	"../models"
	"github.com/gin-gonic/gin"
)

// GetZoneStatusByPostalCodeAndCity creates a new user
func GetZoneStatusByPostalCodeAndCity(context *gin.Context) {

	zoneData := new(models.AccountCreationData)
	jsonDataError := context.BindQuery(&zoneData)

	if jsonDataError != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Invalid json format"})
		return
	}

	if zoneData.PostalCode == "" || zoneData.City == "" {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Postal code and city cannot be empty"})
		return
	}

	selectZoneStatus, dbError := database.GetZoneStatus(zoneData.PostalCode, zoneData.City)

	if dbError != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": "Error db retrieval"})
		return
	}

	context.JSON(http.StatusOK, selectZoneStatus)
}
