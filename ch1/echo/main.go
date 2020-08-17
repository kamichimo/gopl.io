package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := ""
	echo1(os.Args, &result)
	fmt.Println(result)
}

// 標準入力をそのままreturnする
// 参照渡しの練習
func echo1(args []string, result *string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + strconv.Itoa(i) + args[i]
		sep = " "
	}
	fmt.Println(&result)
	*result = s
}

// ブランク区切りで渡された文字列をそのままreturnする
func echo2(name string) string {
	var s, sep string
	for _, arg := range strings.Fields(name) {
		s += sep + arg
		sep = " "
	}
	return s
}

// ブランク区切りで渡された文字列をそのままreturnするワンライナーバージョン
func echo3(name string) string {
	return strings.Join(strings.Fields(name)[1:], " ")
}
