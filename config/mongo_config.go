package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo(cfg *Config) *mongo.Client {
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(fmt.Sprintf("Erro ao conectar com o MongoDB: %v", err))
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		panic(fmt.Sprintf("Erro ao verificar conexão com o MongoDB: %v", err))
	}

	fmt.Println(
		"Conexão com MongoDB estabelecida!",
	)

	return client
}
