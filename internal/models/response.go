package models

// APIResponse represents a standard API response structure
// This is used for Swagger documentation
type APIResponse struct {
	Message        string      `json:"message" example:"Operation successful"`
	Data           interface{} `json:"data,omitempty"`
	Count          int         `json:"count,omitempty" example:"10"`
	ResponseTimeMs int64       `json:"response_time_ms" example:"15"`
	Error          string      `json:"error,omitempty" example:"Error message"`
	Details        string      `json:"details,omitempty" example:"Detailed error information"`
}

// SuccessResponse represents a successful API response
type SuccessResponse struct {
	Message        string      `json:"message" example:"Operation successful"`
	Data           interface{} `json:"data,omitempty"`
	Count          int         `json:"count,omitempty" example:"10"`
	ResponseTimeMs int64       `json:"response_time_ms" example:"15"`
}

// ErrorResponse represents an error API response
type ErrorResponse struct {
	Error          string `json:"error" example:"Error message"`
	Details        string `json:"details,omitempty" example:"Detailed error information"`
	ResponseTimeMs int64  `json:"response_time_ms" example:"15"`
}
