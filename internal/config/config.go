package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read configuration file: %s", err)
	}
}

const (
	DatabaseHost     = "postgres.host"
	DatabasePort     = "postgres.port"
	DatabaseUsername = "postgres.username"
	DatabasePassword = "postgres.password"
	DatabaseName     = "postgres.database"
	DatabaseSslMode  = "postgres.disable"

	BaseUrl = "base_url"

	ControlPanelUsername = "control_panel.username"
	ControlPanelPassword = "control_panel.password"
)
