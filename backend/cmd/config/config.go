// Package config it's an app config
package config

import "os"

// NewAppConfig конструктор
func NewAppConfig() *AppConfig {
	return &AppConfig{
		appPort:        os.Getenv("PORT"),
		petStoragePath: os.Getenv("PET_STORAGE_PATH"),
	}
}

// AppConfig конфиг
type AppConfig struct {
	appPort        string
	petStoragePath string
}

// GetAppPort returns backend port
func (ac AppConfig) GetAppPort() string {
	return ac.appPort
}

// GetAppPort returns storage path
func (ac *AppConfig) GetPetStoragePath() string {
	return ac.petStoragePath
}
