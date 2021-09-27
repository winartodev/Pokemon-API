package main

import (
	controllers "Pokemon-API/Controllers"
	models "Pokemon-API/Models"
	view "Pokemon-API/Views"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
	address  = os.Getenv("DB_ADDRESS")
	port     = os.Getenv("DB_PORT")
	dbname   = os.Getenv("DB_DBNAME")
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, address, port, dbname))
	if err != nil {
		panic(err.Error())
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	pokemonInterface := models.Connect(db)
	pokemonController := controllers.NewPokemonController(pokemonInterface)
	view.EndpointsHandler(pokemonController)

	addr := flag.String("addr", ":8080", "Http Listen And Serve")
	e := http.ListenAndServe(*addr, nil)
	if e == nil {
		log.Fatal("ListenAndServe:", err)
	}
	fmt.Println("ListenAndServe:8080")
}
