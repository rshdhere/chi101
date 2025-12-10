// Package config deals with all the database logic
package config

type DBSecret struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

const (
	secretNameProd    = "prod/gym/postgres"
	secretNameStaging = "stagging/gym/postgres"

	secretNameEnvVar = "DB_SECRET_NAME"

	appEnvVar     = "APP_ENV"
	localEnvVal   = "local"
	stagingEnvVal = "staging"
	prodEnvVal    = "prod"
)
