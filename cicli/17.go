package main

import "fmt"

func main() {
	var x int
	var p int
	var y int
	fmt.Scan(&x, &p, &y)

	i := 0
	for x < y {
		x = x + (x * p / 100)
		i++
	}
	fmt.Println(i)
	// 333
}
