package mset

import "fmt"

type list struct {
	name        string
	description string
	usage       string
}

var ListCMD = list{
	name:        "list",
	description: "Lists the files in the catalog",
	usage:       "mset list",
}

func (c list) Name() string        { return c.name }
func (c list) Description() string { return c.description }
func (c list) Usage() string       { return c.usage }

func (c list) Run([]string) error {
	fmt.Println("danta")
	fmt.Println("lifeway")
	fmt.Println("empty")
	return nil
}
