package main

import "fmt"

func main() {
	var num []int
	var size int
	fmt.Println("What would be the length:")
	fmt.Scanln(&size)
	for len(num) < size {
		var input int
		fmt.Println("Enter a number:")
		fmt.Scanln(&input)
		num = append(num, input)
	}
	//size := len(num)
	//count := 0
	var temp []int
	for i := (size - 1); i >= 0; i-- {
		temp = append(temp, num[i])
	}

	fmt.Println("\n Original List: ")
	for index, value := range num {
		fmt.Printf("Index: %d \t Value: %d \n", index, value)
	}

	fmt.Println("\n Reversed List: ")
	for index, value := range temp {
		fmt.Printf("Index: %d \t Value: %d \n", index, value)
	}
}
