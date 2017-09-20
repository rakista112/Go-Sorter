package main

import (
	"fmt"
)

func isSorted(items []int) bool {
	return items[1] > items[0]
}

func allSorted(items []int) bool {
	for item := range items {
		if item < len(items)-1 && !isSorted(items[item:item+2]) {
			return false
		}
	}
	return true
}

func bubbleSort(arr []int) []int {
	items := 2
	for !allSorted(arr) {
		for item := range arr {
			if item < len(arr)-1 {
				itemsToBeSorted := arr[item : item+items]
				// if the next two items are not sorted
				if !isSorted(itemsToBeSorted) {
					// swap them
					itemsToBeSorted[0], itemsToBeSorted[1] = itemsToBeSorted[1], itemsToBeSorted[0]
				}
			}
		}
	}
	return arr
}

func main() {
	arr := []int{3, 2, 9, 5}
	fmt.Println(arr)
	sorted := bubbleSort(arr)
	fmt.Printf("result %v\n", sorted)
}
