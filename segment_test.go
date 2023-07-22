package main

import "testing"

func TestConvertToSevenSegment(t *testing.T) {
	testCases := []struct {
		name   string
		number string
		scale  int
		output string
	}{
		{
			name:   "zero, scale 1",
			number: "0",
			scale:  1,
			output: " _   \n| |  \n|_|  ",
		},
		{name: "zero, scale 2",
			number: "0",
			scale:  2,
			output: " __   \n|  |  \n|  |  \n|__|  ",
		},
		{name: "zero, eight, scale 2",
			number: "80",
			scale:  2,
			output: " __    __   \n|  |  |  |  \n|__|  |  |  \n|  |  |  |  \n|__|  |__|  ",
		},
		{name: "zero, eight, scale 2",
			number: "3.14",
			scale:  3,
			output: " ___                        \n    |             |  |   |  \n ___|             |  |___|  \n    |             |      |  \n ___|   *         |      |  ",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			app := Application{}

			got := app.ConvertToSevenSegment(test.number, test.scale)
			if got != test.output {
				t.Errorf("Expected:\n%s\nGot:\n%s", test.output, got)
			}
		})
	}
}
