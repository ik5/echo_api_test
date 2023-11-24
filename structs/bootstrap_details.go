package structs

import (
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

// HTTPServerInfo holds information about the HTTP server
type HTTPServerInfo struct {
	// Raw port
	Port uint

	// Raw host
	Host string

	// Raw listen based on both Host and Port
	Listen string

	// web application
	App *echo.Echo
}

// GeneralInfo of the application
type GeneralInfo struct {
	// Logger access
	Logger *slog.Logger

	// Quit channel
	Quit chan bool
}

type DBInfo struct {
	// Port number to be used
	Port uint

	// How many connections to use for connection pool
	MaxConnections uint

	// Host address
	Host string

	// Database name
	DBName string

	// Username to be connect with
	Username string

	// The password for the given username
	Password string

	// Configuration of the pool
	Config *pgxpool.Config

	// Pointer to the pool connection
	Pool *pgxpool.Pool
}

// Context holds information from bootstrap to be used
type Context struct {
	// App is the basic bootstrap information
	App GeneralInfo

	// Information about HTTP Server
	HTTPServer HTTPServerInfo

	// Database information
	DB DBInfo
}
