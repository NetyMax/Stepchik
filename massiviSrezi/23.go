package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	slice := make([]int, N)
	positiveCount := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&slice[i])

		if slice[i] > 0 {
			positiveCount++

		}

	}
	fmt.Print(positiveCount)
}

//Дана последовательность, состоящая из целых чисел. Напишите программу,
//  которая подсчитывает количество положительных чисел среди элементов последовательности.
//Необходимо вывести единственное число - количество положительных элементов в последовательности.
