package types

const DBHOST = "127.0.0.1"

const DBPORT = "5432"

const DBUSER = "postgres"

const DBPASSWORD  = "postgres"

const DBNAME  = "games"

type GameInventory struct {
	GameID      int    `json:"gameid"`
	GameName    string `json:"gamename"`
	GameType    string `json:"gametype"`
	GameStudio string `json:"gamestudio"`
	Platform string `json:"platform"`
}

type AddToInventoryResponse struct {
	Status string `json:"status"`
	Description string `json:"description"`
	Games []GameInventory `json:"game"`
}

type DeleteFromInventory struct {
	Status string `json:"status"`
	Description string `json:"description"`
	Games []GameInventory `json:"game"`
}

type ListInventory struct {
	Status string `json:"status"`
	Description string `json:"description"`
	Games []GameInventory `json:"game"`
}
