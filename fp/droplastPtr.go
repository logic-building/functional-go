package fp

// DropLastPtrInt drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastPtrInt(list []*int) []*int {
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

// DropLastPtrInt64 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastPtrInt64(list []*int64) []*int64 {
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

// DropLastPtrInt32 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastPtrInt32(list []*int32) []*int32 {
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

// DropLastPtrInt16 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastPtrInt16(list []*int16) []*int16 {
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

// DropLastPtrInt8 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastPtrInt8(list []*int8) []*int8 {
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

// DropLastPtrUint drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastPtrUint(list []*uint) []*uint {
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

// DropLastPtrUint64 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastPtrUint64(list []*uint64) []*uint64 {
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

// DropLastPtrUint32 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastPtrUint32(list []*uint32) []*uint32 {
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

// DropLastPtrUint16 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastPtrUint16(list []*uint16) []*uint16 {
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

// DropLastPtrUint8 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastPtrUint8(list []*uint8) []*uint8 {
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

// DropLastPtrStr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastPtrStr(list []*string) []*string {
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

// DropLastPtrBool drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastPtrBool(list []*bool) []*bool {
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
