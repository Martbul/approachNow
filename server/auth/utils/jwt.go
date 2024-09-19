// ! FLAGGED FOR DELETION
package utils

import (
	"context"
	"errors"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/martbul/auth/data"

	"github.com/dgrijalva/jwt-go"
)

// TODO: Put secret key in .env
var jwtKey = []byte("your_jwt_secret_key")
var logger = hclog.Default()

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expires in 24 hours
	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

// GetEmailFromJWT extracts the email from a valid JWT token.
func GetUserIdFromJWT(tokenString string) (int, error) {
	claims, err := ValidateJWT(tokenString)
	if err != nil {
		return -1, err
	}

	row, err := data.Query(context.Background(), `
  		   SELECT user_id
			FROM user_locations
			WHERE ST_DWithin(
    		ST_SetSRID(ST_MakePoint($1), 4326),
    		$2
			)
			`, claims.Email, 5000)

	if err != nil {
		logger.Error("Unable to get DB row", "error", err)

	}

	var userID int
	err = row.Scan(&userID)
	if err != nil {
		logger.Error("Unable to scan a DB row", "error", err)
	}

	return userID, nil
}
