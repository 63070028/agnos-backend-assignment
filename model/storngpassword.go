package model


type StorngPasswordRequest struct {
	Password string `json:"init_password"`
}

type StorngPasswordResponse struct {
	Steps int `json:"num_of_steps"`
}

type ErrorResponse struct {
	TimeStamp string `json:"timestamp"`
	Status int `json:"status"`
	Error string `json:"error"`
	Path string `json:"path"`
}