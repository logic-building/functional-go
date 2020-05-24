package fp

import "sort"
func SortStrs(list []string) []string {
	if len(list) == 0 {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Strings(newList)
	return newList
}

func SortStrsPtr(list []*string) []*string {
	if len(list) == 0 {
		return []*string{}
	}
	newList := make([]string, len(list))
	for i, item := range list {
		newList[i] = *item
	}
	sort.Strings(newList)

	newListPtr := make([]*string, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

func SortStrsDesc(list []string) []string {
	if len(list) == 0 {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Sort(sort.Reverse(sort.StringSlice(newList)))
	return newList
}

func SortStrsDescPtr(list []*string) []*string {
	if len(list) == 0 {
		return []*string{}
	}
	newList := make([]string, len(list))
	for i, item := range list {
		newList[i] = *item
	}

	sort.Sort(sort.Reverse(sort.StringSlice(newList)))

	newListPtr := make([]*string, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}