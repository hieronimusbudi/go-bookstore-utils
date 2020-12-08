package myjwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	resterrors "github.com/hieronimusbudi/go-bookstore-utils/rest_errors"
)

type UserPayload struct {
	Id        int64
	Email     string
	FirstName string
	LastName  string
	Status    string
}

func GenerateToken(user *UserPayload, secret string) (string, resterrors.RestErr) {
	atClaims := jwt.MapClaims{}
	atClaims["id"] = user.Id
	atClaims["email"] = user.Email
	atClaims["first_name"] = user.FirstName
	atClaims["last_name"] = user.LastName
	atClaims["status"] = user.Status
	atClaims["exp"] = time.Now().Add(time.Minute * 180).Unix()

	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := sign.SignedString([]byte(secret))

	if err != nil {
		return "", resterrors.NewInternalServerError("error when trying to create token", err)
	}
	return token, nil
}

func VerifyToken(tokenString string, secret string) (*jwt.Token, resterrors.RestErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != t.Method {
			return nil, resterrors.NewBadRequestError("Unexpected signing method")
		}

		return []byte(secret), nil
	})
	if err != nil {
		return nil, resterrors.NewInternalServerError("error when verify token", err)
	}
	return token, nil
}

func ValidateToken(tokenString string, secret string) (jwt.MapClaims, resterrors.RestErr) {
	token, err := VerifyToken(tokenString, secret)
	if err != nil {
		return nil, resterrors.NewUnauthorizedError(err.Message())
	}

	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		return nil, resterrors.NewUnauthorizedError("Unauthorized")
	}

	return tokenClaims, nil
}
