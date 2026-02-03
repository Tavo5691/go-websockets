package middleware

import (
	"fmt"
	"gochat/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(jwtKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Query("token")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(jwtKey), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "unauthorized"})
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "unauthorized"})
			return
		}

		userId, err := token.Claims.GetSubject()
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "internal error"})
			return
		}

		c.Set("userId", userId)
		c.Next()
	}
}
