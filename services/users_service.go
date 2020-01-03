package services

import (
	"github.com/JErBerlin/bookstore-users-api/domain/users"
	"github.com/JErBerlin/bookstore-users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}

func GetUser() {}
