package sorters

type BubbleSort struct { }

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

func (b BubbleSort) Sort(arr []int) []int {
  return bubbleSort(arr)
}
