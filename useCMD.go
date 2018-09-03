package mset

import "fmt"

type set struct {
	name        string
	description string
	usage       string
}

var SetCMD = set{
	name:        "set",
	description: "Set the <name> as current, copying the settings.xml in ~.m2/",
	usage:       "mset set <name>",
}

func (c set) Name() string        { return c.name }
func (c set) Description() string { return c.description }
func (c set) Usage() string       { return c.usage }

func (c set) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("insufficient arguments. USAGE: %v", c.usage)
	}
	name := args[0]

	fmt.Printf("'%v' set to current\n", name)
	return nil
}
