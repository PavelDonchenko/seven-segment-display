package main

import (
	"fmt"
	"strings"
)

type Receiver interface {
	GetNumber() (string, int, error)
}

type Printer interface {
	PrintSevenSegmentDisplay(output string)
}

type Application struct {
	Receiver
	Printer
}

func NewApplication(receiver Receiver, printer Printer) *Application {
	return &Application{
		Receiver: receiver,
		Printer:  printer,
	}
}

func (a *Application) ConvertToSevenSegment(number string, scale int) string {
	matrixNumber := stringToMatrix(number, scale)

	count := 4 + 2*(scale-2)

	var out []string

	for row := 0; row <= count; row++ {
		var line string
		var outputLine string
		for _, item := range matrixNumber {
			line += item[row] + "  "
			if row != 0 && row != count/2 && row != count {
				outputLine = strings.NewReplacer("0", " ", "1", "|", "2", "-", "3", "*", "4", " ").Replace(line)
			} else {
				outputLine = strings.NewReplacer("0", " ", "1", "|", "2", "-", "3", "*", "4", "_").Replace(line)
			}
		}
		out = append(out, outputLine)
	}

	filtered := make([]string, 0)

	for i, s := range out {
		if strings.Contains(s, "_") {
			filtered = append(filtered, s)
		} else if i%2 != 0 {
			filtered = append(filtered, s)
		}
	}

	res := strings.Join(filtered, "\n")

	return replaceDuplicateStar(res)
}

func stringToMatrix(num string, factor int) [][]string {
	numbers := map[string][]string{
		"0": {"040", "101", "141"},
		"1": {"000", "001", "001"},
		"2": {"040", "041", "140"},
		"3": {"040", "041", "041"},
		"4": {"000", "141", "001"},
		"5": {"040", "140", "041"},
		"6": {"040", "140", "141"},
		"7": {"040", "001", "001"},
		"8": {"040", "141", "141"},
		"9": {"040", "141", "001"},
		".": {"000", "000", "030"},
	}

	var result [][]string
	for _, digit := range strings.Split(num, "") {
		result = append(result, scale(numbers[digit], factor))
	}

	return result
}

func scale(code []string, factor int) []string {
	if factor == 1 {
		return code
	}

	var result []string
	for _, line := range code {
		widen := fmt.Sprintf("%c%s%c", line[0], strings.Repeat(line[1:len(line)-1], factor), line[len(line)-1])
		result = append(result, widen)
	}

	for i := len(result) - 2; i >= 0; i -= 2 {
		stretched := make([]string, factor)
		for j := 0; j < factor; j++ {
			stretched[j] = result[i]
		}
		result = append(result[:i], append(stretched, result[i+1:]...)...)
	}

	for i := 1; i < factor; i++ {
		lastElement := result[len(result)-1]
		if strings.Contains(lastElement, "3") {
			lastElement = strings.ReplaceAll(lastElement, "3", "0")
			result = append([]string{lastElement}, result...)
		} else {
			result = append(result, lastElement)
		}
	}

	return result
}

func replaceDuplicateStar(input string) string {
	characters := []byte(input)

	var result []byte

	starCount := 0
	for _, ch := range characters {
		if ch == '*' {
			starCount++
			if starCount <= 1 {
				result = append(result, ch)
			} else {
				result = append(result, ' ')
			}
		} else {
			result = append(result, ch)
			starCount = 0
		}
	}

	return string(result)
}
