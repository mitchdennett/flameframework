package test

import (
	"net/http"

	"github.com/mitchdennett/flameframework/view"
	"github.com/mitchdennett/httprouter"
)

type TestController struct {
}

func (TestController) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	view.Render("welcome.html", nil)
}
