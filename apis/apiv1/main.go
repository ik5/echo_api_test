package apiv1

import (
	"fmt"

	"github.com/ik5/echo_api_test/apis/apiv1/docs"
	"github.com/ik5/echo_api_test/structs"
	swagger "github.com/swaggo/echo-swagger"
)

var apiV1 *APIv1

// InitAPI generate an API v1 entry
//
// @title API to learn how to use swagger and echo
// @version 1.0
// @description The following project helps to learn echo and how to attach swagger to it.
// @license.name Mozilla Public License 2.0
// @license.url https://www.mozilla.org/en-US/MPL/2.0/
// @produce json
// @accept json
// @schemes http
// @BasePath /v1
func InitAPI(ctx *structs.Context) {
	apiV1 = &APIv1{
		ctx: ctx,
	}

	docs.SwaggerInfo.Host = fmt.Sprintf("http://%s", ctx.HTTPServer.ExternalHost)

	logger := ctx.App.Logger
	app := ctx.HTTPServer.App
	apiGroup := app.Group("/v1")
	apiUser := apiGroup.Group("/users")

	logger.Debug("registering group...")
	apiUser.PUT(
		"/add", apiV1.CreateUser,
	)
	apiUser.GET("/get/by_id/:id", apiV1.GetUserByID)

	apiGroup.GET("/docs/*", swagger.WrapHandler)

	logger.Debug("Finished registering group")
}
