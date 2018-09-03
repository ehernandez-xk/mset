package main

import (
	"fmt"
	"os"

	"github.com/ehernandez-xk/mset"
)

// Command represents the behaviour of a command
type Command interface {
	Name() string
	Usage() string
	Description() string
	Run([]string) error
}

func help(commands []Command) {
	fmt.Println("VERSION:")
	fmt.Printf("\t%v\n", mset.Version)
	fmt.Println("COMMANDS:")
	for _, c := range commands {
		fmt.Printf("%v\t %v\n", c.Name(), c.Description())
		fmt.Printf("\t %v\n", c.Usage())
	}
	fmt.Println("help\t Show this help")
}

func main() {

	commands := []Command{
		mset.AddCMD,
		mset.SetCMD,
		mset.CurrentCMD,
		mset.InitCMD,
		mset.ListCMD,
		mset.RemoveCMD,
	}

	if len(os.Args) == 1 || os.Args[1] == "help" || os.Args[1] == "--help" {
		help(commands)
		return
	}
	for _, cmd := range commands {
		currentCMD := os.Args[1]
		if cmd.Name() == currentCMD {
			err := cmd.Run(os.Args[2:])
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			// add was success set to current
			if currentCMD == "add" {
				mset.SetCMD.Run(os.Args[2:])
			}
			return
		}
	}
	// no valid argument
	help(commands)
	os.Exit(1)
}
