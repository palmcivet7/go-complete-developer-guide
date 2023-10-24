package main

import "fmt"

func main() {
	// 2 alternative ways to create a map
	// var colors map[string]string
	// colors := make(map[string]string)

	// write to a map
	// colors["white"] = "#FFFFFF"

	// delete a mapped value
	// delete(colors, "white")

	colors := map[string]string{
		"red": "#ff0000",
		"green": "#008000",
		"white": "#FFFFFF",
	}

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Hex code for color", key, "is", value)
	}
}