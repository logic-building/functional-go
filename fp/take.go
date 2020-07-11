package fp

// TakeInt returns n items in the list
func TakeInt(n int, list []int) []int {
	if n < 0 {
		return []int{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]int, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeIntPtr returns n items in the list
func TakeIntPtr(n int, list []*int) []*int {
	if n < 0 {
		return []*int{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]*int, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeInt64 returns n items in the list
func TakeInt64(n int, list []int64) []int64 {
	if n < 0 {
		return []int64{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]int64, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeInt64Ptr returns n items in the list
func TakeInt64Ptr(n int, list []*int64) []*int64 {
	if n < 0 {
		return []*int64{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]*int64, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeInt32 returns n items in the list
func TakeInt32(n int, list []int32) []int32 {
	if n < 0 {
		return []int32{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]int32, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeInt32Ptr returns n items in the list
func TakeInt32Ptr(n int, list []*int32) []*int32 {
	if n < 0 {
		return []*int32{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]*int32, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeInt16 returns n items in the list
func TakeInt16(n int, list []int16) []int16 {
	if n < 0 {
		return []int16{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]int16, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeInt16Ptr returns n items in the list
func TakeInt16Ptr(n int, list []*int16) []*int16 {
	if n < 0 {
		return []*int16{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]*int16, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeInt8 returns n items in the list
func TakeInt8(n int, list []int8) []int8 {
	if n < 0 {
		return []int8{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]int8, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeInt8Ptr returns n items in the list
func TakeInt8Ptr(n int, list []*int8) []*int8 {
	if n < 0 {
		return []*int8{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]*int8, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeUint returns n items in the list
func TakeUint(n int, list []uint) []uint {
	if n < 0 {
		return []uint{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]uint, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeUintPtr returns n items in the list
func TakeUintPtr(n int, list []*uint) []*uint {
	if n < 0 {
		return []*uint{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]*uint, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeUint64 returns n items in the list
func TakeUint64(n int, list []uint64) []uint64 {
	if n < 0 {
		return []uint64{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]uint64, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeUint64Ptr returns n items in the list
func TakeUint64Ptr(n int, list []*uint64) []*uint64 {
	if n < 0 {
		return []*uint64{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]*uint64, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeUint32 returns n items in the list
func TakeUint32(n int, list []uint32) []uint32 {
	if n < 0 {
		return []uint32{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]uint32, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeUint32Ptr returns n items in the list
func TakeUint32Ptr(n int, list []*uint32) []*uint32 {
	if n < 0 {
		return []*uint32{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]*uint32, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeUint16 returns n items in the list
func TakeUint16(n int, list []uint16) []uint16 {
	if n < 0 {
		return []uint16{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]uint16, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeUint16Ptr returns n items in the list
func TakeUint16Ptr(n int, list []*uint16) []*uint16 {
	if n < 0 {
		return []*uint16{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]*uint16, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeUint8 returns n items in the list
func TakeUint8(n int, list []uint8) []uint8 {
	if n < 0 {
		return []uint8{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]uint8, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeUint8Ptr returns n items in the list
func TakeUint8Ptr(n int, list []*uint8) []*uint8 {
	if n < 0 {
		return []*uint8{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]*uint8, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeStr returns n items in the list
func TakeStr(n int, list []string) []string {
	if n < 0 {
		return []string{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]string, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeStrPtr returns n items in the list
func TakeStrPtr(n int, list []*string) []*string {
	if n < 0 {
		return []*string{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]*string, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeBool returns n items in the list
func TakeBool(n int, list []bool) []bool {
	if n < 0 {
		return []bool{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]bool, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeBoolPtr returns n items in the list
func TakeBoolPtr(n int, list []*bool) []*bool {
	if n < 0 {
		return []*bool{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]*bool, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeFloat32 returns n items in the list
func TakeFloat32(n int, list []float32) []float32 {
	if n < 0 {
		return []float32{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]float32, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeFloat32Ptr returns n items in the list
func TakeFloat32Ptr(n int, list []*float32) []*float32 {
	if n < 0 {
		return []*float32{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]*float32, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeFloat64 returns n items in the list
func TakeFloat64(n int, list []float64) []float64 {
	if n < 0 {
		return []float64{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]float64, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}

// TakeFloat64Ptr returns n items in the list
func TakeFloat64Ptr(n int, list []*float64) []*float64 {
	if n < 0 {
		return []*float64{}
	}

	newListLen := len(list)

	if n < newListLen {
		newListLen = n
	}
	newList := make([]*float64, newListLen)
	for i := 0; i < newListLen; i++ {
		newList[i] = list[i]
	}
	return newList
}
