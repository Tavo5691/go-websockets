package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenRequest struct {
	UserID string `json:"user_id"`
}

type TokenResponse struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

func (h *Handler) DevToken(c *gin.Context) {
	var req TokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		req.UserID = uuid.New().String()
	}

	if req.UserID == "" {
		req.UserID = uuid.New().String()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   req.UserID,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	tokenString, err := token.SignedString([]byte(h.jwtKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, TokenResponse{
		Token:  tokenString,
		UserID: req.UserID,
	})
}
