package main

import (
	"encoding/json"
	"log"
	"net/http"
	// "fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	// "github.com/lib/pq"
    "github.com/rs/cors"
	"github.com/joho/godotenv"
	"os"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)





type Resource struct {
	gorm.Model

	Link        string
	Name        string
	Author      string
	Description string
}

var db *gorm.DB
var err error


func main() {
	dot := godotenv.Load()
  	if dot != nil {
    log.Fatal("Error loading .env file")
	}
	
	router := mux.NewRouter()
	db, err = gorm.Open(
		"postgres",
		"host="+os.Getenv("HOST")+" port="+os.Getenv("PORT")+" user="+os.Getenv("USER")+
		" dbname="+os.Getenv("DBNAME")+" sslmode=disable password="+os.Getenv("PASSWORD"))
	if err != nil {
	  panic(err)
	}
	defer db.Close()
  
	db.AutoMigrate(&Resource{})

	router.HandleFunc("/resources", GetResources).Methods("GET")
	router.HandleFunc("/resources/{id}", GetResource).Methods("GET")
	router.HandleFunc("/resources", CreateResource).Methods("POST")
	router.HandleFunc("/resources/{id}", DeleteResource).Methods("DELETE")

    handler := cors.Default().Handler(router)
	
	log.Fatal(http.ListenAndServe(":3001", handler))
}

func GetResources(w http.ResponseWriter, r *http.Request) {
	var resources []Resource
	db.Find(&resources)
	json.NewEncoder(w).Encode(&resources)
}

func GetResource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var resource Resource
	db.First(&resource, params["id"])
	json.NewEncoder(w).Encode(&resource)
}

func CreateResource(w http.ResponseWriter, r *http.Request) {
	var resource Resource
	json.NewDecoder(r.Body).Decode(&resource)
	db.Create(&resource)
	json.NewEncoder(w).Encode(&resource)
}

func DeleteResource(w http.ResponseWriter, r *http.Request) {
	//delete
	params := mux.Vars(r)
	var resource Resource
	db.First(&resource, params["id"])
	db.Delete(&resource)

	var resources []Resource
	db.Find(&resources)
	json.NewEncoder(w).Encode(&resources)
}