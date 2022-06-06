package main

import (
	"fmt"
)

func main() {
	var start int
	var finish int

	fmt.Scan(&start)

	for start != 0 {
		start = start / 10
		finish++
		fmt.Println(start)
	}

	fmt.Println(finish)
}
