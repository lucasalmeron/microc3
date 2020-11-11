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
	authClient protoauth.AuthService
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

	router.Path("/users/login").HandlerFunc(handler.LogIn).Methods(http.MethodPost, http.MethodOptions)
	router.Path("/users/paginated").HandlerFunc(handler.GetPaginated).Methods(http.MethodPost, http.MethodOptions)
	router.Path("/users").HandlerFunc(handler.Create).Methods(http.MethodPost, http.MethodOptions)
	router.Path("/users").HandlerFunc(handler.Update).Methods(http.MethodPut, http.MethodOptions)

	router.Path("/users/pushpermission").HandlerFunc(handler.PushPermission).Methods(http.MethodPut, http.MethodOptions)
	router.Path("/users/deletepermission").HandlerFunc(handler.DeletePermission).Methods(http.MethodPut, http.MethodOptions)
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

func (h UsersHandler) LogIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)

	logInReq := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := decoder.Decode(&logInReq); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, "Error unmarshalling request body"})
		return
	}

	response, err := authClient.LogIn(context.TODO(), &protoauth.RequestAuthLogIn{
		Email:    logInReq.Email,
		Password: logInReq.Password,
	})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, errorprovider.ConvertToJSON(err)})
		return
	}

	json.NewEncoder(w).Encode(response)
}

func (h UsersHandler) GetPaginated(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)

	/* EXAMPLE
	{
		"pageNumber":1,
		"registersNumber":10,
		"filters":[
			{"field":"email","value":"luko.ar@gmail.com"}
		]
	}
	*/

	var pageOptions user.PageOptions

	if err := decoder.Decode(&pageOptions); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, "Error unmarshalling request body"})
		return
	}
	var filters []*protousers.RequestPageOptions_Filter
	if len(pageOptions.Filters) > 0 {
		for _, f := range pageOptions.Filters {
			filters = append(filters, &protousers.RequestPageOptions_Filter{
				Field: f.Field,
				Value: f.Value,
			})
		}
	}

	response, err := userClient.GetPaginatedWithQP(context.TODO(), &protousers.RequestPageOptions{
		PageNumber:      pageOptions.PageNumber,
		RegistersNumber: pageOptions.RegistersNumber,
		OrderBy: &protousers.RequestPageOptions_Filter{
			Field: pageOptions.OrderBy.Field,
			Value: pageOptions.OrderBy.Value,
		},
		Filters: filters,
	})

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

func (h UsersHandler) PushPermission(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := struct {
		User       string `json:"user"`
		Permission struct {
			Access map[string]string `json:"access"`
		} `json:"permission"`
	}{}

	//accessMap := make(map[string]string)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, "Error unmarshalling request body"})
		return
	}

	/*
		Read        bool
		Write       bool
		Responsible bool
		Query       bool
		Health      bool
	*/
	accessMap := make(map[string]string)

	// REMEMBER TO VALIDATE QUERYPOINT//
	for key, value := range req.Permission.Access {
		if key == "queryPoint" {
			accessMap["queryPoint"] = value
		}
		if key == "read" {
			accessMap["read"] = value
		}
		if key == "write" {
			accessMap["write"] = value
		}
		if key == "responsible" {
			accessMap["responsible"] = value
		}
		if key == "query" {
			accessMap["query"] = value
		}
		if key == "health" {
			accessMap["health"] = value
		}
	}

	response, err := authClient.PushPermission(context.TODO(), &protoauth.RequestPushPermission{
		UserID: req.User,
		Permission: &protoauth.Permission{
			Access: accessMap,
		},
	})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, errorprovider.ConvertToJSON(err)})
		return
	}

	json.NewEncoder(w).Encode(response)
}

func (h UsersHandler) DeletePermission(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := struct {
		User         string `json:"user"`
		PermissionID string `json:"permission"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, "Error unmarshalling request body"})
		return
	}

	response, err := authClient.DeletePermission(context.TODO(), &protoauth.RequestDeletePermission{
		UserID:       req.User,
		PermissionID: req.PermissionID,
	})
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusInternalServerError, errorprovider.ConvertToJSON(err)})
		return
	}

	json.NewEncoder(w).Encode(response)
}
