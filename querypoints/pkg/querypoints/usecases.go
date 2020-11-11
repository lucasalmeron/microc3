package querypoint

import (
	"context"
	"math"

	user "github.com/lucasalmeron/microc3/users/pkg/users"
)

var repository Repository

func SetRepository(repo Repository) {
	repository = repo
}

func (queryPoint *QueryPoint) Validate() error {

	return nil
}

func (queryPoint *QueryPoint) GetList() ([]QueryPoint, error) {
	return repository.GetList(context.TODO())
}

func (queryPoint *QueryPoint) GetbyID(queryPointID string) (*QueryPoint, error) {
	return repository.GetByID(context.TODO(), queryPointID)
}

func (queryPoint *QueryPoint) GetbyName(queryString string) (*QueryPoint, error) {
	return repository.GetByName(context.TODO(), queryString)
}

func (queryPoint *QueryPoint) GetPaginated(pageOptions *user.PageOptions) (*Page, error) {
	return repository.GetPaginated(context.TODO(), pageOptions)
}

func (queryPoint *QueryPoint) Save() (*QueryPoint, error) {
	if queryPoint.ID == "" {
		return repository.Create(context.TODO(), *queryPoint)
	} else {
		return repository.Update(context.TODO(), *queryPoint)
	}
}

func (user *QueryPoint) Delete(queryPointID string) (*QueryPoint, error) {
	return repository.Delete(context.TODO(), queryPointID)
}

//-----------//PAGINATION//-----------//

func (p *Page) CalcNumberOfPages(pageOptions *user.PageOptions) {
	p.PageNumber = pageOptions.PageNumber
	p.NumberOfPages = int64(math.Ceil(float64(p.Length) / float64(pageOptions.RegistersNumber)))
}
