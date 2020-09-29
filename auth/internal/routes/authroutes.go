package routes

import (
	"github.com/google/uuid"
	path "github.com/lucasalmeron/microc3/auth/pkg/path"
)

var routes = []path.Path{
	{
		ID:          uuid.New().String(),
		Path:        "/api/users/login",
		Method:      "POST",
		Permissions: []string{"any"},
	},
	{
		ID:          uuid.New().String(),
		Path:        "/api/users/login",
		Method:      "POST",
		Permissions: []string{"any"},
	},
	{
		ID:          uuid.New().String(),
		Path:        "/api/users/login",
		Method:      "POST",
		Permissions: []string{"any"},
	},
}
