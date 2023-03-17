package main

import (
	models "auction-house-model/api/Models/databasesConnection"
	routes "auction-house-model/api/Routes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	InitServer()
}

func InitServer() {
	godotenv.Load("../.env")
	server := gin.Default()
	routes.Routes(server)
	models.GetDataBase("postgres")
	server.Run(os.Getenv("PORT"))
}
