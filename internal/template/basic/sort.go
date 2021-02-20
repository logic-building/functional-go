package basic

import "strings"

// Sort is template
func Sort() string {
	return `
type <TYPE>1 []<TYPE>

func (a <TYPE>1) Len() int           { return len(a) }
func (a <TYPE>1) Less(i, j int) bool { return a[i] < a[j] }
func (a <TYPE>1) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Sort<FTYPE>s return new sorted list
func Sort<FTYPE>s(list []<TYPE>) []<TYPE> {
	if len(list) == 0 {
		return []<TYPE>{}
	}
	var newList <TYPE>1
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// Sort<FTYPE>sPtr return new sorted list
func Sort<FTYPE>sPtr(list []*<TYPE>) []*<TYPE> {
	if len(list) == 0 {
		return []*<TYPE>{}
	}
	var newList <TYPE>1
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*<TYPE>, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type <TYPE>2 []<TYPE>

func (a <TYPE>2) Len() int           { return len(a) }
func (a <TYPE>2) Less(i, j int) bool { return a[i] > a[j] }
func (a <TYPE>2) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// Sort<FTYPE>sDesc return new sorted list
func Sort<FTYPE>sDesc(list []<TYPE>) []<TYPE> {
	if len(list) == 0 {
		return []<TYPE>{}
	}
	var newList <TYPE>2
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// Sort<FTYPE>sDescPtr return new sorted list
func Sort<FTYPE>sDescPtr(list []*<TYPE>) []*<TYPE> {
	if len(list) == 0 {
		return []*<TYPE>{}
	}
	var newList <TYPE>2
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*<TYPE>, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}
`
}

// SortInts is template
func SortInts() string {
	return `
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
`
}

// SortFloats64 is template
func SortFloats64() string {
	return `
// SortFloats64 return new sorted list
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

// SortFloats64Ptr return new sorted list
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

// SortFloats64Desc return new sorted list
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

// SortFloats64DescPtr return new sorted list
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
`
}

// SortStrs is template
func SortStrs() string {
	return `
// SortStrs return new sorted list
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

// SortStrsPtr return new sorted list
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

// SortStrsDesc return new sorted list
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

// SortStrsDescPtr return new sorted list
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
`
}

// ReplaceActivitySortingCode  replaces ...
func ReplaceActivitySortingCode(code string) string {
	s1 := `package fp`
	s2 := `package fp

import "sort"`

	code = strings.Replace(code, s1, s2, -1)
	return code

}
