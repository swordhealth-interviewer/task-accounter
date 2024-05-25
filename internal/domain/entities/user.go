package entities

import (
	"errors"
	"net/mail"
)

type Type string

const (
	Manager    Type = "manager"
	Technician Type = "technician"
)

type User struct {
	ID    string
	Name  string
	Email string
	Role  Type
}

func NewUser(name string, address string, role string) (User, error) {
	email, err := validMailAddress(address)
	if err != nil {
		return User{}, errors.New("invalid email address")
	}

	if role != "manager" && role != "technician" {
		return User{}, errors.New("invalid role")
	}

	user := User{
		Name:  name,
		Email: email,
		Role:  Type(role),
	}

	return user, nil
}

func validMailAddress(address string) (string, error) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", err
	}

	return addr.Address, nil
}
