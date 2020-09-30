package routes

import (
	"net/http"

	"github.com/google/uuid"
	path "github.com/lucasalmeron/microc3/auth/pkg/path"
)

var Routes = []path.Path{
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
		Permissions: []string{"admin"},
	},
	{
		ID:          uuid.New().String(),
		Path:        "/api/users/email/{email}",
		Method:      http.MethodGet,
		Permissions: []string{"admin"},
	},
	{
		ID:          uuid.New().String(),
		Path:        "/api/users/create",
		Method:      http.MethodPost,
		Permissions: []string{},
	},
}
