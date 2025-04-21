package auth_usecase

import (
	"medtest/internal/auth"
	cryptoutils "medtest/pkg/crypto_utils"
	"medtest/pkg/logger"
	"medtest/pkg/mail"
	"medtest/pkg/postgres"
	"medtest/pkg/token_cache"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthUC struct {
	DB     *postgres.Postgres
	Cache  *token_cache.TokenCache
	Logger *logger.Logger
	Notify mail.Notify
}

func (Auth *AuthUC) SignIn(c *gin.Context) {
	var ID int
	var Resp auth.SignInResponse
	var TokenID int
	var Mail string
	Auth.Logger.Info("Got sign in request")
	Guid := c.Query("GUID")
	db := Auth.DB.Get()
	row := db.QueryRow("select id, email from users where uniID = $1", Guid)
	if row.Err() != nil {
		Auth.Logger.Debug(row.Err())
		c.JSON(http.StatusNotFound, row.Err())
		return
	}
	err := row.Scan(&ID, &Mail)
	if err != nil {
		Auth.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, err.Error())
		return
	}
	ref := cryptoutils.GenerateSecureToken()
	Resp.RefreshToken = cryptoutils.EncodeToBase64(ref)
	hashed, err := cryptoutils.HashToken(ref)
	if err != nil {
		Auth.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, err.Error())
		return
	}
	rowIns := db.QueryRow("insert into token (Refresh, UserID) values ($1, $2) returning id", hashed, ID).Scan(&TokenID)
	if rowIns != nil {
		Auth.Logger.Debug(rowIns)
		c.JSON(http.StatusOK, rowIns)
		return
	}
	Auth.Logger.Debug(TokenID)
	Claims := jwt.MapClaims{"ip": c.ClientIP(), "id": TokenID, "mail": Mail}
	access, err := cryptoutils.NewJwt(Claims)
	Resp.AccessToken = access
	if err != nil {
		Auth.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, err.Error())
		return
	}
	Auth.Cache.Set(ref, access)
	c.JSON(http.StatusOK, Resp)
}
func (Auth *AuthUC) Refresh(c *gin.Context) {
	var Req auth.RefreshRequest
	var Hashed string
	db := Auth.DB.Get()
	Auth.Logger.Info("Got refresh in request")
	err := c.ShouldBindBodyWithJSON(&Req)
	if err != nil {
		Auth.Logger.Debug(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	Claims, err := cryptoutils.ParseJwt(Req.AccessToken)
	Mapped := Claims.(jwt.MapClaims)
	ID := int(Mapped["id"].(float64))
	IP := Mapped["ip"].(string)
	Mail := Mapped["mail"].(string)
	Auth.Logger.Debug("Mapped claims and got id ip, mail", ID, IP, Mail)
	if IP != c.ClientIP() {
		Auth.Logger.Debug("IP changed!")
		err = Auth.Notify.NewMail(Mail)
		if err != nil {
			Auth.Logger.Debug(err)
		}
	}
	Auth.Logger.Debug(Req.RefreshToken)
	ClearRef, err := cryptoutils.DecodeFromBase64(Req.RefreshToken)
	if err != nil {
		Auth.Logger.Debug(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	Auth.Logger.Debug("Got Encoded Refresh", ClearRef)
	row := db.QueryRow("select Refresh from token where id = $1", ID)
	if row.Err() != nil {
		Auth.Logger.Debug(err)
		c.JSON(http.StatusNotFound, err)
		return
	}
	err = row.Scan(&Hashed)
	if err != nil {
		Auth.Logger.Debug(err)
		c.JSON(http.StatusNotFound, "No token found")
		return
	}
	Auth.Logger.Debug(Hashed)
	err = cryptoutils.CheckTokenHash(ClearRef, Hashed)
	if err != nil {
		Auth.Logger.Debug(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	Auth.Cache.Set(ClearRef, Req.AccessToken)
	//replaced so expiration timer started again
	_, err = db.Exec("delete from token where id = $1", ID)
	if err != nil {
		Auth.Logger.Debug(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, "You Access token has been refreshed, however refresh token is no more")
}

func NewAuth(DB *postgres.Postgres, Cache *token_cache.TokenCache, Logger *logger.Logger, Nofity mail.Notify) auth.Auth {
	return &AuthUC{DB: DB, Cache: Cache, Logger: Logger}
}
