package main

import "fmt"

type Input struct{}

func (i Input) GetNumber() (string, int, error) {
	var number string
	var scale int

	fmt.Print("Enter the number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		return "", 0, err
	}

	fmt.Print("Enter the scale : ")
	_, err = fmt.Scan(&scale)
	if err != nil {
		return "", 0, err
	}

	return number, scale, nil
}
