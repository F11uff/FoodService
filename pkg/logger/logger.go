package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	_const "pet/config/const"
)

const (
	RW             = 0666
	logPath string = "../logs/logs.log"
)

type LoggerService struct {
	Log  *logrus.Logger
	File *os.File
	Path string
}

func NewLogger() *LoggerService {
	return &LoggerService{
		Log:  logrus.New(),
		File: nil,
		Path: _const.PATH_CONFIG_FILE,
	}
}

func (logger *LoggerService) OpenFile() error {
	var err error

	if logger.File != nil {
		logger.Log.Info("File already open (logger.go)")
		return err
	}

	logger.File, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, RW)

	if err != nil {
		logger.Log.Info("Failed to open log file (logger.go)")
		return err
	}

	logger.Log.SetOutput(logger.File)
	logger.Log.SetLevel(logrus.InfoLevel)
	logger.Log.SetFormatter(&logrus.JSONFormatter{})

	return nil
}

func (logger *LoggerService) CloseFile() error {
	if logger.File != nil {
		err := logger.File.Close()

		if err != nil {
			logger.Log.Info("Failed to close log file (logger.go)")
			return err
		}

		logger.File = nil
	}

	return nil
}
