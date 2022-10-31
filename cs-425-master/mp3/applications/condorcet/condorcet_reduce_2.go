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

	counts := make(map[string]int)

	for _, line := range values {
		// Example: line = B,12,0 (A beat B 12 times)
		// Exapmle: line = B,0,15 (B beat A 15 times)
		dataPoints := strings.Split(line, ",")

		if len(dataPoints) < 3 {
			continue
		}

		opp := dataPoints[0]
		mycount, _ := strconv.Atoi(dataPoints[1])
		theircount, _ := strconv.Atoi(dataPoints[2])

		if _, ok := counts[opp]; !ok {
			counts[opp] = 0
		}

		counts[opp] += mycount - theircount
	}

	numEntries := 0
	for _, v := range counts {
		if v > 0 {
			numEntries++
		}
	}

	fmt.Println(key, numEntries)
}
