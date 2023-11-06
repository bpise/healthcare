package main

import (
	"context"
	"errors"
	"net/http"

	"healthcare/cache"
	"healthcare/cronjob"
	"healthcare/db"
	_ "healthcare/docs"
	"healthcare/logger"
	"healthcare/webservice"
)

// @title Underwater Sensors Data Generation and Statistics API
// @version 1.0
// @description ### This program contains 3 main parts: sensor group(s) distribution & control, sensor data generation & transfer, and the sensor statistics APIs.
// @description #### 1. Sensor group kickoff
// @description During the kickoff, the service will generate sensors for each group by greek letter names.
// @description #### 2. Sensor data generation and transfer
// @description In this part, the sensor services will generate the sensor data(fake) by cron jobs. The randomized data will then store the data into the database for testing purposes only. and there may would be a exposed API for such data transfer as well.
// @description #### 3. Sensor statistics APIs
// @description The sensor statistics apis will be exposed to the public. The apis will be used to query the sensor data in the database and return the results.
// @host 127.0.0.1
func main() {
	// Init Logger
	logger.InitLogger()
	defer logger.Sync()

	// Init Redis Client
	cache.InitRedis(context.Background())
	defer cache.Close()

	// Init DB Connection Pool
	db.InitDB(context.Background())
	defer db.Close()

	// Init Web Service Engine and the Routers for the APIs
	webservice.InitWebEngine()
	webservice.InitRouter()

	// Init Cronjob and start it
	cronjob.InitCronJob()
	cronjob.Start()
	defer cronjob.Stop()

	// Start to generate the sensors with simulated data
	go webservice.StartSetupSensors()

	// Start HTTP Web Service
	if err := webservice.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Errorf(err.Error())
	}
}
