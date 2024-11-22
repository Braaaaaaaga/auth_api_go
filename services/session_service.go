package services

import (
	"auth-api/models"
	"auth-api/repositories"
	"fmt"
	"time"
)

type SessionService struct {
	UsuarioRepo          *repositories.UsuarioRepository
	AplicacaoExternaRepo *repositories.AplicacaoExternaRepository
	LoginEvento          *repositories.LoginEventoRepository
}

func NewSessionService(
	usuarioRepo *repositories.UsuarioRepository,
	aplicacaoExternaRepo *repositories.AplicacaoExternaRepository,
	loginEvento *repositories.LoginEventoRepository,
) *SessionService {
	return &SessionService{
		UsuarioRepo:          usuarioRepo,
		AplicacaoExternaRepo: aplicacaoExternaRepo,
		LoginEvento:          loginEvento,
	}
}

func (s *SessionService) Login(credenciais models.Credenciais) {
	event := models.LoginEvento{
		Usuario:   credenciais.Usuario,
		TimeStamp: time.Now(),
	}

	if user, err := s.UsuarioRepo.FindByUsuario(credenciais.Usuario); err == nil && user.Senha == credenciais.Senha {
		event.Status = "sucesso"
		event.Usuario = user.Nome
		s.LoginEvento.Save(event)
		fmt.Printf("Login bem-sucedido para usuário %s\n", user.Nome)
		return
	} else if credenciais.Usuario != "" {
		event := models.LoginEvento{
			Usuario:   credenciais.Usuario,
			TimeStamp: time.Now(),
			Status:    "falha",
		}
		s.LoginEvento.Save(event)
		fmt.Println("Falha no login para usuário:", credenciais.Usuario)
	}

	if app, err := s.AplicacaoExternaRepo.FindByChaveApi(credenciais.ChaveApi); err == nil && app.ChaveApi == credenciais.ChaveApi {
		event.Status = "sucesso"
		event.Usuario = app.Nome
		s.LoginEvento.Save(event)
		fmt.Printf("Login bem-sucedido para aplicação externa %s\n", app.Nome)
		return
	} else if credenciais.ChaveApi != "" {
		event := models.LoginEvento{
			Usuario:   "Chave API fornecida: " + credenciais.ChaveApi,
			TimeStamp: time.Now(),
			Status:    "falha",
		}
		s.LoginEvento.Save(event)
		fmt.Println("Falha no login para chave API:", credenciais.ChaveApi)
	}

}
