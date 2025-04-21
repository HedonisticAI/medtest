package app

import (
	"medtest/config"
	add_usecase "medtest/internal/add/usecase"
	auth_usecase "medtest/internal/auth/usecase"
	httpserver "medtest/pkg/http_server"
	"medtest/pkg/logger"
	"medtest/pkg/mail"
	"medtest/pkg/postgres"
	"medtest/pkg/token_cache"
)

type App struct {
	Server *httpserver.HttpServer
	Logger logger.Logger
	DB     *postgres.Postgres
}

func NewApp(Logger *logger.Logger, Config config.Config) (*App, error) {
	Server := httpserver.NewServer(&Config)
	DB, err := postgres.NewDB(Config)
	if err != nil {
		Logger.Debug(err)
		return nil, err
	}
	Logger.Debug("DB connected")
	Cache := token_cache.NewCache(&Config)
	Nofity := mail.NewNotify(&Config)
	Auth := auth_usecase.NewAuth(DB, Cache, Logger, *Nofity)
	Add := add_usecase.NewAdd(DB, Logger)
	Server.MapPost("/add", Add.AddUser)
	Server.MapGet("/signIn", Auth.SignIn)
	Server.MapGet("/refresh", Auth.Refresh)
	return &App{DB: DB, Logger: *Logger, Server: Server}, nil
}

func (App *App) Run() {
	App.Server.Run()
}
