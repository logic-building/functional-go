package basic

// MinPtr is template
func MinPtr() string {
	return `
// Min<FTYPE>Ptr returns min item from the list.
// Return 0 if the list is either empty or nil
func Min<FTYPE>Ptr(list []*<TYPE>) *<TYPE> {
	if list == nil || len(list) == 0 {
		var zero <TYPE>
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
