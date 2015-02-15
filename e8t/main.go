package main

import (
	"fmt"

	"lonnie.io/e8vm/fs8"
)

func ne(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	ret, e := fs8.AllPackages(".")
	ne(e)

	for _, p := range ret {
		fmt.Println(p)
	}
}
