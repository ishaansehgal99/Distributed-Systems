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
		dataPoints := strings.Split(nextLine, " ")

		A := dataPoints[0]
		count := dataPoints[1]

		fmt.Println(1, A+","+count)
	}
}
