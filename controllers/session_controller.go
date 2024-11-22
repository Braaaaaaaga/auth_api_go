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

func NewSessionController(client *mongo.Client) *SessionController {
	usuarioRepo := repositories.NewUsuarioRepository(client)
	aplicacaoExternaRepo := repositories.NewAplicacaoExternaRepository(client)
	loginEvento := repositories.NewLoginEventoRepository(client)

	sessionService := services.NewSessionService(usuarioRepo, aplicacaoExternaRepo, loginEvento)

	return &SessionController{SessionService: sessionService}
}

func (c *SessionController) Login(credenciais models.Credenciais) {
	c.SessionService.Login(credenciais)
}
