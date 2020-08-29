package main

import "fmt"

const boilingF = 212.0

func main() {
	var boilingPoint string
	boilingPoint = getBoilingPoint()
	fmt.Printf(boilingPoint)
}

func getBoilingPoint() (ret string) {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	ret = fmt.Sprintf("boiling point = %gF or %gC\n", f, c)
	return
}
