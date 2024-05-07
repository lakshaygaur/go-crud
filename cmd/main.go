package main

import (
	"fmt"
	"gocrud/config"
	"gocrud/server"
	"gocrud/storage"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "start",
		Usage: "Start Server",
		Action: func(*cli.Context) error {
			fmt.Println("go-crud start server")
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Starting server...")
	storage.InitDatabase()
	server.StartServer()
	config.SetVars()

	defer storage.HandleDBclose()

}
