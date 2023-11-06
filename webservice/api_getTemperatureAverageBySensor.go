package webservice

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"healthcare/db"

	"github.com/gin-gonic/gin"
)

var tmplTemperatureAverageBySensor = template.Must(template.New("TemperatureAverageBySensor").Parse(db.TemperatureAverageBySensorSQLText))

// Sensor Statistics APIs
// @title getTemperatureAverageBySensor
// @Summary average temperature detected by a particular sensor between the specified date/time pairs (UNIX timestamps)
// @Tags Tested
// @Param	codeName	path	string	true	"code name(e.g:alpha 1)"
// @Param	from		query	int		false	"the specified date/time pairs of from(UNIX timestamps) such as 1699173029"
// @Param	till		query	int		false	"the specified date/time pairs untill(UNIX timestamps) such as 1699175089"
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /sensor/{codeName}/temperature/average [get]
func getTemperatureAverageBySensor(c *gin.Context) {
	from := c.DefaultQuery("from", default_From)
	till := c.DefaultQuery("till", strconv.FormatInt(time.Now().Unix(), 10))

	// Check the param in the request.
	codeName := c.Param("codeName")
	if len(codeName) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "codeName is invalid."})
		return
	}

	average := float32(0)
	err := db.GetSingleRow(tmplTemperatureAverageBySensor, &map[string]interface{}{"CODE_NAME": codeName, "FROM": from, "TILL": till}, &average)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not find the code name from database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"CodeName": codeName, "TemperatureAverage": average})
}
