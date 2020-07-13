package basic

// MaxPtr returns max in the list
func MaxPtr() string {
	return `
// Max<FTYPE>Ptr Returns the greatest of the nums
func Max<FTYPE>Ptr(list []*<TYPE>) *<TYPE> {
	var initialVal <TYPE>
	var max *<TYPE> = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}
`
}
