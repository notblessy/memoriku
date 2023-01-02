package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
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

// CloudinaryCloudName :nodoc:
func CloudinaryCloudName() string {
	return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

// CloudinaryAPIKey :nodoc:
func CloudinaryAPIKey() string {
	return os.Getenv("CLOUDINARY_API_KEY")
}

// CloudinaryAPISecret :nodoc:
func CloudinaryAPISecret() string {
	return os.Getenv("CLOUDINARY_API_SECRET")
}

// CloudinaryUploadFolder :nodoc:
func CloudinaryUploadFolder() string {
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}
