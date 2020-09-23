package querypoint

import (
	"context"
	"errors"
	"math"
	"regexp"

	user "github.com/lucasalmeron/microc3/users/pkg/users"
)

var repository Repository

func SetRepository(repo Repository) {
	repository = repo
}

func (queryPoint *QueryPoint) Validate() error {

	mdoc, err := regexp.MatchString(`^[0-9]{8}$`, queryPoint.DocumentNumber)
	if err != nil {
		return err
	}
	if !mdoc {
		return errors.New("Document Number must be a number of 8 digits")
	}
	memail, err := regexp.MatchString(`^[\w\.]+@([\w]+\.)+[\w]{2,4}$`, queryPoint.Email)
	if err != nil {
		return err
	}
	if !memail {
		return errors.New("Email is invalid")
	}
	return nil
}

func (queryPoint *QueryPoint) GetList() ([]QueryPoint, error) {
	return repository.GetList(context.TODO())
}

func (queryPoint *QueryPoint) GetbyID(queryPointID string) (*QueryPoint, error) {
	return repository.GetByID(context.TODO(), queryPointID)
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
