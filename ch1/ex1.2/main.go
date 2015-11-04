// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Book http://gopl.io Exercise 1.2

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d %s\n", i, arg)
	}
}
