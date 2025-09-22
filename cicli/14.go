package main

import "fmt"

func main() {

	var a int
	var x int
	var c int

	for {
		_, err := fmt.Scan(&a)

		if err != nil {
			break
		}

		if a == 0 {
			break
		}

		if a > x {
			x = a
			c = 1

		} else if a == x {

			c++
		}

	}
	fmt.Println(c)

}
