package mset

import "fmt"

type remove struct {
	name        string
	description string
	usage       string
}

var RemoveCMD = remove{
	name:        "remove",
	description: "Remove the <name> from the catalog",
	usage:       "mset remove <name>",
}

func (c remove) Name() string        { return c.name }
func (c remove) Description() string { return c.description }
func (c remove) Usage() string       { return c.usage }

// Run command removes the file in the catalog
func (c remove) Run([]string) error {
	fmt.Println("file removed")
	return nil
}
