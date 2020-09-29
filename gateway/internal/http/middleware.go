package httphandler

import (
	"fmt"
	"net/http"

	protoauth "github.com/lucasalmeron/microc3/auth/pkg/auth/proto"
)

var (
	authClient protoauth.AuthService
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		fmt.Println(token)

		/*if user, found := amw.tokenUsers[token]; found {
			// We found the token in our map
			log.Printf("Authenticated user %s\n", user)
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}*/
	})
}
