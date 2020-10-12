package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB
var err error

//ConnectDataBase connects to database
func ConnectDataBase() {

	var appConfig map[string]string
	appConfig, err := godotenv.Read()

	if err != nil {
		log.Fatal("Error reading .env file")
	}

	// Ex: user:password@tcp(host:port)/dbname
	mysqlCredentials := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s",
		appConfig["MYSQL_USER"],
		appConfig["MYSQL_PASSWORD"],
		appConfig["MYSQL_HOST"],
		appConfig["MYSQL_PORT"],
		appConfig["MYSQL_DBNAME"],
	)
	fmt.Println(mysqlCredentials)
	DB, err = gorm.Open("mysql", mysqlCredentials)
	//defer DB.Close()
	//DB, err = gorm.Open("mysql", "reetu:Reetu8995@@(localhost:3306)/test")

	if err != nil {
		fmt.Printf("failed to connect to database!")
		os.Exit(1)
	}

	DB.AutoMigrate(&Pingdom{})
}
