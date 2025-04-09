package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"pet/backend/config"
	_const "pet/backend/config/const"
	logger2 "pet/backend/internal/service"
	"syscall"
)

func main() {
	log := logger2.NewLogger()
	err := log.OpenFile()

	if err != nil {
		return
	}

	defer func(log *logger2.LoggerService) {
		err := log.CloseFile()

		if err != nil {
			return
		}
	}(log)

	conf, err := config.NewConfig(_const.PATH_CONFIG_FILE, *log)

	if err != nil {
		log.Log.Info("Error loading config:", err)
		return
	}

	router := gin.Default()

	logFileAPI, err := os.OpenFile(_const.PATH_LOGFILE_API, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, logger2.RW)
	gin.DefaultWriter = logFileAPI
	defer logFileAPI.Close()

	router.Use(gin.Logger())

	log.Log.Info("Сервер запущен на порту : ", conf.HttpServer.Address)

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	go func() {
		log.Log.Fatal(router.Run(conf.HttpServer.Address))
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	sig := <-c
	log.Log.Info("Получен сигнал завершения: ", sig)

	log.Log.Info("Сервер закрыт")
}
