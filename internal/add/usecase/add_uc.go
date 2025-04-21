package add_usecase

import (
	"medtest/internal/add"
	"medtest/pkg/logger"
	"medtest/pkg/postgres"
	"net/http"

	"github.com/beevik/guid"
	"github.com/gin-gonic/gin"
)

type AddUc struct {
	DB     *postgres.Postgres
	Logger *logger.Logger
}

func (AddUc *AddUc) AddUser(c *gin.Context) {
	AddUc.Logger.Info("Got Add request")
	Email := c.Query("Mail")
	Guid := guid.New()
	db := AddUc.DB.Get()
	_, err := db.Exec("insert into users (email, uniID) values ($1, $2)", Email, Guid.String())
	if err != nil {
		AddUc.Logger.Debug(err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "User GUID for future Requests: "+Guid.String())
}

func NewAdd(DB *postgres.Postgres, Logger *logger.Logger) add.AddUser {
	return &AddUc{DB: DB, Logger: Logger}
}
