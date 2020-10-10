package authorization

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/meloalright/guora/conf"
)

// Parse Service
func Parse(tokenString string) (ID int, ProfileID int, err error) {

	SecretString := conf.Config().Secretkey
	Secret := []byte(SecretString)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		ID = claims.ID
		ProfileID = claims.ProfileID
	} else {
		err = errors.New("Token Not Valid")
	}
	return

}
