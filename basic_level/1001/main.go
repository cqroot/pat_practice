package main

import "fmt"

func main() {
    var num int
    var count = 0

    fmt.Scanln(&num)
    for num > 1 {
        if num % 2 == 1 {
            num = (num * 3 + 1) / 2
        } else {
            num = num / 2
        }
        count += 1
    }
    fmt.Println(count)
}
