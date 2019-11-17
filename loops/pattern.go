package main

import "fmt"

func main() {
	row := 5
	for i := 0; i < row; i++ {
		for j := 0; j < row-1-i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
