package main

import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println("Ошибка при вводе:", err)
		return
	}

	if a >= b {
		fmt.Println("Ошибка: Первое число (A) должно быть меньше второго (B).")
		return
	}

	if a <= 0 || b <= 0 {
		fmt.Println("Ошибка: Оба числа должны быть натуральными (больше 0).")
		return
	}

	if a > 100 || b > 100 {
		fmt.Println("Ошибка: Оба числа должны быть не более 100.")
		return
	}

	sum := 0

	for i := a; i <= b; i++ {
		sum += i
	}

	fmt.Println(sum)
}
