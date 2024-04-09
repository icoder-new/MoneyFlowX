package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/icoder-new/MoneyFlowX/pkg/dto"
	"github.com/icoder-new/MoneyFlowX/pkg/service"
	"github.com/icoder-new/MoneyFlowX/utils"
	"net/http"
	"strings"
)

func (h *Handler) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,PATCH,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, auth")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}

		c.Next()
	}
}

func AuthMiddleware(jwtService service.JWT, userService service.User) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, "unrecognized token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) != 2 {
			response := utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, "unrecognized token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		encodedToken := arrayToken[1]
		token, err := jwtService.ValidateToken(encodedToken)
		if err != nil {
			response := utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		payload, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, "not a valid bearer token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := payload["user_id"].(string)

		params := &dto.UserRequestParams{}
		params.UserID = userID
		user, err := userService.GetUser(params)
		if err != nil {
			response := utils.ErrorResponse("Unauthorized", http.StatusUnauthorized, err.Error())
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("user", user)
		c.Next()

	}
}
