package libs

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	ty "gameInventoryAPI/types"
	"net/http"
	"strconv"
)

//Connect to DB
func DBConnect () *sql.DB{
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", ty.DBHOST,ty.DBPORT,ty.DBUSER,ty.DBPASSWORD,ty.DBNAME)
	data, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("DB Connect Failed : %s", err)
	}
	return data
}

func AddToInventory (w http.ResponseWriter, r *http.Request){
	var gameinv ty.GameInventory
	resbody , err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("DB Connect Failed : %s", err)
	}
	_=json.Unmarshal(resbody,&gameinv)

	var response = ty.AddToInventoryResponse{}

	if string(gameinv.GameID) == "" || gameinv.GameName == "" {
		response = ty.AddToInventoryResponse{Status:"Failure",Description:"Game ID/Name is missing",}
	} else {
		db := DBConnect()
		GameID := strconv.Itoa(gameinv.GameID)
		fmt.Println("Inserting game into DB")
		fmt.Println("Inserting new game with ID: " + GameID + " ,name: " + gameinv.GameName + " ,gametype: " + gameinv.GameType + ",GameStudio: " + gameinv.GameStudio + ",Platform: " + gameinv.Platform)

		sqlstatement := `INSERT INTO games(gameID, gameName, gameType, gameStudio, platform) VALUES($1, $2, $3, $4, $5) returning id;`
		fmt.Println(sqlstatement)
		_,err := db.Exec("INSERT INTO games(gameID, gameName, gameType, gameStudio, platform) VALUES($1, $2, $3, $4, $5)",GameID,gameinv.GameName,gameinv.GameType,gameinv.GameStudio,gameinv.Platform)
		if err != nil {
			log.Fatalf("DB insert Failed : %s", err)
		}
		response = ty.AddToInventoryResponse{Status: "SUCCESS", Description: "Game is added to Inventory"}
	}
	json.NewEncoder(w).Encode(response)
}

func DeleteFromInventory (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	gameID := params["gameid"]
	var response = ty.DeleteFromInventory{}
	if gameID == "" {
		response = ty.DeleteFromInventory{Status: "FAILURE", Description: "gameID is not present "}
	} else {
		db := DBConnect()
		fmt.Println("Deleting games from DB")
		_, err := db.Exec("DELETE FROM games where gameID = $1", gameID)
		if err != nil {
			log.Fatalf("DB Connect Failed : %s", err)
		}
		response = ty.DeleteFromInventory{Status: "SUCCESS", Description: "Game is deleted from Inventory"}
	}
	json.NewEncoder(w).Encode(response)
}

func ListInventory (w http.ResponseWriter, r *http.Request){
	db := DBConnect()
	fmt.Println("Getting games...")
	rows, err := db.Query("SELECT * FROM games")
	if err != nil {
		log.Fatalf("DB Connect Failed : %s", err)
	}
	var games []ty.GameInventory
	for rows.Next() {
		var id int
		var gameID int
		var gameName string
		var gameStudio string
		var gameType string
		var platform string
		err = rows.Scan(&id, &gameID, &gameName, &gameStudio, &gameType, &platform)
		if err != nil {
			log.Fatalf("DB Connect Failed : %s", err)
		}
		games = append(games, ty.GameInventory{GameID: gameID, GameName: gameName, GameStudio:gameStudio, GameType:gameType, Platform:platform})
	}
	var response = ty.DeleteFromInventory{Status: "SUCCESS", Description: "Game is Listed", Games:games}
	json.NewEncoder(w).Encode(response)
}