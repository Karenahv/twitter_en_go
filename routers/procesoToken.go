package routers

import (
	"errors"
	"strings"

	"github.com/Karenahv/twitter_en_go/bd"
	"github.com/Karenahv/twitter_en_go/models"
	"github.com/dgrijalva/jwt-go"
)

var Email string
var IDUsuario string

/*ProcesoToken proceso token para extraer sus valores */
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MasterdelDesarrollo_grupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")

	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}
	return claims, false, string(""), err
}