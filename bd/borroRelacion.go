package bd

import (
	"context"

	"github.com/marcegabal/twitterGo/models"
)

func BorroRelacion(t models.Relacion) (bool, error) {
	ctx := context.TODO()
	db := MongoCn.Database(DatabaseName)
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil

}
