package methodchain

// MethodChainBasic is template to generate functional code for different combination of data type
func MethodChainBasic() string {
	return `
type <NEWTYPE>Slice []<TYPE>
type <NEWTYPE>FunctorForMap func(<TYPE>) <TYPE>
type <NEWTYPE>FunctorForFilter func(<TYPE>) bool

type <NEWTYPE>SlicePtr []*<TYPE>
type <NEWTYPE>FunctorForMapPtr func(*<TYPE>) *<TYPE>
type <NEWTYPE>FunctorForFilterPtr func(*<TYPE>) bool

// Make<FTYPE>Slice - creates slice for the functional method such as map, filter
func Make<FTYPE>Slice(values ...<TYPE>) <NEWTYPE>Slice {
	newSlice := <NEWTYPE>Slice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice <NEWTYPE>Slice) Map(functors ...<NEWTYPE>FunctorForMap) <NEWTYPE>Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = Map<FTYPE>(f, tmpSlice)
	}

	return tmpSlice
}

// Make<FTYPE>SlicePtr - creates slice for the functional method such as map, filter
func Make<FTYPE>SlicePtr(values ...*<TYPE>) <NEWTYPE>SlicePtr {
	newSlice := <NEWTYPE>SlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice <NEWTYPE>SlicePtr) MapPtr(functors ...<NEWTYPE>FunctorForMapPtr) <NEWTYPE>SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = Map<FTYPE>Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - filters list based on function passed as argument
func (slice <NEWTYPE>Slice) Filter(functors ...<NEWTYPE>FunctorForFilter) <NEWTYPE>Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = Filter<FTYPE>(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - filters list based on function passed as argument
func (slice <NEWTYPE>SlicePtr) FilterPtr(functors ...<NEWTYPE>FunctorForFilterPtr) <NEWTYPE>SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = Filter<FTYPE>Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Remove - removes the items from the given list based on supplied function and returns new list
func (slice <NEWTYPE>Slice) Remove(functors ...<NEWTYPE>FunctorForFilter) <NEWTYPE>Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = Remove<FTYPE>(f, tmpSlice)
	}

	return tmpSlice
}

// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice <NEWTYPE>SlicePtr) RemovePtr(functors ...<NEWTYPE>FunctorForFilterPtr) <NEWTYPE>SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = Remove<FTYPE>Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// DropWhile - drops the items from the list as long as condition satisfies
func (slice <NEWTYPE>Slice) DropWhile(functors ...<NEWTYPE>FunctorForFilter) <NEWTYPE>Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = DropWhile<FTYPE>(f, tmpSlice)
	}

	return tmpSlice
}

// DropWhilePtr - drops the items from the list as long as condition satisfies
func (slice <NEWTYPE>SlicePtr) DropWhilePtr(functors ...<NEWTYPE>FunctorForFilterPtr) <NEWTYPE>SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = DropWhile<FTYPE>Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Reverse - reverse the list
func (slice <NEWTYPE>Slice) Reverse() <NEWTYPE>Slice {
	return Reverse<FTYPE>s(slice)
}

// ReversePtr - reverse the list
func (slice <NEWTYPE>SlicePtr) ReversePtr() <NEWTYPE>SlicePtr {
	return Reverse<FTYPE>sPtr(slice)
}

// Distinct - removes duplicates
func (slice <NEWTYPE>Slice) Distinct() <NEWTYPE>Slice {
	return Distinct<FTYPE>(slice)
}

// DistinctPtr - removes duplicates
func (slice <NEWTYPE>SlicePtr) DistinctPtr() <NEWTYPE>SlicePtr {
	return Distinct<FTYPE>Ptr(slice)
}
`
}

// MethodChainBasicSort -
func MethodChainBasicSort() string {
	return `
// Sort - sort the list
func (slice <NEWTYPE>Slice) Sort() <NEWTYPE>Slice {
	return Sort<FTYPE>s(slice)
}

// SortDesc - sort the list
func (slice <NEWTYPE>Slice) SortDesc() <NEWTYPE>Slice {
	return Sort<FTYPE>sDesc(slice)
}

// SortPtr - sort the list
func (slice <NEWTYPE>SlicePtr) SortPtr() <NEWTYPE>SlicePtr {
	return Sort<FTYPE>sPtr(slice)
}

// SortDescPtr - sort the list
func (slice <NEWTYPE>SlicePtr) SortDescPtr() <NEWTYPE>SlicePtr {
	return Sort<FTYPE>sDescPtr(slice)
}
`
}

// MethodChainStruct is template to generate functional code for different combination of data type
func MethodChainStruct() string {
	return `
type <NEWTYPE>Slice []<TYPE>
type <NEWTYPE>FunctorForMap func(<TYPE>) <TYPE>
type <NEWTYPE>FunctorForFilter func(<TYPE>) bool

type <NEWTYPE>SlicePtr []*<TYPE>
type <NEWTYPE>FunctorForMapPtr func(*<TYPE>) *<TYPE>
type <NEWTYPE>FunctorForFilterPtr func(*<TYPE>) bool

// Make<FTYPE>Slice - creates slice for the functional method such as map, filter
func Make<FTYPE>Slice(values ...<TYPE>) <NEWTYPE>Slice {
	newSlice := <NEWTYPE>Slice(values)
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice <NEWTYPE>Slice) Map(functors ...<NEWTYPE>FunctorForMap) <NEWTYPE>Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = Map<CONDITIONAL_TYPE>(f, tmpSlice)
	}

	return tmpSlice
}

// Make<FTYPE>SlicePtr - creates slice for the functional method such as map, filter
func Make<FTYPE>SlicePtr(values ...*<TYPE>) <NEWTYPE>SlicePtr {
	newSlice := <NEWTYPE>SlicePtr(values)
	return newSlice
}

// MapPtr - applies the function(1st argument) on each item of the list and returns new list
func (slice <NEWTYPE>SlicePtr) MapPtr(functors ...<NEWTYPE>FunctorForMapPtr) <NEWTYPE>SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = Map<CONDITIONAL_TYPE>Ptr(f, tmpSlice)
	}

	return tmpSlice
}

// Filter - 
func (slice <NEWTYPE>Slice) Filter(functors ...<NEWTYPE>FunctorForFilter) <NEWTYPE>Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = Filter<CONDITIONAL_TYPE>(f, tmpSlice)
	}

	return tmpSlice
}

// FilterPtr - 
func (slice <NEWTYPE>SlicePtr) FilterPtr(functors ...<NEWTYPE>FunctorForFilterPtr) <NEWTYPE>SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = Filter<CONDITIONAL_TYPE>Ptr(f, tmpSlice)
	}

	return tmpSlice
}
`
}

