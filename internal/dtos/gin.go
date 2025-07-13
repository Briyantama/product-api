package dtos

type ResponseDTO struct {
	Success bool `json:"success"`
	Error   any  `json:"error,omitempty"`
	Data    any  `json:"data,omitempty"`
}

type ErrorResponseDTO struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

type TokenResponseDTO struct {
	Token   string `json:"token,omitempty"`
	Message string `json:"message,omitempty"`
}
