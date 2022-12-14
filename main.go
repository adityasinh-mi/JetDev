package main

import (
	"context"
	"flag"
	"jetdev-task/model"
	"jetdev-task/routers/api"
	"jetdev-task/shared/config"
	"jetdev-task/shared/database"
	"jetdev-task/shared/log"
	"jetdev-task/shared/utils"
	"os"
)

//Execution starts from main function
func main() {
	configFile := flag.String("c", "", "configuration file without extension. For config.toml then put \" -c config-development\"")
	flag.Parse()

	pwd, err := os.Getwd()
	if err != nil {
		os.Exit(1)
	}
	var cf config.IConfig
	if *configFile == "" {
		cf = config.NewRealtimeConfig("config", pwd)
	} else {
		cf = config.NewConfig(*configFile)
	}

	log.Init("app", cf.AppVersion(), cf.Info().Path, cf.Info().Level, cf.Info().MaxAge)
	log.GetLog().Info("", "App service start! %s", cf.AppVersion())

	database.Init(cf)
	log.GetLog().Info("", "DB connected")
	model.AutoMigrate()
	rt := api.NewRouter(cf)
	rt.Setup()

	go rt.Run()

	utils.GracefulStop(log.GetLog(), func(ctx context.Context) error {
		var err error
		if err = rt.Close(ctx); err != nil {
			return err
		}
		if err = database.Close(); err != nil {
			return err
		}
		return nil
	})
}
