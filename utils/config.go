package utils

import (
	"fr33d0mz/moneyflowx/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

var AppSettings models.Settings

func PutAdditionalSettings() {
	AppSettings.AppParams.LogDebug = "./logs/debug.log"
	AppSettings.AppParams.LogInfo = "./logs/info.log"
	AppSettings.AppParams.LogWarning = "./logs/warning.log"
	AppSettings.AppParams.LogError = "./logs/error.log"

	AppSettings.AppParams.LogCompress = true
	AppSettings.AppParams.LogMaxSize = 10
	AppSettings.AppParams.LogMaxAge = 100
	AppSettings.AppParams.LogMaxBackups = 100
	AppSettings.AppParams.AppVersion = "v1"
}

func ReadSettings() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Couldn't open config file. Error is: ", err.Error())
	}

	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Failed to read `config.yaml` file. Error is: ", err.Error())
	}

	setup()
	setupPostgres()
}

func setup() {
	AppSettings.AppParams.ServerName = "localhost"
	AppSettings.AppParams.ServerURL = "127.0.0.1"
	AppSettings.AppParams.PortRun = viper.GetString("port")
	AppSettings.AppParams.SecretKey = os.Getenv("SECRET_API")
	AppSettings.AppParams.TokenTTL = cast.ToInt(os.Getenv("TOKEN_LIFESPAN"))
	AppSettings.AppParams.JWTIssuer = os.Getenv("JWT_ISSUER")
}

func setupPostgres() {
	AppSettings.PostgresParams.Server = viper.GetString("db.host")
	AppSettings.PostgresParams.Port = viper.GetString("db.port")
	AppSettings.PostgresParams.User = viper.GetString("db.username")
	AppSettings.PostgresParams.Database = viper.GetString("db.dbname")
	AppSettings.PostgresParams.SSLMode = viper.GetString("db.sslmode")
	AppSettings.PostgresParams.Password = os.Getenv("DB_PASSWORD")
}
