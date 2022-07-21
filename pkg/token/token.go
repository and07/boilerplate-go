package token

import "github.com/golang-jwt/jwt"

// RefreshTokenCustomClaims specifies the claims for refresh token
type RefreshTokenCustomClaims struct {
	UserID    string `json:"user_id,omitempty"`
	CustomKey string `json:"custom_key,omitempty"`
	KeyType   string `json:"key_type,omitempty"`
	jwt.StandardClaims
}

// AccessTokenCustomClaims specifies the claims for access token
type AccessTokenCustomClaims struct {
	UserID  string   `json:"user_id,omitempty"`
	KeyType string   `json:"key_type,omitempty"`
	Roles   []string `json:"roles,omitempty"`
	jwt.StandardClaims
}
