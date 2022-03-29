package handlers

//https://medium.com/@bnprashanth256/oauth2-with-google-account-gmail-in-go-golang-1372c237d25e

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"

	"golang.org/x/oauth2"
	goog "golang.org/x/oauth2/google"
)

// Endpoint is Google's OAuth 2.0 endpoint.
var Endpoint = goog.Endpoint

const endpointProfile string = "https://www.googleapis.com/oauth2/v2/userinfo"

type googleUser struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	Link          string `json:"link"`
	Picture       string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}

func newConfig(clientKey string, secret string, callbackURL string, scopes ...string) *oauth2.Config {
	c := &oauth2.Config{
		ClientID:     clientKey,
		ClientSecret: secret,
		RedirectURL:  callbackURL,
		Endpoint:     Endpoint,
		Scopes:       []string{},
	}

	if len(scopes) > 0 {
		c.Scopes = append(c.Scopes, scopes...)
	} else {
		c.Scopes = []string{"email"}
	}
	return c
}

// SetState sets the state string associated with the given request.
// If no state string is associated with the request, one will be generated.
// This state is sent to the provider and can be retrieved during the
// callback.
var SetState = func(req *http.Request) string {
	state := req.URL.Query().Get("state")
	if len(state) > 0 {
		return state
	}

	// If a state query param is not passed in, generate a random
	// base64-encoded nonce so that the state on the auth URL
	// is unguessable, preventing CSRF attacks, as described in
	//
	// https://auth0.com/docs/protocols/oauth2/oauth-state#keep-reading
	nonceBytes := make([]byte, 64)
	_, err := io.ReadFull(rand.Reader, nonceBytes)
	if err != nil {
		panic("gothic: source of randomness unavailable: " + err.Error())
	}
	return base64.URLEncoding.EncodeToString(nonceBytes)
}

// GetState gets the state returned by the provider during the callback.
// This is used to prevent CSRF attacks, see
// http://tools.ietf.org/html/rfc6749#section-10.12
var GetState = func(req *http.Request) string {
	params := req.URL.Query()
	if params.Encode() == "" && req.Method == http.MethodPost {
		return req.FormValue("state")
	}
	return params.Get("state")
}
