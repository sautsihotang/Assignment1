package main

import (
	"assignment1/data"
	"os"
)

func main() {

	input := os.Args[1]

	data.SearchStudents(input)
}
