package bd

import (
	"context"
	"fmt"

	"github.com/marcegabal/twitterGo/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCn *mongo.Client
var DatabaseName string

func ConectarBd(ctx context.Context) error {
	fmt.Println("Empiezo a conectar bd")
	user := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	connString := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)

	var clientOptions = options.Client().ApplyURI(connString)

	client, err := mongo.Connect(ctx, clientOptions)
	fmt.Println("Llego hasta aca con la bd???", user, password, host, connString)
	if err != nil {
		fmt.Println("Error 1 --->" + err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error 2 --->" + err.Error())
		return err
	}

	fmt.Println("Conexion exitosa")
	MongoCn = client
	DatabaseName = ctx.Value(models.Key("database")).(string)
	fmt.Println("database-->>> " + DatabaseName)
	return nil
}

func BaseConectada() bool {
	err := MongoCn.Ping(context.TODO(), nil)
	return err == nil
}
