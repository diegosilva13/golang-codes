package main

import (
	"fmt"
)

func main() {
	colors := make(map[int]string)

	colors[1] = "#ffffff"
	colors[2] = "#cccccc"

	printMap(colors)
}

func printMap(m map[int]string) {
	for value, index := range m {
		fmt.Println(value, index)
	}
}
