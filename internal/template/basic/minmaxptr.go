package basic

// MinMaxPtr is template
func MinMaxPtr() string {
	return `
// MinMax<FTYPE>Ptr returns min and max items from the list.
// Return 0,0 if the list is either empty or nil
func MinMax<FTYPE>Ptr(list []*<TYPE>) (*<TYPE>, *<TYPE>) {
	var zero <TYPE>

	if list == nil || len(list) == 0 {
		return &zero, &zero
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if *v < *min {
			min = v
		} else if *v > *max {
			max = v
		}
	}
	return min, max
}
`
}
