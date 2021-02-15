package main

import "fmt"

func main() {
	var (
		num int
	)
	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		var (
			a_cnt1 int  = 0
			a_cnt2 int  = 0
			a_cnt3 int  = 0
			result bool = true
			flag   int  = 0
			input  string
		)

		fmt.Scanln(&input)

		for _, str := range input {
			if !result {
				continue
			}

			if str == 'P' {
				if flag != 0 {
					result = false
				} else {
					flag += 1
				}
			} else if str == 'T' {
				if flag != 1 {
					result = false
				} else {
					flag += 1
				}
			} else if str == 'A' {
				if flag == 0 {
					a_cnt1 += 1
				} else if flag == 1 {
					a_cnt2 += 1
				} else if flag == 2 {
					a_cnt3 += 1
				}
			} else {
				result = false
			}
		}

		if flag != 2 || a_cnt2 == 0 || a_cnt1*a_cnt2 != a_cnt3 {
			result = false
		}
		if !result {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}
	}
}
