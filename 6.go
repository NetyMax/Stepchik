package main

import "fmt"

func main() {

	var d int = 90
	fmt.Scan(&d)
	m := d * 2
	h := m / 60
	m %= 60

	fmt.Printf("It is %d hours %d minutes.", h, m)

}
