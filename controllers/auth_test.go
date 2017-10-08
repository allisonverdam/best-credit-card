package controllers

import (
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	router := newRouter()
	router.Post("/auth", Auth("secret"))
	runAPITests(t, router, []apiTestCase{
		{"t1 - successful login", "POST", "/auth", `{"username":"allisonverdam", "password":"1234"}`, http.StatusOK, ""},
		{"t2 - unsuccessful login", "POST", "/auth", `{"username":"demo", "password":"bad"}`, http.StatusUnauthorized, ""},
		{"t3 - unsuccessful login", "POST", "/auth", `{"username":"demo", "password":"bad"}`, http.StatusUnauthorized, ""},
	})
}
