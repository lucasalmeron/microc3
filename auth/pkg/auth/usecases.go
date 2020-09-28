package auth

import (
	"context"
)

var repository Repository

func SetRepository(repo Repository) {
	repository = repo
}

func (auth *Auth) Validate() error {

	return nil
}

func (auth *Auth) GetList() ([]Auth, error) {
	return repository.GetList(context.TODO())
}

func (auth *Auth) GetbyID(authID string) (*Auth, error) {
	return repository.GetByID(context.TODO(), authID)
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
