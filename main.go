package main

import (
	"context"
	"errors"
	"net/http"

	"healthcare/cache"
	"healthcare/db"
	"healthcare/logger"
	"healthcare/webservice"
)

func main() {
	// Init Logger
	logger.InitLogger()
	defer logger.Sync()

	// Init DB Connection Pool
	cache.InitRedis(context.Background())
	defer cache.Close()

	// Init DB Connection Pool
	db.InitDB(context.Background())
	defer db.Close()

	// Init Web Service Engine and the Routers
	webservice.InitWebEngine()
	webservice.InitRouter()

	// Start HTTP Web Service
	if err := webservice.Run(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Errorf(err.Error())
	}
}
