package user

import (
	"context"
)

type User struct {
	ID             string `json:"id" bson:"_id,omitempty"`
	FirstName      string `json:"firstName" bson:"firstName"`
	LastName       string `json:"lastName" bson:"lastName"`
	DocumentNumber string `json:"documentNumber" bson:"documentNumber"`
	Password       string `json:"password" bson:"password"`
	RePassword     string `json:"repassword"`
	Email          string `json:"email" bson:"email"`
	PhoneNumber    string `json:"phoneNumber" bson:"phoneNumber"`
	GDEUser        string `json:"GDEUser" bson:"GDEUser"`
	Position       string `json:"position" bson:"position"`

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
	GetList(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, userID string) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetPaginated(ctx context.Context, pageOpt *PageOptions) (*Page, error)
	Create(ctx context.Context, user User) (*User, error)
	Update(ctx context.Context, user User) (*User, error)
	Delete(ctx context.Context, userID string) (*User, error)
}

//-----------//PAGINATION//-----------//

type Page struct {
	Data          []User `json:"data" bson:"data"`
	PageNumber    int64  `json:"pageNumber" bson:"pageNumber"`
	NumberOfPages int64  `json:"registersNumber" bson:"numberOfPages"`
	Length        int64  `json:"length" bson:"length"`
}

type PageOptions struct {
	PageNumber      int64    `json:"pageNumber"`
	RegistersNumber int64    `json:"registersNumber"`
	OrderBy         Filter   `json:"orderBy"`
	Filters         []Filter `json:"filters"`
}

type Filter struct {
	Field string `json:"field"`
	Value string `json:"value"`
}
