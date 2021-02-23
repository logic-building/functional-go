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

// Filter - 
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

// FilterPtr - 
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
