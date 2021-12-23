package main

import (
	"github.com/isalikov/goservice/internal/env"
	"github.com/isalikov/goservice/internal/server"
	"github.com/isalikov/goservice/internal/utils"
	"log"
)

func main()  {
	envConfig := &env.Config{}

	err := env.Parse(envConfig)
	if err != nil {
		log.Fatalln(err, "Parsing environment")
	}

	logger := &utils.Logger{Debug: envConfig.Env != "release"}
	logger.Init(envConfig.SentryDsn, envConfig.Release)

	err = server.Serve(envConfig)
	if err != nil {
		logger.Error(err, "gRPC")
	}
}
