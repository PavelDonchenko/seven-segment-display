package main

import (
	"errors"
	"fmt"
)

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

	if scale < 1 {
		return "", 0, errors.New("scale number should be bigger than 0")
	}

	//if scale%2 == 0 {
	//	scale += 1
	//}

	return number, scale, nil
}
