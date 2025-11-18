// Arr sort ASC
package main

import "fmt"

func main() {
	fmt.Println(selectionSort([]int{1011, 20, 30, 5, 100}))
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 1; i < len(arr); i++ { // можно начинать с 1, т.к. 0 уже учтён
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(arr []int) []int {
	var sorted []int
	workArr := make([]int, len(arr))
	copy(workArr, arr) // работаем с копией, чтобы не менять исходный

	for len(workArr) > 0 {
		smallestIndex := findSmallest(workArr)
		sorted = append(sorted, workArr[smallestIndex])
		// Удаляем элемент по индексу
		workArr = append(workArr[:smallestIndex], workArr[smallestIndex+1:]...)
	}
	return sorted
}
