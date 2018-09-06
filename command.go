package mset

import (
	"fmt"
	"os"
)

// Command represents the behaviour of a command
type Command interface {
	Name() string
	Usage() string
	Description() string
	Run([]string) error
}

// print the help
func help(commands []Command) {
	fmt.Printf("VERSION:\n\t%v\nCOMMANDS:\n", Version)
	for _, c := range commands {
		fmt.Printf("%v\t %v\n", c.Name(), c.Description())
		fmt.Printf("\t %v\n", c.Usage())
	}
	fmt.Println("help\t Show this help")
}

// Exec execs the args in the commands
func Exec(args []string) error {
	commands := []Command{
		AddCMD,
		SetCMD,
		CurrentCMD,
		InitCMD,
		ListCMD,
		RemoveCMD,
	}

	if len(os.Args) == 1 || os.Args[1] == "help" || os.Args[1] == "--help" {
		help(commands)
		return nil
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
				SetCMD.Run(os.Args[2:])
			}
			return nil
		}
	}
	help(commands)
	return fmt.Errorf("invalid argument")
}
