package router

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func createMonster(c *gin.Context) {
	ServerIDString := c.Param("serverID")

	fmt.Println(ServerIDString)
}
