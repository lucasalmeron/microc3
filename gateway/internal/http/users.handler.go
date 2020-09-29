package httphandler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/gorilla/mux"
	"github.com/micro/go-micro/v2/client"

	protoauth "github.com/lucasalmeron/microc3/auth/pkg/auth/proto"
	protousers "github.com/lucasalmeron/microc3/users/pkg/users/proto"

	user "github.com/lucasalmeron/microc3/users/pkg/users"

	errorprovider "github.com/lucasalmeron/microc3/gateway/internal/helper"
)

var (
	userClient protousers.UsersService
)

type UsersHandler struct{}

func InitUserHandler(router *mux.Router) {
	//create gRPC clients//
	userClient = protousers.NewUsersService("go.micro.service.users", client.DefaultClient)
	authClient = protoauth.NewAuthService("go.micro.service.auth", client.DefaultClient)

	handler := new(UsersHandler)

	router.Path("/users/list").HandlerFunc(handler.GetList).Methods(http.MethodGet, http.MethodOptions)
	router.Path("/users/id/{userID:[0-9a-fA-F]{24}}").HandlerFunc(handler.GetByID).Methods(http.MethodGet, http.MethodOptions)
	router.Path("/users/email/{email}").HandlerFunc(handler.GetByEmail).Methods(http.MethodGet, http.MethodOptions)

	router.Path("/users/create").HandlerFunc(handler.Create).Methods(http.MethodPost, http.MethodOptions)
	router.Path("/users/update").HandlerFunc(handler.Update).Methods(http.MethodPut, http.MethodOptions)

}

func (h UsersHandler) GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response, err := userClient.GetUsers(context.TODO(), &empty.Empty{})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, errorprovider.ConvertToJSON(err)})
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (h UsersHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userID := mux.Vars(r)["userID"]

	response, err := userClient.GetUserByID(context.TODO(), &protousers.RequestUserID{Id: userID})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, errorprovider.ConvertToJSON(err)})
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (h UsersHandler) GetByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userEmail := mux.Vars(r)["email"]

	response, err := userClient.GetUserByEmail(context.TODO(), &protousers.RequestUserEmail{Email: userEmail})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, errorprovider.ConvertToJSON(err)})
		return
	}
	json.NewEncoder(w).Encode(response)
}

func (h UsersHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)

	var user user.User

	if err := decoder.Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, "Error unmarshalling request body"})
		return
	}

	response, err := userClient.CreateUser(context.TODO(), &protousers.RequestCreateUser{
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		DocumentNumber: user.DocumentNumber,
		Password:       user.Password,
		Repassword:     user.RePassword,
		Email:          user.Email,
		PhoneNumber:    user.PhoneNumber,
		GDEUser:        user.GDEUser,
		Position:       user.Position,
	})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, errorprovider.ConvertToJSON(err)})
		return
	}

	json.NewEncoder(w).Encode(response)
}

func (h UsersHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)

	var user user.User

	if err := decoder.Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, "Error unmarshalling request body"})
		return
	}

	response, err := userClient.UpdateUser(context.TODO(), &protousers.RequestUpdateUser{
		Id:             user.ID,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		DocumentNumber: user.DocumentNumber,
		Password:       user.Password,
		Email:          user.Email,
		PhoneNumber:    user.PhoneNumber,
		GDEUser:        user.GDEUser,
		Position:       user.Position,
	})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, errorprovider.ConvertToJSON(err)})
		return
	}

	json.NewEncoder(w).Encode(response)
}
