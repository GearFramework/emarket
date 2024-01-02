package models

import (
	"context"
	"github.com/GearFramework/emarket/internal/entities"
	"net/http"
)

type Scenario string

const (
	ScenarioPhone    Scenario = "phone"
	ScenarioEmail    Scenario = "email"
	ScenarioUsername Scenario = "form"
	ScenarioVkID     Scenario = "vk"
)

type Status string

const (
	Ok    Status = "OK"
	Error Status = "ERROR"
)

type RequestLogin struct {
	Scenario Scenario `json:"scenario"`
	Phone    string   `json:"phone,omitempty"`
	Email    string   `json:"email,omitempty"`
	Password string   `json:"password"`
}

type RequestRegister struct {
	Scenario Scenario `json:"scenario"`
	Phone    string   `json:"phone,omitempty"`
	Email    string   `json:"email,omitempty"`
	Password string   `json:"password"`
}

type Response struct {
	Status Status
}

type ResponseLogin struct {
	Response
	Token string
}

type ResponseInvalidLogin struct {
	Response
	ErrorCode    uint32 `json:"code"`
	ErrorMessage string `json:"message,omitempty"`
}

type ResponseInvalidRegister struct {
	Response
	ErrorCode    uint32 `json:"code"`
	ErrorMessage string `json:"message,omitempty"`
}

type Identifier interface {
	IdentityByCookie(c *http.Cookie) (string, error)
	Login(ctx context.Context, r *RequestLogin)
	Register(ctx context.Context, r *RequestRegister) (*entities.Customer, error)
}
