// ! FLAGGED FOR DELETION
package utils

import (
	"context"
	"errors"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v4"
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
	logger.Info(claims.Email)



	row, err := data.Query(context.Background(), `
	SELECT id
	FROM users
	WHERE email = $1
 `, claims.Email)
			

	if err != nil {
		logger.Error("Unable to get DB row", "error", err)

	}

	var id int
err = row.Scan(&id)
if err != nil {
    if err == pgx.ErrNoRows {
        // Handle the case where no user was found
        logger.Error("No user found with the provided email", "email", claims.Email)
        return 0, nil // or return an appropriate error value
    }
    logger.Error("Unable to scan a DB row", "error", err)
    return 0, err // Return the error for further handling
}


	return id, nil
}
