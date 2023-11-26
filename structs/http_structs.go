package structs

// HTTPError holds information regarding HTTP request
type HTTPError struct {
	// HTTP Status Code
	StatusCode int `example:"400" json:"status_code" validate:"required,min=400,max=599"`

	// Reason for the error
	Message string `example:"Something went wrong" json:"message" validate:"required"`

	// Additional information to provide (if existd)
	Info map[string]any `json:"info,omitempty" validate:"omitempty,omitnil"`
}
