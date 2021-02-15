package main

import "fmt"

func main() {
	var (
		t  int
		n1 int
		n2 int
		n3 int
	)
	fmt.Scanln(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n1, &n2, &n3)
		if n1+n2 > n3 {
			fmt.Printf("Case #%d: true\n", i+1)
		} else {
			fmt.Printf("Case #%d: false\n", i+1)
		}
	}
}
