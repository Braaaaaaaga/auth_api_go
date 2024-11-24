package controllers

import (
	"auth-api/models"
	"auth-api/repositories"
	"auth-api/services"

	"go.mongodb.org/mongo-driver/mongo"
)

type SessionController struct {
	SessionService *services.SessionService
}

func NewSessionController(client *mongo.Client, jwtSecret string) *SessionController {
	usuarioRepo := repositories.NewUsuarioRepository(client)
	aplicacaoExternaRepo := repositories.NewAplicacaoExternaRepository(client)
	loginEvento := repositories.NewLoginEventoRepository(client)

	sessionService := services.NewSessionService(usuarioRepo, aplicacaoExternaRepo, loginEvento, jwtSecret)

	return &SessionController{SessionService: sessionService}
}

func (c *SessionController) Login(credenciais models.Credenciais) (string, error) {
	return c.SessionService.Login(credenciais)
}
