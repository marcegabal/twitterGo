package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/marcegabal/twitterGo/bd"
	"github.com/marcegabal/twitterGo/models"
)

func Registro(ctx context.Context) models.RespApi {
	var t models.Usuario
	var r models.RespApi

	r.Status = 400

	fmt.Println("Entro a registro")
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Debe especificar mail"
		fmt.Println(r.Message)
		return r
	}

	if len(t.Password) < 6 {
		r.Message = "La password debe tener minimo 6 caracteres"
		fmt.Println(r.Message)
		return r
	}

	_, encontrado, _ := bd.ChequeYaExisteUsuario(t.Email)
	if encontrado {
		r.Message = "Ya existe un usuario con ese email"
		fmt.Println(r.Message)
		return r
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		r.Message = "Error en registro de usuario " + err.Error()
		fmt.Println(r.Message)
		return r
	}

	if !status {
		r.Message = "No se logro el registro de usuario "
		fmt.Println(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Registro Ok"
	fmt.Println(r.Message)
	return r
}
