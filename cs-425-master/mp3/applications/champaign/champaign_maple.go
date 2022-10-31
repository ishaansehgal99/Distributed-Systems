package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// MapleJuice application: Number of active, residential, and mailable addresses per zipcode in Champaign, IL
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
		dataPoints := SplitAtCommas(nextLine)

		if len(dataPoints) < 41 {
			continue
		}

		zipCope := dataPoints[40]
		isResidential := dataPoints[13] == "Y"
		isMailable := dataPoints[14] == "Y"
		isActive := dataPoints[18] == "Active"

		if isResidential && isMailable && isActive {
			fmt.Println(zipCope, 1)
		}
	}
}

// SplitAtCommas - https://stackoverflow.com/questions/59297737/go-split-string-by-comma-but-ignore-comma-within-double-quotes
func SplitAtCommas(s string) []string {
	res := []string{}
	var beg int
	var inString bool

	for i := 0; i < len(s); i++ {
		if s[i] == ',' && !inString {
			res = append(res, s[beg:i])
			beg = i + 1
		} else if s[i] == '"' {
			if !inString {
				inString = true
			} else if i > 0 && s[i-1] != '\\' {
				inString = false
			}
		}
	}
	return append(res, s[beg:])
}

func removeSpecialChars(input string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	return reg.ReplaceAllString(input, "")
}
