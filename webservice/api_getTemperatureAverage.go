package webservice

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"healthcare/cache"
	"healthcare/db"
	"healthcare/logger"

	"github.com/gin-gonic/gin"
)

var tmplTemperatureAverage = template.Must(template.New("TemperatureAverage").Parse(db.TemperatureAverageSQLText))

// getTemperatureAverage - get the current average temperature inside the group
func getTemperatureAverage(c *gin.Context) {
	// Check the param in the request.
	groupName := c.Param("groupName")
	if len(groupName) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "groupName is invalid."})
		return
	}

	// Get the cached TemperatureAverage for a group
	// if succeed, return the result directly
	cacheKey := groupName + "TempAvg"
	result, err := cache.Get(c, cacheKey).Result()
	if err == nil {
		cachedAverage, err := strconv.ParseFloat(result, 32)
		if err == nil {
			logger.Debugf("cachedAverage:%f", cachedAverage)
			c.JSON(http.StatusOK, gin.H{"GroupName": groupName, "TemperatureAverage": float32(cachedAverage)})
			return
		}
	}

	// Calculate TemperatureAverage in database by a given group name
	average := float32(0)
	err = db.GetSingleRow(tmplTemperatureAverage, &map[string]interface{}{"GROUP_NAME": groupName}, &average)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not find the group name from database"})
		return
	}

	// Store the TemperatureAverage into redis for caching
	_, err = cache.Set(c, cacheKey, average, time.Second*10).Result()
	if err != nil {
		logger.Errorf(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"GroupName": groupName, "TemperatureAverage": average})
}
