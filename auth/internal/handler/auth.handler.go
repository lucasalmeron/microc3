package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protoauth "github.com/lucasalmeron/microc3/auth/pkg/auth/proto"
	protoqp "github.com/lucasalmeron/microc3/querypoints/pkg/querypoints/proto"
	protousers "github.com/lucasalmeron/microc3/users/pkg/users/proto"

	authroutes "github.com/lucasalmeron/microc3/auth/internal/routes"
	auth "github.com/lucasalmeron/microc3/auth/pkg/auth"
	queryPoint "github.com/lucasalmeron/microc3/querypoints/pkg/querypoints"
)

var (
	pubCreated  micro.Event
	pubMofidied micro.Event
	pubDeleted  micro.Event
	userClient  protousers.UsersService
	qpClient    protoqp.QueryPointsService
	tokenSecret = "565985%$#fjgSAS" ///ADD FROM ENVIRONMENT
)

func InitEvents(c client.Client) {
	pubCreated = micro.NewEvent("go.micro.auth.created", c)
	pubMofidied = micro.NewEvent("go.micro.auth.modified", c)
	pubDeleted = micro.NewEvent("go.micro.auth.deleted", c)
	//create gRPC clients//
	userClient = protousers.NewUsersService("go.micro.service.users", c)
	qpClient = protoqp.NewQueryPointsService("go.micro.service.querypoints", c)
}

func buildProtoPermission(auth auth.Auth) ([]*protoauth.Permission, error) {
	permissions := make([]*protoauth.Permission, 0)

	for _, p := range auth.Permissions {
		parsedMap := make(map[string]string)
		for key, value := range p.Access {
			if key == "queryPoint" {
				respQP, err := qpClient.GetByID(context.TODO(), &protoqp.RequestQueryPointID{
					Id: fmt.Sprintf("%v", value),
				})
				if err != nil {
					log.Error(err)
					return nil, err
				}
				var queryPoint queryPoint.QueryPoint
				queryPoint.ID = respQP.Id
				queryPoint.Name = respQP.Name
				queryPoint.Address = respQP.Address
				queryPoint.District = respQP.District
				queryPoint.Department = respQP.Department
				queryPoint.Phone = respQP.Phone
				/* more fields ....*/
				parsedQP, err := json.Marshal(queryPoint)
				if err != nil {
					log.Error(err)
					return nil, err
				}
				parsedMap[key] = string(parsedQP)
			} else {
				parsedMap[key] = fmt.Sprintf("%v", value)
			}
		}
		permissions = append(permissions, &protoauth.Permission{
			Id:     p.ID,
			Access: parsedMap,
		})
	}

	return permissions, nil

	/*return &protoauth.ResponseAuth{
		Id:          auth.ID,
		User:        auth.User,
		Permissions: permissions,
		CreatedAt:   auth.CreatedAt,
		ModifiedAt:  auth.ModifiedAt,
		DeletedAt:   auth.DeletedAt,
	}*/
}

type AuthHandler struct{}

func (e *AuthHandler) LogIn(ctx context.Context, req *protoauth.RequestAuthLogIn, res *protoauth.ResponseLogIn) error {
	log.Info("Received auth.LogIn request")

	respUser, err := userClient.LogIn(context.TODO(), &protousers.RequestUserLogIn{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		log.Error(err)
		return err
	}
	userAuth := &auth.Auth{
		User: respUser.Id,
	}
	userPermissions, err := userAuth.GetByUserID(userAuth.User)
	if err != nil {
		log.Error(err)
		return err
	}

	signedTime := time.Now()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":          respUser.Id,
		"name":        respUser.FirstName + " " + respUser.LastName,
		"email":       respUser.Email,
		"admin":       userPermissions.Admin,
		"permissions": userPermissions.Permissions,
		"iat":         signedTime.Unix(),
		"exp":         signedTime.Add(8 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	//RESPONSE
	res.Token = tokenString

	return nil
}

func (e *AuthHandler) AuthPath(ctx context.Context, req *protoauth.RequestAuthPath, res *protoauth.ResponseAuthPath) error {
	log.Info("Received auth.AuthPath request")
	request := &http.Request{
		Method: req.Method,
		URL:    &url.URL{Path: req.Path},
	}

	for _, r := range authroutes.Routes {
		muxrouter := mux.NewRouter()
		muxrouter.NewRoute().Path(r.Path).Methods(r.Method)

		if muxrouter.Match(request, &mux.RouteMatch{}) {
			fmt.Println("MATCH ROUTE")
			fmt.Println(r)
			if len(r.Permissions) > 0 { //if the route require permissions
				if req.Token == "" {
					res.Authorized = false
					return nil
				}
				splitedToken := strings.Split(req.Token, " ")
				if splitedToken[0] != "Bearer" {
					log.Error("Invalid Token format")
					return status.Error(codes.PermissionDenied, "Invalid Token format")
				}
				token, err := jwt.Parse(splitedToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
					}
					return []byte(tokenSecret), nil
				})
				if err != nil {
					log.Error(err)
					return status.Error(codes.PermissionDenied, err.Error())
				}
				if !token.Valid {
					log.Error("Invalid Token")
					return status.Error(codes.PermissionDenied, "Invalid Token")
				}

				//IF ADMIN IS AUTHORIZED
				if token.Claims.(jwt.MapClaims)["admin"] == true {
					res.Authorized = true
					return nil
				}

				///////////UNMARSHALING PERMISSIONS//////////
				// convert token to struct
				tokenByte, err := json.Marshal(token.Claims.(jwt.MapClaims)["permissions"])
				if err != nil {
					log.Error(err)
					return status.Error(codes.Internal, err.Error())
				}
				var permissions []auth.Permission
				err = json.Unmarshal(tokenByte, &permissions)
				if err != nil {
					log.Error(err)
					return status.Error(codes.Internal, err.Error())
				}
				// convert token to struct
				var Authorize bool
				for _, p := range permissions {

					for _, permissionRequired := range r.Permissions {
						if _, ok := p.Access[permissionRequired]; ok {
							Authorize = true
						}
					}
				}
				if Authorize {
					res.Authorized = true
				} else {
					res.Authorized = false
				}
				return nil
			}

			res.Authorized = true
			return nil
		}
	}
	return status.Error(codes.NotFound, "Path doesn't exist")
}

