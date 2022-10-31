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
	key := args[2]

	file, err := os.Open(filename)

	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	values := make([]string, 0)
	for scanner.Scan() {
		nextLine := scanner.Text()

		if strings.TrimSpace(nextLine) != "" {
			values = append(values, nextLine)
		}
	}

	result := len(values)

	fmt.Println(key, result)
}
