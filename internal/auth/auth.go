package auth

import "github.com/gin-gonic/gin"

type Auth interface {
	SignIn(c *gin.Context)
	Refresh(c *gin.Context)
}
