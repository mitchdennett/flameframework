package middleware

import (
	"net/http"
)

type Middleware interface {
	Before(w http.ResponseWriter, r *http.Request) bool
	After(w http.ResponseWriter, r *http.Request)
}