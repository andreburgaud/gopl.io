// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Exercice 1.3

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const MAX = 1000000

type fntest func()

func test1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func test2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func test3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func timeit(fn fntest) {
	//start := time.Now()
	for i := 0; i < MAX; i++ {
		fn()
	}
	//fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func main() {
	times := ""
	start := time.Now()
	timeit(test1)
	times += fmt.Sprintf("%.2fs elapsed", time.Since(start).Seconds()) + ", "
	start = time.Now()
	timeit(test2)
	times += fmt.Sprintf("%.2fs elapsed", time.Since(start).Seconds()) + ", "
	start = time.Now()
	timeit(test3)
	times += fmt.Sprintf("%.2fs elapsed", time.Since(start).Seconds())
	fmt.Println(times)
}
