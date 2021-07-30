package v1

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func init() {
	pg := postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", "user_microservice_t_pgdb_1", "postgres", "password", "users_microservice_db", "5432"),
	})
	var err error
	db, err = gorm.Open(pg, &gorm.Config{})
	if err != nil {
		panic("db cannot connect")
	}
}
