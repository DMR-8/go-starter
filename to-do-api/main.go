package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/DMR-8/to-do-api/controller"
	"github.com/DMR-8/to-do-api/model"
	_ "github.com/go-sql-driver/mysql" // mysql driver
)

func main() {
	mux := controller.Register()
	db := model.Connect()
	defer db.Close()
	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}