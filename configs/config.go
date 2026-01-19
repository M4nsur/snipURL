package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db DbConfig
}

type DbConfig struct {
	DNS string
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
	}
}