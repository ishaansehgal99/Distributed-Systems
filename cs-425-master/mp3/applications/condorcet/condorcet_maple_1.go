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
		line := strings.Split(nextLine, " ")

		if len(line) < 2 {
			continue
		}

		order := line[1]

		dataPoints := strings.Split(order, ",")

		if len(dataPoints) < 3 {
			continue
		}

		candidate1 := dataPoints[0]
		candidate2 := dataPoints[1]
		candidate3 := dataPoints[2]

		fmt.Println(candidate1+candidate2, 1)
		fmt.Println(candidate1+candidate3, 1)
		fmt.Println(candidate2+candidate3, 1)
	}
}
