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

var tmplTransparencyAverage = template.Must(template.New("TransparencyAverage").Parse(db.TransparencyAverageSQLText))

// getTransparencyAverage - Retrieves the current average transparency inside the group.
func getTransparencyAverage(c *gin.Context) {
	// Check the param in the request.
	groupName := c.Param("groupName")
	if len(groupName) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "groupName is invalid."})
		return
	}

	// Get the cached TransparencyAverage for a group
	// if successful, return the result directly
	cacheKey := groupName + "TranAvg"
	result, err := cache.Get(c, cacheKey).Result()
	if err == nil {
		// If the result was found in the cache, parse it and return it directly
		cachedAverage, err := strconv.ParseFloat(result, 32)
		if err == nil {
			logger.Debugf("cachedAverage:%f", cachedAverage)
			c.JSON(http.StatusOK, gin.H{"GroupName": groupName, "TransparencyAverage": float32(cachedAverage)})
			return
		}
	}

	// Calculate TransparencyAverage in database by a given group name
	average := float32(0)
	err = db.GetSingleRow(tmplTransparencyAverage, &map[string]interface{}{"GROUP_NAME": groupName}, &average)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not find the group name from database"})
		return
	}

	// Store the TransparencyAverage into Redis for caching
	_, err = cache.Set(c, cacheKey, average, time.Second*10).Result()
	if err != nil {
		logger.Errorf(err.Error())
	}

	// Return the calculated TransparencyAverage
	c.JSON(http.StatusOK, gin.H{"GroupName": groupName, "TransparencyAverage": average})
}
