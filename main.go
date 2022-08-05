package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	lb "inventory/libs"
)

func main() {

	// Init the mux router
	API := mux.NewRouter()

	// List All Games
	API.HandleFunc("/login/", lb.Login).Methods("POST")

	// List All Games
	API.HandleFunc("/games/", lb.ListInventory).Methods("GET")

	// Add A Game To Inventory
	API.HandleFunc("/game/", lb.AddToInventory).Methods("POST")

	// Delete A Game From Inventory
	API.HandleFunc("/games/{gameid}", lb.DeleteFromInventory).Methods("DELETE")

	// serve the app
	fmt.Println("Server at 9678")
	log.Fatal(http.ListenAndServe(":9678", API))
}