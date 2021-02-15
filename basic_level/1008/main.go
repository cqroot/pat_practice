package main

import "fmt"

func main() {
	var (
		n     int
		m     int
		s     []int
		first bool = true
	)
	fmt.Scanln(&n, &m)
	s = make([]int, n)
	m = m % n

	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	for i := n - m; i < n; i++ {
		if first {
			first = false
		} else {
			fmt.Print(" ")
		}
		fmt.Print(s[i])
	}
	for i := 0; i < n-m; i++ {
		if first {
			first = false
		} else {
			fmt.Print(" ")
		}
		fmt.Print(s[i])
	}
	fmt.Println()
}
