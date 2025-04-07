package main

import (
	"fmt"
	"pet/config"
	_const "pet/config/const"
	logger2 "pet/pkg/logger"
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

	fmt.Println(conf)
	fmt.Println(conf.Db)
	fmt.Println(conf.HttpServer)
	//fmt.Println(log)
	log.Log.Debug("rigjie")

	log.Log.Info("Сервер запущен")
}
