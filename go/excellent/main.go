package main

import "fmt"

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		fmt.Print("even OK")
		return "even!!"
	}else{
		return "odd"
	}
}
