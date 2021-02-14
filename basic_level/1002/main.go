package main

import (
	"fmt"
	// "strconv"
)

func main() {
	num_list := []string{"ling", "yi", "er", "san", "si", "wu", "liu", "qi", "ba", "jiu"}

	var (
		num_str string
		sum     int32 = 0
	)
	fmt.Scanln(&num_str)

	for _, str := range num_str {
		sum += str - '0'
	}

	if sum < 10 {
		fmt.Println(num_list[sum])
	} else if sum < 100 {
		fmt.Println(num_list[sum/10], num_list[sum%10])
	} else {
		fmt.Println(num_list[sum/100], num_list[sum/10%10], num_list[sum%10])
	}
}
