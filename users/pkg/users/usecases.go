package user

import "context"

var repository Repository

func SetRepository(repo Repository) {
	repository = repo
}

func (user *User) Validate() error {
	return nil
}

func (user *User) GetUsers() ([]User, error) {
	return repository.GetUsers(context.Background())
}

func (user *User) GetUser(userID string) (*User, error) {
	return repository.GetUser(context.Background(), userID)
}

func (user *User) CreateUser(newUser User) (*User, error) {
	return repository.CreateUser(context.Background(), newUser)
}

func (user *User) UpdateUser(reqUser User) (*User, error) {
	return repository.UpdateUser(context.Background(), reqUser)
}

func (user *User) DeleteUser(userID string) error {
	return repository.DeleteUser(context.Background(), userID)
}
