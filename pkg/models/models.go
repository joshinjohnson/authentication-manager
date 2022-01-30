package models

type LoginSuccessResponse struct {
    Message string `json:"Message"`
    Token   string `json:"Token"`
}

type RegistrationSuccessResponse struct {
    Message string `json:"Message"`
}

type HomeSuccessResponse struct {
    Message string `json:"Message"`
}
