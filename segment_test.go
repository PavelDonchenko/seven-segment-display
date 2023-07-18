package main

import "testing"

func TestConvertToSevenSegment(t *testing.T) {
	testCases := []struct {
		name   string
		number float64
		scale  int
		output string
	}{
		{
			name:   "zero, scale 1",
			number: 0,
			scale:  1,
			output: " -\n| |\n| |\n - ",
		},
		{name: "zero, scale 2",
			number: 0,
			scale:  2,
			output: " --\n|  |\n|  |\n|  |\n|  |\n -- ",
		},
		{name: "zero, eight, scale 2",
			number: 80,
			scale:  2,
			output: " --\n|  |\n|  |\n --\n|  |\n|  | \n --\n --\n|  |\n|  |\n|  |\n|  |\n --",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			app := Application{
				Number: 0,
				Scale:  1,
			}

			got := app.ConvertToSevenSegment(app.Number, app.Scale)
			if got != test.output {
				t.Errorf("Expected:\n%s\nGot:\n%s", test.output, got)
			}
		})
	}
}
