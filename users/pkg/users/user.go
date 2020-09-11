package user

import (
	"context"
	"time"
)

type User struct {
	ID             string `json:"id" bson:"_id,omitempty"`
	FirstName      string `json:"firstName" bson:"firstName"`
	LastName       string `json:"lastName" bson:"lastName"`
	DocumentNumber string `json:"documentNumber" bson:"documentNumber"`
	Password       string `json:"password" bson:"password"`
	Email          string `json:"email" bson:"email"`
	PhoneNumber    string `json:"phoneNumber" bson:"phoneNumber"`
	GDEUser        string `json:"GDEUser" bson:"GDEUser"`
	Position       string `json:"position" bson:"position"`

	Admin       bool `json:"admin" bson:"admin"`
	Permissions []struct {
		Read        bool `json:"read" bson:"read"`
		Write       bool `json:"write" bson:"write"`
		Responsible bool `json:"responsible" bson:"responsible"`
		Query       bool `json:"query" bson:"query"`
		Health      bool `json:"health" bson:"health"`
		QueryPoint  bool `json:"queryPoint" bson:"queryPoint"`
	} `json:"permissions" bson:"permissions"`

	ModifierUser string    `json:"modifierUser" bson:"modifierUser"`
	CreatedAt    time.Time `json:"createdAt" bson:"createdAt"`
	ModifiedAt   time.Time `json:"modifiedAt" bson:"modifiedAt"`
	DeletedAt    time.Time `json:"deletedAt" bson:"deletedAt"`
}

type Repository interface {
	GetUsers(ctx context.Context) ([]User, error)
	GetUser(ctx context.Context, userID string) (*User, error)
	CreateUser(ctx context.Context, user User) (*User, error)
	UpdateUser(ctx context.Context, user User) (*User, error)
	DeleteUser(ctx context.Context, userID string) error
}
