package server

type Response struct {
	Status  string `json:"status"`            // success | error | failed
	Message string `json:"message,omitempty"` // message
	Error   string `json:"error,omitempty"`   // error
	Data    any    `json:"data,omitempty"`    // data
}
