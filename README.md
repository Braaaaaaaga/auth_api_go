
# 🚀 Sistema de Autenticação com Go, MongoDB, JWT e Gin

![Go](https://img.shields.io/badge/Go-1.20-blue?logo=go)
![MongoDB](https://img.shields.io/badge/MongoDB-v6.0-green?logo=mongodb)
![Gin](https://img.shields.io/badge/Gin-1.9.0-lightgrey?logo=go)
![JWT](https://img.shields.io/badge/JWT-HS256-orange?logo=jsonwebtokens)
![Viper](https://img.shields.io/badge/Viper-1.12-blueviolet)

## 📚 Sobre o Projeto

Este é um sistema de autenticação que utiliza **Golang**, **MongoDB**, **JWT**, e o framework **Gin** para criar um backend rápido e escalável. O projeto permite:

- Login de usuários e aplicações externas.
- Geração e validação de tokens JWT.
- Armazenamento de logs de login no MongoDB.
- Configuração dinâmica com variáveis de ambiente usando **Viper**.

---

## 🛠️ Tecnologias Usadas

- **[Go](https://go.dev/):** Linguagem de programação moderna e performática.
- **[Gin](https://gin-gonic.com/):** Framework para criação de APIs HTTP.
- **[MongoDB](https://www.mongodb.com/):** Banco de dados NoSQL para armazenamento.
- **[JWT](https://jwt.io/):** Tokens para autenticação segura.
- **[Viper](https://github.com/spf13/viper):** Gerenciador de variáveis de ambiente e configurações.

---

## 🌟 Funcionalidades

- **Autenticação de usuários e aplicações externas:**
  - Login por credenciais (usuário e senha).
  - Login por chave API (aplicações externas).
- **Geração de tokens JWT:**
  - Tokens válidos por 24 horas.
  - Validação para acessar recursos protegidos.
- **Logs de login no MongoDB:**
  - Armazena quem logou, quando logou e o status (sucesso ou falha).
- **Endpoints para gerenciamento:**
  - Adicionar usuários e aplicações externas.

---

## ⚙️ Configuração

- **Pré-requisitos:**
  - Golang: Versão 1.20 ou superior.
  - MongoDB: Versão 6.0 ou superior.
  - Postman: Para testar os endpoints (opcional).

### Passo a Passo

1. **Clone o repositório:**
   ```bash
   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio
   ```

2. **Instale as dependências:**
   ```bash
   go mod tidy
   ```

3. **Configure o arquivo `.env`:**
   Crie um arquivo `.env` na raiz do projeto com os valores abaixo:
   ```plaintext
   MONGO_URI=mongodb://localhost:27017
   MONGO_DATABASE=sistema
   JWT_SECRET=sua-chave-secreta
   ```

4. **Inicie o MongoDB:**
   Certifique-se de que o MongoDB está rodando localmente:
   ```bash
   mongod
   ```

5. **Execute o servidor:**
   ```bash
   go run main.go
   ```

6. **URL base:**  
   Acesse o servidor em `http://localhost:8080`.

---

## 🌐 Endpoints

### 1️⃣ **Login**

- **Rota:** `/login`
- **Método:** `POST`
- **Descrição:** Faz login e retorna um token JWT.

#### **Body de Exemplo:**

**Usuário:**
```json
{
  "usuario": "admin",
  "senha": "senha123"
}
```

**Aplicação Externa:**
```json
{
  "chaveApi": "minha-chave-api"
}
```

#### **Resposta de Sucesso:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

#### **Resposta de Erro:**
```json
{
  "error": "Credenciais inválidas"
}
```

---

### 2️⃣ **Adicionar Usuário**

- **Rota:** `/usuarios`
- **Método:** `POST`
- **Descrição:** Adiciona um novo usuário.

#### **Body de Exemplo:**
```json
{
  "nome": "Admin",
  "usuario": "admin",
  "senha": "senha123",
  "nivel_acesso": "admin"
}
```

#### **Resposta de Sucesso:**
```json
{
  "message": "Usuário criado com sucesso"
}
```

#### **Resposta de Erro:**
```json
{
  "error": "Erro ao inserir usuário"
}
```

---

### 3️⃣ **Adicionar Aplicação Externa**

- **Rota:** `/aplicacoes`
- **Método:** `POST`
- **Descrição:** Adiciona uma nova aplicação externa.

#### **Body de Exemplo:**
```json
{
  "nome": "Minha Aplicação",
  "chaveApi": "minha-chave-api"
}
```

#### **Resposta de Sucesso:**
```json
{
  "message": "Aplicação externa criada com sucesso"
}
```

#### **Resposta de Erro:**
```json
{
  "error": "Erro ao inserir aplicação externa"
}
```

---

## 🛡️ Segurança

- **Mantenha o arquivo `.env` fora do controle de versão.**
- **Use uma chave secreta forte para o JWT (`JWT_SECRET`).**
- **Implemente HTTPS em produção.**

---

## 🧪 Testes

Use o **Postman** ou outra ferramenta para testar os endpoints. Certifique-se de enviar os dados no formato JSON e na URL correta.

---

## 📋 To-Do List

- [x] Autenticação com JWT.
- [x] Configuração com Viper.
- [x] Endpoints para login.
- [x] Registro de logs de login.
- [x] Adicionar suporte ao Docker.
- [ ] Implementar endpoint para listar logs de login.


---

## 👩‍💻 Contribuição

Sinta-se à vontade para abrir issues ou pull requests para sugerir melhorias. 🚀

---

## 📄 Licença

Este projeto está licenciado sob a **MIT License**. Consulte o arquivo `LICENSE` para mais detalhes.
