package server

import (
	"net/http"
	"reflect"
	"fmt"
	"github.com/julienschmidt/httprouter"
	. "github.com/mitchdennett/flameframework"
	. "github.com/mitchdennett/flameframework/middleware"
	"github.com/mitchdennett/flameframework/routes"
	
)

type myHandler struct {
	router *httprouter.Router
}

var middlewareMap map[string][]Middleware

// NewMux makes a new empty Mux.
func NewHandler() *myHandler {
	return &myHandler{router: httprouter.New()}
}

func (mux *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	router := mux.router
	path := r.URL.Path

	if handle, ps, _ := router.Lookup(r.Method, path); handle != nil {

		//Running before middleware
		abortExecution := false
		middlewareList := middlewareMap[r.Method+"::"+path]
		for _, middleware := range middlewareList {
			retBool := middleware.Before(w, r)

			if !abortExecution {
				abortExecution = retBool
			}
		}

		Current.SetResponse(w)

		if !abortExecution {
			handle(w, r, ps)
		}

		for _, middleware := range middlewareList {
			middleware.After(w, r)
		}

		Current.SetResponse(nil)
		return
	}

	http.NotFound(w, r)
	return
}

func ListenAndServe(addr string) {
	handler := NewHandler()
	server := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	registerRoutes(handler, routes.WebRoutes)

	server.ListenAndServe()
}

func registerRoutes(handler *myHandler, routeList []routes.Route) {

	fmt.Println(routeList)

	middlewareMap = make(map[string][]Middleware)

	for _, route := range routeList {
		if route.Route_type == "GET" {
			handler.router.GET(route.Url, route.Controller)
		} else if route.Route_type == "POST" {
			handler.router.POST(route.Url, route.Controller)
		}

		t := []Middleware{}
		for _, middleware := range route.MiddlewareList {
			ms := reflect.New(middleware).Elem().Interface().(Middleware)
			t = append(t, ms)
		}
		middlewareMap[route.Route_type+"::"+route.Url] = t
	}

}
