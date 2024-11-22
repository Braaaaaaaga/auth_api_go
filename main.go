package main

import (
	"auth-api/config"
	"auth-api/controllers"
	"auth-api/models"
	"fmt"
)

func main() {
	client := config.ConnectMongo()
	defer client.Disconnect(nil)

	sessionController := controllers.NewSessionController(client)

	// Teste de login como usuário
	fmt.Println("Tentando login como usuário...")
	sessionController.Login(models.Credenciais{
		Usuario: "admin",
		Senha:   "senha123",
	})

	// Teste de login como aplicação externa
	fmt.Println("Tentando login como aplicação externa...")
	sessionController.Login(models.Credenciais{
		ChaveApi: "minha-chave-api",
	})
}
