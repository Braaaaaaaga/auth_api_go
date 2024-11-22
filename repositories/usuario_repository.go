package repositories

import (
	"auth-api/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UsuarioRepository struct {
	Collection *mongo.Collection
}

func NewUsuarioRepository(client *mongo.Client) *UsuarioRepository {
	return &UsuarioRepository{
		Collection: client.Database("sistema").Collection("usuarios"),
	}
}

func (r *UsuarioRepository) FindByUsuario(usuario string) (*models.Usuario, error) {
	filter := bson.M{"usuario": usuario}
	var user models.Usuario
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&user)
	return &user, err
}
