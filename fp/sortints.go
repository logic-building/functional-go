package fp

import "sort"

// SortInts return new sorted list
func SortInts(list []int) []int {
	if len(list) == 0 {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Ints(newList)
	return newList
}

// SortIntsPtr return new sorted list
func SortIntsPtr(list []*int) []*int {
	if len(list) == 0 {
		return []*int{}
	}
	newList := make([]int, len(list))
	for i, item := range list {
		newList[i] = *item
	}
	sort.Ints(newList)

	newListPtr := make([]*int, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

// SortIntsDesc return new sorted list
func SortIntsDesc(list []int) []int {
	if len(list) == 0 {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Sort(sort.Reverse(sort.IntSlice(newList)))
	return newList
}

// SortIntsDescPtr return new sorted list
func SortIntsDescPtr(list []*int) []*int {
	if len(list) == 0 {
		return []*int{}
	}
	newList := make([]int, len(list))
	for i, item := range list {
		newList[i] = *item
	}

	sort.Sort(sort.Reverse(sort.IntSlice(newList)))

	newListPtr := make([]*int, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}
