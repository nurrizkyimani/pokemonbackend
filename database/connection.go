package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/nurrizkyimani/pahamifybackend/model"
)

var (
	//DBConn is xxx
	DBConn *gorm.DB
)

//InitDatabase is a function to initialize the connection with PSQL in Heroku PSQL
func InitDatabase() {
	var err error
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	dbURL := os.Getenv("db_url")

	fmt.Println(dbURL)

	DBConn, err = gorm.Open("postgres", dbURL)
	if err != nil {
		fmt.Print(err)
		panic("failed to conenct ")
	}

	fmt.Println("Connection opened database")

	DBConn.AutoMigrate(&model.Pokemon{}, &model.Poketype{})

	fmt.Println("database migrated")
}
