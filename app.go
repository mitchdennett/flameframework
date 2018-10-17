package flame

import (
    "net/http"
    "sync"
)

type app struct {
    response http.ResponseWriter
}

var Current *app
var once sync.Once

func init() {
	Current = GetInstance()
}

func GetInstance() *app {
    once.Do(func() {
        Current = &app{}
    })
    return Current
}

func (a *app) SetResponse(r http.ResponseWriter) {
    a.response = r
    return
}

func (a *app) GetResponse() (http.ResponseWriter) {
    return a.response
}
