package service

import (
	"github.com/and07/boilerplate-go/pkg/data"
	"github.com/and07/boilerplate-go/pkg/utils"
	"github.com/hashicorp/go-hclog"
	"golang.org/x/crypto/bcrypt"
)

// Authentication interface lists the methods that our authentication service should implement
type Authentication interface {
	Authenticate(reqUser *data.User, user *data.User) bool
}

// AuthService is the implementation of our Authentication
type AuthService struct {
	logger  hclog.Logger
	configs *utils.Configurations
}

// NewAuthService returns a new instance of the auth service
func NewAuthService(logger hclog.Logger, configs *utils.Configurations) *AuthService {
	return &AuthService{logger, configs}
}

// Authenticate checks the user credentials in request against the db and authenticates the request
func (auth *AuthService) Authenticate(reqUser *data.User, user *data.User) bool {

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password)); err != nil {
		auth.logger.Debug("password hashes are not same")
		return false
	}
	return true
}
