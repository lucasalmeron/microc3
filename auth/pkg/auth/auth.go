package auth

import (
	"context"
)

/*type Permission struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	Read        bool   `json:"read" bson:"read"`
	Write       bool   `json:"write" bson:"write"`
	Responsible bool   `json:"responsible" bson:"responsible"`
	Query       bool   `json:"query" bson:"query"`
	Health      bool   `json:"health" bson:"health"`
	QueryPoint  string `json:"queryPoint" bson:"queryPoint"`
}*/
type Permission struct {
	ID          string            `json:"id" bson:"_id,omitempty"`
	Access      map[string]string `json:"access" bson:"access"`
	Read        bool              `json:"read" bson:"read"`
	Write       bool              `json:"write" bson:"write"`
	Responsible bool              `json:"responsible" bson:"responsible"`
	Query       bool              `json:"query" bson:"query"`
	Health      bool              `json:"health" bson:"health"`
	QueryPoint  string            `json:"queryPoint" bson:"queryPoint"`
}

type Auth struct {
	ID   string `json:"id" bson:"_id,omitempty"`
	User string `json:"user" bson:"user"`

	Admin       bool         `json:"admin" bson:"admin"`
	Permissions []Permission `json:"permissions" bson:"permissions"`

	ModifierUser string `json:"modifierUser" bson:"modifierUser"`
	CreatedAt    int64  `json:"createdAt" bson:"createdAt"`
	ModifiedAt   int64  `json:"modifiedAt" bson:"modifiedAt"`
	DeletedAt    int64  `json:"deletedAt" bson:"deletedAt"`
}

type Repository interface {
	GetByID(ctx context.Context, authID string) (*Auth, error)
	GetByUserID(ctx context.Context, userID string) (*Auth, error)
	Create(ctx context.Context, auth Auth) (*Auth, error)
	Update(ctx context.Context, auth Auth) (*Auth, error)
	Delete(ctx context.Context, authID string) (*Auth, error)
	PushPermission(ctx context.Context, userID string, permission Permission) (*Auth, error)
	DeletePermission(ctx context.Context, userID string, permissionID string) (*Auth, error)
}
