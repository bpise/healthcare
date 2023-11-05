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

// getTemperatureInRegionMin - current minimum temperature inside the region
func getTemperatureInRegionMin(c *gin.Context) {
	params := getRegionParams(c)

	min := new(float32)
	err := db.GetSingleRow(tmplTemperatureInRegionMin, params, &min)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"region": params, "TemperatureInRegionMin": min})
}

// getTemperatureInRegionMax - current maximum temperature inside the region
func getTemperatureInRegionMax(c *gin.Context) {
	params := getRegionParams(c)

	max := new(float32)
	err := db.GetSingleRow(tmplTemperatureInRegionMax, params, &max)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"region": params, "TemperatureInRegionMax": max})
}

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