func (e *AuthHandler) GetByID(ctx context.Context, req *protoauth.RequestAuthID, res *protoauth.ResponseAuth) error {
	log.Info("Received auth.GetByID request")
	reqAuth := new(auth.Auth)
	foundAuth, err := reqAuth.GetByID(req.Id)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	permissions, err := buildProtoPermission(*foundAuth)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	//RESPONSE+
	res.Id = foundAuth.ID
	res.User = foundAuth.User
	res.Permissions = permissions
	res.CreatedAt = foundAuth.CreatedAt
	res.ModifiedAt = foundAuth.ModifiedAt
	res.DeletedAt = foundAuth.DeletedAt

	return nil
}

func (e *AuthHandler) GetByUserID(ctx context.Context, req *protoauth.RequestUserID, res *protoauth.ResponseAuth) error {
	log.Info("Received auth.GetByUserID request")
	reqAuth := new(auth.Auth)
	foundAuth, err := reqAuth.GetByUserID(req.User)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	permissions, err := buildProtoPermission(*foundAuth)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE+
	res.Id = foundAuth.ID
	res.User = foundAuth.User
	res.Admin = foundAuth.Admin
	res.Permissions = permissions
	res.CreatedAt = foundAuth.CreatedAt
	res.ModifiedAt = foundAuth.ModifiedAt
	res.DeletedAt = foundAuth.DeletedAt

	return nil
}

func (e *AuthHandler) GetQueryPointsByUserID(ctx context.Context, req *protoauth.RequestUserID, res *protoauth.ResponseAuthWithQueryPoint) error {
	log.Info("Received auth.GetQueryPointsByUserID request")
	reqAuth := new(auth.Auth)
	foundAuth, err := reqAuth.GetByUserID(req.User)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}
	queryPoints := []string{}
	for _, p := range foundAuth.Permissions {
		if qp, ok := p.Access["queryPoint"]; ok {
			queryPoints = append(queryPoints, fmt.Sprintf("%v", qp))
		}
	}

	//RESPONSE+
	res.Id = foundAuth.ID
	res.User = foundAuth.User
	res.Admin = foundAuth.Admin
	res.QueryPointsIDs = queryPoints
	return nil
}

func (e *AuthHandler) Create(ctx context.Context, req *protoauth.RequestCreateAuth, res *protoauth.ResponseAuth) error {
	log.Info("Received auth.Create request")
	reqAuth := &auth.Auth{
		User: req.User,
	}
	for _, p := range req.Permissions {
		parsedMap := make(map[string]interface{})
		for key, value := range p.Access {
			parsedMap[key] = value
		}
		reqAuth.Permissions = append(reqAuth.Permissions, auth.Permission{
			Access: parsedMap,
		})
	}

	err := reqAuth.Validate()
	if err != nil {
		log.Error(err)
		return status.Error(codes.InvalidArgument, err.Error())
	}

	user, err := userClient.GetUserByID(ctx, &protousers.RequestUserID{Id: reqAuth.User})
	if err != nil {
		log.Error(err)
		return err
	}

	if user.Id == "" {
		log.Error("User doesn't exist")
		return status.Error(codes.AlreadyExists, "User doesn't exist")
	}

	createdAuth, err := reqAuth.Save()
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	permissions, err := buildProtoPermission(*createdAuth)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = createdAuth.ID
	res.Permissions = permissions
	res.CreatedAt = createdAuth.CreatedAt
	res.ModifiedAt = createdAuth.ModifiedAt
	res.DeletedAt = createdAuth.DeletedAt

	err = pubCreated.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}

