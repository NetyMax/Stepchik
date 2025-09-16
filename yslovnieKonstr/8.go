package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	a := num / 100
	b := (num / 10) % 10
	c := num % 10

	if a != b && b != c && a != c {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
