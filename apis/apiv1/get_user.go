package apiv1

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/ik5/echo_api_test/models"
	"github.com/labstack/echo/v4"
)

func (api *APIv1) GetUserByID(ctx echo.Context) error {
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

		return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
	}

	slog.Debug("converted id number", "id", userID)

	user := &models.Users{}
	user.SetConfig(api.ctx.DB.Pool, api.ctx.App.Logger)

	err = user.FindByID(userID)
	if err != nil {
		logger.Error("Unable to find user", "err", err)

		return echo.NewHTTPError(
			http.StatusInternalServerError, "Unable to look for user",
		)
	}

	if user.ID == 0 {
		logger.Warn("user was not found", "id", userID)

		return echo.NewHTTPError(http.StatusNotFound, "user was not found")
	}

	err = ctx.JSON(
		http.StatusOK, user,
	)
	if err != nil {
		logger.Error("Unable to send JSON", "err", err)

		return echo.NewHTTPError(
			http.StatusInternalServerError,
			"User found, but cannot send output",
		)
	}

	return nil
}
