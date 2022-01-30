package models

type LoginSuccessResponse struct {
    Message string `json:"message"`
    Token   string `json:"token"`
}

type RegistrationSuccessResponse struct {
    Message string `json:"message"`
}
