package apiv1

import "github.com/ik5/echo_api_test/structs"

var apiV1 *APIv1

// InitAPI generate an API v1 entry
func InitAPI(ctx *structs.Context) {
	apiV1 = &APIv1{
		ctx: ctx,
	}

	logger := ctx.App.Logger
	app := ctx.HTTPServer.App
	apiGroup := app.Group("/v1")
	apiUser := apiGroup.Group("/users")

	logger.Debug("registering group...")
	apiUser.PUT(
		"/add", apiV1.CreateUser,
	)
	apiUser.GET("/get/by_id/:id", apiV1.GetUserByID)

	logger.Debug("Finished registering group")
}
