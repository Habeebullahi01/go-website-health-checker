package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Solp",
		Usage: "A tool for checking website health.",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "The domain or url of the website to check.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "The port to check.",
				Required: false,
			},
		},
		Action: func(ctx *cli.Context) error {
			port := ctx.String("port")
			if port == "" {
				port = "80"
			}
			status := Check(ctx.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
