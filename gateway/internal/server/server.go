package server

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/micro/go-micro/v2/web"

	httphandler "github.com/lucasalmeron/microc3/gateway/internal/http"
	httpmiddlewares "github.com/lucasalmeron/microc3/gateway/internal/http/middlewares"
)

type Server struct {
	service web.Service
	router  *mux.Router
}

func (srv *Server) Init() {

	srv.router = mux.NewRouter().StrictSlash(true)
	apiPath := srv.router.PathPrefix("/api").Subrouter()
	srv.router.Use(httpmiddlewares.AuthMiddleware)

	// Only matches if domain is "www.example.com".
	//router.Host("www.example.com")
	httphandler.InitUserHandler(apiPath)
	httphandler.InitQueryPointsHandler(apiPath)

	// create new web service
	srv.service = web.NewService(
		web.Name("go.micro.api.api"),
		web.Version("latest"),
		web.Handler(srv.router),
	)

	// initialise service
	if err := srv.service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	//service.Handle("/api", srv.router)

	// run service
	if err := srv.service.Run(); err != nil {
		log.Fatal(err)
	}

}
