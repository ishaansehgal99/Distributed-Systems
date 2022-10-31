package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	filename := args[1]

	file, err := os.Open(filename)

	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nextLine := scanner.Text()
		// Example: line = AC 13
		line := strings.Split(nextLine, " ")
		A, B, count := line[0][0], line[0][1], line[1]

		fmt.Println(string(A), string(B)+","+count+",0")
		fmt.Println(string(B), string(A)+",0,"+count)
	}
}
