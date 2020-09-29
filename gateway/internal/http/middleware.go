package httphandler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	protoauth "github.com/lucasalmeron/microc3/auth/pkg/auth/proto"
	errorprovider "github.com/lucasalmeron/microc3/gateway/internal/helper"
)

var (
	authClient protoauth.AuthService
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		response, err := authClient.AuthPath(context.TODO(), &protoauth.RequestAuthPath{
			Token:  r.Header.Get("Authorization"),
			Host:   r.Host,
			Path:   r.RequestURI,
			Method: r.Method,
		})

		if err != nil {
			log.Print(err)
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
