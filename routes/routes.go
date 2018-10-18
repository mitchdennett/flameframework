package routes

import (
	"reflect"
	"runtime"
	"fmt"
	"strings"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

var WebRoutes []Route
var parentRoutes map[string]string

func init() {
	WebRoutes = []Route{}
	parentRoutes = make(map[string]string)
}

type Route struct {
	Route_type string
	Controller func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
	Url string
	MiddlewareList []reflect.Type
}

func RouteFilePrefix(prefix string) {
	if file, ok := getCallerFile(3); ok {
		parentRoutes[file] = prefix
	}	
}

func Get(url string, middleware string, controller func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)) {
	define("GET", url, middleware, controller)
}

func Post(url string, middleware string, controller func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)) {
	define("POST", url, middleware, controller)
}

func define(routeType string, url string, middleware string, controller func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)) {
	if file, ok := getCallerFile(4); ok {
		if parentRoutes[file] != "" {
			if url == "/"{
				url = parentRoutes[file]
			}else{
				url = parentRoutes[file] + url
			}
		}
	}

	r := Route{}
	r.Route_type = routeType
	r.Controller = controller
	r.Url = url
	WebRoutes = append(WebRoutes, r)
}

func getCallerFile(depth int) (string, bool) {
	fpcs := make([]uintptr, 1)
    // Skip 'depth' levels to get the caller
    n := runtime.Callers(depth, fpcs)
    if n == 0 {
        fmt.Println("MSG: NO CALLER")
    }
 
    caller := runtime.FuncForPC(fpcs[0]-1)
    if caller == nil {
        fmt.Println("MSG CALLER WAS NIL")
	} else {
		// Print the file name and line number
		filePath, _ := caller.FileLine(fpcs[0]-1)
		fileSlice := strings.Split(filePath, " ")
		file := fileSlice[0]

		return file, true;
	}

	return "", false
}

func (r Route) Middleware(middleware ...func()(reflect.Type)) Route {
	for _, ty := range middleware {
		midd := ty()
		r.MiddlewareList = append(r.MiddlewareList, midd)
	}
	return r
}
