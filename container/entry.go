package container

import (
	"encoding/json"

	"github.com/caarlos0/env/v11"
	_ "github.com/joho/godotenv/autoload"
)

// AppConfig
type AppConfig struct {
	Port        int64  `env:"PORT" envDefault:"9090"`
	FrontendDir string `env:"FRONTEND_DIR"`
	DocDir      string `env:"DOC_DIR"`
	Database    string `env:"DATABASE,file,expand"`
	Users       string `env:"USERS,file,expand"`
	Domain      string `env:"DOMAIN"`
	Salt        string `env:"SALT"`
	LoginKey    string `env:"LOGIN_KEY"`
}

var (
	cfg *AppConfig
)

func Initialize() error {
	cfg = &AppConfig{}
	if err := env.Parse(cfg); err != nil {
		return err
	}
	databases := map[string]string{}
	if err := json.Unmarshal([]byte(cfg.Database), &databases); err != nil {
		panic("json unmarshal error:" + err.Error())
	}

	for name, dsn := range databases {
		InitializeDatabase(name, dsn)
	}
	return nil
}

func GetAppConfig() *AppConfig {
	return cfg
}
