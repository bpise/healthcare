package webservice

import (
	"os"

	"healthcare/logger"

	"github.com/gin-gonic/gin"
)

const default_PORT = "8080"

var webEngine *gin.Engine

func InitWebEngine() {
	if webEngine != nil {
		return
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = default_PORT
	}
	logger.Infof("Starting HTTP service on port %v ... ", port)

	webEngine = gin.Default()
}

func Run() error {
	if webEngine != nil {
		return webEngine.Run()
	}
	return nil
}

func middlewareAuth(c *gin.Context) {
	logger.Infof("exec middleware for authentication...")
	c.Next()
}

func InitRouter() {
	group := webEngine.Group("/group", middlewareAuth)
	group.POST("/:groupName", setupGroup)
	group.GET("/:groupName/transparency/average", getTransparencyAverage)
	group.GET("/:groupName/temperature/average", getTemperatureAverage)
	group.GET("/:groupName/species", getSpecies)
	group.GET("/:groupName/species/top/:n", getSpeciesTopN)

	region := webEngine.Group("/region", middlewareAuth)
	region.GET("/temperature/min", getTemperatureInRegionMin)
	region.GET("/temperature/max", getTemperatureInRegionMax)

	sensor := webEngine.Group("/sensor", middlewareAuth)
	sensor.GET("/:codeName/temperature/average", getTemperatureAverageBySensor)
}
