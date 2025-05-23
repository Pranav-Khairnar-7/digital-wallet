package models

import (
	customError "myapp/error"
	"myapp/util"
	"time"
)

type User struct {
	Base
	Name     string
	Age      int
	Email    string
	Password string
	Accounts []*Account
}

func NewUser(name string, password string, age int, email string) (*User, error) {

	var user = User{
		Base: Base{
			ID:        util.GetUniqueUserID(),
			CreatedAt: time.Now(),
		},
		Name:     name,
		Password: password,
		Age:      age,
		Email:    email,
	}

	err := user.Validate()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *User) Validate() error {

	if !util.IsValidEmailString(u.Email) {
		return customError.NewValidationError("email doesn't contain @")
	}

	if util.IsStringEmpty(u.Name) {
		return customError.NewValidationError("Name is empty.")
	}

	if u.Age < 0 || u.Age > 150 {
		return customError.NewValidationError("Age is less than 0 or greater than 150.")
	}
	return nil
}
