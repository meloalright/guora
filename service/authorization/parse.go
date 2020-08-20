package authorization

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/meloalright/guora/conf"
)

func Parse(tokenString string) (ID int, ProfileID int, err error) {

	SecretString := conf.Config().Secretkey

	Secret := []byte(SecretString)

	token, err := jwt.ParseWithClaims(tokenString, &AuthorizationClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})

	if claims, ok := token.Claims.(*AuthorizationClaims); ok && token.Valid {
		ID = claims.ID
		ProfileID = claims.ProfileID
		return
	} else {
		err = errors.New("token not valid")
		return
	}

}
