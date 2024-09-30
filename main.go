package main

import (
	"github.com/douglasmg7/gin_rest_api.git/db"
	"github.com/douglasmg7/gin_rest_api.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDb()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
