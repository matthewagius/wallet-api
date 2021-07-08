# Wallet API

## About the Project
The goal is to write a JSON API in Golang to get the balance and manage credit/debit
operations on the players wallets. For example, you might receive calls on your API to get
the balance of the wallet with id 123, or to credit the wallet with id 456 by 10.00 ?. The
storage mechanism to use will be MySQL.

## Built With
- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [MySQL](https://github.com/go-gorm/gorm)
- [JWT](https://github.com/dgrijalva/jwt-go)
- [Numbers](https://github.com/shopspring/decimal)
- [Logger](https://github.com/sirupsen/logrus)
- [Assert](https://github.com/stretchr/testify/assert)
- [Swaggo](https://github.com/swaggo/gin-swagger)
- [Docker](https://www.docker.com)

## Prerequisites

You need to have the following installed on your machine to run this application:
- Go 
- Docker 

## Getting Started

To load the mySQL docker container:
- open git bash on root folder
- run: $docker-compose -f docker-compose.yml up

To run the API:
- run $go run main.go
- browse to: http://localhost:8080/swagger/index.html - full api documentation can be viewed here

## Browse the API using Swagger UI:
- execute api/v1/auth endpoint. Please use the following credentials 
{
  "email": "user.wallet@email.com",
  "password": "gue55myp4ss"
}
- once executed the api will return an access token. Copy this access token
- Click on the Authorize button at the top right corner
- type in Bearer followed by your copied access token and click on Authorize

You may now use swagger to browse the API. 
