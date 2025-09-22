package main

import "fmt"

func main() {
	for {
		var x int

		_, err := fmt.Scan(&x)
		if err != nil {
			break
		}

		if x < 10 {
			continue
		}
		if x > 100 {
			break
		}

		fmt.Println(x)
	}
}
