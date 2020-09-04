package main

import (
	"log"
	"net/http"
	"os"

	"github.com/sandblox-official/api/players"

	"github.com/sandblox-official/api/worlds"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "zane:5245@tcp(127.0.0.1:3306)/gamedb?charset=utf8mb4&parseTime=True&loc=Local"
var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func main() {

	//Set up and connect database
	if err != nil {
		log.Fatalln("Failed to connect database: ", err)
	}
	db.AutoMigrate(&worlds.Worlds{})
	db.AutoMigrate(&players.Players{})

	//Router
	r := mux.NewRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
