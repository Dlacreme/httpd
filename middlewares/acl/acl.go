// Package acl provides http.Handlers to prevent access to pages for
// authenticated users and for non-authenticated users.
package acl

import (
	"net/http"

	"github.com/Dlacreme/httpd/back/flight"
)

func rejectRequest(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusUnauthorized)
}

// ForAdmin gives access to admin only
func ForAdmin(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := flight.Context(w, r)

		// If user is not authenticated, don't allow them to access the page
		t := c.Sess.Values["UserTypeId"]
		if t == nil {
			rejectRequest(w, r)
			return
		}
		tt := t.(int)
		if tt > 1 {
			rejectRequest(w, r)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// ForStaff gives access to admin only
func ForStaff(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := flight.Context(w, r)

		// If user is not authenticated, don't allow them to access the page
		t := c.Sess.Values["UserTypeId"]
		if t == nil {
			rejectRequest(w, r)
			return
		}
		tt := t.(int)
		if tt > 2 {
			rejectRequest(w, r)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// ForMember gives access to admin only
func ForMember(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := flight.Context(w, r)

		// If user is not authenticated, don't allow them to access the page
		t := c.Sess.Values["UserTypeId"]
		if t == nil {
			rejectRequest(w, r)
			return
		}
		tt := t.(int)
		if tt > 3 {
			rejectRequest(w, r)
			return
		}

		h.ServeHTTP(w, r)
	})
}

// DisallowAuth does not allow authenticated users to access the page.
func DisallowAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c := flight.Context(w, r)

		// If user is authenticated, don't allow them to access the page
		if c.Sess.Values["UserId"] != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		h.ServeHTTP(w, r)
	})
}
