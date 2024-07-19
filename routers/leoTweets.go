package routers

import (
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/marcegabal/twitterGo/bd"
	"github.com/marcegabal/twitterGo/models"
)

func LeoTweets(request events.APIGatewayProxyRequest) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	pagina := request.QueryStringParameters["pagina"]

	if len(ID) < 1 {
		r.Message = "El par ID es obligatorio"
		return r
	}

	if len(pagina) < 1 {
		pagina = "1"
	}

	pag, err := strconv.Atoi(pagina)
	if err != nil {
		r.Message = "Debe enviar valor numerico " + err.Error()
		return r
	}

	tweets, correcto := bd.LeoTweets(ID, int64(pag))
	if !correcto {
		r.Message = "Error leyendo tweets"
		return r
	}

	respJson, err := json.Marshal(tweets)
	if err != nil {
		r.Status = 500
		r.Message = "Error formateando tweets como json"
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}
