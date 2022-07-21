package httpserver

import (
	"github.com/and07/boilerplate-go/pkg/data"
	"github.com/and07/boilerplate-go/pkg/service"
	"github.com/and07/boilerplate-go/pkg/service/mail"
	"github.com/and07/boilerplate-go/pkg/token"
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

// AuthHandler wraps instances needed to perform operations on user object
type AuthHandler struct {
	logger      hclog.Logger
	configs     *utils.Configurations
	validator   *data.Validation
	repo        data.AuthRepository
	authService service.Authentication
	jwtManager  *token.JWTManager
	mailService mail.Service
	oauthConfGl *oauth2.Config
}

// Option ...
type Option func(*AuthHandler)

// WithGoogleAuth ...
func WithGoogleAuth(clientKey string, secret string, callbackURL string, scopes ...string) Option {
	return func(a *AuthHandler) {
		a.oauthConfGl = newConfig(clientKey, secret, callbackURL, scopes...)
	}
}

// NewAuthHandler returns a new UserHandler instance
func NewAuthHandler(l hclog.Logger, c *utils.Configurations, v *data.Validation, r data.AuthRepository, auth service.Authentication, jwtManager *token.JWTManager, mail mail.Service, opts ...Option) *AuthHandler {
	a := &AuthHandler{
		logger:      l,
		configs:     c,
		validator:   v,
		repo:        r,
		authService: auth,
		mailService: mail,
		jwtManager:  jwtManager,
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

// TokenResponse below data types are used for encoding and decoding b/t go types and json
type TokenResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

// AuthResponse ...
type AuthResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	Username     string `json:"username"`
}

// UsernameUpdate ...
type UsernameUpdate struct {
	Username string `json:"username"`
}

// CodeVerificationReq ...
type CodeVerificationReq struct {
	Code string `json:"code"`
	Type string `json:"type"`
}

// PasswordResetReq ...
type PasswordResetReq struct {
	Password   string `json:"password"`
	PasswordRe string `json:"password_re"`
	Code       string `json:"code"`
}

// ErrUserAlreadyExists ...
var ErrUserAlreadyExists = "User already exists with the given email"

// ErrUserNotFound ...
var ErrUserNotFound = "No user account exists with given email. Please sign in first"

// UserCreationFailed ...
var UserCreationFailed = "Unable to create user.Please try again later"

// PgDuplicateKeyMsg ...
var PgDuplicateKeyMsg = "duplicate key value violates unique constraint"

// PgNoRowsMsg ...
var PgNoRowsMsg = "no rows in result set"
