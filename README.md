# API's for game inventory managment

REST API's for insert ,delete and list games from Inventory

## Installation

cd src/gameInventoryAPI/
go build main.go
./main

postgres needs to be installed
Query for postgres:
psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'games'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE games" &&\ psql -c "ALTER USER postgres WITH ENCRYPTED PASSWORD 'postgres';" &&\ psql -U postgres -c CREATE TABLE games (gameID varchar(50) NOT NULL,gameName varchar(50) NOT NULL,gameType varchar(50) NOT NULL,gameStudio varchar(50) NOT NULL,platform varchar(50) NOT NULL,PRIMARY KEY (id))

## Usage

GET : http://localhost:9678/games/
DELETE : http://localhost:9678/games/<gameid>/
LIST : http://localhost:9678/game/
