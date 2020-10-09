package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	port := 8787

	router := gin.Default()

	router.GET("/ping", PingPongHandler)

	routerErr := router.Run(fmt.Sprintf(":%d", port))
	if routerErr != nil {
		log.Fatal(routerErr)
	}
}

func PingPongHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
