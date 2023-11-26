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
//
// @Summary Create a new user.
// @Description Create a new user and returns it's record on success.
// @Tags Users
// @Accept json
// @Produce json
// @Param user_details body structs.User true "User object"
// @Success 200 {object} models.Users "User details as returned by the saved record"
// @Failure 400 {object} structs.HTTPError "Something in the request is wrong/unexpected"
// @Failure 500 {object} structs.HTTPError "Something in internal operation was bad"
// @Router /users/add [put]
func (api *APIv1) CreateUser(ctx echo.Context) error {
	request := ctx.Request()

	defer request.Body.Close() // close on exiting the function...

	logger := api.ctx.App.Logger
	logger.Debug("In CreateUser")

	if request.ContentLength <= 0 {
		logger.Error("ContentLength is <= 0 ", "ContentLength", request.ContentLength)

		return echo.NewHTTPError(
			http.StatusBadRequest,
			structs.HTTPError{
				StatusCode: http.StatusBadRequest,
				Message:    ErrUnableToReadContentLength.Error(),
				Info: map[string]any{
					"ContentLength": request.ContentLength,
				},
			},
		)
	}

	body, err := io.ReadAll(request.Body)

	if err != nil || len(body) == 0 {
		logger.Error("Unable to read body", "err", err)

		return echo.NewHTTPError(
			http.StatusBadRequest,
			structs.HTTPError{
				StatusCode: http.StatusBadRequest,
				Message:    ErrUnableToReadContentLength.Error(),
			},
		)
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
				structs.HTTPError{
					StatusCode: http.StatusInternalServerError,
					Message:    "Record created, but unable to send output",
					Info: map[string]any{
						"answer": user,
					},
				},
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
			structs.HTTPError{
				StatusCode: http.StatusBadRequest,
				Message:    "Unable to parse or validate the content",
				Info: map[string]any{
					"payload": string(payload),
				},
			},
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
			return nil, echo.NewHTTPError(
				http.StatusBadRequest,
				structs.HTTPError{
					StatusCode: http.StatusBadRequest,
					Message:    "username already exists",
					Info: map[string]any{
						"payload": userData,
					},
				},
			)
		}

		return nil, echo.NewHTTPError(
			http.StatusInternalServerError,
			structs.HTTPError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Unable to save user",
				Info: map[string]any{
					"payload": userData,
				},
			},
		)
	}

	logger.Debug("have user model", "user", *user)

	return user, nil
}
