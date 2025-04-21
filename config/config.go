package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DB        DB
	Cache     Cache
	Server    Server
	Mailer    Mailer
	BuildType string
}

type Mailer struct {
	From     string
	PWD      string
	SMTPhost string
	SMTPport string
}

type Cache struct {
	Duration int
}

type DB struct {
	DBName string
	DBPort string
	DBHost string
	DBPwd  string
	DBUser string
}

type Server struct {
	Port string
}

func NewServer() *Server {
	Port, exist := os.LookupEnv("SERVER_PORT")
	if !exist {
		return nil
	}
	return &Server{Port: Port}
}

func NewCache() *Cache {
	TimeStr, exist := os.LookupEnv("TOKEN_EXP")
	if !exist {
		return nil
	}
	Time, err := strconv.Atoi(TimeStr)
	if err != nil {
		return nil
	}
	return &Cache{Duration: Time}
}
func NewMail() *Mailer {
	SMTPport, exist := os.LookupEnv("SMTP_PORT")
	if !exist {
		return nil
	}
	From, exist := os.LookupEnv("MAIL_ADDRESS")
	if !exist {
		return nil
	}
	Pwd, exist := os.LookupEnv("MAIL_PWD")
	if !exist {
		return nil
	}
	SMTPhost, exist := os.LookupEnv("SMTP_HOST")
	if !exist {
		return nil
	}
	return &Mailer{PWD: Pwd, From: From, SMTPhost: SMTPhost, SMTPport: SMTPport}
}
func NewDB() *DB {
	DBHost, exist := os.LookupEnv("DB_HOST")
	if !exist {
		return nil
	}
	DBPort, exist := os.LookupEnv("DB_PORT")
	if !exist {
		return nil
	}
	DBUser, exist := os.LookupEnv("DB_USER")
	if !exist {
		return nil
	}
	DBPwd, exist := os.LookupEnv("DB_PWD")
	if !exist {
		return nil
	}
	DBName, exist := os.LookupEnv("DB_NAME")
	if !exist {
		return nil
	}
	return &DB{DBName: DBName, DBPort: DBPort, DBHost: DBHost, DBPwd: DBPwd, DBUser: DBUser}
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
		return nil
	}
	Type, exist := os.LookupEnv("BUILD_TYPE")
	if !exist {
		return nil
	}
	DB := NewDB()
	if DB == nil {
		return nil
	}

	Server := NewServer()
	if Server == nil {
		return nil
	}
	Cache := NewCache()
	if Cache == nil {
		return nil
	}
	Mailer := NewMail()
	if Mailer == nil {
		return nil
	}
	return &Config{DB: *DB, Cache: *Cache, Server: *Server, Mailer: *Mailer, BuildType: Type}
}
