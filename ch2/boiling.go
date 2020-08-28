package main

import "fmt"

const boilingF = 212.0

func main() {
	var boilingPoint string
	boilingPoint = getBoilingPoint()
	fmt.Printf(boilingPoint)
}

func getBoilingPoint() string {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	return fmt.Sprintf("boiling point = %gF or %gC\n", f, c)
}
