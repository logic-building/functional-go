package fp

// DedupeInt Returns a new list removing consecutive duplicates in list.
func DedupeInt(list []int) []int {
	var newList []int

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeIntPtr Returns a new list removing consecutive duplicates in list.
func DedupeIntPtr(list []*int) []*int {
	var newList []*int

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && *list[i] == *list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeInt64 Returns a new list removing consecutive duplicates in list.
func DedupeInt64(list []int64) []int64 {
	var newList []int64

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeInt64Ptr Returns a new list removing consecutive duplicates in list.
func DedupeInt64Ptr(list []*int64) []*int64 {
	var newList []*int64

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && *list[i] == *list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeInt32 Returns a new list removing consecutive duplicates in list.
func DedupeInt32(list []int32) []int32 {
	var newList []int32

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeInt32Ptr Returns a new list removing consecutive duplicates in list.
func DedupeInt32Ptr(list []*int32) []*int32 {
	var newList []*int32

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && *list[i] == *list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeInt16 Returns a new list removing consecutive duplicates in list.
func DedupeInt16(list []int16) []int16 {
	var newList []int16

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeInt16Ptr Returns a new list removing consecutive duplicates in list.
func DedupeInt16Ptr(list []*int16) []*int16 {
	var newList []*int16

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && *list[i] == *list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeInt8 Returns a new list removing consecutive duplicates in list.
func DedupeInt8(list []int8) []int8 {
	var newList []int8

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeInt8Ptr Returns a new list removing consecutive duplicates in list.
func DedupeInt8Ptr(list []*int8) []*int8 {
	var newList []*int8

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && *list[i] == *list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeUint Returns a new list removing consecutive duplicates in list.
func DedupeUint(list []uint) []uint {
	var newList []uint

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeUintPtr Returns a new list removing consecutive duplicates in list.
func DedupeUintPtr(list []*uint) []*uint {
	var newList []*uint

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && *list[i] == *list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeUint64 Returns a new list removing consecutive duplicates in list.
func DedupeUint64(list []uint64) []uint64 {
	var newList []uint64

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeUint64Ptr Returns a new list removing consecutive duplicates in list.
func DedupeUint64Ptr(list []*uint64) []*uint64 {
	var newList []*uint64

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && *list[i] == *list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeUint32 Returns a new list removing consecutive duplicates in list.
func DedupeUint32(list []uint32) []uint32 {
	var newList []uint32

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeUint32Ptr Returns a new list removing consecutive duplicates in list.
func DedupeUint32Ptr(list []*uint32) []*uint32 {
	var newList []*uint32

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && *list[i] == *list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeUint16 Returns a new list removing consecutive duplicates in list.
func DedupeUint16(list []uint16) []uint16 {
	var newList []uint16

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeUint16Ptr Returns a new list removing consecutive duplicates in list.
func DedupeUint16Ptr(list []*uint16) []*uint16 {
	var newList []*uint16

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && *list[i] == *list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeUint8 Returns a new list removing consecutive duplicates in list.
func DedupeUint8(list []uint8) []uint8 {
	var newList []uint8

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeUint8Ptr Returns a new list removing consecutive duplicates in list.
func DedupeUint8Ptr(list []*uint8) []*uint8 {
	var newList []*uint8

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && *list[i] == *list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeStr Returns a new list removing consecutive duplicates in list.
func DedupeStr(list []string) []string {
	var newList []string

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeStrPtr Returns a new list removing consecutive duplicates in list.
func DedupeStrPtr(list []*string) []*string {
	var newList []*string

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && *list[i] == *list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeBool Returns a new list removing consecutive duplicates in list.
func DedupeBool(list []bool) []bool {
	var newList []bool

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeBoolPtr Returns a new list removing consecutive duplicates in list.
func DedupeBoolPtr(list []*bool) []*bool {
	var newList []*bool

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && *list[i] == *list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeFloat32 Returns a new list removing consecutive duplicates in list.
func DedupeFloat32(list []float32) []float32 {
	var newList []float32

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeFloat32Ptr Returns a new list removing consecutive duplicates in list.
func DedupeFloat32Ptr(list []*float32) []*float32 {
	var newList []*float32

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && *list[i] == *list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeFloat64 Returns a new list removing consecutive duplicates in list.
func DedupeFloat64(list []float64) []float64 {
	var newList []float64

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && list[i] == list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}

// DedupeFloat64Ptr Returns a new list removing consecutive duplicates in list.
func DedupeFloat64Ptr(list []*float64) []*float64 {
	var newList []*float64

	lenList := len(list)
	for i := 0; i < lenList; i++ {
		if i+1 < lenList && *list[i] == *list[i+1] {
			continue
		}
		newList = append(newList, list[i])
	}
	return newList
}
