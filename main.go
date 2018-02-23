package main

import (
	"os"

	"gopkg.in/urfave/cli.v1"
	"sort"
	"fmt"
	"strconv"
	src "github.com/AlexeyArno/go-comand-helper/src"
	"io/ioutil"
	"log"
)


func createProject(name string, path string, templatePath string){
	fname := path+`\`+name
	os.MkdirAll(fname, os.ModePerm)
	fmt.Println("Directory :", fname, " created")

	files, err := ioutil.ReadDir(templatePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir(){
			src.CopyFile(templatePath+`\`+f.Name(), fname+`\`+f.Name())
			fmt.Println("Created file: ",f.Name())
		}
	}

	fmt.Println(fname," ready!!!")

}

func main() {
	app := cli.NewApp()
	var count int
	var embarcadero bool
	var dev bool
	var path string



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
				cli.StringFlag{
					Name: "path, p",
					Value: "C:",
					Destination: &path,
					Usage: "Path to directory for projects",
				},
			},
			Action:  func(c *cli.Context) error {
				fmt.Println("Count: ", count)
				fmt.Println("Type: ")
				if dev {
					for i:= 1;i<=count;i++{
						createProject(strconv.FormatInt(int64(i),10), path, "./examples/dev");
					}
					fmt.Println("All Work Done!!!")
				} else if embarcadero{
					for i:= 1;i<=count;i++{
						createProject(strconv.FormatInt(int64(i),10), path, "./examples/emb");
					}
					fmt.Println("All Work Done!!!")
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