package main

import "fmt"

func main() {
	//method 1
	// color := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }
	// color["white"] = "ffffff"
	// fmt.Println(color)

	//method 2
	//var colors map[string]string

	//method3
	colors := make(map[string]string)
	colors["Red"] = "#ff0000"
	colors["white"] = "#ffffff"
	colors["green"] = "#4bf745"
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("The color is ", color, "and the hex code is ", hex)
	}
}
