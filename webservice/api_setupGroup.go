package webservice

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strings"

	"healthcare/db"
	"healthcare/logger"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const (
	default_SensorNumber   = 3
	default_OutputRate     = 30
	default_3d_max         = 100000
	default_OutputRate_max = 60
)

type group struct {
	GroupName    string `form:"groupName" json:"groupName" uri:"groupName"`
	SensorNumber int    `form:"number" json:"number"`
	OutputRate   int    `form:"rate" json:"rate"`
}

var (
	tmplCreateSensor = template.Must(template.New("CreateSensor").Parse(db.CreateSensorSQLText))
	tmplSensorIdMax  = template.Must(template.New("SensorIdMax").Parse(db.SensorIdMaxSQLText))
)

// setupGroup generate the sensors based on the given group details(group name, number of sensors and the output rate)
func setupGroup(c *gin.Context) {
	// Check the param in the request.
	groupName := strings.ToLower(c.Param("groupName"))
	if len(groupName) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "groupName is invalid, should be a greek letter."})
		return
	}

	// Setup the GroupName, SensorNumber, and OutputRate of Sensor
	var newGroup group
	if err := c.ShouldBindWith(&newGroup, binding.Form); err != nil {
		c.String(http.StatusNotFound, err.Error())
	}
	newGroup.GroupName = groupName

	// Generate the related sensors for a specified group
	if err := generateSensorForGroup(newGroup); err != nil {
		logger.Errorf(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"GroupName": newGroup, "message": "Setup Group is completed."})
}

// generateSensorForGroup - Generate sensors for a group based on the provided details.
func generateSensorForGroup(newGroup group) error {
	if newGroup.SensorNumber <= 0 {
		newGroup.SensorNumber = default_SensorNumber
	}
	if newGroup.OutputRate <= 0 {
		newGroup.OutputRate = default_OutputRate
	}

	// Generate sensors in the group based on the provided details.
	for i := 0; i < newGroup.SensorNumber; i++ {
		if err := createSensor(newGroup); err != nil {
			logger.Errorf(err.Error())
			return err
		}
	}

	logger.Infof("Setup Group is completed.")
	return nil
}

// createSensor - Create a sensor for the specified group.
func createSensor(g group) error {
	maxIdx := int(0)
	err := db.GetSingleRow(tmplSensorIdMax, &map[string]interface{}{"GROUP_NAME": g.GroupName}, &maxIdx)
	if err != nil {
		maxIdx = 0
	}

	maxIdx++
	_, err = db.DoInsert(tmplCreateSensor, &map[string]interface{}{
		"GROUP_NAME": g.GroupName,
		"CODE_NAME":  fmt.Sprintf("%s %d", g.GroupName, maxIdx),
		"IDX":        maxIdx,
		"X_3D":       rand.Int63() % default_3d_max,
		"Y_3D":       rand.Int63() % default_3d_max,
		"Z_3D":       rand.Int63() % default_3d_max,
		"RATE":       g.OutputRate,
	})
	if err != nil {
		return err
	}
	return nil
}
