package services

import (
	"github.com/gorilla/sessions"
	"net/http"
)

var Session_Name = []byte("lalalalaa")

var SessionStore = sessions.NewCookieStore(Session_Name)

func GetSession(r *http.Request, name string) (*sessions.Session, error) {
	return SessionStore.Get(r, name)
}
