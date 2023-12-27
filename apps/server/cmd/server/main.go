package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/acerohernan/meet/pkg/config"
	"github.com/acerohernan/meet/pkg/service"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Meet SFU",
		Usage: "Selective Forwading Unit for video conferencing app",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "config",
				Usage:    "path to yaml config file",
				Required: true,
			},
		},
		Action: startServer,
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}

func startServer(ctx *cli.Context) error {
	configFilePath := ctx.String("config")

	if configFilePath == "" {
		return errors.New("yaml config file required")
	}

	configFile, err := os.ReadFile(configFilePath)

	if err != nil {
		return err
	}

	conf, err := config.NewConfig(string(configFile))

	if err != nil {
		return err
	}

	server := service.NewServer(conf)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-sigChan
		fmt.Print("exit requersted, shutting down. ", "signal: ", sig)
		server.Stop()
	}()

	server.Start()

	return nil
}
