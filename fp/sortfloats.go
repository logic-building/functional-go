package fp

import "sort"
func SortFloats64(list []float64) []float64 {
	if len(list) == 0 {
		return []float64{}
	}
	newList := make([]float64, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Float64s(newList)
	return newList
}

func SortFloats64Ptr(list []*float64) []*float64 {
	if len(list) == 0 {
		return []*float64{}
	}
	newList := make([]float64, len(list))
	for i, item := range list {
		newList[i] = *item
	}
	sort.Float64s(newList)

	newListPtr := make([]*float64, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

func SortFloats64Desc(list []float64) []float64 {
	if len(list) == 0 {
		return []float64{}
	}
	newList := make([]float64, len(list))
	for i, item := range list {
		newList[i] = item
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(newList)))
	return newList
}

func SortFloats64DescPtr(list []*float64) []*float64 {
	if len(list) == 0 {
		return []*float64{}
	}
	newList := make([]float64, len(list))
	for i, item := range list {
		newList[i] = *item
	}

	sort.Sort(sort.Reverse(sort.Float64Slice(newList)))

	newListPtr := make([]*float64, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}