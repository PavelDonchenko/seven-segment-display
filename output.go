package main

import "fmt"

type Output struct{}

func (o Output) PrintSevenSegmentDisplay(output string) {
	fmt.Println(output)
}
