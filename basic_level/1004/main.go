package main

import "fmt"

func main() {
	var (
		num int

		current_name  string
		current_id    string
		current_score int

		max_name  string
		max_id    string
		max_score int = -1

		min_name  string
		min_id    string
		min_score int = 101
	)

	fmt.Scanln(&num)

	for i := 0; i < num; i++ {
		fmt.Scanln(&current_name, &current_id, &current_score)

		if current_score > max_score {
			max_name = current_name
			max_id = current_id
			max_score = current_score
		}

		if current_score < min_score {
			min_name = current_name
			min_id = current_id
			min_score = current_score
		}
	}
	fmt.Println(max_name, max_id)
	fmt.Println(min_name, min_id)
}
