package auth

import (
	"context"
)

type Auth struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	User string `json:"user" bson:"user"`

	Admin       bool `json:"admin" bson:"admin"`
	Permissions []struct {
		Read        bool   `json:"read" bson:"read"`
		Write       bool   `json:"write" bson:"write"`
		Responsible bool   `json:"responsible" bson:"responsible"`
		Query       bool   `json:"query" bson:"query"`
		Health      bool   `json:"health" bson:"health"`
		QueryPoint  string `json:"queryPoint" bson:"queryPoint"`
	} `json:"permissions" bson:"permissions"`

	ModifierUser string `json:"modifierUser" bson:"modifierUser"`
	CreatedAt    int64  `json:"createdAt" bson:"createdAt"`
	ModifiedAt   int64  `json:"modifiedAt" bson:"modifiedAt"`
	DeletedAt    int64  `json:"deletedAt" bson:"deletedAt"`
}

type Repository interface {
	GetList(ctx context.Context) ([]Auth, error)
	GetByID(ctx context.Context, authID string) (*Auth, error)
	Create(ctx context.Context, auth Auth) (*Auth, error)
	Update(ctx context.Context, auth Auth) (*Auth, error)
	Delete(ctx context.Context, authID string) (*Auth, error)
}
