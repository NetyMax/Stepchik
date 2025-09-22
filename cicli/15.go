package main

import "fmt"

func main() {

	var n int
	var c int
	var d int
	fmt.Scan(&n, &c, &d)

	for i := 1; i <= n; i++ {
		if i%c == 0 && c%d != 0 {
			fmt.Println(i)
			return
		}

	}

}
