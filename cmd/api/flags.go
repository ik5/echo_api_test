package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultPort            uint   = 5000
	defaultPGPort          uint   = 5432
	defaultPGMaxConnection uint   = 5
	defaultHost            string = ""
	defaultPGHost          string = "localhost"
	defaultPGDB            string = "test_api"
	defaultPGUserName      string = "postgres"
	defaultPGPassword      string = ""
)

type settings struct {
	Port             uint
	PGPort           uint
	PGMaxConnections uint
	Host             string
	PGHost           string
	PGDB             string
	PGUserName       string
	PGPassword       string
}

func initSettings() settings {
	host := flag.String(
		"h", defaultHost, "host to listen on for HTTP server",
	)

	port := flag.Uint(
		"p", defaultPort, "The port number to listen to on the HTTP server",
	)

	pgHost := flag.String(
		"pg-host", defaultPGHost, "The address for postgres",
	)

	pgPort := flag.Uint(
		"pg-port", defaultPGPort, "The port for postgres",
	)

	pgMaxConnections := flag.Uint(
		"pg-conn", defaultPGMaxConnection, "Number of connections to have in connection pool",
	)

	pgDB := flag.String(
		"pg-db", defaultPGDB, "The name of the database",
	)

	pgUserName := flag.String(
		"pg-username", defaultPGUserName, "Postgres username",
	)

	pgPassword := flag.String(
		"pg-pass", defaultPGPassword, "Postgres password for provided username",
	)

	flag.Parse()

	return settings{
		Host:             *host,
		Port:             *port,
		PGHost:           *pgHost,
		PGPort:           *pgPort,
		PGMaxConnections: *pgMaxConnections,
		PGDB:             *pgDB,
		PGUserName:       *pgUserName,
		PGPassword:       *pgPassword,
	}
}

// HTTPListen generates a listen string for HTTP
func (s settings) HTTPListen() string {
	return net.JoinHostPort(s.Host, fmt.Sprintf("%d", s.Port))
}

// PGDSN generate dsn for configuration
func (s settings) PGDSN() string {
	const defaultMinDv uint = 2

	items := []string{}

	elements := map[string]string{
		"user":                          s.PGUserName,
		"password":                      s.PGPassword,
		"host":                          s.PGHost,
		"port":                          fmt.Sprintf("%d", s.PGPort),
		"dbname":                        s.PGDB,
		"pool_max_conns":                fmt.Sprintf("%d", s.PGMaxConnections),
		"pool_min_conns":                fmt.Sprintf("%d", s.PGMaxConnections/defaultMinDv),
		"pool_max_conn_lifetime":        "5m",
		"pool_max_conn_idle_time":       "2m30s",
		"pool_health_check_period":      "1m",
		"pool_max_conn_lifetime_jitter": "3m",
		"connect_timeout":               "30",
		"application_name":              filepath.Base(os.Args[0]),
		"keepalives":                    "1",
		"sslmode":                       "prefer",
		"sslcertmode":                   "allow",
	}

	for key, value := range elements {
		items = append(items, fmt.Sprintf("%s=%s", key, value))
	}

	return strings.Join(items, " ")
}