func (e *AuthHandler) Update(ctx context.Context, req *protoauth.RequestUpdateAuth, res *protoauth.ResponseAuth) error {
	log.Info("Received auth.Update request")
	if req.Id == "" {
		log.Error("Invalid ID")
		return status.Error(codes.InvalidArgument, "Invalid ID")
	}
	reqAuth := &auth.Auth{
		ID:   req.Id,
		User: req.User,
	}

	for _, p := range req.Permissions {
		parsedMap := make(map[string]interface{})
		for key, value := range p.Access {
			parsedMap[key] = value
		}
		reqAuth.Permissions = append(reqAuth.Permissions, auth.Permission{
			Access: parsedMap,
		})
	}

	err := reqAuth.Validate()
	if err != nil {
		log.Error(err)
		return status.Error(codes.InvalidArgument, err.Error())
	}

	updatedAuth, err := reqAuth.Save()
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	permissions, err := buildProtoPermission(*updatedAuth)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = updatedAuth.ID
	res.Permissions = permissions
	res.CreatedAt = updatedAuth.CreatedAt
	res.ModifiedAt = updatedAuth.ModifiedAt
	res.DeletedAt = updatedAuth.DeletedAt

	err = pubMofidied.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}

func (e *AuthHandler) Delete(ctx context.Context, req *protoauth.RequestAuthID, res *protoauth.ResponseAuth) error {
	log.Info("Received auth.Delete request")
	reqauth := new(auth.Auth)
	deletedAuth, err := reqauth.Delete(req.Id)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	permissions, err := buildProtoPermission(*deletedAuth)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = deletedAuth.ID
	res.Permissions = permissions
	res.CreatedAt = deletedAuth.CreatedAt
	res.ModifiedAt = deletedAuth.ModifiedAt
	res.DeletedAt = deletedAuth.DeletedAt

	err = pubDeleted.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}

func (e *AuthHandler) PushPermission(ctx context.Context, req *protoauth.RequestPushPermission, res *protoauth.ResponseAuth) error {
	log.Info("Received auth.PushPermission request")

	if req.UserID == "" {
		log.Error("Invalid userID")
		return status.Error(codes.InvalidArgument, "Invalid userID")
	}

	reqAuth := &auth.Auth{
		User: req.UserID,
	}

	parsedMap := make(map[string]interface{})
	for key, value := range req.Permission.Access {
		parsedMap[key] = value
	}

	newPermission := &auth.Permission{
		Access: parsedMap,
	}

	err := newPermission.Validate()
	if err != nil {
		log.Error(err)
		return status.Error(codes.InvalidArgument, err.Error())
	}

	updatedAuth, err := reqAuth.PushPermission(*newPermission)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	permissions, err := buildProtoPermission(*updatedAuth)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = updatedAuth.ID
	res.User = updatedAuth.User
	res.Permissions = permissions
	res.CreatedAt = updatedAuth.CreatedAt
	res.ModifiedAt = updatedAuth.ModifiedAt
	res.DeletedAt = updatedAuth.DeletedAt

	err = pubMofidied.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}

func (e *AuthHandler) UpdatePermission(ctx context.Context, req *protoauth.RequestPushPermission, res *protoauth.ResponseAuth) error {
	log.Info("Received auth.UpdatePermission request")

	if req.UserID == "" {
		log.Error("Invalid userID")
		return status.Error(codes.InvalidArgument, "Invalid userID")
	}

	reqAuth := &auth.Auth{
		User: req.UserID,
	}

	parsedMap := make(map[string]interface{})
	for key, value := range req.Permission.Access {
		parsedMap[key] = value
	}

	newPermission := &auth.Permission{
		ID:     req.Permission.Id,
		Access: parsedMap,
	}

	err := newPermission.Validate()
	if err != nil {
		log.Error(err)
		return status.Error(codes.InvalidArgument, err.Error())
	}

	updatedAuth, err := reqAuth.UpdatePermission(*newPermission)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	permissions, err := buildProtoPermission(*updatedAuth)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = updatedAuth.ID
	res.User = updatedAuth.User
	res.Permissions = permissions
	res.CreatedAt = updatedAuth.CreatedAt
	res.ModifiedAt = updatedAuth.ModifiedAt
	res.DeletedAt = updatedAuth.DeletedAt

	err = pubMofidied.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}

	return nil
}

func (e *AuthHandler) DeletePermission(ctx context.Context, req *protoauth.RequestDeletePermission, res *protoauth.ResponseAuth) error {
	log.Info("Received auth.DeletePermission request")
	if req.UserID == "" {
		log.Error("Invalid userID")
		return status.Error(codes.InvalidArgument, "Invalid userID")
	}
	if req.PermissionID == "" {
		log.Error("Invalid permissionID")
		return status.Error(codes.InvalidArgument, "Invalid permissionID")
	}
	reqAuth := &auth.Auth{
		User: req.UserID,
	}

	updatedAuth, err := reqAuth.DeletePermission(req.PermissionID)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	permissions, err := buildProtoPermission(*updatedAuth)
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	//RESPONSE
	res.Id = updatedAuth.ID
	res.User = updatedAuth.User
	res.Permissions = permissions
	res.CreatedAt = updatedAuth.CreatedAt
	res.ModifiedAt = updatedAuth.ModifiedAt
	res.DeletedAt = updatedAuth.DeletedAt

	err = pubMofidied.Publish(ctx, res)
	if err != nil {
		log.Error(err)
	}
	return nil
}
