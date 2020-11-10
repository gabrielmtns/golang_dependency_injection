package main

import (
	"fmt"
	"os"
	"./database"
	"./runner"
)

func main() {
	runner := runner.NewRunner(database.NewFileDatabase("test.txt"))
	if err := runner.Run(os.Stdout, os.Args); err !=nil {
		fmt.Println(err)
	}

}
