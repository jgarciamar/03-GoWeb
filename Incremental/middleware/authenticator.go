package middleware

import "net/http"

type Authenticator struct {
	token string
}

func NewAuthenticator(token string) *Authenticator {
	if token == "" {
		token = "example-token"
	}

	return &Authenticator{
		token: token,
	}
}

func (a *Authenticator) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authToken := r.Header.Get("Authorization")

		if authToken != a.token {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
		//After

	})
}
