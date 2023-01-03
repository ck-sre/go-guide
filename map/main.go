package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#sgq13g",
		"white": "#ffffff",
	}

	// fmt.Println(colors)
	printMap(colors)

}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println("Hex code for", k, "is", v)
	}
}
