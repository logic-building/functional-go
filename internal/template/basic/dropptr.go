package basic

// DropPtr returns a new list after dropping the given item
func DropPtr() string {
	return `
// Drop<FTYPE>Ptr returns a new list after dropping the given item
func Drop<FTYPE>Ptr(num *<TYPE>, list []*<TYPE>) []*<TYPE> {
	var newList []*<TYPE>
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// Drop<FTYPE>sPtr returns a new list after dropping the given items
func Drop<FTYPE>sPtr(nums []*<TYPE>, list []*<TYPE>) []*<TYPE> {
	var newList []*<TYPE>
	for _, v := range list {
		if !Exists<FTYPE>Ptr(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}
`
}
