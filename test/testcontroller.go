package test

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/mitchdennett/flameframework/view"
)

type TestController struct {
}

func (TestController) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	view.Render("welcome.html", nil)
}
