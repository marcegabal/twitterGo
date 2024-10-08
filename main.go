package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/marcegabal/twitterGo/awsgo"
	"github.com/marcegabal/twitterGo/bd"
	"github.com/marcegabal/twitterGo/handlers"
	"github.com/marcegabal/twitterGo/models"
	"github.com/marcegabal/twitterGo/secretmanager"
)

func main() {
	fmt.Println("Empieza programa")
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse
	fmt.Println("Ejecuto el lambda")
	awsgo.InicializoAws()
	fmt.Println("Termino InicializoAws")
	if !ValidoParametros() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en variables de entorno",
			Headers: map[string]string{
				"Content-type": "application/json",
			},
		}

		return res, nil

	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en lectura de secret " + err.Error(),
			Headers: map[string]string{
				"Content-type": "application/json",
			},
		}

		return res, nil
	}
	fmt.Println("Body>>>" + request.Body)
	fmt.Println(">>>" + request.HTTPMethod)
	fmt.Println(">>>" + request.Path)
	fmt.Println(">>>" + os.Getenv("UrlPrefix"))

	path := strings.Replace(request.PathParameters["twittergo"], os.Getenv("UrlPrefix"), "", -1)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))
	fmt.Println("Termino context y params")
	//Chequeo conexion a bd
	err = bd.ConectarBd(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error en conexion a bd " + err.Error(),
			Headers: map[string]string{
				"Content-type": "application/json",
			},
		}

		return res, nil
	}
	fmt.Println("Termino conectar bd")
	//parte final handlers
	respApi := handlers.Manejadores(awsgo.Ctx, request)
	if respApi.CustomResp == nil {
		fmt.Println("aaa ---> " + respApi.Message)
		res = &events.APIGatewayProxyResponse{
			StatusCode: respApi.Status,
			Body:       respApi.Message,
			Headers: map[string]string{
				"Content-type": "application/json",
			},
		}

		return res, nil
	} else {
		fmt.Println("bbb")
		return respApi.CustomResp, nil
	}
}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("BucketName")
	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}

	return true
}
