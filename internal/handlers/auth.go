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

	var input models.User
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

	var input models.UserLogin

	if err := c.ShouldBindJSON(&input); err != nil {
		h.Logger.Error("Failed to bind JSON: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}

	token, err := h.App.SignIn(input)
	if err != nil {
		h.Logger.Error("Authentication failed: ", err)
		newErrorResponse(c, http.StatusInternalServerError, errorAuthorize)
		return
	}

	fmt.Println(token)
	header := c.GetHeader("Authorization")
	fmt.Println(header)
	c.JSON(http.StatusOK, token)
}

func (h *Handlers) JWTAuth(c *gin.Context) {
	const BearerSchema = "Bearer "
	const op = "handlers.JWTAuth"

	header := c.GetHeader("Authorization")
	if header == "" {
		h.Logger.Error("Missing Authorization Header", op)
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}

	if !strings.HasPrefix(header, BearerSchema) {
		h.Logger.Error("Invalid Authorization Header", op)
		newErrorResponse(c, http.StatusBadRequest, incorrectData)
		return
	}

	tokenStr := header[len(BearerSchema):]
	claims := &auth.TokenClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return auth.JwtKey, nil
	})

	if err != nil {
		h.Logger.Error("Invalid Token: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, errorParseToken)
		return
	}

	if !token.Valid {
		h.Logger.Error("Invalid Token: ", fmt.Errorf("%s: %v", op, err))
		newErrorResponse(c, http.StatusBadRequest, userUnauthorized)
		return
	}

	c.Set("username", claims.Username)
	c.Set("role", claims.Role)
	c.Next()
}

func (h *Handlers) isAdminMiddleware(c *gin.Context) {
	const op = "handlers.isAdminMiddleware"

	role, ok := c.Get("role")
	if !ok {
		h.Logger.Error("Error get role: ", op)
		newErrorResponse(c, http.StatusUnauthorized, userUnauthorized)
		return
	}

	if role != "admin" {
		h.Logger.Error("User is not admin: ", op)
		newErrorResponse(c, http.StatusForbidden, userAccessDenied)
		return
	}
	c.Next()
}
