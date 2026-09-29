package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	return append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {
	for i, element := range slice {
		slice[i] = f(element)
	}
}

func mapArray(f func(a int) int, array [3]int) [3]int{
	tempArray := [3]int{}
	for i, element := range array {
		tempArray[i] = f(element)
	}
	return tempArray
}

func main() {
	intsSlice := []int{1, 2, 3, 4, 5}
	
	fmt.Println(double(intsSlice))
}
