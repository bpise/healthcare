package webservice

import (
	"os"

	"healthcare/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// default_PORT - Default port for the HTTP service.
const default_PORT = "8080"

var webEngine *gin.Engine

// InitWebEngine - Initializes the HTTP service if it's not already initialized.
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

// Run - Runs the HTTP service if it's initialized.
func Run() error {
	if webEngine != nil {
		return webEngine.Run()
	}
	return nil
}

// middlewareAuth - Authentication middleware for handling requests.
func middlewareAuth(c *gin.Context) {
	logger.Infof("exec middleware for authentication...")
	c.Next()
}

// InitRouter - Initializes the sensor statistics APIs and sets up routes.
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

	// supports swagger docs
	webEngine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
