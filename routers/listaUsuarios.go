package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/marcegabal/twitterGo/bd"
	"github.com/marcegabal/twitterGo/models"
)

func ListaUsuarios(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	page := request.QueryStringParameters["page"]
	typeUser := request.QueryStringParameters["type"]
	search := request.QueryStringParameters["search"]
	IDUsuario := claim.ID.Hex()

	if len(page) == 0 {
		page = "1"
	}

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		r.Message = "Debe enviar page como entero mayor a 0 " + err.Error()
		return r
	}

	usuarios, status := bd.LeoUsuariosTodos(IDUsuario, int64(pagTemp), search, typeUser)
	if !status {
		r.Message = "Error leyendo usuarios"
		return r
	}

	respJson, err := json.Marshal(usuarios)
	if err != nil {
		r.Status = 500
		r.Message = "Error formateando usuarios a json"
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
