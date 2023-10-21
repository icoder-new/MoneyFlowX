package logger

import (
	"fmt"
	"fr33d0mz/moneyflowx/utils"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Info  *log.Logger
	Error *log.Logger
	Warn  *log.Logger
	Debug *log.Logger
)

func Init() {
	fileInfo, err := os.OpenFile(utils.AppSettings.AppParams.LogInfo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	fileError, err := os.OpenFile(utils.AppSettings.AppParams.LogError, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	fileWarn, err := os.OpenFile(utils.AppSettings.AppParams.LogWarning, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	fileDebug, err := os.OpenFile(utils.AppSettings.AppParams.LogDebug, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}

	Info = log.New(fileInfo, "", log.Ldate|log.Lmicroseconds)
	Error = log.New(fileError, "", log.Ldate|log.Lmicroseconds)
	Warn = log.New(fileWarn, "", log.Ldate|log.Lmicroseconds)
	Debug = log.New(fileDebug, "", log.Ldate|log.Lmicroseconds)

	lumberLogInfo := &lumberjack.Logger{
		Filename:   utils.AppSettings.AppParams.LogInfo,
		MaxSize:    utils.AppSettings.AppParams.LogMaxSize, // megabytes
		MaxBackups: utils.AppSettings.AppParams.LogMaxBackups,
		MaxAge:     utils.AppSettings.AppParams.LogMaxAge,   //days
		Compress:   utils.AppSettings.AppParams.LogCompress, // disabled by default
		LocalTime:  true,
	}

	lumberLogError := &lumberjack.Logger{
		Filename:   utils.AppSettings.AppParams.LogError,
		MaxSize:    utils.AppSettings.AppParams.LogMaxSize, // megabytes
		MaxBackups: utils.AppSettings.AppParams.LogMaxBackups,
		MaxAge:     utils.AppSettings.AppParams.LogMaxAge,   //days
		Compress:   utils.AppSettings.AppParams.LogCompress, // disabled by default
		LocalTime:  true,
	}

	lumberLogWarn := &lumberjack.Logger{
		Filename:   utils.AppSettings.AppParams.LogWarning,
		MaxSize:    utils.AppSettings.AppParams.LogMaxSize, // megabytes
		MaxBackups: utils.AppSettings.AppParams.LogMaxBackups,
		MaxAge:     utils.AppSettings.AppParams.LogMaxAge,   //days
		Compress:   utils.AppSettings.AppParams.LogCompress, // disabled by default
		LocalTime:  true,
	}

	lumberLogDebug := &lumberjack.Logger{
		Filename:   utils.AppSettings.AppParams.LogDebug,
		MaxSize:    utils.AppSettings.AppParams.LogMaxSize, // megabytes
		MaxBackups: utils.AppSettings.AppParams.LogMaxBackups,
		MaxAge:     utils.AppSettings.AppParams.LogMaxAge,   //days
		Compress:   utils.AppSettings.AppParams.LogCompress, // disabled by default
		LocalTime:  true,
	}

	gin.DefaultWriter = io.MultiWriter(os.Stdout, lumberLogInfo)

	Info.SetOutput(gin.DefaultWriter)
	Error.SetOutput(lumberLogError)
	Warn.SetOutput(lumberLogWarn)
	Debug.SetOutput(lumberLogDebug)
}
