package drivers

import (
	"fmt"
	"net/http"

	"github.com/mitchdennett/flameframework/encrypt"
)

type SessionCookieDriver struct {
	w http.ResponseWriter
	r *http.Request
}

func (s *SessionCookieDriver) SetRequest(req *http.Request) {
	s.r = req
}

func (s *SessionCookieDriver) SetResponse(w http.ResponseWriter) {
	s.w = w
}

func (s *SessionCookieDriver) Get(name string) string {
	fmt.Println(s.r)
	var cookie, err = s.r.Cookie(name)
	if err == nil {
		value, erro := encrypt.Decrypt(cookie.Value)
		if erro != nil {
			return value
		}
	}

	return ""
}

func (s *SessionCookieDriver) Has(name string) bool {

	if s.Get(name) != "" {
		return true
	}

	return false
}

func (s *SessionCookieDriver) Set(name string, value string) {

	encoded, err := encrypt.Encrypt(value)

	if err != nil {
		return
	}

	cookie := http.Cookie{
		Name:     name,
		Value:    encoded,
		HttpOnly: true,
	}

	http.SetCookie(s.w, &cookie)
}

func (s *SessionCookieDriver) Delete(name string) {

	cookie := http.Cookie{
		Name:     name,
		Value:    "",
		MaxAge:   0,
		HttpOnly: true,
	}

	http.SetCookie(s.w, &cookie)
}

func (s *SessionCookieDriver) All() {
}

func (s *SessionCookieDriver) Flash(key string, value string) {
}

func (s *SessionCookieDriver) Reset(flashOnly bool) {}
