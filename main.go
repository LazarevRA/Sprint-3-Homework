package main

import (
	"fmt"
	"math"
)

// Average возвращает среднее арифметическое элементов слайса []int.
// Напишите код функции
// ...
func Average(list []int) float64 {
	if len(list) == 0 {
		return 0
	}
	sum := 0
	count := 0
	for _, ch := range list {
		sum += ch
		count++
	}

	return float64(sum) / float64(count)
}

// Range возвращает размах числовой последовательности.
// Напишите код функции
// ...

func Range(list []int) int {
	if len(list) == 0 {
		return 0
	}
	max := list[0]
	min := max
	for _, ch := range list {
		if ch > max {
			max = ch
		}
		if ch < min {
			min = ch
		}
	}
	return max - min
}

func main() {
	lists := [][]int{
		{},
		{57},
		{78, -7},
		{99, 200, 0},
		{4, 4, 4, 3},
		{102, -7, 44, -7, 102},
		{82, -23, 1, 5, 98, 100},
		{100000, 90000, 20000, 20000, 20000, 22000, 25500, 22000},
	}
	averages := []float64{
		0, 57, 36, 100, 4, 47, 44, 39938,
	}
	ranges := []int{
		0, 0, 85, 200, 1, 109, 123, 80000,
	}

	for i, list := range lists {
		if average := math.Round(Average(list)); average != averages[i] {
			fmt.Printf("average %d: %.2f != %.2f\n", i, averages[i], average)
		}
		if r := Range(list); r != ranges[i] {
			fmt.Printf("range %d: %d != %d\n", i, ranges[i], r)
		}
	}
	fmt.Println("Тестирование завершено")
}
