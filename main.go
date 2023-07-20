package main

import "fmt"

func main() {
	input := Input{}

	output := Output{}

	app := NewApplication(input, output)

	number, scale, err := app.GetNumber()
	if err != nil {
		fmt.Println(err)
		return
	}

	result := app.ConvertToSevenSegment(number, scale)

	app.PrintSevenSegmentDisplay(result)

}
