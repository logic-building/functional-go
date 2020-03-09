package basic

// Max<FTYPE>Ptr returns max in the list
func MaxPtr() string {
	return `
   func Max<FTYPE>Ptr(list []*<TYPE>) *<TYPE> {
	var initialVal <TYPE> = 0
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
