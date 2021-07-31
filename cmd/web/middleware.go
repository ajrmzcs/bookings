package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the Page")
		next.ServeHTTP(w, r)
	})
}

func NoSufr(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		Path:       "/",
		Secure:     false,
		HttpOnly:   true,
		SameSite:   http.SameSiteLaxMode,
	})

	return csrfHandler
}
