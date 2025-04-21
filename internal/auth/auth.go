package auth

import (
	"errors"

	"github.com/gin-gonic/gin"
)

var ErrBadPair = errors.New("despite refresh token being valid, pair you provided doesnt exist")
var ErrNoPair = errors.New("no refresh token")

type Auth interface {
	SignIn(c *gin.Context)
	Refresh(c *gin.Context)
}

type SignInResponse struct {
	AccessToken  string `json:"Access"`
	RefreshToken string `json:"Refresh"`
	// HashedRef    string `json:"Hashed"`
}

// They are the same, yep
type RefreshRequest struct {
	AccessToken  string `json:"Access"`
	RefreshToken string `json:"Refresh"`
}
