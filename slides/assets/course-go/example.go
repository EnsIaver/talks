package main

import "fmt"

func main() {
	slice := []string{"one", "two", "three", "four"}

	for index := range slice {
		fmt.Println(index)
	}

	fmt.Println()

	for index, item := range slice {
		fmt.Println(index, item)
	}

	fmt.Println()

	for _, item := range slice {
		fmt.Println(item)
	}
}
