package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	values := make([]string, 0)
	for scanner.Scan() {
		nextLine := scanner.Text()

		if strings.TrimSpace(nextLine) != "" {
			values = append(values, nextLine)
		}
	}

	maxCount := -1
	maxNames := make([]string, 0)

	for _, line := range values {
		lineSplit := strings.Split(line, ",")

		if len(lineSplit) < 2 {
			continue
		}

		name := lineSplit[0]
		count, _ := strconv.Atoi(lineSplit[1])
		if count > maxCount {
			maxCount = count
			maxNames = make([]string, 0)
			maxNames = append(maxNames, name)
		} else if count == maxCount {
			maxNames = append(maxNames, name)
		}
	}

	returnNames := strings.Join(maxNames[:], ",")

	fmt.Println("Winner(s): " + returnNames)
}
