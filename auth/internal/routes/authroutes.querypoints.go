package routes

import (
	"net/http"

	"github.com/google/uuid"
	path "github.com/lucasalmeron/microc3/auth/pkg/path"
)

func LoadQueryPointsRoutes() {
	queryPointRoutes := []path.Path{
		{
			ID:          uuid.New().String(),
			Path:        "/api/querypoints/list",
			Method:      http.MethodGet,
			Permissions: []string{"admin"},
		},
		{
			ID:          uuid.New().String(),
			Path:        "/api/querypoints/id/{queryPointID:[0-9a-fA-F]{24}}",
			Method:      http.MethodGet,
			Permissions: []string{"admin", "query"},
		},
		{
			ID:          uuid.New().String(),
			Path:        "/api/querypoints/create",
			Method:      http.MethodPost,
			Permissions: []string{},
		},
		{
			ID:          uuid.New().String(),
			Path:        "/api/querypoints/update",
			Method:      http.MethodPut,
			Permissions: []string{},
		},
	}
	Routes = append(Routes, queryPointRoutes...)
}
