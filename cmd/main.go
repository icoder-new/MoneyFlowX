package main

import (
	"context"
	"github.com/icoder-new/MoneyFlowX"
	"github.com/icoder-new/MoneyFlowX/db"
	"github.com/icoder-new/MoneyFlowX/logger"
	"github.com/icoder-new/MoneyFlowX/models"
	"github.com/icoder-new/MoneyFlowX/pkg/handler"
	"github.com/icoder-new/MoneyFlowX/pkg/repository"
	"github.com/icoder-new/MoneyFlowX/pkg/service"
	"github.com/icoder-new/MoneyFlowX/utils"
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
	err := _db.AutoMigrate(&models.User{}, &models.PasswordReset{}, &models.Wallet{}, &models.Transaction{})
	if err != nil {
		logger.Error.Fatalf("[MAIN] error while auto migrating: %v", err.Error())
		return
	}
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
