package main

import (
	"github.com/gin-gonic/gin"
	"github.com/GreyHood-Studio/ai_server/router"
)

func mainHTTPServer(port string) {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	router.SetAPIRoute(r)

	r.Run(port)
}

func main() {

	mainHTTPServer(":4000")
}