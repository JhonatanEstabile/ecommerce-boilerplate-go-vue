package main

import (
	"ecommerce/pkg/database"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.GetConnection()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalln("Erro ao pegar driver", err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://pkg/database/migrations",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatalln("Erro ao criar migracao", err.Error())
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalln("Erro ao rodar migracao", err.Error())
	}

	log.Println("✅ Migrations aplicadas com sucesso!")
}
