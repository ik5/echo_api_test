package structs

import (
	"encoding/json"
	"fmt"

	"github.com/ik5/echo_api_test/utils/runtimeutils"
)

type User struct {
	// User's first name
	FirstName string `json:"first_name" validate:"required,alphaunicode,max=15,min=2"`

	// User's middle name (if exists)
	MiddleName string `json:"middle_name,omitempty" validate:"min=0,max=15,omitempty,omitnil"`

	// User's last name
	LastName string `json:"last_name" validate:"required,alphaunicode,max=20,min=2"`

	// User's username that is in use
	Username string `json:"username" validate:"required,alphanumunicode|containsrune=.-_0x20100x20110x2015,max=24,min=2"`
}

// SetUser parse JSON request and validate the content
func SetUser(buf []byte) (*User, error) {
	funcName := runtimeutils.GetCallerFunctionName()

	var user User

	err := json.Unmarshal(buf, &user)
	if err != nil {
		return nil, fmt.Errorf("[%s] %w", funcName, err)
	}

	err = validate.Struct(user)
	if err != nil {
		err = fmt.Errorf("[%s] %w", funcName, err)
	}

	return &user, err
}
