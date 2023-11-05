package webservice

import (
	"html/template"
	"math/rand"

	"healthcare/db"
	"healthcare/logger"
)

// Default constants for generating data
const (
	default_count_max = 1000
	default_temp_min  = -30
	default_temp_max  = 30
	default_tran_max  = 100
)

// fishSpecies - List of fish species
var fishSpecies = []string{
	"Atlantic Bluefin Tuna", "Atlantic Cod", "Atlantic Goliath Grouper",
	"Atlantic Salmon", "Beluga Sturgeon", "Blue Marlin",
	"Blue Tang", "Bluebanded Goby", "Bluehead Wrasse",
}

// SQL templates
var (
	tmplActivatedSensors         = template.Must(template.New("ActivatedSensors").Parse(db.ActivatedSensorsSQLText))
	tmplCreateFishSpecieData     = template.Must(template.New("CreateFishSpecieData").Parse(db.CreateFishSpecieDataSQLText))
	tmplNearbySensorTransparency = template.Must(template.New("NearbySensorTransparency").Parse(db.NearbySensorTransparencySQLText))
)

// generateSensorData - Generates sensor data(randomized) for a sensor with the given sensor ID, depth, and nearby sensor's transparency.
func generateSensorData(sensorId string, deep int64, tran int) error {
	spec := rand.Int() % len(fishSpecies)
	count := rand.Int() % default_count_max
	p := float64(deep) / float64(default_3d_max)
	temp := default_temp_max - ((default_temp_max - default_temp_min) * (p))

	affected, err := db.DoInsert(tmplCreateFishSpecieData, &map[string]interface{}{
		"NAME":  fishSpecies[spec],
		"COUNT": count,
		"TEMP":  float32(temp),
		"TRAN":  tran,
		"ID":    sensorId,
	})
	if err != nil {
		return err
	}
	logger.Infof("DoInsert affected rows: %d", affected)
	return nil
}

// getNearbySensorTransparency - Retrieves transparency data for a nearby sensor or generates random data if unavailable
func getNearbySensorTransparency(sensorId string, x, y, z int64) (tran int) {
	// Get transparency from the database
	err := db.GetSingleRow(tmplNearbySensorTransparency, &map[string]interface{}{
		"ID":   sensorId,
		"X_3D": x,
		"Y_3D": y,
		"Z_3D": z,
	}, &tran)
	if err != nil {
		return rand.Int() % default_tran_max
	}

	// Adjust transparency with a random change
	minChange := rand.Int() % 3
	tran += minChange
	if tran >= default_tran_max {
		tran -= minChange
	}

	return
}
