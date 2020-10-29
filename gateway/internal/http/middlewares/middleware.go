package httpmiddlewares

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	protoauth "github.com/lucasalmeron/microc3/auth/pkg/auth/proto"
	errorprovider "github.com/lucasalmeron/microc3/gateway/internal/helper"
	"github.com/micro/go-micro/v2/client"
)

var (
	authClient = protoauth.NewAuthService("go.micro.service.auth", client.DefaultClient)
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		response, err := authClient.AuthPath(context.TODO(), &protoauth.RequestAuthPath{
			Token:  r.Header.Get("Authorization"),
			Host:   r.Host,
			Path:   r.RequestURI,
			Method: r.Method,
		})

		if err != nil {
			log.Print(err)
			/*errStatus, _ := status.FromError(err)
			fmt.Println(errStatus.Message())
			fmt.Println(errStatus.Code())*/

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusForbidden, errorprovider.ConvertToJSON(err)})
			return
		}

		if response.Authorized {
			next.ServeHTTP(w, r)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&errorprovider.HttpError{http.StatusUnauthorized, "Unauthorized"})

	})
}
