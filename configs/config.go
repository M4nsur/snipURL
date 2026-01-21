package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db DbConfig
	Auth AuthConfig 
}

type DbConfig struct {
	DNS string
}

type AuthConfig struct {
	Secret string
}


func LoadConfig () *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err, "err .env")
		return nil
	}
	return &Config{
		Db: DbConfig{
			DNS: os.Getenv("DSN"),
			
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
}