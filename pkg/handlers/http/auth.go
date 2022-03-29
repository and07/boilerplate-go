package handlers

import (
	"fmt"

	"github.com/and07/boilerplate-go/pkg/data"
	"github.com/and07/boilerplate-go/pkg/service"
	"github.com/and07/boilerplate-go/pkg/service/mail"
	"github.com/and07/boilerplate-go/pkg/utils"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/oauth2"
)

// UserKey is used as a key for storing the User object in context at middleware
type UserKey struct{}

// UserIDKey is used as a key for storing the UserID in context at middleware
type UserIDKey struct{}

// VerificationDataKey is used as the key for storing the VerificationData in context at middleware
type VerificationDataKey struct{}

// UserHandler wraps instances needed to perform operations on user object
type AuthHandler struct {
	logger      hclog.Logger
	configs     *utils.Configurations
	validator   *data.Validation
	repo        data.Repository
	authService service.Authentication
	mailService mail.MailService
	oauthConfGl *oauth2.Config
}

// Option ...
type Option func(*AuthHandler)

// WithDebugPort ...
func WithGoogleAuth(clientKey string, secret string, callbackURL string, scopes ...string) Option {
	return func(a *AuthHandler) {
		a.oauthConfGl = newConfig(clientKey, secret, callbackURL, scopes...)
	}
}

// NewUserHandler returns a new UserHandler instance
func NewAuthHandler(l hclog.Logger, c *utils.Configurations, v *data.Validation, r data.Repository, auth service.Authentication, mail mail.MailService, opts ...Option) *AuthHandler {
	a := &AuthHandler{
		logger:      l,
		configs:     c,
		validator:   v,
		repo:        r,
		authService: auth,
		mailService: mail,
	}
	for _, opt := range opts {
		// Call the option giving the instantiated
		opt(a)
	}
	return a
}

// GenericResponse is the format of our response
type GenericResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Errors []string `json:"errors"`
}

// Below data types are used for encoding and decoding b/t go types and json
type TokenResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type AuthResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Username     string `json:"username"`
}

type UsernameUpdate struct {
	Username string `json:"username"`
}

type CodeVerificationReq struct {
	Code string `json: "code"`
	Type string `json" "type"`
}

type PasswordResetReq struct {
	Password   string `json: "password"`
	PasswordRe string `json: "password_re"`
	Code       string `json: "code"`
}

var ErrUserAlreadyExists = fmt.Sprintf("User already exists with the given email")
var ErrUserNotFound = fmt.Sprintf("No user account exists with given email. Please sign in first")
var UserCreationFailed = fmt.Sprintf("Unable to create user.Please try again later")

var PgDuplicateKeyMsg = "duplicate key value violates unique constraint"
var PgNoRowsMsg = "no rows in result set"
