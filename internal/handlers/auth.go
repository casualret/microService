package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"microService/internal/auth"
	"microService/internal/models"
	"net/http"
	"strings"
)

func (h *Handlers) SignUp(c *gin.Context) {
	const op = "handlers.SignUp"

	var input models.CreateUserReq
	err := c.BindJSON(&input)
	if err != nil {
		h.Logger.Error("Failed to bind JSON: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}

	err = h.App.SignUp(input)
	if err != nil {
		h.Logger.Error("Error creating User: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusInternalServerError, internalServerError)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *Handlers) SignIn(c *gin.Context) {

	const op = "handlers.SignIn"

	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		h.Logger.Error("Failed to bind JSON: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}

	token, err := h.App.SignIn(input)
	if err != nil {
		h.Logger.Error("Authentication failed: ", err)
		newErrorResponse(c, http.StatusInternalServerError, internalServerError)
		return
	}

	fmt.Println(token)
	header := c.GetHeader("Authorization")
	fmt.Println(header)
	c.JSON(http.StatusOK, token)
}

func (h *Handlers) JWTAuth(c *gin.Context) {
	const BearerSchema = "Bearer "
	header := c.GetHeader("Authorization")
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization Header"})
		c.Abort()
		return
	}

	if !strings.HasPrefix(header, BearerSchema) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization Header"})
		c.Abort()
		return
	}

	tokenStr := header[len(BearerSchema):]
	claims := &auth.Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return auth.JwtKey, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		c.Abort()
		return
	}

	c.Set("username", claims.Username)
	c.Next()
}
