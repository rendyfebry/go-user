package rest

// ErrorResponse ...
type ErrorResponse struct {
	Message  string `json:"message"`
	HTTPCode int    `json:"-"`
}
