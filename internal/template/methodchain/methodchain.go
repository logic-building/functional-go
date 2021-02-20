package methodchain

// MethodChain is template to generate functional code for different combination of data type
func MethodChain() string {
	return `
type <NEWTYPE>Slice []<TYPE>
type <NEWTYPE>Functor func(<TYPE>) <TYPE>

// Make<FTYPE>Slice - creates slice for the functional method such as map, filter
func Make<FTYPE>Slice(values ...<TYPE>) <NEWTYPE>Slice {
	newSlice := <NEWTYPE>Slice(values)
	return newSlice
}

func mapCore<NEWTYPE>(f <NEWTYPE>Functor, slice <NEWTYPE>Slice) <NEWTYPE>Slice {
	newSlice := make(<NEWTYPE>Slice, len(slice))
	for i, v := range slice {
		newSlice[i] = f(v)
	}
	return newSlice
}

// Map - applies the function(1st argument) on each item of the list and returns new list
func (slice <NEWTYPE>Slice) Map(functors ...<NEWTYPE>Functor) <NEWTYPE>Slice {

	tmpSlice := slice
	
	for _, f := range functors {
		tmpSlice = mapCore<NEWTYPE>(f, tmpSlice)
	}

	return tmpSlice
}
`
}
