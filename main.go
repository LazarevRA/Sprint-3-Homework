package main

import (
	"fmt"
	"sort" // или "slices"
)

// Mode возвращает моды числовой последовательности.
// Напишите код функции
// ...

func Mode(list []int) ([]int, int) {
	if len(list) == 0 {
		return nil, 1
	}
	sorted := list
	sort.Ints(sorted)
	var result []int
	colvoChisel := make(map[int]int)
	for _, ch := range sorted {
		_, ok := colvoChisel[ch]
		if !ok {
			colvoChisel[ch] = 1
		} else {
			colvoChisel[ch]++
		}
	}

	max := 0
	for _, value := range colvoChisel {
		if value > max {
			max = value
		}
	}

	for key, colvo := range colvoChisel {
		if colvo == max {
			result = append(result, key)
		}
	}

	if max <= 1 {
		return nil, max
	}
	sortedRes := result
	sort.Ints(sortedRes)

	return sortedRes, max

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
	modes := [][]int{
		{}, {}, {}, {},
		{4},
		{-7, 102}, {},
		{20000},
	}
	mcount := []int{
		1, 1, 1, 1, 3, 2, 1, 3,
	}
	for i, list := range lists {
		mode, count := Mode(list)
		if len(mode) != len(modes[i]) {
			fmt.Printf("len mode %d: %v != %v\n", i, modes[i], mode)
		} else {
			for j, v := range mode {
				if v != modes[i][j] {
					fmt.Printf("mode %d: %v != %v\n", i, modes[i], mode)
				}
			}
		}
		if count != mcount[i] {
			fmt.Printf("mcount %d: %d != %d\n", i, mcount[i], count)
		}
	}
	fmt.Println("Тестирование завершено")
}
