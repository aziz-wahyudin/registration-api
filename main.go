package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aziz-wahyudin/registration-api/config"
	"github.com/aziz-wahyudin/registration-api/factory"
	"github.com/aziz-wahyudin/registration-api/utils/database/mysql"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.DBMigration(db)

	factory.InitFactory(db)

	fmt.Println("starting at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
