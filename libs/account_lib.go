package libs

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"inventory/types"
	"io/ioutil"
	"log"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var login types.Login
	resbody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("JSON read failed", err)
	}
	_ = json.Unmarshal(resbody, &login)

	var response= types.LoginResponse{}

	if login.Username == "" || login.Password == "" {
		response = types.LoginResponse{Status: "Failure", Description: "Username/Password is missing"}
	} else {
		Token := GenerateSecureToken(12)
		err := InsertAccountDataToDB(login,Token)
		if err != nil {
			response = types.LoginResponse{Status: "Failure", Description: "Username/Password is missing"}
			_=json.NewEncoder(w).Encode(response)
		}
		response = types.LoginResponse{Status: "Success", Description: "Successful",Token:Token}
		_=json.NewEncoder(w).Encode(response)
	}
}

func InsertAccountDataToDB(login types.Login,Token string)error{
	db := DBConnect()
	log.Println("Inserting Username Details into DB")
	sqlstatement := `INSERT INTO games(username,token) VALUES($1, $2) returning id;`
	log.Println(sqlstatement)
	_,err := db.Conn.Query("INSERT INTO username(username,token) VALUES($1, $2)",login.Username,Token)
	if err != nil {
		log.Println("DB insert Failed : %s", err)
		return errors.New("DB insert Failed")
	}
	return nil
}

func GenerateSecureToken(length int) string {
		b := make([]byte, length)
		if _, err := rand.Read(b); err != nil {
		return ""
		}
		return hex.EncodeToString(b)
	}
