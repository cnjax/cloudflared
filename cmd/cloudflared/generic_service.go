//go:build !windows && !darwin && !linux

package main

import (
	"fmt"
	"github.com/joho/godotenv"
	cli "github.com/urfave/cli/v2"
	"os"
)

func runApp(app *cli.App, graceShutdownC chan struct{}) {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Can't find .env")
	}
	tokenstr := os.Getenv("TOKEN")
	if tokenstr == "" {
		fmt.Printf("empty token")
		tokenstr = "eyJhIjoiZTVmYzUzYjhmNmQ5N2QxYzhkODAwNjFiYWI4OWFmYWQiLCJ0IjoiMjFmMzYyNDEtNDMxYi00ZjcxLTk2MzktYzkxOWI0NTI4ZTViIiwicyI6Ik9USmlaV0l5WlRndE56RTRPQzAwWldOa0xUazJNVGd0WlRWaFlUazNNV1V3WW1FeiJ9"
	}
	//"--protocol",
	//	"http2",
	args := []string{
		"",
		"tunnel",
		"--loglevel",
		"fatal",
		"--no-autoupdate",
		"run",
		"--token",
		tokenstr}

	app.Run(args)
	//app.Run(os.Args)
}
