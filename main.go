package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/williamzhou8818/gifts-miaoshia/database"
	"github.com/williamzhou8818/gifts-miaoshia/handler"
	"github.com/williamzhou8818/gifts-miaoshia/util"
)

func Init() {
	util.InitLog("log")

	database.GetGiftDBConnection()

}

func main() {

	Init()
	//gin
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	//
	router := gin.Default()

	router.Static("/js", "views/js") // 在URL是访问目录/js相当于访问文件系统中的views/js目录
	router.Static("/img", "views/img")
	router.StaticFile("/favicon.ico", "views/img/dqq.png")
	router.LoadHTMLFiles("views/lottery.html")

	//-------------------------
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "lottery.html", nil)
	})

	router.GET("/gifts", handler.GetAllGifts) //获取所有奖品信息
	router.Run("localhost:7777")
}

// go run ./main.go
