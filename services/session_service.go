package services

import (
	"auth-api/models"
	"auth-api/repositories"
	"auth-api/utils"
	"fmt"
	"time"
)

type SessionService struct {
	UsuarioRepo          *repositories.UsuarioRepository
	AplicacaoExternaRepo *repositories.AplicacaoExternaRepository
	LoginEvento          *repositories.LoginEventoRepository
	JWTSecret            string
}

func NewSessionService(
	usuarioRepo *repositories.UsuarioRepository,
	aplicacaoExternaRepo *repositories.AplicacaoExternaRepository,
	loginEvento *repositories.LoginEventoRepository,
	JWTSecret string,
) *SessionService {
	return &SessionService{
		UsuarioRepo:          usuarioRepo,
		AplicacaoExternaRepo: aplicacaoExternaRepo,
		LoginEvento:          loginEvento,
		JWTSecret:            JWTSecret,
	}
}

func (s *SessionService) Login(credenciais models.Credenciais) (string, error) {
	var token string

	event := models.LoginEvento{
		Usuario:   credenciais.Usuario,
		TimeStamp: time.Now(),
	}

	if user, err := s.UsuarioRepo.FindByUsuario(credenciais.Usuario); err == nil && user.Senha == credenciais.Senha {
		event.Status = "sucesso"
		event.Usuario = user.Nome
		s.LoginEvento.Save(event)

		token, err = utils.GerarToken(user.Nome, s.JWTSecret)
		if err != nil {
			return "", fmt.Errorf("erro ao gerar token JWT: %v", err)
		}

		fmt.Printf("Login bem-sucedido para usuário %s\n", user.Nome)
		return token, nil
	}

	event.Status = "falha"
	s.LoginEvento.Save(event)
	fmt.Println("Falha no login")
	return "", fmt.Errorf("credenciais inválidas")
}
