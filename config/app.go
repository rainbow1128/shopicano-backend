package config

import (
	"github.com/spf13/viper"
)

type LogLevel string

const (
	Debug    LogLevel = "debug"
	Info     LogLevel = "info"
	Warning  LogLevel = "warning"
	Critical LogLevel = "critical"
)

// Application holds the application configuration
type Application struct {
	Base          string
	Port          int
	LogLevel      LogLevel
	BackendUrl    string
	FrontStoreUrl string
	DashboardUrl  string
	JWTKey        string
}

// app is the default application configuration
var app Application

// App returns the default application configuration
func App() *Application {
	return &app
}

// LoadApp loads application configuration
func LoadApp() {
	mu.Lock()
	defer mu.Unlock()

	app = Application{
		Base:          viper.GetString("app.host"),
		Port:          viper.GetInt("app.port"),
		LogLevel:      LogLevel(viper.GetString("app.log_level")),
		BackendUrl:    viper.GetString("app.backend_url"),
		FrontStoreUrl: viper.GetString("app.front_store_url"),
		DashboardUrl:  viper.GetString("app.dashboard_url"),
		JWTKey:        viper.GetString("app.jwt_key"),
	}
}
