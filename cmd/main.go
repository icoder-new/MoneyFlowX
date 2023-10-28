package main

import (
	"context"
	"fr33d0mz/moneyflowx"
	"fr33d0mz/moneyflowx/db"
	"fr33d0mz/moneyflowx/logger"
	"fr33d0mz/moneyflowx/pkg/handler"
	"fr33d0mz/moneyflowx/pkg/repository"
	"fr33d0mz/moneyflowx/pkg/service"
	"fr33d0mz/moneyflowx/utils"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	utils.ReadSettings()
	utils.PutAdditionalSettings()

	logger.Init()
	logger.Info.Println("[MAIN] logger working")

	db.StartDbConnection()
	_db := db.GetDBConn()
	db.AutoMigrate(_db)
	defer db.DisconnectDB(_db)

	repository := repository.NewRepository(_db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	srv := new(moneyflowx.Server)

	go func() {
		if err := srv.Run(utils.AppSettings.AppParams.PortRun, handler.InitRoutes()); err != nil {
			logger.Error.Fatalf("[MAIN] error occurred while running server: %v", err.Error())
			return
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Error.Fatalf("[MAIN] error occurred while shutdown server: %v", err.Error())
		return
	}
}
