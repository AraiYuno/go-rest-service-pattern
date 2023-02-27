package services

import (
	"errors"
)

var ErrFailedValidation = errors.New("failed validation")

type Service struct {
	Driver string
}

type RegisterUserInput struct {
	Username         string            `json:"username"`
	Password         string            `json:"password"`
	ValidationErrors map[string]string `json:"-"`
}

func (s *Service) RegisterUser(input *RegisterUserInput) (map[string]string, error) {
	input.ValidationErrors = make(map[string]string)

	if input.Username == "" {
		input.ValidationErrors["username"] = "must be provided"
	}

	// And any other validation checks...

	if len(input.ValidationErrors) > 0 {
		return nil, ErrFailedValidation
	}

	// create map of string key and value
	user := make(map[string]string)
	user["username"] = input.Username
	user["password"] = input.Password
	user["email"] = "kyle.ahn@gantry.io"
	return user, nil
}
