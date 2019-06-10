package controllers

import (
	"net/http"

	"github.com/mitchdennett/httprouter"
)

type BaseController interface {
	Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}
