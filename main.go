package main

import (
	"bufio"
	"fmt"
	"gopkg.in/urfave/cli.v1" // imports as package "cli"
	"log"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "tcsv"
	app.Usage = "transpose stdin and join by comma"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "separator, s",
			Value: ",",
			Usage: "separator to join strings",
		},
	}
	app.Action = func(c *cli.Context) error {
		stdin := bufio.NewScanner(os.Stdin)
		values := make([]string, 0)
		for stdin.Scan() {
			text := stdin.Text()
			values = append(values, text)
		}
		fmt.Println(strings.Join(values, c.String("separator")))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
