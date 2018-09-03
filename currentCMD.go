package mset

import "fmt"

type current struct {
	name        string
	description string
	usage       string
}

var CurrentCMD = current{
	name:        "current",
	description: "Shows the current file",
	usage:       "mset current",
}

func (c current) Name() string        { return c.name }
func (c current) Description() string { return c.description }
func (c current) Usage() string       { return c.usage }

func (c current) Run([]string) error {
	fmt.Println("running Current command")
	return nil
}
