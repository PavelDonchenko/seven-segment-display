package main

func main() {
	app := NewApplication()

	err := app.GetNumber()
	if err != nil {
		return
	}
}
