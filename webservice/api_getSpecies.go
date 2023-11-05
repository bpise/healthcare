package webservice

import (
	"html/template"
	"net/http"
	"strconv"
	"time"

	"healthcare/db"

	"github.com/gin-gonic/gin"
)

var (
	tmplFishSpecies     = template.Must(template.New("FishSpecies").Parse(db.FishSpeciesSQLText))
	tmplFishSpeciesTopN = template.Must(template.New("FishSpeciesTopN").Parse(db.FishSpeciesTopNSQLText))
)

const default_From = "1699056000"

// getSpecies - full list of species (with counts) currently detected inside the group
func getSpecies(c *gin.Context) {
	// Check the param in the request.
	groupName := c.Param("groupName")
	if len(groupName) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "groupName is invalid."})
		return
	}

	result, err := db.GetRows(tmplFishSpecies, &map[string]interface{}{"GROUP_NAME": groupName})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not find the group name from database"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// getSpeciesTopN - list of top N species (with counts) currently detected inside the group
func getSpeciesTopN(c *gin.Context) {
	// Check the param in the request.
	groupName := c.Param("groupName")
	topN := c.Param("n")
	if len(groupName) == 0 || len(topN) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "groupName/topN is invalid."})
		return
	}
	strconv.FormatInt(time.Now().Unix(), 10)

	//
	from := c.Query("from")
	till := c.Query("till")
	validFromTill := false
	if len(from) > 0 || len(till) > 0 {
		validFromTill = true
		if len(from) == 0 {
			from = default_From
		}
		if len(till) == 0 {
			till = strconv.FormatInt(time.Now().Unix(), 10)
		}
	}

	result, err := db.GetRows(tmplFishSpeciesTopN, &map[string]interface{}{"GROUP_NAME": groupName, "TOP_N": topN, "isValidFromTill": validFromTill, "FROM": from, "TILL": till})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not find the group name from database"})
		return
	}

	c.JSON(http.StatusOK, result)
}
