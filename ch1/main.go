package main

import (
	"fmt"
	"os"
)

func main() {

	result := ""
	fmt.Println(&result)
	echo1(os.Args, &result)
	//fetch(os.Args, result)

	fmt.Println(result)
}
