package main

import (
	"fmt"
)

func main() {
	// var colours2 map[string]string
	// colours3 := make(map[string]string)

	colours := map[string]string{
		"red": "#ff0000",
		"green": "#00ff00",
		"blue": "0000ff",
	}

	colours["white"] = "#ffffff"
	delete(colours, "white")

	fmt.Println(colours)
	printMap(colours)
}

func printMap(c map[string]string) {
	for color, hex := range(c) {
		fmt.Println(color, hex)
	}
}
