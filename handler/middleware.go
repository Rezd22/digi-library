package handler

import (
	"fmt"
	"net/http"
	"paa/utils"

	"github.com/gin-gonic/gin"
)

func (h *Handler) IsLogin(c *gin.Context) {
	cookie, err := c.Request.Cookie("jwt")
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	// Extract token string
	tokenString := cookie.Value

	// Validate token
	token, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return
	}
	login := fmt.Sprintf("%v", token)
	c.Set("username", login)
	c.Next()
}