// MethodChainStructForRemove is template to generate functional code for different combination of data type
func MethodChainStructForRemove() string {
	return `
// Remove - removes the items from the given list based on supplied function and returns new list
func (slice <NEWTYPE>Slice) Remove(functors ...<NEWTYPE>FunctorForFilter) <NEWTYPE>Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = Remove<CONDITIONAL_TYPE>(f, tmpSlice)
	}

	return tmpSlice
}
`
}

// MethodChainStructForRemove is template to generate functional code for different combination of data type
func MethodChainStructForRemovePtr() string {
	return `
// RemovePtr - removes the items from the given list based on supplied function and returns new list
func (slice <NEWTYPE>SlicePtr) RemovePtr(functors ...<NEWTYPE>FunctorForFilterPtr) <NEWTYPE>SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = Remove<CONDITIONAL_TYPE>Ptr(f, tmpSlice)
	}

	return tmpSlice
}
`
}

// MethodChainStructForDropWhile is template to generate functional code for different combination of data type
func MethodChainStructForDropWhile() string {
	return `
// DropWhile - drops the items from the list as long as condition satisfies
func (slice <NEWTYPE>Slice) DropWhile(functors ...<NEWTYPE>FunctorForFilter) <NEWTYPE>Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = DropWhile<CONDITIONAL_TYPE>(f, tmpSlice)
	}

	return tmpSlice
}
`
}

// MethodChainStructForDropWhilePtr is template to generate functional code for different combination of data type
func MethodChainStructForDropWhilePtr() string {
	return `
// DropWhilePtr - drops the items from the list as long as condition satisfies
func (slice <NEWTYPE>SlicePtr) DropWhilePtr(functors ...<NEWTYPE>FunctorForFilterPtr) <NEWTYPE>SlicePtr {

	tmpSlice := slice
	
	for _, f := range functors {
		if f == nil {
			continue
		}
		tmpSlice = DropWhile<CONDITIONAL_TYPE>Ptr(f, tmpSlice)
	}

	return tmpSlice
}
`
}

// MethodChainStructForReverse is template to generate functional code for different combination of data type
func MethodChainStructForReverse() string {
	return `
// Reverse - reverse the list
func (slice <NEWTYPE>Slice) Reverse() <NEWTYPE>Slice {
	return Reverse<CONDITIONAL_TYPE>s(slice)
}
`
}

// MethodChainStructForReversePtr is template to generate functional code for different combination of data type
func MethodChainStructForReversePtr() string {
	return `
// ReversePtr - reverse the list
func (slice <NEWTYPE>SlicePtr) ReversePtr() <NEWTYPE>SlicePtr {
	return Reverse<CONDITIONAL_TYPE>sPtr(slice)
}
`
}

// MethodChainStructSort -
func MethodChainStructSort() string {
	return `
// SortBy<FFIELD_NAME> - sort the list
func (slice <NEWTYPE>Slice) SortBy<FFIELD_NAME>() <NEWTYPE>Slice {
	return Sort<FSTRUCT_NAME>By<FFIELD_NAME>(slice)
}

// SortBy<FFIELD_NAME>Desc - sort the list
func (slice <NEWTYPE>Slice) SortBy<FFIELD_NAME>Desc() <NEWTYPE>Slice {
	return Sort<FSTRUCT_NAME>By<FFIELD_NAME>Desc(slice)
}
`
}

// MethodChainStructSortPtr -
func MethodChainStructSortPtr() string {
	return `
// SortBy<FFIELD_NAME> - sort the list
func (slice <NEWTYPE>SlicePtr) SortBy<FFIELD_NAME>Ptr() <NEWTYPE>SlicePtr {
	return Sort<FSTRUCT_NAME>By<FFIELD_NAME>Ptr(slice)
}

// SortBy<FFIELD_NAME>DescPtr - sort the list
func (slice <NEWTYPE>SlicePtr) SortBy<FFIELD_NAME>DescPtr() <NEWTYPE>SlicePtr {
	return Sort<FSTRUCT_NAME>By<FFIELD_NAME>DescPtr(slice)
}
`
}

// MethodChainStructDistinct -
func MethodChainStructDistinct() string {
	return `
// Distinct - removes duplicates
func (slice <NEWTYPE>Slice) Distinct() <NEWTYPE>Slice {
	return Distinct<CONDITIONAL_TYPE>(slice)
}
`
}

// MethodChainStructDistinct -
func MethodChainStructDistinctPtr() string {
	return `
// DistinctPtr - removes duplicates
func (slice <NEWTYPE>SlicePtr) DistinctPtr() <NEWTYPE>SlicePtr {
	return Distinct<CONDITIONAL_TYPE>Ptr(slice)
}
`
}
