package repositories

import (
	"auth-api/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type LoginEventoRepository struct {
	Collection *mongo.Collection
}

func NewLoginEventoRepository(client *mongo.Client) *LoginEventoRepository {
	return &LoginEventoRepository{
		Collection: client.Database("sistema").Collection("logins"),
	}
}

func (r *LoginEventoRepository) Save(event models.LoginEvento) error {
	_, err := r.Collection.InsertOne(context.TODO(), event)
	return err
}
