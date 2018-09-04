package mset

import "fmt"

type list struct {
	name        string
	description string
	usage       string
}

var ListCMD = list{
	name:        "list",
	description: "List files available in the catalog",
	usage:       "mset list",
}

func (c list) Name() string        { return c.name }
func (c list) Description() string { return c.description }
func (c list) Usage() string       { return c.usage }

func (c list) Run([]string) error {
	files, err := filesInCatalog()
	if err != nil {
		return err
	}
	current, _ := getCurrentName()
	for k := range files {
		fmt.Print(k)
		if k == current {
			fmt.Print("\t(current)")
		}
		fmt.Println("")
	}

	return nil
}
