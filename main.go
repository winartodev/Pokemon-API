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

func main() {
	// Mysql Configuration
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_ADDRESS"), os.Getenv("DB_PORT"), os.Getenv("DB_DBNAME")))

	if err != nil {
		panic(err.Error())
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	pokemonInterface := models.Connect(db)
	pokemonController := controllers.NewPokemonController(pokemonInterface)
	view.EndpointsHandler(pokemonController)

	// Array Configuration
	// var pokemonStruct = &pokemon.Entity{}
	// pokemonInterface := models.WithStruct(pokemonStruct)
	// pokemonController := controllers.NewPokemonController(pokemonInterface)
	// view.EndpointsHandler(pokemonController)

	log.Println("started at : http://127.0.0.1:8080/")
	var addr = flag.String("addr", ":8080", "Http Listen And Serve")
	e := http.ListenAndServe(*addr, nil)
	if e != nil {
		log.Fatal("ListenAndServe:", e)
	}
}
