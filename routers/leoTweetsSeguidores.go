package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/marcegabal/twitterGo/bd"
	"github.com/marcegabal/twitterGo/models"
)

func LeoTweetsSeguidores(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400
	IDUsuario := claim.ID.Hex()

	pagina := request.QueryStringParameters["pagina"]
	if len(pagina) < 1 {
		pagina = "1"
	}

	pag, err := strconv.Atoi(pagina)
	if err != nil {
		r.Message = "Debe enviar valor numerico " + err.Error()
		return r
	}

	tweets, correcto := bd.LeoTweetsSeguidores(IDUsuario, pag)
	if !correcto {
		r.Message = "Error al leer los tweets"
		return r
	}

	respJson, err := json.Marshal(tweets)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear tweets de seguidores"
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
