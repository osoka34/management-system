package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	UID      string `json:"uid"`
	jwt.RegisteredClaims
}

var (
	jwtSecret        = []byte("your_secret_key")
	jwtRefreshSecret = []byte("your_refresh_key")
)


type KeyAccessClaims struct{}
type KeyRefreshClaims struct{}


type AccessClaims struct {
	Username string `json:"username"`
	UID      string `json:"uid"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	Username string `json:"username"`
	UID      string `json:"uid"`
	jwt.RegisteredClaims
}

const (
	EmptyArgs Error = "empty args"
)

func GenerateTokens(username, uid string) (accessToken, refreshToken string, err error) {
	if username == "" || uid == "" {
		return "", "", EmptyArgs
	}

	accessClaims := AccessClaims{
		Username: username,
		UID:      uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).
		SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	refreshClaims := RefreshClaims{
		Username: username,
		UID:      uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(1 * 24 * time.Hour),
			),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).
		SignedString(jwtRefreshSecret)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func ValidateAccessToken(tokenStr string) (*AccessClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&AccessClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*AccessClaims)
	if ok && token.Valid {
		return claims, nil
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, jwt.ErrTokenExpired
	}

	return nil, err
}

func ValidateRefreshToken(tokenStr string) (*RefreshClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&RefreshClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtRefreshSecret, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*RefreshClaims)
	if ok && token.Valid {
		return claims, nil
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, jwt.ErrTokenExpired
	}

	return nil, err
}

func ExtractRefreshClaims(token string) (*RefreshClaims, error) {
	claims := &RefreshClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtRefreshSecret, nil
	})
	return claims, err
}

func ExtractAccessClaims(token string) (*AccessClaims, error) {
	claims := &AccessClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	return claims, err
}
