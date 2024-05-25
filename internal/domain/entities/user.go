package entities

import (
	"errors"
	"net/mail"
)

type UserRole string
type UserError string

const (
	UserRoleManager    UserRole = "manager"
	UserRoleTechnician UserRole = "technician"

	invalidEmailError UserError = "invalid email address"
	invalidRoleError  UserError = "invalid role"
)

type User struct {
	ID    string
	Name  string
	Email string
	Role  UserRole
}

func NewUser(name string, address string, role string) (User, error) {
	email, err := validMailAddress(address)
	if err != nil {
		return User{}, errors.New(string(invalidEmailError))
	}

	if role != string(UserRoleManager) && role != string(UserRoleTechnician) {
		return User{}, errors.New(string(invalidRoleError))
	}

	user := User{
		Name:  name,
		Email: email,
		Role:  UserRole(role),
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
