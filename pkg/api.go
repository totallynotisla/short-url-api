package pkg

type ApiResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
