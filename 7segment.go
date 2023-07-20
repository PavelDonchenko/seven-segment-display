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

	var out []string

	for row := 0; row < 5+2*(scale-1); row++ {
		var line string
		var outputLine string
		for _, item := range matrixNumber {
			line += item[row] + "  "
			outputLine = strings.NewReplacer("0", " ", "1", "|", "2", "-", "3", "*", "4", "_").Replace(line)
			//fmt.Print(item[row], " ")
		}
		out = append(out, outputLine)
	}

	res := strings.Join(out, "\n")

	return replaceDuplicateStar(res)
}

func stringToMatrix(num string, factor int) [][]string {
	numbers := map[string][]string{
		"0": {"020", "101", "000", "101", "020"},
		"1": {"000", "001", "000", "001", "000"},
		"2": {"020", "001", "020", "100", "020"},
		"3": {"020", "001", "020", "001", "020"},
		"4": {"000", "101", "020", "001", "000"},
		"5": {"020", "100", "020", "001", "020"},
		"6": {"020", "100", "020", "101", "020"},
		"7": {"020", "001", "000", "001", "000"},
		"8": {"020", "101", "020", "101", "020"},
		"9": {"020", "101", "020", "001", "000"},
		".": {"000", "000", "000", "000", "030"},
		"-": {"000", "000", "040", "000", "000"},
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

	return result
}

func replaceDuplicateStar(input string) string {
	characters := []byte(input)

	var result []byte

	starCount := 0
	for _, ch := range characters {
		if ch == '*' || ch == '_' {
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
