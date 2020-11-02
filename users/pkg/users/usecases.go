package user

import (
	"context"
	"errors"
	"math"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

var repository Repository

func SetRepository(repo Repository) {
	repository = repo
}

func (user *User) Validate() error {

	mdoc, err := regexp.MatchString(`^[0-9]{8}$`, user.DocumentNumber)
	if err != nil {
		return err
	}
	if !mdoc {
		return errors.New("Document Number must be a number of 8 digits")
	}
	memail, err := regexp.MatchString(`^[\w\.]+@([\w]+\.)+[\w]{2,4}$`, user.Email)
	if err != nil {
		return err
	}
	if !memail {
		return errors.New("Email is invalid")
	}
	return nil
}

func (user *User) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}

func (user *User) ComparePasswords(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (user *User) GetList() ([]User, error) {
	return repository.GetList(context.TODO())
}

func (user *User) GetbyID(userID string) (*User, error) {
	return repository.GetByID(context.TODO(), userID)
}
func (user *User) GetbyEmail(email string) (*User, error) {
	return repository.GetByEmail(context.TODO(), email)
}

func (user *User) GetPaginated(pageOptions *PageOptions) (*Page, error) {
	return repository.GetPaginated(context.TODO(), pageOptions)
}

func (user *User) Save() (*User, error) {
	if user.ID == "" {
		return repository.Create(context.TODO(), *user)
	} else {
		return repository.Update(context.TODO(), *user)
	}
}

func (user *User) Delete(userID string) (*User, error) {
	return repository.Delete(context.TODO(), userID)
}

//-----------//PAGINATION//-----------//

func (p *PageOptions) Validate() error {
	if p.PageNumber < 1 {
		return errors.New("pageNumber must be greater than 1")
	}
	if p.RegistersNumber < 2 || p.RegistersNumber > 10 {
		p.RegistersNumber = 10
	}
	if p.OrderBy.Field != "" || p.OrderBy.Value != "" {
		if p.OrderBy.Field == "" {
			return errors.New("OrderBy Field must be a string")
		}
		if p.OrderBy.Value == "" {
			return errors.New("OrderBy Value must be a string")
		}
		if p.OrderBy.Value != "asc" && p.OrderBy.Value != "desc" {
			return errors.New("OrderBy Value must be asc or desc")
		}
	} else {
		//set Order by default
		p.OrderBy.Field = "createdAt"
		p.OrderBy.Value = "desc"
	}
	if len(p.Filters) > 0 {
		for _, filter := range p.Filters {
			if filter.Field != "" || filter.Value != "" {
				if filter.Field == "" {
					return errors.New("Filter Field must be a string")
				}
				if filter.Value == "" {
					return errors.New("Filter Value must be a string")
				}
			}
		}
	}
	return nil
}

func (p *Page) CalcNumberOfPages(pageOptions *PageOptions) {
	p.PageNumber = pageOptions.PageNumber
	p.NumberOfPages = int64(math.Ceil(float64(p.Length) / float64(pageOptions.RegistersNumber)))
}
