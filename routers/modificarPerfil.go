package routers

import (
	"context"
	"encoding/json"

	"github.com/marcegabal/twitterGo/bd"
	"github.com/marcegabal/twitterGo/models"
)

func ModificarPerfil(ctx context.Context, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	var t models.Usuario
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		r.Message = "Datos incorrectos " + err.Error()
		return r
	}

	status, err := bd.ModificoRegistro(t, claim.ID.Hex())
	if err != nil {
		r.Message = "Error al modificar registro " + err.Error()
		return r
	}
	if !status {
		r.Message = "No se modifico registro"
		return r
	}

	r.Status = 200
	r.Message = "Se modifico sin problemas"
	return r

}
