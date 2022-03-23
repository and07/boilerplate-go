package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/and07/boilerplate-go/internal/pkg/template"
	"github.com/markbates/goth/gothic"
)

func hiHandler(ctx context.Context, tpl *template.Template) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		//span, _ := opentracing.StartSpanFromContext(ctx, "Scratch.hiHandler")
		//defer span.Finish()
		counter.Inc()

		xRealIP := r.Header.Get("X-Real-Ip")
		xForwardedFor := r.Header.Get("X-Forwarded-For")
		remoteAddr := r.RemoteAddr

		tpl.RenderTemplate(w, "main.html", fmt.Sprintf("X-Real-Ip:%s X-Forwarded-For:%s RemoteAddr:%s", xRealIP, xForwardedFor, remoteAddr))
	}
}

func userHandler(ctx context.Context, tpl *template.Template) func(res http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		tpl.RenderTemplate(w, "user.html", user)
	}
}

func logoutHandler(ctx context.Context, tpl *template.Template) func(res http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		gothic.Logout(w, r)
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
}

func providerHandler(ctx context.Context, tpl *template.Template) func(res http.ResponseWriter, req *http.Request) {
	// try to get the user without re-authenticating
	return func(w http.ResponseWriter, r *http.Request) {
		if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
			tpl.RenderTemplate(w, "user.html", gothUser)
		} else {
			gothic.BeginAuthHandler(w, r)
		}
	}
}

func publicHandle(ctx context.Context, tpl *template.Template) *http.ServeMux {
	/*
		rPublic := pat.New()
		rPublic.Get("/auth/{provider}/callback", userHandler(ctx, tpl))
		rPublic.Get("/logout/{provider}", logoutHandler(ctx, tpl))
		rPublic.Get("/auth/{provider}", providerHandler(ctx, tpl))
		rPublic.Get("/auth", authHandler(ctx, tpl))
		rPublic.Get("/", hiHandler(ctx, tpl))
	*/

	rPublic := http.NewServeMux()
	rPublic.HandleFunc("/", hiHandler(ctx, tpl))
	rPublic.HandleFunc("/auth/google/callback", userHandler(ctx, tpl))
	rPublic.HandleFunc("/logout", logoutHandler(ctx, tpl))
	rPublic.HandleFunc("/auth/", providerHandler(ctx, tpl))
	rPublic.HandleFunc("/auth", authHandler(ctx, tpl))

	return rPublic
}
