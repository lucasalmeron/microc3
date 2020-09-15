package user

import (
	"context"
)

var repository Repository

func SetRepository(repo Repository) {
	repository = repo
}

func (user *User) Validate() error {
	return nil
}

func (user *User) GetUsers() ([]User, error) {
	return repository.GetUsers(context.TODO())
}

func (user *User) GetUser(userID string) (*User, error) {
	return repository.GetUser(context.TODO(), userID)
}

func (user *User) Save() (*User, error) {
	if user.ID == "" {
		return repository.Create(context.TODO(), *user)
	} else {
		return repository.Update(context.TODO(), *user)
	}
}

func (user *User) Delete(userID string) (*User, error) {
	return repository.Delete(context.TODO(), userID)
}
