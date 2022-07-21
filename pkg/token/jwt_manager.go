package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"time"

	"github.com/and07/boilerplate-go/pkg/data"
	"github.com/and07/boilerplate-go/pkg/utils"
	"github.com/golang-jwt/jwt"
	"github.com/hashicorp/go-hclog"
)

// Authentication interface lists the methods that our authentication service should implement
type TokenManager interface {
	GenerateAccessToken(user *data.User) (string, error)
	GenerateRefreshToken(user *data.User) (string, error)
	GenerateCustomKey(userID string, password string) string
	ValidateAccessToken(token string) (string, error)
	ValidateRefreshToken(token string) (string, string, error)
}

type User struct {
	Username       string
	HashedPassword string
	Roles          []string
}

type JWTManager struct {
	logger  hclog.Logger
	configs *utils.Configurations
}

func NewJWTManager(logger hclog.Logger, configs *utils.Configurations) *JWTManager {
	return &JWTManager{logger, configs}
}

// GenerateRefreshToken generate a new refresh token for the given user
func (j *JWTManager) GenerateRefreshToken(user *data.User) (string, error) {

	cusKey := j.GenerateCustomKey(user.ID, user.TokenHash)
	tokenType := "refresh"

	claims := RefreshTokenCustomClaims{
		user.ID,
		cusKey,
		tokenType,
		jwt.StandardClaims{
			Issuer: "auth.service",
		},
	}

	signBytes, err := ioutil.ReadFile(j.configs.RefreshTokenPrivateKeyPath)
	if err != nil {
		j.logger.Error("unable to read private key", "error", err)
		return "", errors.New("could not generate refresh token. please try again later")
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		j.logger.Error("unable to parse private key", "error", err)
		return "", errors.New("could not generate refresh token. please try again later")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(signKey)
}

// GenerateAccessToken generates a new access token for the given user
func (j *JWTManager) GenerateAccessToken(user *data.User) (string, error) {

	userID := user.ID
	tokenType := "access"

	claims := AccessTokenCustomClaims{
		UserID:  userID,
		KeyType: tokenType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(j.configs.JwtExpiration)).Unix(),
			Issuer:    "auth.service",
		},
	}

	signBytes, err := ioutil.ReadFile(j.configs.AccessTokenPrivateKeyPath)
	if err != nil {
		j.logger.Error("unable to read private key", "error", err)
		return "", errors.New("could not generate access token. please try again later")
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		j.logger.Error("unable to parse private key", "error", err)
		return "", errors.New("could not generate access token. please try again later")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(signKey)
}

// GenerateCustomKey creates a new key for our jwt payload
// the key is a hashed combination of the userID and user tokenhash
func (j *JWTManager) GenerateCustomKey(userID string, tokenHash string) string {

	// data := userID + tokenHash
	h := hmac.New(sha256.New, []byte(tokenHash))
	h.Write([]byte(userID))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

// ValidateAccessToken parses and validates the given access token
// returns the userId present in the token payload
func (j *JWTManager) ValidateAccessToken(tokenString string) (*AccessTokenCustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			j.logger.Error("Unexpected signing method in auth token")
			return nil, errors.New("unexpected signing method in auth token")
		}
		verifyBytes, err := ioutil.ReadFile(j.configs.AccessTokenPublicKeyPath)
		if err != nil {
			j.logger.Error("unable to read public key", "error", err)
			return nil, err
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			j.logger.Error("unable to parse public key", "error", err)
			return nil, err
		}

		return verifyKey, nil
	})

	if err != nil {
		j.logger.Error("unable to parse claims", "error", err)
		return nil, err
	}

	claims, ok := token.Claims.(*AccessTokenCustomClaims)
	if !ok || !token.Valid || claims.UserID == "" || claims.KeyType != "access" {
		return nil, errors.New("invalid token: authentication failed")
	}
	return claims, nil
}

// ValidateRefreshToken parses and validates the given refresh token
// returns the userId and customkey present in the token payload
func (j *JWTManager) ValidateRefreshToken(tokenString string) (*RefreshTokenCustomClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			j.logger.Error("Unexpected signing method in auth token")
			return nil, errors.New("unexpected signing method in auth token")
		}
		verifyBytes, err := ioutil.ReadFile(j.configs.RefreshTokenPublicKeyPath)
		if err != nil {
			j.logger.Error("unable to read public key", "error", err)
			return nil, err
		}

		verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
		if err != nil {
			j.logger.Error("unable to parse public key", "error", err)
			return nil, err
		}

		return verifyKey, nil
	})

	if err != nil {
		j.logger.Error("unable to parse claims", "error", err)
		return nil, err
	}

	claims, ok := token.Claims.(*RefreshTokenCustomClaims)
	j.logger.Debug("ok", ok)
	if !ok || !token.Valid || claims.UserID == "" || claims.KeyType != "refresh" {
		j.logger.Debug("could not extract claims from token")
		return nil, errors.New("invalid token: authentication failed")
	}
	return claims, nil
}
