package pkg

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config ConfigRes

func GetEnv(filename string) error {
	err := godotenv.Load(filename)
	return err
}

func InitConfig() error {
	envMode := os.Getenv("ENV")
	if !isValidEnv(envMode) {
		return errors.New("invalid env mode")
	}

	Config.Env = EnvMode(envMode)
	Config.DbUrl = os.Getenv("DATABASE_URL")

	port := os.Getenv("PORT")
	portInt, err := strconv.Atoi(port)

	if err != nil {
		return err
	}

	mailPort := os.Getenv("MAIL_PORT")
	mailPortInt, err := strconv.Atoi(mailPort)

	if err != nil {
		return err
	}

	Config.Port = int(portInt)
	Config.Mail = MailConfig{
		Host:     os.Getenv("MAIL_HOST"),
		Port:     mailPortInt,
		User:     os.Getenv("MAIL_USER"),
		Password: os.Getenv("MAIL_PASS"),
	}

	return nil
}

func isValidEnv(env string) bool {
	if env == string(EnvDevelopment) || env == string(EnvProduction) {
		return true
	}

	return false
}
