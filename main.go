package main

import (
	"ddd-go/configs"
	"ddd-go/modules/servers"
	"ddd-go/pkg/databases"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err.Error())
	}
	cfg := new(configs.Configs)

	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("DB_PORT")
	cfg.PostgreSQL.Protocol = os.Getenv("DB_PROTOCOL")
	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")

	db, err := databases.NewPostgreSQLDBConnection(cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer db.Close()

	s := servers.NewServer(cfg, db)
	s.Start()
}
