package main

import (
	"auth-api/config"
	"auth-api/controllers"
	"auth-api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	cfg := config.LoadConfig()

	client := config.ConnectMongo(cfg)
	defer client.Disconnect(nil)

	sessionController := controllers.NewSessionController(client, cfg.JWTSecret)

	router := gin.Default()

	router.POST(
		"/login",
		func(c *gin.Context) {
			var credenciais models.Credenciais
			if err := c.ShouldBindJSON(&credenciais); err != nil {
				c.JSON(400, gin.H{"error": err.Error()})
				return
			}

			token, err := sessionController.Login(credenciais)
			if err != nil {
				c.JSON(401, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"token": token})
		},
	)

	router.POST("/usuarios", func(c *gin.Context) {
		var usuario models.Usuario
		if err := c.ShouldBindJSON(&usuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
			return
		}

		collection := client.Database(cfg.MongoDatabase).Collection("usuarios")
		_, err := collection.InsertOne(c, usuario)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir usuário"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso"})
	})

	router.POST("/aplicacoes", func(c *gin.Context) {
		var aplicacao models.AplicacaoExterna
		if err := c.ShouldBindJSON(&aplicacao); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
			return
		}

		collection := client.Database(cfg.MongoDatabase).Collection("aplicacoes_externas")
		_, err := collection.InsertOne(c, aplicacao)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir aplicação externa"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Aplicação externa criada com sucesso"})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Servidor rodando!",
		})
	})

	fmt.Println("Iniciando o servidor na porta 8080...")
	err := router.Run(":8080")
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}
