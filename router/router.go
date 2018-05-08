package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("Health Play Server"))
}

// router package는 gameserver들을 컨트롤 하기 위한 로직
func SetAPIRoute(router *gin.Engine) {
	// 서버의 상태를 가지고 오는 정보
	//router.GET("/floors", getServerStatus)
	router.GET("/ping", healthCheck)
	
	// play server에 agent 형태로 붙거나, 나가는 형태
	floor := router.Group("/floor")
	{
		floor.POST("/:serverID", accessFloor)
		floor.DELETE("/:serverID", exitFloor)
	}

	// 몬스터 리젠, 보스 출현 etc...
	monster := router.Group("/monster")
	{
		monster.POST("/:serverID", createMonster)
	}
}