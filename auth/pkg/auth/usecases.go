package auth

import (
	"context"
)

var repository Repository

func SetRepository(repo Repository) {
	repository = repo
}

//Validate permission
func (permission *Permission) Validate() error {

	return nil
}

//Validate Auth
func (auth *Auth) Validate() error {

	return nil
}

func (auth *Auth) GetByID(authID string) (*Auth, error) {
	return repository.GetByID(context.TODO(), authID)
}

func (auth *Auth) GetByUserID(userID string) (*Auth, error) {
	return repository.GetByUserID(context.TODO(), userID)
}

func (auth *Auth) Save() (*Auth, error) {
	if auth.ID == "" {
		return repository.Create(context.TODO(), *auth)
	} else {
		return repository.Update(context.TODO(), *auth)
	}
}

func (auth *Auth) Delete(authID string) (*Auth, error) {
	return repository.Delete(context.TODO(), authID)
}

func (auth *Auth) PushPermission(permission Permission) (*Auth, error) {
	return repository.PushPermission(context.TODO(), auth.User, permission)
}

func (auth *Auth) UpdatePermission(permission Permission) (*Auth, error) {
	return repository.UpdatePermission(context.TODO(), auth.User, permission)
}

func (auth *Auth) DeletePermission(permissionID string) (*Auth, error) {
	return repository.DeletePermission(context.TODO(), auth.User, permissionID)
}
