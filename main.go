package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muhamadnaufan37/go-mrt-schedule.git/modules/station"
)

func main() {
	InitRouter()
}

func InitRouter() {
	var (
		router = gin.Default()
		api    = router.Group("/v1/api")
	)

	station.Initiate(api)

	router.Run("localhost:8080")

}
