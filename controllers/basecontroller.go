package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)


type BaseController interface {
	Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}