package add

import "github.com/gin-gonic/gin"

// package to add data and testing, replace Type
type AddUser interface {
	AddUser(c *gin.Context)
}
