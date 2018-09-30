package mset

import "fmt"

type add struct {
	name        string
	description string
	usage       string
}

var AddCMD = add{
	name:        "add",
	description: "Add a new entry in the catalog, copying <file> in a new settings.xml",
	usage:       "mset add <name> <file>",
}

func (c add) Name() string        { return c.name }
func (c add) Description() string { return c.description }
func (c add) Usage() string       { return c.usage }

// Run command adds a new settings.xml file to the catalog
// it does some checks before save it as valid name
func (c add) Run(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("insufficient arguments. USAGE: %v", c.usage)
	}
	name := args[0]
	sourceFile := args[1]

	if !isValidName(name) {
		return fmt.Errorf("invalid name alphanumeric and '-' allowed. e.g. 'danta'")
	}
	if isNameAlreadyTaken(name) {
		return fmt.Errorf("name '%v' already taken", name)
	}

	if err := fileExist(sourceFile); err != nil {
		return err
	}

	if err := saveFileToCatalog(name, sourceFile); err != nil {
		return err
	}
	fmt.Printf("file '%v' saved\n", name)
	return nil
}
