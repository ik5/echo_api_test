package apiv1

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/ik5/echo_api_test/models"
	"github.com/ik5/echo_api_test/structs"
	"github.com/labstack/echo/v4"
)

// CreateUser is the Exported API request to create a user
func (api *APIv1) CreateUser(ctx echo.Context) error {
	logger := api.ctx.App.Logger
	logger.Debug("In CreateUser")

	request := ctx.Request()

	if request.ContentLength <= 0 {
		return echo.NewHTTPError(http.StatusBadRequest, ErrUnableToReadContentLength)
	}

	body, err := io.ReadAll(request.Body)
	_ = request.Body.Close()

	if err != nil || len(body) == 0 {
		logger.Error("Unable to read body", "err", err)

		return echo.NewHTTPError(http.StatusBadRequest, ErrUnableToReadContentLength)
	}

	logger.Debug("have payload", "body", slog.StringValue(string(body)))

	user, err := api.handleNewUserPayload(body)
	if err != nil {
		return err
	}

	if user != nil {
		err = ctx.JSON(
			http.StatusOK, *user,
		)

		if err != nil {
			logger.Error("Unable to send JSON", "err", err)

			return echo.NewHTTPError(
				http.StatusInternalServerError,
				"Record created, but unable to send output",
			)
		}
	}

	return nil
}

func (api *APIv1) handleNewUserPayload(payload []byte) (*models.Users, error) {
	logger := api.ctx.App.Logger

	userData, err := structs.SetUser(payload)
	if err != nil {
		logger.Error("error parsing body", "err", err)

		return nil, echo.NewHTTPError(
			http.StatusBadRequest,
			"Unable to parse or validate the content",
		)
	}

	logger.Debug("have user", "user", userData)

	user := &models.Users{}
	user.SetConfig(api.ctx.DB.Pool, api.ctx.App.Logger)

	err = user.Insert(userData)
	if err != nil {
		logger.Error("Unable to insert user", "err", err)

		unwrappedErr := errors.Unwrap(err)

		logger.Debug("checking for specific errors", "unwrappedErr", unwrappedErr)

		if strings.HasPrefix(
			unwrappedErr.Error(),
			"ERROR: duplicate key value violates unique constraint",
		) {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "username already exists")
		}

		return nil, echo.NewHTTPError(
			http.StatusInternalServerError, "Unable to save user",
		)
	}

	logger.Debug("have user model", "user", *user)

	return user, nil
}
