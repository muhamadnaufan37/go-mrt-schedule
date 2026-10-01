package station

import (
	"net/http"

	"github.com/gin-gonic/gin"
	response "github.com/muhamadnaufan37/go-mrt-schedule.git/modules/common/reseponse"
)

func Initiate(api *gin.RouterGroup) {
	StationService := NewService()

	station := api.Group("/station")
	station.GET("/", func(c *gin.Context) {
		GetAllStation(c, StationService)
	})
}

func GetAllStation(c *gin.Context, service Service) {
	datas, err := service.GetAllStation()
	if err != nil {
		c.JSON(http.StatusBadRequest,
			response.APIResponse{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})
		return
	}

	c.JSON(http.StatusOK,
		response.APIResponse{
			Success: true,
			Message: "Data berhasil diambil",
			Data:    datas,
		})

}
