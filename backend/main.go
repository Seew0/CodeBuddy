package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/seew0/DoubtBuddy/server"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	port := os.Getenv("port")
	engine := gin.Default()
	Server := server.NewServer(port,engine)
	Server.Run()
}
