package basic

// Max<FTYPE>Ptr returns max in the list
func MinPtr() string {
	return `
// MinInt returns min item from the list.
// Return 0 if the list is either empty or nil
func Min<FTYPE>Ptr(list []*<TYPE>) *<TYPE> {
	if list == nil || len(list) == 0 {
		var zero <TYPE> = 0
		return &zero
	}
	min := list[0]
	for _, v := range list {
		if *v < *min {
			min = v
		}
	}
	return min
}
`
}
