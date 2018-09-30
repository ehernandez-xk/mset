package mset

import "fmt"

type current struct {
	name        string
	description string
	usage       string
}

var CurrentCMD = current{
	name:        "current",
	description: "Show the name of the current settings.xml file",
	usage:       "mset current",
}

func (c current) Name() string        { return c.name }
func (c current) Description() string { return c.description }
func (c current) Usage() string       { return c.usage }

// Run command finds and shows the current settings.xml file in the catalog
// the current file is placed in .m2/settings.xml
func (c current) Run([]string) error {
	name, err := getCurrentName()
	if err != nil {
		return err
	}
	fmt.Println(name)
	return nil
}
