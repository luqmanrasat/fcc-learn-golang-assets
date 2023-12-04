package main

import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		isMultipleOf3 := i%3 == 0
		isMultipleOf5 := i%5 == 0

		if isMultipleOf3 && isMultipleOf5 {
			fmt.Println("fizzbuzz")
			continue
		}

		if isMultipleOf3 {
			fmt.Println("fizz")
			continue
		}

		if isMultipleOf5 {
			fmt.Println("buzz")
			continue
		}

		fmt.Println(i)
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
