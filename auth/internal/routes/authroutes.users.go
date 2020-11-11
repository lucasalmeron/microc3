package routes

import (
	"net/http"

	"github.com/google/uuid"
	path "github.com/lucasalmeron/microc3/auth/pkg/path"
)

//Routes is an array of all auth routes
var Routes []path.Path

func LoadUsersRoutes() {
	userRoutes := []path.Path{
		{
			ID:          uuid.New().String(),
			Path:        "/api/users/login",
			Method:      http.MethodPost,
			Permissions: []string{},
		},
		{
			ID:          uuid.New().String(),
			Path:        "/api/users/id/{userID:[0-9a-fA-F]{24}}",
			Method:      http.MethodGet,
			Permissions: []string{"admin", "query"},
		},
		{
			ID:          uuid.New().String(),
			Path:        "/api/users/email/{email}",
			Method:      http.MethodGet,
			Permissions: []string{"admin"},
		},
		{
			ID:          uuid.New().String(),
			Path:        "/api/users/paginated",
			Method:      http.MethodPost,
			Permissions: []string{"admin"},
		},
		{
			ID:          uuid.New().String(),
			Path:        "/api/users",
			Method:      http.MethodPost,
			Permissions: []string{},
		},
		{
			ID:          uuid.New().String(),
			Path:        "/api/users",
			Method:      http.MethodPut,
			Permissions: []string{},
		},
		{
			ID:          uuid.New().String(),
			Path:        "/api/users/pushpermission",
			Method:      http.MethodPut,
			Permissions: []string{},
		},
		{
			ID:          uuid.New().String(),
			Path:        "/api/users/deletepermission",
			Method:      http.MethodPut,
			Permissions: []string{},
		},
	}
	Routes = append(Routes, userRoutes...)
}
