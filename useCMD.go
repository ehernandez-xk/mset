package mset

import "fmt"

type use struct {
	name        string
	description string
	usage       string
}

var UseCMD = use{
	name:        "use",
	description: "Changes the current file",
	usage:       "mset use <name>",
}

func (c use) Name() string        { return c.name }
func (c use) Description() string { return c.description }
func (c use) Usage() string       { return c.usage }

func (c use) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("insufficient arguments. USAGE: %v", c.usage)
	}
	name := args[0]

	fmt.Printf("'%v' set to current\n", name)
	return nil
}
