package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

// LoadENV :nodoc:
func LoadENV() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(fmt.Sprintf("ERROR: %s", err))
	}

	return err
}

// ENV :nodoc:
func ENV() string {
	return os.Getenv("ENV")
}

// HTTPPort :nodoc:
func HTTPPort() string {
	return os.Getenv("PORT")
}

// MysqlHost :nodoc:
func MysqlHost() string {
	return os.Getenv("DB_HOST")
}

// MysqlUser :nodoc:
func MysqlUser() string {
	return os.Getenv("DB_USER")
}

// MysqlPassword :nodoc:
func MysqlPassword() string {
	return os.Getenv("DB_PASS")
}

// MysqlDBName :nodoc:
func MysqlDBName() string {
	return os.Getenv("DB_NAME")
}

// JWTSecret :nodoc:
func JWTSecret() string {
	return os.Getenv("JWT_SECRET")
}
