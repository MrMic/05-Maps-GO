package main

import "fmt"

func main() {
	// colors := map[string]string

	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#FF0000",
	// 	"green": "#00FF00",
	// 	"blue":  "#0000FF",
	// }

	colors["white"] = "#FFFFFF"
	delete(colors, "white")

	fmt.Println(colors)
}
