package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"unsafe"
)

// 指定したURLにある内容を取得・表示する
func fetch(args []string, result string) {
	for _, url := range args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		result = *(*string)(unsafe.Pointer(&b))
		//fmt.Printf("%s", b)
	}
}
