package router

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/GreyHood-Studio/ai_server/network"
	"strconv"
	"net/http"
)

var Agents map[int]network.TCPClient

func accessFloor(c *gin.Context) {
	serverIDString := c.Param("serverID")
	serverIP := c.PostForm("ip")
	portString := c.PostForm("port")
	serverId, _ := strconv.Atoi(serverIDString)
	serverPort, _ := strconv.Atoi(portString)

	client := network.NewClient(serverId, serverIP, serverPort)
	client.Run()

	fmt.Println(serverIDString, serverIP, portString)

	c.String(http.StatusOK, fmt.Sprintln("complete access floor ", serverIDString, serverIP, serverPort))
}

func exitFloor(c *gin.Context) {
	ServerIDString := c.Param("serverID")
	serverIP := c.PostForm("ip")
	portString := c.PostForm("port")

	fmt.Println(ServerIDString, serverIP, portString)
}
