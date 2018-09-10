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
func (c remove) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("insufficient arguments. USAGE: %v", c.usage)
	}
	name := args[0]

	path, err := findFilePathInCatalog(name)
	if err != nil {
		return err
	}
	if err := removeFile(path); err != nil {
		return err
	}

	fmt.Printf("file '%v' removed\n", name)
	return nil
}
