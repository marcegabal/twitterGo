package routers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/marcegabal/twitterGo/bd"
	"github.com/marcegabal/twitterGo/models"
)

func EliminarTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.RespApi {
	var r models.RespApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El iID es obligatorio"
		return r
	}

	err := bd.BorroTweet(ID, claim.ID.Hex())
	if err != nil {
		r.Message = "Error al borrar tweet " + err.Error()
		return r
	}

	r.Message = "Tweet borrado"
	r.Status = 200
	return r
}
