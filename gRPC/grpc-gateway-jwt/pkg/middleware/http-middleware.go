package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

// Middleware handles all jwt parsing and validation automatically when used
type Middleware struct {
	// embed the validator to make token calls cleaner
	Validator
}

// NewMiddleware creates a new middleware that validates using the
// given public key file
func NewMiddleware(publicKeyPath string) (*Middleware, error) {
	validator, err := NewValidator(publicKeyPath)
	if err != nil {
		return nil, fmt.Errorf("unable to create validator: %w", err)
	}

	return &Middleware{
		Validator: *validator,
	}, nil
}

func (m *Middleware) HandleHTTP(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check IsLoginRequest(r.Context())
		if r.URL.Path == "/login" {
			// log.Println("HTTP SERVER middleware - /login")
			fmt.Println("2. start /login")
			// check basic auth
			// _, _, ok := r.BasicAuth()
			// if !ok {
			// 	w.WriteHeader(http.StatusUnauthorized)
			// 	w.Write([]byte("missing basic auth")) //nolint
			// 	return
			// }

			// var v map[string]string
			// err := json.NewDecoder(r.Body).Decode(&v)
			// if err != nil {
			// 	w.WriteHeader(http.StatusBadRequest)
			// 	w.Write([]byte("invalid json body")) //nolint
			// 	return
			// }

			// fmt.Println(v["username"], v["password"])
			// if v["username"] == "" || v["password"] == "" {
			// 	w.WriteHeader(http.StatusUnauthorized)
			// 	w.Write([]byte("invalid credentials")) //nolint
			// 	return
			// }

			// r.SetBasicAuth(v["username"], v["password"])
			// // r.Header.Add("Content-Type", "application/json")

			// log.Printf("it is here 1")
			h.ServeHTTP(w, r)
			// log.Printf("it is here 2")

			fmt.Println("2. finish /login")
			return
		}

		// fmt.Println(r.URL.Path)
		fmt.Println("2. start")

		parts := strings.Split(r.Header.Get("Authorization"), " ")
		if len(parts) < 2 || parts[0] != "Bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("missing or invalid authorization header")) //nolint
			return
		}
		tokenString := parts[1]

		token, err := m.GetToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid token: " + err.Error())) //nolint
			return
		}

		// Get a new context with the parsed token
		ctx := ContextWithToken(r.Context(), token)

		// fmt.Println("* HTTP SERVER middleware validated and set set token")

		// call the next handler with the updated context
		h.ServeHTTP(w, r.WithContext(ctx))

		fmt.Println("2. finish")
	}
}
