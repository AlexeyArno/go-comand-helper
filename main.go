package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"
	"sort"
	"fmt"
)

func main() {
	app := cli.NewApp()
	var count int
	var embarcadero bool
	var dev bool



	app.Commands = []cli.Command{
		{
			Name:    "create",
			Aliases: []string{"c"},
			Usage:   "Create All Projects",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "dev, d",
					Usage: "Create projects for dev-c++",
					Destination: &dev,
				},
				cli.BoolFlag{
					Name: "embarcadero, e",
					Usage: "Create projects for embarcadero-studio",
					Destination: &embarcadero,
				},
				cli.IntFlag{
					Name: "count, c",
					Value: 1,
					Destination: &count,
					Usage: "Create projects for embarcadero-studio",
				},
			},
			Action:  func(c *cli.Context) error {
				fmt.Println("Count: ", count)
				fmt.Println("Type: ")
				if dev {
					fmt.Print("Dev c++")
				} else if embarcadero{
					fmt.Print("Embarcadero")
				}else{
					fmt.Print("-")
				}
				return nil
			},
		},
	}

	//sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}