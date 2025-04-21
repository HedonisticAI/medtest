package usecase

import (
	"medtest/internal/auth"
	"medtest/pkg/logger"
	"medtest/pkg/postgres"
	"medtest/pkg/token_cache"

	"github.com/gin-gonic/gin"
)

type AuthUC struct {
	DB     *postgres.Postgres
	Cache  *token_cache.TokenCache
	Logger *logger.Logger
}

func (Auth *AuthUC) SignIn(c *gin.Context) {

}
func (Auth *AuthUC) Refresh(c *gin.Context) {}

func NewAuth(DB *postgres.Postgres, Cache *token_cache.TokenCache, Logger *logger.Logger) auth.Auth {
	return &AuthUC{DB: DB, Cache: Cache, Logger: Logger}
}
