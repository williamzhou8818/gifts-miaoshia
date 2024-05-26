package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/williamzhou8818/gifts-miaoshia/database"
)

// 获取所有奖品信息，用于初始化轮盘
func GetAllGifts(ctx *gin.Context) {
	gifts := database.GetAllGiftsV1()
	if len(gifts) == 0 {
		ctx.JSON(http.StatusInternalServerError, nil)
	} else {
		//抹掉敏感信息
		for _, gift := range gifts {
			gift.Count = 0
		}
		ctx.JSON(http.StatusOK, gifts)
	}
}
