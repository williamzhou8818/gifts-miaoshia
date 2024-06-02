package main

import (
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/williamzhou8818/gifts-miaoshia/database"
	"github.com/williamzhou8818/gifts-miaoshia/handler"
	"github.com/williamzhou8818/gifts-miaoshia/util"
)

var (
	writeOrderFinish bool
)

func listenSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	for {
		sig := <-c //阻塞，直到信号的到来
		if writeOrderFinish {
			util.LogRus.Infof("receive signal %s, exit", sig.String())
			os.Exit(0)
		} else {
			util.LogRus.Infof("receive signal %s, but not exit", sig.String())
		}
	}
}

func Init() {
	util.InitLog("log")
	database.InitGiftInventory()

	if err := database.ClearOrders(); err != nil {
		panic(err)
	} else {
		util.LogRus.Info("clear table orders")
	}

	handler.InitChannel()
	//
	go listenSignal()
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
	router.GET("/lucky", handler.Lottery)     ////点击抽奖按钮
	router.Run("localhost:7777")
}

// go run ./main.go
