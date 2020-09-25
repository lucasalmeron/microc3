package httphandler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/micro/go-micro/v2/client"

	protousers "github.com/lucasalmeron/microc3/users/pkg/users/proto"
)

var (
	userClient protousers.UsersService
)

type httpError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UsersHandler struct{}

func InitUserHandler(router *mux.Router) {

	userClient = protousers.NewUsersService("go.micro.service.users", client.DefaultClient)
	handler := new(UsersHandler)

	router.Path("/users").HandlerFunc(handler.GetList).Methods(http.MethodGet, http.MethodOptions)
	router.Path("/users/{usersID:[0-9a-fA-F]{24}}").HandlerFunc(handler.GetByID).Methods(http.MethodGet, http.MethodOptions)

	router.Path("/users").HandlerFunc(handler.Create).Methods(http.MethodPost, http.MethodOptions)
	router.Path("/users").HandlerFunc(handler.Update).Methods(http.MethodPut, http.MethodOptions)

}

func (h UsersHandler) GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("list")
}

func (h UsersHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := mux.Vars(r)["usersID"]

	response, err := userClient.GetUserByID(context.TODO(), &protousers.RequestUserID{Id: userID})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&httpError{http.StatusInternalServerError, err.Error()})
		return
	}
	log.Print(response)
	json.NewEncoder(w).Encode(response)
}

func (h UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("create")
}

func (h UsersHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode("update")
}
