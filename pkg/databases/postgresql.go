package databases

import (
	"ddd-go/configs"
	"ddd-go/pkg/utils"
	"log"

	"github.com/jmoiron/sqlx"
)

func NewPostgreSQLDBConnection(cfg *configs.Configs) (*sqlx.DB, error) {
	postgresURL, err := utils.ConnectionUrlBuilder("postgresql", cfg)
	if err != nil {
		return nil, err
	}

	db, err := sqlx.Connect("pgx", postgresURL)
	if err != nil {
		defer db.Close()
		log.Printf("error, can't connect to database, %s", err.Error())
		return nil, err
	}

	log.Println("postgreSQL database has been connected")
	return db, nil
}
