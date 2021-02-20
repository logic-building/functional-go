package fp

// DropLastInt drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastInt(list []int) []int {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []int{}
	}

	newList := make([]int, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastInt64 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastInt64(list []int64) []int64 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []int64{}
	}

	newList := make([]int64, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastInt32 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastInt32(list []int32) []int32 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []int32{}
	}

	newList := make([]int32, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastInt16 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastInt16(list []int16) []int16 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []int16{}
	}

	newList := make([]int16, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastInt8 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastInt8(list []int8) []int8 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []int8{}
	}

	newList := make([]int8, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastUint drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastUint(list []uint) []uint {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []uint{}
	}

	newList := make([]uint, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastUint64 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastUint64(list []uint64) []uint64 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []uint64{}
	}

	newList := make([]uint64, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastUint32 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastUint32(list []uint32) []uint32 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []uint32{}
	}

	newList := make([]uint32, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastUint16 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastUint16(list []uint16) []uint16 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []uint16{}
	}

	newList := make([]uint16, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastUint8 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastUint8(list []uint8) []uint8 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []uint8{}
	}

	newList := make([]uint8, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastStr drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastStr(list []string) []string {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []string{}
	}

	newList := make([]string, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastBool drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastBool(list []bool) []bool {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []bool{}
	}

	newList := make([]bool, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastFloat32 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastFloat32(list []float32) []float32 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []float32{}
	}

	newList := make([]float32, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

// DropLastFloat64 drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastFloat64(list []float64) []float64 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []float64{}
	}

	newList := make([]float64, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}
