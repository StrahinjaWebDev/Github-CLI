package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println(div(1, 0))
	fmt.Println(safeDiv(1, 0))
}

// named return values
func safeDiv(a, b int) (q int, err error) {
	// q and err are local variables in safediv
	// just like a and b
	defer func() {
		// e's type is any
		if e := recover(); e != nil {
			log.Println("error:", e)
			err = fmt.Errorf("%v", e)
		}
	}()
	return a / b, nil
	// q = a / b
}

func div(a, b int) int {
	return a / b
}
