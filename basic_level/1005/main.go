package main

import "fmt"

func main() {
	var (
		visit [101]int
		k     int
		num   int
		first bool = true
	)

	fmt.Scanln(&k)
	for i := 0; i < k; i++ {
		fmt.Scan(&num)

		if visit[num] != 1 {
			visit[num] = 2
		} else {
			continue
		}
		for num != 1 {
			if num%2 == 0 {
				num /= 2
			} else {
				num = (num*3 + 1) / 2
			}
			if num < 101 {
				visit[num] = 1
			}
		}
	}
	for i := 100; i >= 0; i-- {
		if visit[i] == 2 {
			if first {
				first = false
			} else {
				fmt.Print(" ")
			}
			fmt.Print(i)
		}
	}
	fmt.Println()
}
