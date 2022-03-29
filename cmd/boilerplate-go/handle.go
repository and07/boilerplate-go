package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/and07/boilerplate-go/internal/pkg/template"
	"github.com/and07/boilerplate-go/pkg/data"
	handlers "github.com/and07/boilerplate-go/pkg/handlers/httpserver"
	"github.com/and07/boilerplate-go/pkg/service"
	"github.com/and07/boilerplate-go/pkg/service/mail"
	"github.com/and07/boilerplate-go/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	"github.com/jmoiron/sqlx"
	"github.com/markbates/goth/gothic"
	//"google.golang.org/api/idtoken"
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

func callbackProviderHandler(ctx context.Context, tpl *template.Template) func(res http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		/*
			// Get authorization code from URL
			authCode := r.URL.Query().Get("code")

			// Get ID token and access token through authorization code
			oauth2Token, err := oauth2Config.Exchange(ctx, authCode)
			if err != nil {
				http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
				return
			}
			rawIDToken, ok := oauth2Token.Extra("id_token").(string)
			if !ok {
				http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
				return
			}
		*/
		expirationTime := time.Now().Add(5 * time.Minute)
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   user.IDToken,
			Expires: expirationTime,
		})

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

func loginProviderHandler(ctx context.Context, tpl *template.Template) func(res http.ResponseWriter, req *http.Request) {
	// try to get the user without re-authenticating
	return func(w http.ResponseWriter, r *http.Request) {
		gothic.BeginAuthHandler(w, r)
	}
}

/*
func profileHandler(ctx context.Context, tpl *template.Template) func(res http.ResponseWriter, req *http.Request) {
	// try to get the user without re-authenticating
	return func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Get the JWT string from the cookie
		tknStr := c.Value

		payload, err := idtoken.Validate(context.Background(), tknStr, "328290909614-3casaiclr5c4ftspb91kun5ckl1av5om.apps.googleusercontent.com")
		if err != nil {
			fmt.Fprintf(w, "%s", err)
			return
		}

		fmt.Fprintf(w, "%#v", payload)

	}
}
*/

func publicHandle(ctx context.Context, tpl *template.Template, db *sqlx.DB, configs *utils.Configurations, logger hclog.Logger) http.Handler {
	/*
			rPublic := pat.New()
			rPublic.Get("/auth/{provider}/callback", userHandler(ctx, tpl))
			rPublic.Get("/logout/{provider}", logoutHandler(ctx, tpl))
			rPublic.Get("/auth/{provider}", providerHandler(ctx, tpl))
			rPublic.Get("/auth", authHandler(ctx, tpl))
			rPublic.Get("/", hiHandler(ctx, tpl))


		rPublic := http.NewServeMux()
		rPublic.HandleFunc("/", hiHandler(ctx, tpl))
		rPublic.HandleFunc("/auth/google/callback", userHandler(ctx, tpl))
		rPublic.HandleFunc("/logout/:provider", logoutHandler(ctx, tpl))
		rPublic.HandleFunc("/login", providerHandler(ctx, tpl))
		rPublic.HandleFunc("/auth", authHandler(ctx, tpl))
		rPublic.HandleFunc("/profile", profileHandler(ctx, tpl))

		return rPublic

	*/

	// validator contains all the methods that are need to validate the user json in request
	validator := data.NewValidation()

	// creation of user table.
	if db != nil {
		db.MustExec(userSchema)
		db.MustExec(verificationSchema)
	}

	// repository contains all the methods that interact with DB to perform CURD operations for user.
	//repository := data.NewPostgresRepository(db, logger)

	repositoryMemory := data.NewMemoryRepository(logger)

	// authService contains all methods that help in authorizing a user request
	authService := service.NewAuthService(logger, configs)

	// mailService contains the utility methods to send an email
	mailService := mail.NewMGMailService(logger, configs)

	googleKEY := os.Getenv("GOOGLE_KEY")
	if googleKEY == "" {
		googleKEY = "328290909614-3casaiclr5c4ftspb91kun5ckl1av5om.apps.googleusercontent.com"
	}
	googleSECRET := os.Getenv("GOOGLE_SECRET")
	if googleSECRET == "" {
		googleSECRET = "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"
	}
	// UserHandler encapsulates all the services related to user
	uh := handlers.NewAuthHandler(logger, configs, validator, repositoryMemory, authService, mailService, handlers.WithGoogleAuth(googleKEY, googleSECRET, "http://localhost:8080/auth/google/callback"))

	// create a serve mux
	sm := mux.NewRouter()

	// register handlers
	mailR := sm.PathPrefix("/verify").Methods(http.MethodPost).Subrouter()
	mailR.HandleFunc("/mail", uh.VerifyMail)
	mailR.HandleFunc("/password-reset", uh.VerifyPasswordReset)
	mailR.Use(uh.MiddlewareValidateVerificationData)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/signup", uh.Signup)
	postR.HandleFunc("/login", uh.Login)
	postR.Use(uh.MiddlewareValidateUser)

	// used the PathPrefix as workaround for scenarios where all the
	// get requests must use the ValidateAccessToken middleware except
	// the /refresh-token request which has to use ValidateRefreshToken middleware
	refToken := sm.PathPrefix("/refresh-token").Subrouter()
	refToken.HandleFunc("", uh.RefreshToken)
	refToken.Use(uh.MiddlewareValidateRefreshToken)

	get := sm.Methods(http.MethodGet).Subrouter()
	get.HandleFunc("/", hiHandler(ctx, tpl))
	//get.HandleFunc("/auth/{provider}/callback", callbackProviderHandler(ctx, tpl))
	get.HandleFunc("/auth/google/callback", uh.GoogleLogin)
	get.HandleFunc("/logout/{provider}", logoutHandler(ctx, tpl))
	get.HandleFunc("/login/google", uh.GoogleOAuth)
	get.HandleFunc("/auth", authHandler(ctx, tpl))

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/greet", uh.Greet)
	getR.HandleFunc("/profile", uh.Profile)
	getR.HandleFunc("/get-password-reset-code", uh.GeneratePassResetCode)
	getR.Use(uh.MiddlewareValidateAccessToken)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/update-username", uh.UpdateUsername)
	putR.HandleFunc("/reset-password", uh.ResetPassword)
	putR.Use(uh.MiddlewareValidateAccessToken)

	return sm

}
