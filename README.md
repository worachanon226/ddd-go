# ddd-go

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)

## Description

This is an example of implementation of Clean Architecture in Go (Golang) project referenced from [medium@rayato](https://medium.com/@rayato159/มาเขียน-rest-api-โดยใช้-clean-architecture-ใน-golang-กันเถอะ-b47ce99c3297), [medium@manakuro](https://manakuro.medium.com/clean-architecture-with-go-bce409427d31).

## Advantages of Clean Architecture

![The Clean Architecture Image](https://miro.medium.com/v2/resize:fit:1400/format:webp/1*B7LkQDyDqLN3rRSrNYkETA.jpeg)

## The Layers

![Layers Image](https://miro.medium.com/v2/resize:fit:1400/format:webp/1*yASGoy2V6F560OoSU470Tg.png)

## Project Structure

```bash
.
├── README.md
├── assets
│ └── logs
├── configs
│ └── config.go
├── go.mod
├── go.sum
├── main.go
├── modules
│ ├── entities
│ │ ├── response.go
│ │ └── users.go
│ ├── servers
│ │ ├── handler.go
│ │ └── server.go
│ └── users
│ ├── controllers
│ │ └── users_controllers.go
│ ├── repositories
│ │ └── users_repositories.go
│ └── usecases
│ └── users_usecases.go
└── pkg
├── databases
│ ├── migrations
│ └── postgresql.go
├── middlewares
└── utils
└── connection_url_builder.go

```
