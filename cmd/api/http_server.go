package main

import (
	"context"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	compressionLevel int    = 5
	maxBodySize      string = "2M"
)

func initHTTP(logger *slog.Logger) *echo.Echo {
	app := echo.New()

	app.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: compressionLevel,
	}))

	app.Use(middleware.BodyLimit(maxBodySize))

	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:        true,
		LogURI:           true,
		LogError:         true,
		LogHost:          true,
		LogMethod:        true,
		LogLatency:       true,
		LogURIPath:       true,
		LogReferer:       true,
		LogProtocol:      true,
		LogRemoteIP:      true,
		LogRoutePath:     true,
		LogRequestID:     true,
		LogUserAgent:     true,
		LogResponseSize:  true,
		LogContentLength: true,
		HandleError:      true,
		LogValuesFunc: func(c echo.Context, loggerValues middleware.RequestLoggerValues) error {
			if loggerValues.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("RemoteIP", loggerValues.RemoteIP),
					slog.String("protocol", loggerValues.Protocol),
					slog.String("method", loggerValues.Method),
					slog.String("uri", loggerValues.URI),
					slog.String("path", loggerValues.URIPath),
					slog.Int("status", loggerValues.Status),
					slog.String("content-length", loggerValues.ContentLength),
				)

				return nil
			}
			logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
				slog.String("RemoteIP", loggerValues.RemoteIP),
				slog.String("protocol", loggerValues.Protocol),
				slog.String("method", loggerValues.Method),
				slog.String("uri", loggerValues.URI),
				slog.String("path", loggerValues.URIPath),
				slog.Int("status", loggerValues.Status),
				slog.String("content-length", loggerValues.ContentLength),
				slog.String("err", loggerValues.Error.Error()),
			)

			return nil
		},
	}))

	return app
}
