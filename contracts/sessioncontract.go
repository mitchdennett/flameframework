package contracts

import "net/http"

type SessionContract interface {
	Get(key string) string
	Set(key string, value string)
	Has(key string) bool
	All()
	Delete(key string)
	Flash(key string, value string)
	Reset(flashOnly bool)
	SetRequest(r *http.Request)
	SetResponse(w http.ResponseWriter)
}
