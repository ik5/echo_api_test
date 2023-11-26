package apiv1

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/ik5/echo_api_test/models"
	"github.com/ik5/echo_api_test/structs"
	"github.com/labstack/echo/v4"
)

// GetUserByID finds a user by a given user id.
//
// @Summary Finds a user by a given user id, and returning the record
// @Description Finds a user by a given user id, and returning the record
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.Users "User details as returned by the saved record"
// @Failure 400 {object} structs.HTTPError "Something in the request is wrong/unexpected"
// @Failure 404 {object} structs.HTTPError "User was not found by provided ID"
// @Failure 500 {object} structs.HTTPError "Something in internal operation was bad"
// @Router /users/get/by_id/{id} [GET]
func (api *APIv1) GetUserByID(ctx echo.Context) error {
	_ = ctx.Request().Body.Close() // Always close body :/

	logger := api.ctx.App.Logger
	strID := ctx.Param("id")

	logger.Debug("get user by id", "id", strID)

	userID, err := strconv.Atoi(strID)
	if err != nil {
		logger.Error(
			"Unable to convert strID",
			"strID", strID,
			"err", err,
		)

		return echo.NewHTTPError(
			http.StatusBadRequest,
			structs.HTTPError{
				StatusCode: http.StatusBadRequest,
				Message:    "Invalid id",
				Info: map[string]any{
					"id": strID,
				},
			},
		)
	}

	slog.Debug("converted id number", "id", userID)

	user := &models.Users{}
	user.SetConfig(api.ctx.DB.Pool, api.ctx.App.Logger)

	err = user.FindByID(userID)
	if err != nil {
		logger.Error("Unable to find user", "err", err)

		return echo.NewHTTPError(
			http.StatusInternalServerError,
			structs.HTTPError{
				StatusCode: http.StatusInternalServerError,
				Message:    "Unable to look for user",
				Info: map[string]any{
					"id": strID,
				},
			},
		)
	}

	if user.ID == 0 {
		logger.Warn("user was not found", "id", userID)

		return echo.NewHTTPError(
			http.StatusNotFound,
			structs.HTTPError{
				StatusCode: http.StatusNotFound,
				Message:    "user was not found",
				Info: map[string]any{
					"id": strID,
				},
			},
		)
	}

	err = ctx.JSON(
		http.StatusOK, user,
	)
	if err != nil {
		logger.Error("Unable to send JSON", "err", err)

		return echo.NewHTTPError(
			http.StatusInternalServerError,
			structs.HTTPError{
				StatusCode: http.StatusInternalServerError,
				Message:    "User found, but cannot send output",
				Info: map[string]any{
					"id": strID,
				},
			},
		)
	}

	return nil
}
