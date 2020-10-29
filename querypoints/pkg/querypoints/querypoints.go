package querypoint

import (
	"context"

	user "github.com/lucasalmeron/microc3/users/pkg/users"
)

type QueryPoint struct {
	ID         string `json:"id" bson:"_id,omitempty"`
	Name       string `json:"name" bson:"name"`
	Phone      string `json:"phone" bson:"phone"`
	Address    string `json:"address" bson:"address"`
	District   string `json:"district" bson:"district"`
	Department string `json:"department" bson:"department"`

	Actions []string `json:"actions" bson:"actions"`

	ModifierUser string `json:"modifierUser" bson:"modifierUser"`
	CreatedAt    int64  `json:"createdAt" bson:"createdAt"`
	ModifiedAt   int64  `json:"modifiedAt" bson:"modifiedAt"`
	DeletedAt    int64  `json:"deletedAt" bson:"deletedAt"`
}

type Repository interface {
	GetList(ctx context.Context) ([]QueryPoint, error)
	GetByID(ctx context.Context, queryPointID string) (*QueryPoint, error)
	GetPaginated(ctx context.Context, pageOpt *user.PageOptions) (*Page, error)
	Create(ctx context.Context, queryPoint QueryPoint) (*QueryPoint, error)
	Update(ctx context.Context, queryPoint QueryPoint) (*QueryPoint, error)
	Delete(ctx context.Context, queryPointID string) (*QueryPoint, error)
}

//-----------//PAGINATION//-----------//

type Page struct {
	Data          []QueryPoint `json:"data" bson:"data"`
	PageNumber    int64        `json:"pageNumber" bson:"pageNumber"`
	NumberOfPages int64        `json:"registersNumber" bson:"numberOfPages"`
	Length        int64        `json:"length" bson:"length"`
}
