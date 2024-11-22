<!--
This file contains the README documentation for the auth_api_go project.
-->
# auth_api_go

## Overview
auth_api_go is a Go-based authentication API that provides secure and efficient user authentication and authorization mechanisms.

## Features
- User registration and login
- Password hashing and verification
- JWT token generation and validation
- Middleware for protected routes
- Role-based access control

## Installation
To install the dependencies, run:
```sh
go mod tidy
```

## Usage
To start the server, run:
```sh
go run main.go
```

## API Endpoints
### Authentication
- `POST /register` - Register a new user
- `POST /login` - Authenticate a user and return a JWT token

### User
- `GET /user` - Get user details (protected route)

## Configuration
Configure the application using environment variables:
- `DB_HOST` - Database host
- `DB_USER` - Database user
- `DB_PASSWORD` - Database password
- `JWT_SECRET` - Secret key for JWT

## Contributing
Contributions are welcome! Please open an issue or submit a pull request.

## License
This project is licensed under the MIT License.