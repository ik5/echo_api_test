package apiv1

import (
	"errors"

	"github.com/ik5/echo_api_test/structs"
)

var (
	ErrUnableToReadContentLength = errors.New("Unable to read content length")
	ErrUnableToReadPayload       = errors.New("Unable to read payload content")
)

// APIv1 supporting internal details
type APIv1 struct {
	ctx *structs.Context
}
