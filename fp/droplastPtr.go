package fp

// DropLastIntPtr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastIntPtr(list []*int) []*int {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []*int{}
	}

	newList := make([]*int, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastInt64Ptr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastInt64Ptr(list []*int64) []*int64 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []*int64{}
	}

	newList := make([]*int64, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastInt32Ptr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastInt32Ptr(list []*int32) []*int32 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []*int32{}
	}

	newList := make([]*int32, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastInt16Ptr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastInt16Ptr(list []*int16) []*int16 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []*int16{}
	}

	newList := make([]*int16, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastInt8Ptr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastInt8Ptr(list []*int8) []*int8 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []*int8{}
	}

	newList := make([]*int8, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastUintPtr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastUintPtr(list []*uint) []*uint {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []*uint{}
	}

	newList := make([]*uint, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastUint64Ptr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastUint64Ptr(list []*uint64) []*uint64 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []*uint64{}
	}

	newList := make([]*uint64, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastUint32Ptr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastUint32Ptr(list []*uint32) []*uint32 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []*uint32{}
	}

	newList := make([]*uint32, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastUint16Ptr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastUint16Ptr(list []*uint16) []*uint16 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []*uint16{}
	}

	newList := make([]*uint16, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastUint8Ptr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastUint8Ptr(list []*uint8) []*uint8 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []*uint8{}
	}

	newList := make([]*uint8, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastStrPtr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastStrPtr(list []*string) []*string {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []*string{}
	}

	newList := make([]*string, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastBoolPtr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastBoolPtr(list []*bool) []*bool {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []*bool{}
	}

	newList := make([]*bool, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastFloat32Ptr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastFloat32Ptr(list []*float32) []*float32 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []*float32{}
	}

	newList := make([]*float32, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastFloat64Ptr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastFloat64Ptr(list []*float64) []*float64 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []*float64{}
	}

	newList := make([]*float64, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}
