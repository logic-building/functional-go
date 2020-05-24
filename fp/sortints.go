package fp

import "sort"
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