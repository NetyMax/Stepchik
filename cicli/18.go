package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	db := make(map[int]bool)
	t := b
	for t > 0 {
		d := t % 10
		db[d] = true
		t /= 10

	}

	t = a
	var result []int
	for t > 0 {
		d := t % 10
		if db[d] {
			result = append(result, d)
		}
		t /= 10
	}
	for i := len(result) - 1; i >= o; i-- {

		fmt.Print(result[i], " ")
	}

}
