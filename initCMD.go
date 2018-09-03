package mset

import "fmt"

type initc struct {
	name        string
	description string
	usage       string
}

var InitCMD = initc{
	name:        "init",
	description: fmt.Sprintf("Creates an empty directory %v to store the files", getCatalogPath()),
	usage:       "mset init",
}

func (c initc) Name() string        { return c.name }
func (c initc) Description() string { return c.description }
func (c initc) Usage() string       { return c.usage }

func (c initc) Run([]string) error {
	err := createCatalogDirectory()
	if err != nil {
		return err
	}
	fmt.Println("mset initialized")
	return nil
}
