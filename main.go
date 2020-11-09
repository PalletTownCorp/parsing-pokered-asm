package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Function to parse file to search string and blank line are deleted
func parseFileToSearchString(file *string, pattern *string) {
	data, err := ioutil.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n\t")
	for index, line := range lines {
		if strings.Contains(line, *pattern) {
			fmt.Printf("%d : %s\r\n", index, line)
		}
	}
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Error : arguments are needed")
	} else {
		file := args[1]
		pattern := ""
		if len(args) == 3 {
			pattern = args[2]
		}
		parseFileToSearchString(&file, &pattern)
	}
}
