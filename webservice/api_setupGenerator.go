package webservice

import (
	"fmt"
	"math/rand"
	"time"

	"healthcare/cronjob"
	"healthcare/db"
	"healthcare/logger"
)

var groups = map[string]interface{}{"alpha": true, "beta": true, "gamma": true, "delta": true, "epsilon": true, "zeta": true, "eta": true, "theta": true, "iota": true, "kappa": true, "lambda": true, "mu": true, "nu": true, "xi": true, "omicron": true, "pi": true, "rho": true, "sigma": true, "tau": true, "upsilon": true, "phi": true, "chi": true, "psi": true, "omega": true}

// startSetupSensors for generate the sensors and fake data
func StartSetupSensors() {
	logger.Infof("Setup Sensors is starting......")
	time.Sleep(time.Second * 10)

	// generate the sensor groups and the sensors in each group.
	if err := autoGenerateSensorForGroup(); err != nil {
		logger.Errorf("Auto Setup Group error:%v", err)
	}

	logger.Infof("Auto Setup Group is completed.")
	time.Sleep(time.Second * 10)

	// generate randomized fake data
	if err := startToGenerateSensorData(); err != nil {
		logger.Errorf("StartToGenerateSensorData error:%v", err)
	}

	logger.Infof("Setup the randomized fake data generators is completed, and generators will be started very soon...")
}

// autoGenerateSensorForGroup
func autoGenerateSensorForGroup() error {
	for groupName := range groups {
		if err := generateSensorForGroup(group{
			GroupName:  groupName,
			OutputRate: default_OutputRate + (rand.Int() % default_OutputRate_max),
		}); err != nil {
			logger.Errorf(err.Error())
		}
	}
	return nil
}

// startToGenerateSensorData
func startToGenerateSensorData() error {
	result, err := db.GetRows(tmplActivatedSensors, nil)
	if err != nil {
		return err
	}
	logger.Infof("Activated Sensors:%d", len(*result))

	for _, sensor := range *result {
		logger.Infof("sensor info:%v", sensor)

		uuid, ok1 := sensor["id"].(string)
		x3d, ok2 := sensor["x_3d"].(int64)
		y3d, ok3 := sensor["y_3d"].(int64)
		z3d, ok4 := sensor["z_3d"].(int64)
		rate, ok5 := sensor["output_rate_sec"].(int32)

		if ok1 && ok2 && ok3 && ok4 && ok5 {
			jobSpec := fmt.Sprintf("*/%d * * * * *", rate)
			logger.Debugf("jobSpec:%s", jobSpec)

			// AddFunc in the globleCronJob
			_, err := cronjob.AddFunc(jobSpec, "", func() {
				tran := getNearbySensorTransparency(uuid, x3d, y3d, z3d)
				_ = generateSensorData(uuid, z3d, tran)
			})
			if err != nil {
				logger.Errorf(err.Error())
			}
		}
	}

	return nil
}
