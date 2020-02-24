package main

import (
	"log"
	"os"

	"github.com/FZambia/viper-lite"

	"lrn/configs"
	"lrn/internal/services/logger"
)

func main() {
	if err := run(); err != nil {
		log.Println("error:", err)
		os.Exit(1)
	}
}

func run() error {

	configs.InitConfig()

	logger := logger.NewZap(viper.GetBool(configs.ConfigDebug))

	return nil
}
