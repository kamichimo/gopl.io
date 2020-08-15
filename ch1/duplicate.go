package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// 標準入力から２回以上現れる行を出現回数と一緒に表示する
func dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	// input.Err()からのエラーの可能性を無視している
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// ２回以上現れる行を出現回数と一緒に表示する
//　標準入力から読み込むか、名前が指定されたファイルの一覧から読み込む
func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

// ２回以上現れる行を出現回数と一緒に表示する
//　名前が指定されたファイルの一覧から読み込む
func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
