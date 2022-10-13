package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	svr := http.Server{
		Addr:    ":7070",
		Handler: startEngine(),
	}
	if err := svr.ListenAndServe(); err != nil {
		log.Fatal("start server failed ", err.Error())
	}
}

func startEngine() (engine *gin.Engine) {
	engine = gin.Default()
	engine.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "PONG")
	})
	return
}