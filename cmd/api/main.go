package main

import (
	"log/slog"

	"github.com/ik5/echo_api_test/structs"
)

var (
	quit        = make(chan bool, 1)
	logger      = initLogger()
	cliSettings settings
	ctx         structs.Context
)

func main() {
	cliSettings = initSettings()
	// TODO: remove after development is done
	logger.Debug("starting new application", "settings", slog.AnyValue(cliSettings))

	ctx = structs.Context{
		App: structs.GeneralInfo{
			Logger: logger,
			Quit:   quit,
		},
		HTTPServer: structs.HTTPServerInfo{
			Port:   cliSettings.Port,
			Host:   cliSettings.Host,
			Listen: cliSettings.HTTPListen(),
			// App:    &echo.Echo{},
		},
		DB: structs.DBInfo{
			Port:           cliSettings.PGPort,
			MaxConnections: cliSettings.PGMaxConnections,
			Host:           cliSettings.PGHost,
			DBName:         cliSettings.PGDB,
			Username:       cliSettings.PGUserName,
			Password:       cliSettings.PGPassword,
		},
	}

	dbPool, dbConf, err := InitPGPool(&cliSettings)
	if err != nil {
		panic(err)
	}

	ctx.DB.Config = dbConf
	ctx.DB.Pool = dbPool

	go signalling()
	<-quit
	dbPool.Close()
	logger.Info("bye")
}
