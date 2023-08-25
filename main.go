package main

import (
	"assignment-1/lib"
	"fmt"
	"os"
)

func main() {
	name, err := lib.ParseByArgs(os.Args)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	student, err := lib.SearchStudentByName(name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	lib.PrintOutput(student)
}