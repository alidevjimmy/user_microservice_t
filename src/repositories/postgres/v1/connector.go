package repositories

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func postgresConnector() *gorm.DB {
	dialector := postgres.New(postgres.Config{
		DSN: "postgres_db",
		DriverName: "postgres",
		PreferSimpleProtocol: true,
	})
	db , err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}