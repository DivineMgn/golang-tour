// https://go-tour-ru-ru.appspot.com/moretypes/26

package main

import "fmt"

func main() {
	fib := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x1, x2 := 0, 1

	return func() int {
		temp := x1
		x1, x2 = x2, x2+x1
		return temp
	}
}
