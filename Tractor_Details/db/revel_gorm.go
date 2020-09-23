package postgresql

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func init() {
	psqlTractorServiceUsername := os.Getenv("psql_tractor_service_username")
	psqlTractorServicePassword := os.Getenv("psql_tractor_service_password")
	psqlTractorServiceHost := os.Getenv("psql_tractor_service_host")
	psqlTractorServiceSchema := os.Getenv("psql_tractor_service_schema")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		psqlTractorServiceHost,
		psqlTractorServiceUsername,
		psqlTractorServicePassword,
		psqlTractorServiceSchema,
		"5433")

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

}
