package models

import (
	"context"
	"net/http"
)

type Scenario string

const (
	ScenarioPhone    Scenario = "phone"
	ScenarioEmail    Scenario = "email"
	ScenarioUsername Scenario = "form"
)

type Status string

const (
	Ok    Status = "OK"
	Error Status = "ERROR"
)

const ()

type RequestLogin struct {
	Scenario Scenario `json:"scenario"`
	Username string   `json:"username,omitempty"`
	Password string   `json:"password"`
	Phone    string   `json:"phone,omitempty"`
	Email    string   `json:"email,omitempty"`
}

type Response struct {
	Status Status
}

type ResponseInvalidLogin struct {
	Response
	ErrorCode    uint32 `json:"code"`
	ErrorMessage string `json:"message,omitempty"`
}

type Identifier interface {
	IdentityByCookie(c *http.Cookie) (string, error)
	Login(ctx context.Context, r *RequestLogin)
}
