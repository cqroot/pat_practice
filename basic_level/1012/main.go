package main

import (
	"fmt"
)

func main() {
	var (
		num     int
		current int
		a1_sum  int = 0
		a1_len  int = 0
		a2_flag int = 1
		a2_sum  int = 0
		a2_len  int = 0
		a3_len  int = 0
		a4_sum  int = 0
		a4_len  int = 0
		a5_max  int = 0
		a5_len  int = 0
	)

	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		fmt.Scan(&current)

		switch current % 5 {
		case 0:
			if current%2 == 0 {
				a1_sum += current
				a1_len++
			}
		case 1:
			a2_sum += current * a2_flag
			a2_flag = 0 - a2_flag
			a2_len++
		case 2:
			a3_len += 1
		case 3:
			a4_sum += current
			a4_len++
		case 4:
			if current > a5_max {
				a5_max = current
				a5_len++
			}
		}
	}

	if a1_len != 0 {
		fmt.Printf("%d ", a1_sum)
	} else {
		fmt.Print("N ")
	}
	if a2_len != 0 {
		fmt.Printf("%d ", a2_sum)
	} else {
		fmt.Print("N ")
	}
	if a3_len != 0 {
		fmt.Printf("%d ", a3_len)
	} else {
		fmt.Print("N ")
	}
	if a4_len != 0 {
		fmt.Printf("%.1f ", float64(a4_sum)/float64(a4_len))
	} else {
		fmt.Print("N ")
	}
	if a5_len != 0 {
		fmt.Printf("%d", a5_max)
	} else {
		fmt.Print("N")
	}
	fmt.Println()
}
