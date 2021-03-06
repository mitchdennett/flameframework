package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mitchdennett/flameframework/routes"
	"github.com/mitchdennett/flameframework/test"
)

func TestServer(t *testing.T) {
	var routeList = []routes.Route{
		routes.Get().Define("/", test.TestController{}),
	}

	handler := NewHandler()
	registerRoutes(handler, routeList)

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `<html>
    <head>
    </head>
    <body>
        <h1>Welcome</h1>
    </body>
</html>`

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
