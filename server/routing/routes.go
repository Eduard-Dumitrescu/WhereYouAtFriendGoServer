package routing

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter setup routing
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/createCitizen", controllers.CreateUser)

	authorized := r.Group("/general")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(controllers.AuthMiddleWare)
	{
		authorized.GET("/zoneStatus", controllers.GetZoneStatusByPostalCodeAndCity)
		authorized.PUT("/updateIsInsideStatus", controllers.UpdateIsInsideStatus)
	}

	return r
}
