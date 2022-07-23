package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// secret key being used to sign tokens
var (
	secretKey = []byte("secret")
)

func GetSecretKey() []byte {
	return secretKey
}

// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password" mapstructure:"password"`
	Email    string `json:"email" mapstructure:"email"`
}

// GenerateToken generates a jwt token and assign a email to it's claims and return it
func GenerateToken(email, password string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	/* Create a map to store our claims */
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims */
	claims["email"] = email
	claims["password"] = password
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(GetSecretKey())
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken parses a jwt token and returns the email in it's claims
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return GetSecretKey(), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email := claims["email"].(string)
		return email, nil
	} else {
		return "", err
	}
}

func RefreshToken(tokenStr string, expirationTime time.Time) (string, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return GetSecretKey(), nil
	})

	if err != nil {
		return "", err
	}

	if !tkn.Valid {
		return "", fmt.Errorf("invalid token")
	}

	// Now, create a new token for the current use, with a renewed expiration time
	// expirationTime := time.Now().Add(refreshTime * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(GetSecretKey())
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
