package main

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/ik5/echo_api_test/apis/apiv1"
	"github.com/ik5/echo_api_test/structs"
)

var (
	logger      *slog.Logger
	quit        chan bool
	cliSettings settings
	ctx         structs.Context
)

func main() {
	quit = make(chan bool, 1)
	exeName := filepath.Base(os.Args[0])

	logger = initLogger()

	fmt.Printf(
		"[%s] %s - v%s built at: %s\n",
		exeName, gitBranch, gitVersion, buildTime,
	)

	cliSettings = initSettings()

	/*
		logger.Debug(
			"Initialized configuration",
			slog.String("app name", exeName),
			slog.String("branch", gitBranch),
			slog.String("version", gitVersion),
			slog.String("build time", buildTime),
		)
	*/

	// TODO: remove after development is done
	logger.Debug("starting new application", "settings", slog.AnyValue(cliSettings))

	ctx = structs.Context{
		App: structs.GeneralInfo{
			Logger: logger,
			Quit:   &quit,
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

	dbPool, dbConf, err := initPGPool(&cliSettings)
	if err != nil {
		panic(err)
	}

	webAPP := initHTTP(logger)

	ctx.DB.Config = dbConf
	ctx.DB.Pool = dbPool
	ctx.HTTPServer.App = webAPP

	go signalling()

	go ctx.InitServer(logger)

	apiv1.InitAPI(&ctx)

	<-quit

	err = ctx.Shutdown()
	if err != nil {
		logger.Warn("Unable to shutdown the HTTP server", "err", err)
	}

	dbPool.Close()

	logger.Info("bye")
}
