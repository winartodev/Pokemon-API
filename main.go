package main

import (
	controllers "Pokemon-API/Controllers"
	models "Pokemon-API/Models"
	view "Pokemon-API/Views"
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/db_pokemon")
	if err != nil {
		panic(err.Error())
	}
	pokemonInterface := models.Connect(db)
	pokemonController := controllers.NewPokemonController(pokemonInterface)
	view.EndpointsHandler(pokemonController)

	addr := flag.String("addr", ":8080", "Http Listen And Serve")
	e := http.ListenAndServe(*addr, nil)
	if e == nil {
		log.Fatal("ListenAndServe:", err)
	}
}