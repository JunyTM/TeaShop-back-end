package infrastructure

import (
	"log"
	"teashop/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Openconnect Open session using db
func openConnection() (*gorm.DB, error) {
	// connectSQL := "host=" + dbHost +
	// 	" user=" + dbUser +
	// 	" dbname=" + dbName +
	// 	" password=" + dbPassword +
	// 	" sslmode=disable"
	connectSQL := dbURL

	db, err := gorm.Open(postgres.Open(connectSQL), &gorm.Config{})

	if err != nil {
		log.Printf("Has problem at connection database: %+v\n", err)
		return nil, err
	}
	return db, nil
}

func InitDatabase(allowMigrate bool) error {
	var err error
	db, err = openConnection()
	if err != nil {
		log.Printf("Has problem at creat database")
	}
	if allowMigrate {
		log.Println("Migrating database...")
		db.AutoMigrate(
			&model.User{},
		)
	}
	return nil
}
