package composites

import (
	"github.com/joho/godotenv"
	"github.com/m-a-r-a-t/go-rest-api/pkg/pg_database"
	"log"
	"os"
)

func initEnvVariables() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

type Config struct {
	DatabaseConf pg_database.DatabaseConfig
}

func Settings() *Config {
	initEnvVariables()

	return &Config{
		DatabaseConf: pg_database.DatabaseConfig{
			Host: getEnv("DATABASE_HOST_APP", ""),
			Port: getEnv("DATABASE_PORT_APP", ""),
			User: getEnv("DATABASE_USER_APP", ""),
			Pass: getEnv("DATABASE_PASS_APP", ""),
			Name: getEnv("DATABASE_NAME_APP", ""),
		},
	}

}
