package main

import (
	"fmt"
)

type SevenSegment interface {
	GetNumber() error
	PrintSevenSegmentDisplay(output string)
}

type Application struct {
	Number float64
	Scale  int
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) GetNumber() error {
	var number float64
	var scale int

	fmt.Print("Enter the number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		return err
	}

	fmt.Print("Enter the scale : ")
	_, err = fmt.Scan(&scale)
	if err != nil {
		return err
	}

	a.Number = number
	a.Scale = scale

	return nil
}

func (a *Application) PrintSevenSegmentDisplay(output string) {
	fmt.Println(output)
}

func (a *Application) ConvertToSevenSegment(num float64, scale int) string {
	panic("implement me")
}
