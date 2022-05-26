package services

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// GenerateToken generates only an access token
func GenerateToken(signMethod *jwt.SigningMethodHMAC, claims jwt.MapClaims, secret *string) (*string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(signMethod, claims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(*secret))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func GenerateClaims(email string) (jwt.MapClaims, jwt.MapClaims) {
	log.Println("generate  claim function", email)
	accessClaims := jwt.MapClaims{
		"user_email": email,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}
	refreshClaims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"sub": 1,
	}

	return accessClaims, refreshClaims
}
