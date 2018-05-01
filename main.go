package main

import (
	"fmt"
	"github.com/fatih/color"
	rsp "github.com/ru-lai/randomsp"
	"github.com/urfave/cli"
	"os"
)

func main() {
	errCol := color.New(color.FgRed).SprintFunc()

	app := cli.NewApp()
	app.Name = "randomsp-cli"
	app.Version = "0.1.0"
	app.Usage = "a cli tool to randomly query stocks from a few major indices"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Tyler Boright",
			Email: "ru.lai.development@gmail.com",
		},
	}

	app.Action = func(c *cli.Context) error {
		color.Magenta("Get a random stock pick!")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "randomStock",
			Aliases: []string{"rs"},
			Usage:   "Return a random stock from any of the supported indices",
			Action: func(c *cli.Context) error {
				stock, err := rsp.GetRandomIndexStock()
				if err != nil {
					return fmt.Errorf(errCol("The randomsp failed with following error: %s. \nCheck your internet connection and try again."), err)
				}

				color.Magenta(fmt.Sprintf("Your randomly selected stock symbol is:\n%s", stock.Symbol))
				color.Magenta(fmt.Sprintf("%s is a stock on the %s stock index", stock.Symbol, stock.Index))

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		return
	}
}
