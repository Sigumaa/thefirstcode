package main

import "fmt"

func Fizzbuzz(value int) string {
	if value%15 == 0 {
		return "FizzBuzz"
	} else if value%5 == 0 {
		return "Buzz"
	} else if value%3 == 0 {
		return "Fizz"
	} else {
		return fmt.Sprint(value)
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(Fizzbuzz(i))
	}
}
