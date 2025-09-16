package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)

	q := (a / 10) % 10
	w := a % 10
	e := (a / 100) % 10
	r := (a / 1000) % 10
	t := (a / 10000) % 10
	y := (a / 100000) % 10

	if y+t+r == q+w+e {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}
