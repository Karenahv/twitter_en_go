package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Karenahv/twitter_en_go/bd"
	"github.com/Karenahv/twitter_en_go/jwt"
	"github.com/Karenahv/twitter_en_go/models"
)

/*Login realiza el login del usuario si existe en la bd*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña inválidos "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}
	documento, existe := bd.ItentoLogin(t.Email, t.Password)
	if existe == false {
		http.Error(w, "Usuario y/o Contraseña inválidos "+err.Error(), 400)
		return
	}
	jwtkey, err := jwt.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar generar el Token correspondiente "+err.Error(), 400)
		return
	}
	resp := models.RespuestaLogin{
		Token: jwtkey,
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtkey,
		Expires: expirationTime,
	})
}
