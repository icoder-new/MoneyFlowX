package handler

import (
	"fr33d0mz/moneyflowx/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(h.CORSMiddleware())
	router.Use(gin.Recovery())

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ping - pong",
		})
	})

	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.Register)
			auth.POST("/sign-in", h.Login)
			auth.POST("/forgot-password", h.ForgotPassword)
			auth.POST("/reset-password", h.ResetPassword)
		}

		verify := api.Group("/verify")
		{
			verify.GET("/send", AuthMiddleware(h.service.JWT, h.service.User), h.SendToken)
			verify.GET("/:verifyToken", h.VerifyUser)
		}

		user := api.Group("/user")
		{
			user.GET("/profiles", AuthMiddleware(h.service.JWT, h.service.User), h.Profile)
		}

		transaction := api.Group("/transactions")
		{
			transaction.Use(AuthMiddleware(h.service.JWT, h.service.User))
			transaction.GET("/", h.GetTransactions)
			//TODO: add top-up transaction
			transaction.POST("/transfer", h.Transfer)
		}
	}

	return router
}
