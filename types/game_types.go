package types

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

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Status string `json:"status"`
	Description string `json:"description"`
	Token string `json:"token"`
}