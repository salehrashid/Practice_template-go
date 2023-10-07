package main

import (
	"fmt"
	"go-template/constants"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"log"
	"net/http"
)

type Cars struct {
	gorm.Model
	Id     int
	Name   string
	Engine string
}

func main() {
	httpHandler()
}

func httpHandler() {
	fmt.Println("server nya lagi jalan nih bang, http://localhost:7000")

	http.HandleFunc("/car", car)

	log.Fatal(http.ListenAndServe(":7000", nil))
}

func car(writer http.ResponseWriter, _ *http.Request) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		constants.Host, constants.Port, constants.User, constants.Password, constants.Dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	var cars []Cars
	db.Find(&cars)

	tmplt := template.Must(template.ParseFiles("template/car-table.html"))
	if err := tmplt.Execute(writer, cars); err != nil {
		panic(err)
	}
}
