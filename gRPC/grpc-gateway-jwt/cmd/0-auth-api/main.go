package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/QuanDN22/BE/gRPC/grpc-gateway-jwt/pkg/middleware"
)

func main() {
	issuer, err := middleware.NewIssuer("./auth.ed")
	if err != nil {
		panic(err)
	}

	auth, err := NewAuthAPI(issuer)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/login-auth-api", auth.HandleLogin)

	fmt.Println("Listening on :5000")
	err = http.ListenAndServe(":5000", mux)
	if err != nil {
		panic(err)
	}
}

// AuthService handles authentication and issues tokens
type AuthAPI struct {
	issuer *middleware.Issuer
}

// NewAuthService creates a new service using the given issuer
func NewAuthAPI(issuer *middleware.Issuer) (*AuthAPI, error) {
	if issuer == nil {
		return nil, errors.New("issuer is required")
	}

	return &AuthAPI{
		issuer: issuer,
	}, nil
}

func (a *AuthAPI) HandleLogin(w http.ResponseWriter, r *http.Request) {
	// check basic auth
	user, pass, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("missing basic auth")) //nolint
		return
	}

	// This is a bad idea in anything real. But this isn't an auth methods talk
	// this talk is about JWTs so we only have trivial auth checking here
	if user != "admin" || pass != "pass" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("invalid credentials")) //nolint
		return
	}

	tokenString, err := a.issuer.IssueToken("admin", []string{"admin", "basic"})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to issue token:" + err.Error())) //nolint
		return
	}

	_, _ = w.Write([]byte(tokenString + "\n"))
}
