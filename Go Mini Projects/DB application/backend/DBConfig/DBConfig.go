package DBConfig

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Postgres connect configuration
func connect() *gorm.DB {
	// Create a dsn string to talk to postgres. 
	// This dns is the key that opens the door of postgres db world. 
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	// Turn the dns key and open the postgres door. 
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config)
	// if the door doesn't open, you are locked out. It's time to freaking panic :D !
	if err != nil {
		log.Print("Ahhhhhrrrggg!!! The door doesn't open! You are locked out of your own DB! :D ")
		panic(err)
	}

	// if the door opens:
	log.Print("The door opened. You are welcome to your own DB!")

	// Create tables in the Postgres DB system.
	db.AutoMigrate(
		&UserModel.User{},
		&UserModel.Band{},
	)

	return db
}