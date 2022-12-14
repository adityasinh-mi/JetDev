package config

import (
	"github.com/spf13/viper"
)

type App struct {
	Port string // App.Port
}

func (r *RealtimeConfig) reloadApp() {
	r.app.Port = viper.GetString("App.Port")
	r.testApp()
}

func (r *RealtimeConfig) testApp() {
	testEmptyString(r.app, "Port")
}
