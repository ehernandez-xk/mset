package main

import (
	"fmt"
	"os"

	"github.com/ehernandez-xk/mset"
)

func main() {

	err := mset.Exec(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
