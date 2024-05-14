## Game Inventory Management REST API
This project implements a RESTful API for managing game inventory. It allows users to perform CRUD (Create, Read, Update, Delete) operations on game items stored in the inventory.

## Features
**Create** : Add new game items to the inventory.\
**Read** : Retrieve information about existing game items.\
**Update** : Modify details of game items already in the inventory.\
**Delete** : Remove game items from the inventory.
## Technologies Used
**Go Language** : Programming language used for backend development and REST API.\
**Docker** : running applications.\
**Docker Compose** : allows you to define and manage multi-container applications in a single YAML file.\
**PostgreSQL** : Database management system used for storing game inventory data.
## Installation
Clone the repository:

```sh   
git clone https://github.com/ankitnayaDa/Game-Inventory-API.git
```
Navigate to the project directory:

```sh
cd game-inventory-management
```

Install a docker and docker compose:

```sh
https://docs.docker.com/engine/install/
https://docs.docker.com/compose/install/linux/#install-using-the-repository
```

Run the app :

```sh
docker-compose build
docker-compose up -d
```

## API Endpoints
GET http://localhost:9678/games/: Retrieve a list of all game items.\
POST http://localhost:9678/games/: Create a new game item.\
PUT http://localhost:9678/games/{id}: Update details of a specific game item by ID.\
DELETE http://localhost:9678/games/{id}: Delete a game item by ID.

## Usage
Use a tool like Postman or cURL to send HTTP requests to the API endpoints.

Example requests:
```sh
http
Copy code
GET http://localhost:9678/games/
http
Copy code
POST /items
Content-Type: application/json

{
  "gameID": "Game Name",
  "gameName": "mass effect 1",
  "gameType": "Action",
  "gameStudio": "EA",
  "platform ": "xbox"
}
```
## Authentication
This API does not currently require authentication. However, for production use, consider implementing authentication mechanisms such as JWT (JSON Web Tokens) or OAuth.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
