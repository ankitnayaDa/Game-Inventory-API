package libs

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	ty "inventory/types"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Database struct {
	Conn *sql.DB
}

//Connect to DB
func DBConnect() Database {
	db := Database{}
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "PostgreSQL", ty.DBPORT, ty.DBUSER, ty.DBPASSWORD, ty.DBNAME)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("DB Open Failed : %s", err)
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		log.Fatalf("DB Ping Failed : %s")
	}
	log.Println("DB Connect Successfully")
	return db
}

func AddToInventory(w http.ResponseWriter, r *http.Request) {
	var gameinv ty.GameInventory
	resbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("JSON read failed", err)
	}
	_ = json.Unmarshal(resbody, &gameinv)

	var response = ty.AddToInventoryResponse{}

	if string(gameinv.GameID) == "" || gameinv.GameName == "" {
		response = ty.AddToInventoryResponse{Status: "Failure", Description: "Game ID/Name is missing"}
	} else {
		db := DBConnect()
		GameID := strconv.Itoa(gameinv.GameID)
		log.Println("Inserting game into DB")
		log.Println("Inserting new game with ID: " + GameID + " ,name: " + gameinv.GameName + " ,gametype: " + gameinv.GameType + ",GameStudio: " + gameinv.GameStudio + ",Platform: " + gameinv.Platform)

		sqlstatement := `INSERT INTO games(gameID, gameName, gameType, gameStudio, platform) VALUES($1, $2, $3, $4, $5) returning id;`
		log.Println(sqlstatement)
		_, err = db.Conn.Query("INSERT INTO games(gameID, gameName, gameType, gameStudio, platform) VALUES($1, $2, $3, $4, $5)", GameID, gameinv.GameName, gameinv.GameType, gameinv.GameStudio, gameinv.Platform)
		if err != nil {
			log.Println("DB insert Failed : %s", err)
			response = ty.AddToInventoryResponse{Status: "FAILURE", Description: "Game is addition failed"}
		}
		response = ty.AddToInventoryResponse{Status: "SUCCESS", Description: "Game is added to Inventory"}
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteFromInventory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gameID := params["gameid"]
	var response = ty.DeleteFromInventory{}
	if gameID == "" {
		response = ty.DeleteFromInventory{Status: "FAILURE", Description: "gameID is not present "}
	} else {
		db := DBConnect()
		log.Println("Deleting games from DB")
		_, err := db.Conn.Query("DELETE FROM games where gameID = $1", gameID)
		if err != nil {
			log.Println("DB delete failed : %s", err)
			response = ty.DeleteFromInventory{Status: "FAILURE", Description: "Game is deletion failed"}
		}
		response = ty.DeleteFromInventory{Status: "SUCCESS", Description: "Game is deleted from Inventory"}
	}
	json.NewEncoder(w).Encode(response)
}

func ListInventory(w http.ResponseWriter, r *http.Request) {
	db := DBConnect()
	var response = ty.DeleteFromInventory{}
	log.Println("Getting games...")
	rows, err := db.Conn.Query("SELECT * FROM games")
	if err != nil {
		log.Println("DB Query Failed : %s", err)
	}
	var games []ty.GameInventory
	for rows.Next() {
		var gameID int
		var gameName string
		var gameStudio string
		var gameType string
		var platform string
		err = rows.Scan(&gameID, &gameName, &gameStudio, &gameType, &platform)
		if err != nil {
			log.Println("DB list failed : %s", err)
		}
		games = append(games, ty.GameInventory{GameID: gameID, GameName: gameName, GameStudio: gameStudio, GameType: gameType, Platform: platform})
	}
	if games == nil {
		response = ty.DeleteFromInventory{Status: "SUCCESS", Description: "Game is Not Added", Games: games}
	}
	response = ty.DeleteFromInventory{Status: "SUCCESS", Description: "Game is Listed", Games: games}
	json.NewEncoder(w).Encode(response)
}
