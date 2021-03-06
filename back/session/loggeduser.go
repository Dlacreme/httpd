package session

import (
	"github.com/gorilla/sessions"
)

var (
	loggedUserBuilder func(*sessions.Session) interface{}
)

func RegisterLoggedUserBuilder(fn func(*sessions.Session) interface{}) {
	loggedUserBuilder = fn
}

func BuildUser(sess *sessions.Session) interface{} {
	if loggedUserBuilder == nil {
		return nil
	}
	return loggedUserBuilder(sess)
}
