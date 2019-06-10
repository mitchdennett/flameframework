package session

import (
	"net/http"
	"os"

	"github.com/mitchdennett/flameframework/contracts"
	"github.com/mitchdennett/flameframework/drivers"
)

var sessionDriver contracts.SessionContract

func init() {
	driver := os.Getenv("SESSION_DRIVER")
	if driver == "cookie" {
		sessionDriver = &drivers.SessionCookieDriver{}
	} else {

	}

	sessionDriver = &drivers.SessionCookieDriver{}
}

func Retrieve(w http.ResponseWriter, r *http.Request) contracts.SessionContract {
	sessionDriver.SetRequest(r)
	sessionDriver.SetResponse(w)
	return sessionDriver
}
