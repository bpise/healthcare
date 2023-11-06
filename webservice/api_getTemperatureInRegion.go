package webservice

import (
	"html/template"
	"net/http"

	"healthcare/db"

	"github.com/gin-gonic/gin"
)

const (
	default_Min = "0"
	default_MAx = "1000000"
)

var (
	tmplTemperatureInRegionMin = template.Must(template.New("TemperatureInRegionMin").Parse(db.TemperatureInRegionMinSQLText))
	tmplTemperatureInRegionMax = template.Must(template.New("TemperatureInRegionMax").Parse(db.TemperatureInRegionMaxSQLText))
)

// Sensor Statistics APIs
// @title getTemperatureInRegionMin
// @Summary Retrieves the current minimum temperature inside the region.
// @Tags Tested
// @Param	xMin	query	int		false	"minimum of X in region(3D-coordinates) such as 10"
// @Param	xMax	query	int		false	"maximum of X in region(3D-coordinates) such as 10"
// @Param	yMin	query	int		false	"minimum of Y in region(3D-coordinates) such as 10"
// @Param	yMax	query	int		false	"maximum of Y in region(3D-coordinates) such as 10"
// @Param	zMin	query	int		false	"minimum of Z in region(3D-coordinates) such as 10"
// @Param	zMax	query	int		false	"maximum of Z in region(3D-coordinates) such as 10"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /region/temperature/min [get]
func getTemperatureInRegionMin(c *gin.Context) {
	// Get region parameters from the request.
	params := getRegionParams(c)

	// Query the database for the minimum temperature in the region.
	min := new(float32)
	err := db.GetSingleRow(tmplTemperatureInRegionMin, params, &min)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	// Return the minimum temperature in the region as JSON response.
	c.JSON(http.StatusOK, gin.H{"region": params, "TemperatureInRegionMin": min})
}

// Sensor Statistics APIs
// @title getTemperatureInRegionMax
// @Summary Retrieves the current maximum temperature inside the region.
// @Tags Tested
// @Param	xMin	query	int		false	"minimum of X in region(3D-coordinates) such as 10"
// @Param	xMax	query	int		false	"maximum of X in region(3D-coordinates) such as 10"
// @Param	yMin	query	int		false	"minimum of Y in region(3D-coordinates) such as 10"
// @Param	yMax	query	int		false	"maximum of Y in region(3D-coordinates) such as 10"
// @Param	zMin	query	int		false	"minimum of Z in region(3D-coordinates) such as 10"
// @Param	zMax	query	int		false	"maximum of Z in region(3D-coordinates) such as 10"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /region/temperature/max [get]
func getTemperatureInRegionMax(c *gin.Context) {
	// Get region parameters from the request.
	params := getRegionParams(c)

	// Query the database for the maximum temperature in the region.
	max := new(float32)
	err := db.GetSingleRow(tmplTemperatureInRegionMax, params, &max)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	// Return the maximum temperature in the region as JSON response.
	c.JSON(http.StatusOK, gin.H{"region": params, "TemperatureInRegionMax": max})
}

// getRegionParams - Parses region parameters from the request and returns them as a map.
func getRegionParams(c *gin.Context) *map[string]string {
	param := make(map[string]string)
	param["xMin"] = c.DefaultQuery("xMin", default_Min)
	param["xMax"] = c.DefaultQuery("xMax", default_MAx)
	param["yMin"] = c.DefaultQuery("yMin", default_Min)
	param["yMax"] = c.DefaultQuery("yMax", default_MAx)
	param["zMin"] = c.DefaultQuery("zMin", default_Min)
	param["zMax"] = c.DefaultQuery("zMax", default_MAx)

	return &param
}
