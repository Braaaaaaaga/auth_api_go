package repositories

import (
	"auth-api/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AplicacaoExternaRepository struct {
	Collection *mongo.Collection
}

func NewAplicacaoExternaRepository(client *mongo.Client) *AplicacaoExternaRepository {
	return &AplicacaoExternaRepository{
		Collection: client.Database("sistema").Collection("aplicacoes_externas"),
	}
}

func (r *AplicacaoExternaRepository) FindByChaveApi(chaveApi string) (*models.AplicacaoExterna, error) {
	filter := bson.M{"chave_api": chaveApi}
	var app models.AplicacaoExterna
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&app)
	return &app, err
}
