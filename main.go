package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sandblox-official/database-api/players"

	"github.com/sandblox-official/database-api/worlds"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "zane:52455245@tcp(localhost:3306)/gamedb?charset=utf8mb4&parseTime=True&loc=Local"
var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

func main() {
	f, err := os.OpenFile("./logs/main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
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
	//API Routes
	r.HandleFunc("/worlds", worldHandler)
	r.HandleFunc("/worlds/{first}", worldHandler)
	r.HandleFunc("/worlds/{first}/{second}", worldHandler)
	r.HandleFunc("/players", playerHandler)
	r.HandleFunc("/players/{first}/{second}", playerHandler)
	log.Println("Serving at", "http://localhost:"+port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}

func worldHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	switch r.Method {
	case http.MethodGet:

		retWorld := &worlds.Worlds{}
		retWorlds := &[]worlds.Worlds{}
		switch vars["first"] {
		case "by":
			db.Where("owner = ?", vars["second"]).
				Find(&retWorlds)

		case "popular":
			db.Where("name = ?", vars["second"]).
				Find(&retWorlds)

		default:
			if vars["first"] == "" {
				db.Find(&retWorlds)
			} else {
				db.Where("name = ?", vars["first"]).
					First(&retWorld)
			}
		}

		if retWorld.Name == "" {
			jsonRetWorlds, err := json.Marshal(retWorlds)
			if err != nil {
				fmt.Fprintln(w, err)
			} else {

				fmt.Fprintln(w, string(jsonRetWorlds))
			}
		} else {
			jsonRetWorld, err := json.Marshal(retWorld)
			if err != nil {
				fmt.Fprintln(w, err)
			} else {
				fmt.Fprintln(w, string(jsonRetWorld))
			}
		}

	case http.MethodPost:
		inWorld := &worlds.Worlds{}
		searchWorld := worlds.Worlds{}
		err := json.NewDecoder(r.Body).Decode(&inWorld)
		if err != nil {
			fmt.Fprintln(w, err)
		}
		db.Where("name = ?", inWorld.Name).
			Find(&searchWorld)
		if searchWorld.Name != "" {
			//exists!
			fmt.Fprintln(w, "world exists")
		} else {
			log.Println(inWorld.Name, "created by", inWorld.Owner)
			db.Create(&inWorld)
		}
	//case http.MethodDelete:
	//case http.MethodPut:
	default:
		fmt.Fprintf(w, "Method not allowed!")
	}
}
func playerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "GET")
	case http.MethodPost:
		fmt.Fprintf(w, "GET")
	default:
		fmt.Fprintf(w, "Method not allowed!")
	}
}
