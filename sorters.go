package sorters

type Sorter interface {
	Sort([]int) []int
}

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
