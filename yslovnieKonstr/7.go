package main

import "fmt"

func main() {
	a := 5
	fmt.Scan(&a)

	if a > 0 {
		fmt.Println("Число положительное")
	}

	if a < 0 {
		fmt.Println("Число отрицательное")
	}

	if a == 0 {
		fmt.Println("Ноль")
	}
}

// com
