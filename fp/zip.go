package fp

// ZipInt takes two inputs: first list of type: []int, second list of type: []int.
// Then it merges two list and returns a new map of type: map[int]int
func ZipInt(list1 []int, list2 []int) map[int]int {
	newMap := make(map[int]int)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipIntInt64 takes two inputs: first list of type: []int, second list of type: []int64.
// Then it merges two list and returns a new map of type: map[int]int64
func ZipIntInt64(list1 []int, list2 []int64) map[int]int64 {
	newMap := make(map[int]int64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipIntInt32 takes two inputs: first list of type: []int, second list of type: []int32.
// Then it merges two list and returns a new map of type: map[int]int32
func ZipIntInt32(list1 []int, list2 []int32) map[int]int32 {
	newMap := make(map[int]int32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipIntInt16 takes two inputs: first list of type: []int, second list of type: []int16.
// Then it merges two list and returns a new map of type: map[int]int16
func ZipIntInt16(list1 []int, list2 []int16) map[int]int16 {
	newMap := make(map[int]int16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipIntInt8 takes two inputs: first list of type: []int, second list of type: []int8.
// Then it merges two list and returns a new map of type: map[int]int8
func ZipIntInt8(list1 []int, list2 []int8) map[int]int8 {
	newMap := make(map[int]int8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipIntUint takes two inputs: first list of type: []int, second list of type: []uint.
// Then it merges two list and returns a new map of type: map[int]uint
func ZipIntUint(list1 []int, list2 []uint) map[int]uint {
	newMap := make(map[int]uint)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipIntUint64 takes two inputs: first list of type: []int, second list of type: []uint64.
// Then it merges two list and returns a new map of type: map[int]uint64
func ZipIntUint64(list1 []int, list2 []uint64) map[int]uint64 {
	newMap := make(map[int]uint64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipIntUint32 takes two inputs: first list of type: []int, second list of type: []uint32.
// Then it merges two list and returns a new map of type: map[int]uint32
func ZipIntUint32(list1 []int, list2 []uint32) map[int]uint32 {
	newMap := make(map[int]uint32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipIntUint16 takes two inputs: first list of type: []int, second list of type: []uint16.
// Then it merges two list and returns a new map of type: map[int]uint16
func ZipIntUint16(list1 []int, list2 []uint16) map[int]uint16 {
	newMap := make(map[int]uint16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipIntUint8 takes two inputs: first list of type: []int, second list of type: []uint8.
// Then it merges two list and returns a new map of type: map[int]uint8
func ZipIntUint8(list1 []int, list2 []uint8) map[int]uint8 {
	newMap := make(map[int]uint8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipIntStr takes two inputs: first list of type: []int, second list of type: []string.
// Then it merges two list and returns a new map of type: map[int]string
func ZipIntStr(list1 []int, list2 []string) map[int]string {
	newMap := make(map[int]string)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipIntBool takes two inputs: first list of type: []int, second list of type: []bool.
// Then it merges two list and returns a new map of type: map[int]bool
func ZipIntBool(list1 []int, list2 []bool) map[int]bool {
	newMap := make(map[int]bool)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipIntFloat32 takes two inputs: first list of type: []int, second list of type: []float32.
// Then it merges two list and returns a new map of type: map[int]float32
func ZipIntFloat32(list1 []int, list2 []float32) map[int]float32 {
	newMap := make(map[int]float32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipIntFloat64 takes two inputs: first list of type: []int, second list of type: []float64.
// Then it merges two list and returns a new map of type: map[int]float64
func ZipIntFloat64(list1 []int, list2 []float64) map[int]float64 {
	newMap := make(map[int]float64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt64Int takes two inputs: first list of type: []int64, second list of type: []int.
// Then it merges two list and returns a new map of type: map[int64]int
func ZipInt64Int(list1 []int64, list2 []int) map[int64]int {
	newMap := make(map[int64]int)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt64 takes two inputs: first list of type: []int64, second list of type: []int64.
// Then it merges two list and returns a new map of type: map[int64]int64
func ZipInt64(list1 []int64, list2 []int64) map[int64]int64 {
	newMap := make(map[int64]int64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt64Int32 takes two inputs: first list of type: []int64, second list of type: []int32.
// Then it merges two list and returns a new map of type: map[int64]int32
func ZipInt64Int32(list1 []int64, list2 []int32) map[int64]int32 {
	newMap := make(map[int64]int32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt64Int16 takes two inputs: first list of type: []int64, second list of type: []int16.
// Then it merges two list and returns a new map of type: map[int64]int16
func ZipInt64Int16(list1 []int64, list2 []int16) map[int64]int16 {
	newMap := make(map[int64]int16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt64Int8 takes two inputs: first list of type: []int64, second list of type: []int8.
// Then it merges two list and returns a new map of type: map[int64]int8
func ZipInt64Int8(list1 []int64, list2 []int8) map[int64]int8 {
	newMap := make(map[int64]int8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt64Uint takes two inputs: first list of type: []int64, second list of type: []uint.
// Then it merges two list and returns a new map of type: map[int64]uint
func ZipInt64Uint(list1 []int64, list2 []uint) map[int64]uint {
	newMap := make(map[int64]uint)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt64Uint64 takes two inputs: first list of type: []int64, second list of type: []uint64.
// Then it merges two list and returns a new map of type: map[int64]uint64
func ZipInt64Uint64(list1 []int64, list2 []uint64) map[int64]uint64 {
	newMap := make(map[int64]uint64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt64Uint32 takes two inputs: first list of type: []int64, second list of type: []uint32.
// Then it merges two list and returns a new map of type: map[int64]uint32
func ZipInt64Uint32(list1 []int64, list2 []uint32) map[int64]uint32 {
	newMap := make(map[int64]uint32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt64Uint16 takes two inputs: first list of type: []int64, second list of type: []uint16.
// Then it merges two list and returns a new map of type: map[int64]uint16
func ZipInt64Uint16(list1 []int64, list2 []uint16) map[int64]uint16 {
	newMap := make(map[int64]uint16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt64Uint8 takes two inputs: first list of type: []int64, second list of type: []uint8.
// Then it merges two list and returns a new map of type: map[int64]uint8
func ZipInt64Uint8(list1 []int64, list2 []uint8) map[int64]uint8 {
	newMap := make(map[int64]uint8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt64Str takes two inputs: first list of type: []int64, second list of type: []string.
// Then it merges two list and returns a new map of type: map[int64]string
func ZipInt64Str(list1 []int64, list2 []string) map[int64]string {
	newMap := make(map[int64]string)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt64Bool takes two inputs: first list of type: []int64, second list of type: []bool.
// Then it merges two list and returns a new map of type: map[int64]bool
func ZipInt64Bool(list1 []int64, list2 []bool) map[int64]bool {
	newMap := make(map[int64]bool)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt64Float32 takes two inputs: first list of type: []int64, second list of type: []float32.
// Then it merges two list and returns a new map of type: map[int64]float32
func ZipInt64Float32(list1 []int64, list2 []float32) map[int64]float32 {
	newMap := make(map[int64]float32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt64Float64 takes two inputs: first list of type: []int64, second list of type: []float64.
// Then it merges two list and returns a new map of type: map[int64]float64
func ZipInt64Float64(list1 []int64, list2 []float64) map[int64]float64 {
	newMap := make(map[int64]float64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt32Int takes two inputs: first list of type: []int32, second list of type: []int.
// Then it merges two list and returns a new map of type: map[int32]int
func ZipInt32Int(list1 []int32, list2 []int) map[int32]int {
	newMap := make(map[int32]int)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt32Int64 takes two inputs: first list of type: []int32, second list of type: []int64.
// Then it merges two list and returns a new map of type: map[int32]int64
func ZipInt32Int64(list1 []int32, list2 []int64) map[int32]int64 {
	newMap := make(map[int32]int64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt32 takes two inputs: first list of type: []int32, second list of type: []int32.
// Then it merges two list and returns a new map of type: map[int32]int32
func ZipInt32(list1 []int32, list2 []int32) map[int32]int32 {
	newMap := make(map[int32]int32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt32Int16 takes two inputs: first list of type: []int32, second list of type: []int16.
// Then it merges two list and returns a new map of type: map[int32]int16
func ZipInt32Int16(list1 []int32, list2 []int16) map[int32]int16 {
	newMap := make(map[int32]int16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt32Int8 takes two inputs: first list of type: []int32, second list of type: []int8.
// Then it merges two list and returns a new map of type: map[int32]int8
func ZipInt32Int8(list1 []int32, list2 []int8) map[int32]int8 {
	newMap := make(map[int32]int8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt32Uint takes two inputs: first list of type: []int32, second list of type: []uint.
// Then it merges two list and returns a new map of type: map[int32]uint
func ZipInt32Uint(list1 []int32, list2 []uint) map[int32]uint {
	newMap := make(map[int32]uint)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt32Uint64 takes two inputs: first list of type: []int32, second list of type: []uint64.
// Then it merges two list and returns a new map of type: map[int32]uint64
func ZipInt32Uint64(list1 []int32, list2 []uint64) map[int32]uint64 {
	newMap := make(map[int32]uint64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt32Uint32 takes two inputs: first list of type: []int32, second list of type: []uint32.
// Then it merges two list and returns a new map of type: map[int32]uint32
func ZipInt32Uint32(list1 []int32, list2 []uint32) map[int32]uint32 {
	newMap := make(map[int32]uint32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt32Uint16 takes two inputs: first list of type: []int32, second list of type: []uint16.
// Then it merges two list and returns a new map of type: map[int32]uint16
func ZipInt32Uint16(list1 []int32, list2 []uint16) map[int32]uint16 {
	newMap := make(map[int32]uint16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt32Uint8 takes two inputs: first list of type: []int32, second list of type: []uint8.
// Then it merges two list and returns a new map of type: map[int32]uint8
func ZipInt32Uint8(list1 []int32, list2 []uint8) map[int32]uint8 {
	newMap := make(map[int32]uint8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt32Str takes two inputs: first list of type: []int32, second list of type: []string.
// Then it merges two list and returns a new map of type: map[int32]string
func ZipInt32Str(list1 []int32, list2 []string) map[int32]string {
	newMap := make(map[int32]string)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt32Bool takes two inputs: first list of type: []int32, second list of type: []bool.
// Then it merges two list and returns a new map of type: map[int32]bool
func ZipInt32Bool(list1 []int32, list2 []bool) map[int32]bool {
	newMap := make(map[int32]bool)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt32Float32 takes two inputs: first list of type: []int32, second list of type: []float32.
// Then it merges two list and returns a new map of type: map[int32]float32
func ZipInt32Float32(list1 []int32, list2 []float32) map[int32]float32 {
	newMap := make(map[int32]float32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt32Float64 takes two inputs: first list of type: []int32, second list of type: []float64.
// Then it merges two list and returns a new map of type: map[int32]float64
func ZipInt32Float64(list1 []int32, list2 []float64) map[int32]float64 {
	newMap := make(map[int32]float64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt16Int takes two inputs: first list of type: []int16, second list of type: []int.
// Then it merges two list and returns a new map of type: map[int16]int
func ZipInt16Int(list1 []int16, list2 []int) map[int16]int {
	newMap := make(map[int16]int)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt16Int64 takes two inputs: first list of type: []int16, second list of type: []int64.
// Then it merges two list and returns a new map of type: map[int16]int64
func ZipInt16Int64(list1 []int16, list2 []int64) map[int16]int64 {
	newMap := make(map[int16]int64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt16Int32 takes two inputs: first list of type: []int16, second list of type: []int32.
// Then it merges two list and returns a new map of type: map[int16]int32
func ZipInt16Int32(list1 []int16, list2 []int32) map[int16]int32 {
	newMap := make(map[int16]int32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt16 takes two inputs: first list of type: []int16, second list of type: []int16.
// Then it merges two list and returns a new map of type: map[int16]int16
func ZipInt16(list1 []int16, list2 []int16) map[int16]int16 {
	newMap := make(map[int16]int16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt16Int8 takes two inputs: first list of type: []int16, second list of type: []int8.
// Then it merges two list and returns a new map of type: map[int16]int8
func ZipInt16Int8(list1 []int16, list2 []int8) map[int16]int8 {
	newMap := make(map[int16]int8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt16Uint takes two inputs: first list of type: []int16, second list of type: []uint.
// Then it merges two list and returns a new map of type: map[int16]uint
func ZipInt16Uint(list1 []int16, list2 []uint) map[int16]uint {
	newMap := make(map[int16]uint)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt16Uint64 takes two inputs: first list of type: []int16, second list of type: []uint64.
// Then it merges two list and returns a new map of type: map[int16]uint64
func ZipInt16Uint64(list1 []int16, list2 []uint64) map[int16]uint64 {
	newMap := make(map[int16]uint64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt16Uint32 takes two inputs: first list of type: []int16, second list of type: []uint32.
// Then it merges two list and returns a new map of type: map[int16]uint32
func ZipInt16Uint32(list1 []int16, list2 []uint32) map[int16]uint32 {
	newMap := make(map[int16]uint32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt16Uint16 takes two inputs: first list of type: []int16, second list of type: []uint16.
// Then it merges two list and returns a new map of type: map[int16]uint16
func ZipInt16Uint16(list1 []int16, list2 []uint16) map[int16]uint16 {
	newMap := make(map[int16]uint16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt16Uint8 takes two inputs: first list of type: []int16, second list of type: []uint8.
// Then it merges two list and returns a new map of type: map[int16]uint8
func ZipInt16Uint8(list1 []int16, list2 []uint8) map[int16]uint8 {
	newMap := make(map[int16]uint8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt16Str takes two inputs: first list of type: []int16, second list of type: []string.
// Then it merges two list and returns a new map of type: map[int16]string
func ZipInt16Str(list1 []int16, list2 []string) map[int16]string {
	newMap := make(map[int16]string)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt16Bool takes two inputs: first list of type: []int16, second list of type: []bool.
// Then it merges two list and returns a new map of type: map[int16]bool
func ZipInt16Bool(list1 []int16, list2 []bool) map[int16]bool {
	newMap := make(map[int16]bool)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt16Float32 takes two inputs: first list of type: []int16, second list of type: []float32.
// Then it merges two list and returns a new map of type: map[int16]float32
func ZipInt16Float32(list1 []int16, list2 []float32) map[int16]float32 {
	newMap := make(map[int16]float32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt16Float64 takes two inputs: first list of type: []int16, second list of type: []float64.
// Then it merges two list and returns a new map of type: map[int16]float64
func ZipInt16Float64(list1 []int16, list2 []float64) map[int16]float64 {
	newMap := make(map[int16]float64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt8Int takes two inputs: first list of type: []int8, second list of type: []int.
// Then it merges two list and returns a new map of type: map[int8]int
func ZipInt8Int(list1 []int8, list2 []int) map[int8]int {
	newMap := make(map[int8]int)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt8Int64 takes two inputs: first list of type: []int8, second list of type: []int64.
// Then it merges two list and returns a new map of type: map[int8]int64
func ZipInt8Int64(list1 []int8, list2 []int64) map[int8]int64 {
	newMap := make(map[int8]int64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt8Int32 takes two inputs: first list of type: []int8, second list of type: []int32.
// Then it merges two list and returns a new map of type: map[int8]int32
func ZipInt8Int32(list1 []int8, list2 []int32) map[int8]int32 {
	newMap := make(map[int8]int32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt8Int16 takes two inputs: first list of type: []int8, second list of type: []int16.
// Then it merges two list and returns a new map of type: map[int8]int16
func ZipInt8Int16(list1 []int8, list2 []int16) map[int8]int16 {
	newMap := make(map[int8]int16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt8 takes two inputs: first list of type: []int8, second list of type: []int8.
// Then it merges two list and returns a new map of type: map[int8]int8
func ZipInt8(list1 []int8, list2 []int8) map[int8]int8 {
	newMap := make(map[int8]int8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt8Uint takes two inputs: first list of type: []int8, second list of type: []uint.
// Then it merges two list and returns a new map of type: map[int8]uint
func ZipInt8Uint(list1 []int8, list2 []uint) map[int8]uint {
	newMap := make(map[int8]uint)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt8Uint64 takes two inputs: first list of type: []int8, second list of type: []uint64.
// Then it merges two list and returns a new map of type: map[int8]uint64
func ZipInt8Uint64(list1 []int8, list2 []uint64) map[int8]uint64 {
	newMap := make(map[int8]uint64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt8Uint32 takes two inputs: first list of type: []int8, second list of type: []uint32.
// Then it merges two list and returns a new map of type: map[int8]uint32
func ZipInt8Uint32(list1 []int8, list2 []uint32) map[int8]uint32 {
	newMap := make(map[int8]uint32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt8Uint16 takes two inputs: first list of type: []int8, second list of type: []uint16.
// Then it merges two list and returns a new map of type: map[int8]uint16
func ZipInt8Uint16(list1 []int8, list2 []uint16) map[int8]uint16 {
	newMap := make(map[int8]uint16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt8Uint8 takes two inputs: first list of type: []int8, second list of type: []uint8.
// Then it merges two list and returns a new map of type: map[int8]uint8
func ZipInt8Uint8(list1 []int8, list2 []uint8) map[int8]uint8 {
	newMap := make(map[int8]uint8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt8Str takes two inputs: first list of type: []int8, second list of type: []string.
// Then it merges two list and returns a new map of type: map[int8]string
func ZipInt8Str(list1 []int8, list2 []string) map[int8]string {
	newMap := make(map[int8]string)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt8Bool takes two inputs: first list of type: []int8, second list of type: []bool.
// Then it merges two list and returns a new map of type: map[int8]bool
func ZipInt8Bool(list1 []int8, list2 []bool) map[int8]bool {
	newMap := make(map[int8]bool)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt8Float32 takes two inputs: first list of type: []int8, second list of type: []float32.
// Then it merges two list and returns a new map of type: map[int8]float32
func ZipInt8Float32(list1 []int8, list2 []float32) map[int8]float32 {
	newMap := make(map[int8]float32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipInt8Float64 takes two inputs: first list of type: []int8, second list of type: []float64.
// Then it merges two list and returns a new map of type: map[int8]float64
func ZipInt8Float64(list1 []int8, list2 []float64) map[int8]float64 {
	newMap := make(map[int8]float64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUintInt takes two inputs: first list of type: []uint, second list of type: []int.
// Then it merges two list and returns a new map of type: map[uint]int
func ZipUintInt(list1 []uint, list2 []int) map[uint]int {
	newMap := make(map[uint]int)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUintInt64 takes two inputs: first list of type: []uint, second list of type: []int64.
// Then it merges two list and returns a new map of type: map[uint]int64
func ZipUintInt64(list1 []uint, list2 []int64) map[uint]int64 {
	newMap := make(map[uint]int64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUintInt32 takes two inputs: first list of type: []uint, second list of type: []int32.
// Then it merges two list and returns a new map of type: map[uint]int32
func ZipUintInt32(list1 []uint, list2 []int32) map[uint]int32 {
	newMap := make(map[uint]int32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUintInt16 takes two inputs: first list of type: []uint, second list of type: []int16.
// Then it merges two list and returns a new map of type: map[uint]int16
func ZipUintInt16(list1 []uint, list2 []int16) map[uint]int16 {
	newMap := make(map[uint]int16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUintInt8 takes two inputs: first list of type: []uint, second list of type: []int8.
// Then it merges two list and returns a new map of type: map[uint]int8
func ZipUintInt8(list1 []uint, list2 []int8) map[uint]int8 {
	newMap := make(map[uint]int8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint takes two inputs: first list of type: []uint, second list of type: []uint.
// Then it merges two list and returns a new map of type: map[uint]uint
func ZipUint(list1 []uint, list2 []uint) map[uint]uint {
	newMap := make(map[uint]uint)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUintUint64 takes two inputs: first list of type: []uint, second list of type: []uint64.
// Then it merges two list and returns a new map of type: map[uint]uint64
func ZipUintUint64(list1 []uint, list2 []uint64) map[uint]uint64 {
	newMap := make(map[uint]uint64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUintUint32 takes two inputs: first list of type: []uint, second list of type: []uint32.
// Then it merges two list and returns a new map of type: map[uint]uint32
func ZipUintUint32(list1 []uint, list2 []uint32) map[uint]uint32 {
	newMap := make(map[uint]uint32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUintUint16 takes two inputs: first list of type: []uint, second list of type: []uint16.
// Then it merges two list and returns a new map of type: map[uint]uint16
func ZipUintUint16(list1 []uint, list2 []uint16) map[uint]uint16 {
	newMap := make(map[uint]uint16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUintUint8 takes two inputs: first list of type: []uint, second list of type: []uint8.
// Then it merges two list and returns a new map of type: map[uint]uint8
func ZipUintUint8(list1 []uint, list2 []uint8) map[uint]uint8 {
	newMap := make(map[uint]uint8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUintStr takes two inputs: first list of type: []uint, second list of type: []string.
// Then it merges two list and returns a new map of type: map[uint]string
func ZipUintStr(list1 []uint, list2 []string) map[uint]string {
	newMap := make(map[uint]string)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUintBool takes two inputs: first list of type: []uint, second list of type: []bool.
// Then it merges two list and returns a new map of type: map[uint]bool
func ZipUintBool(list1 []uint, list2 []bool) map[uint]bool {
	newMap := make(map[uint]bool)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUintFloat32 takes two inputs: first list of type: []uint, second list of type: []float32.
// Then it merges two list and returns a new map of type: map[uint]float32
func ZipUintFloat32(list1 []uint, list2 []float32) map[uint]float32 {
	newMap := make(map[uint]float32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUintFloat64 takes two inputs: first list of type: []uint, second list of type: []float64.
// Then it merges two list and returns a new map of type: map[uint]float64
func ZipUintFloat64(list1 []uint, list2 []float64) map[uint]float64 {
	newMap := make(map[uint]float64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint64Int takes two inputs: first list of type: []uint64, second list of type: []int.
// Then it merges two list and returns a new map of type: map[uint64]int
func ZipUint64Int(list1 []uint64, list2 []int) map[uint64]int {
	newMap := make(map[uint64]int)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint64Int64 takes two inputs: first list of type: []uint64, second list of type: []int64.
// Then it merges two list and returns a new map of type: map[uint64]int64
func ZipUint64Int64(list1 []uint64, list2 []int64) map[uint64]int64 {
	newMap := make(map[uint64]int64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint64Int32 takes two inputs: first list of type: []uint64, second list of type: []int32.
// Then it merges two list and returns a new map of type: map[uint64]int32
func ZipUint64Int32(list1 []uint64, list2 []int32) map[uint64]int32 {
	newMap := make(map[uint64]int32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint64Int16 takes two inputs: first list of type: []uint64, second list of type: []int16.
// Then it merges two list and returns a new map of type: map[uint64]int16
func ZipUint64Int16(list1 []uint64, list2 []int16) map[uint64]int16 {
	newMap := make(map[uint64]int16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint64Int8 takes two inputs: first list of type: []uint64, second list of type: []int8.
// Then it merges two list and returns a new map of type: map[uint64]int8
func ZipUint64Int8(list1 []uint64, list2 []int8) map[uint64]int8 {
	newMap := make(map[uint64]int8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint64Uint takes two inputs: first list of type: []uint64, second list of type: []uint.
// Then it merges two list and returns a new map of type: map[uint64]uint
func ZipUint64Uint(list1 []uint64, list2 []uint) map[uint64]uint {
	newMap := make(map[uint64]uint)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint64 takes two inputs: first list of type: []uint64, second list of type: []uint64.
// Then it merges two list and returns a new map of type: map[uint64]uint64
func ZipUint64(list1 []uint64, list2 []uint64) map[uint64]uint64 {
	newMap := make(map[uint64]uint64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint64Uint32 takes two inputs: first list of type: []uint64, second list of type: []uint32.
// Then it merges two list and returns a new map of type: map[uint64]uint32
func ZipUint64Uint32(list1 []uint64, list2 []uint32) map[uint64]uint32 {
	newMap := make(map[uint64]uint32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint64Uint16 takes two inputs: first list of type: []uint64, second list of type: []uint16.
// Then it merges two list and returns a new map of type: map[uint64]uint16
func ZipUint64Uint16(list1 []uint64, list2 []uint16) map[uint64]uint16 {
	newMap := make(map[uint64]uint16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint64Uint8 takes two inputs: first list of type: []uint64, second list of type: []uint8.
// Then it merges two list and returns a new map of type: map[uint64]uint8
func ZipUint64Uint8(list1 []uint64, list2 []uint8) map[uint64]uint8 {
	newMap := make(map[uint64]uint8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint64Str takes two inputs: first list of type: []uint64, second list of type: []string.
// Then it merges two list and returns a new map of type: map[uint64]string
func ZipUint64Str(list1 []uint64, list2 []string) map[uint64]string {
	newMap := make(map[uint64]string)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint64Bool takes two inputs: first list of type: []uint64, second list of type: []bool.
// Then it merges two list and returns a new map of type: map[uint64]bool
func ZipUint64Bool(list1 []uint64, list2 []bool) map[uint64]bool {
	newMap := make(map[uint64]bool)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint64Float32 takes two inputs: first list of type: []uint64, second list of type: []float32.
// Then it merges two list and returns a new map of type: map[uint64]float32
func ZipUint64Float32(list1 []uint64, list2 []float32) map[uint64]float32 {
	newMap := make(map[uint64]float32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint64Float64 takes two inputs: first list of type: []uint64, second list of type: []float64.
// Then it merges two list and returns a new map of type: map[uint64]float64
func ZipUint64Float64(list1 []uint64, list2 []float64) map[uint64]float64 {
	newMap := make(map[uint64]float64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint32Int takes two inputs: first list of type: []uint32, second list of type: []int.
// Then it merges two list and returns a new map of type: map[uint32]int
func ZipUint32Int(list1 []uint32, list2 []int) map[uint32]int {
	newMap := make(map[uint32]int)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint32Int64 takes two inputs: first list of type: []uint32, second list of type: []int64.
// Then it merges two list and returns a new map of type: map[uint32]int64
func ZipUint32Int64(list1 []uint32, list2 []int64) map[uint32]int64 {
	newMap := make(map[uint32]int64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint32Int32 takes two inputs: first list of type: []uint32, second list of type: []int32.
// Then it merges two list and returns a new map of type: map[uint32]int32
func ZipUint32Int32(list1 []uint32, list2 []int32) map[uint32]int32 {
	newMap := make(map[uint32]int32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint32Int16 takes two inputs: first list of type: []uint32, second list of type: []int16.
// Then it merges two list and returns a new map of type: map[uint32]int16
func ZipUint32Int16(list1 []uint32, list2 []int16) map[uint32]int16 {
	newMap := make(map[uint32]int16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint32Int8 takes two inputs: first list of type: []uint32, second list of type: []int8.
// Then it merges two list and returns a new map of type: map[uint32]int8
func ZipUint32Int8(list1 []uint32, list2 []int8) map[uint32]int8 {
	newMap := make(map[uint32]int8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint32Uint takes two inputs: first list of type: []uint32, second list of type: []uint.
// Then it merges two list and returns a new map of type: map[uint32]uint
func ZipUint32Uint(list1 []uint32, list2 []uint) map[uint32]uint {
	newMap := make(map[uint32]uint)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint32Uint64 takes two inputs: first list of type: []uint32, second list of type: []uint64.
// Then it merges two list and returns a new map of type: map[uint32]uint64
func ZipUint32Uint64(list1 []uint32, list2 []uint64) map[uint32]uint64 {
	newMap := make(map[uint32]uint64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint32 takes two inputs: first list of type: []uint32, second list of type: []uint32.
// Then it merges two list and returns a new map of type: map[uint32]uint32
func ZipUint32(list1 []uint32, list2 []uint32) map[uint32]uint32 {
	newMap := make(map[uint32]uint32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint32Uint16 takes two inputs: first list of type: []uint32, second list of type: []uint16.
// Then it merges two list and returns a new map of type: map[uint32]uint16
func ZipUint32Uint16(list1 []uint32, list2 []uint16) map[uint32]uint16 {
	newMap := make(map[uint32]uint16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint32Uint8 takes two inputs: first list of type: []uint32, second list of type: []uint8.
// Then it merges two list and returns a new map of type: map[uint32]uint8
func ZipUint32Uint8(list1 []uint32, list2 []uint8) map[uint32]uint8 {
	newMap := make(map[uint32]uint8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint32Str takes two inputs: first list of type: []uint32, second list of type: []string.
// Then it merges two list and returns a new map of type: map[uint32]string
func ZipUint32Str(list1 []uint32, list2 []string) map[uint32]string {
	newMap := make(map[uint32]string)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint32Bool takes two inputs: first list of type: []uint32, second list of type: []bool.
// Then it merges two list and returns a new map of type: map[uint32]bool
func ZipUint32Bool(list1 []uint32, list2 []bool) map[uint32]bool {
	newMap := make(map[uint32]bool)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint32Float32 takes two inputs: first list of type: []uint32, second list of type: []float32.
// Then it merges two list and returns a new map of type: map[uint32]float32
func ZipUint32Float32(list1 []uint32, list2 []float32) map[uint32]float32 {
	newMap := make(map[uint32]float32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint32Float64 takes two inputs: first list of type: []uint32, second list of type: []float64.
// Then it merges two list and returns a new map of type: map[uint32]float64
func ZipUint32Float64(list1 []uint32, list2 []float64) map[uint32]float64 {
	newMap := make(map[uint32]float64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint16Int takes two inputs: first list of type: []uint16, second list of type: []int.
// Then it merges two list and returns a new map of type: map[uint16]int
func ZipUint16Int(list1 []uint16, list2 []int) map[uint16]int {
	newMap := make(map[uint16]int)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint16Int64 takes two inputs: first list of type: []uint16, second list of type: []int64.
// Then it merges two list and returns a new map of type: map[uint16]int64
func ZipUint16Int64(list1 []uint16, list2 []int64) map[uint16]int64 {
	newMap := make(map[uint16]int64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint16Int32 takes two inputs: first list of type: []uint16, second list of type: []int32.
// Then it merges two list and returns a new map of type: map[uint16]int32
func ZipUint16Int32(list1 []uint16, list2 []int32) map[uint16]int32 {
	newMap := make(map[uint16]int32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint16Int16 takes two inputs: first list of type: []uint16, second list of type: []int16.
// Then it merges two list and returns a new map of type: map[uint16]int16
func ZipUint16Int16(list1 []uint16, list2 []int16) map[uint16]int16 {
	newMap := make(map[uint16]int16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint16Int8 takes two inputs: first list of type: []uint16, second list of type: []int8.
// Then it merges two list and returns a new map of type: map[uint16]int8
func ZipUint16Int8(list1 []uint16, list2 []int8) map[uint16]int8 {
	newMap := make(map[uint16]int8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint16Uint takes two inputs: first list of type: []uint16, second list of type: []uint.
// Then it merges two list and returns a new map of type: map[uint16]uint
func ZipUint16Uint(list1 []uint16, list2 []uint) map[uint16]uint {
	newMap := make(map[uint16]uint)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint16Uint64 takes two inputs: first list of type: []uint16, second list of type: []uint64.
// Then it merges two list and returns a new map of type: map[uint16]uint64
func ZipUint16Uint64(list1 []uint16, list2 []uint64) map[uint16]uint64 {
	newMap := make(map[uint16]uint64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint16Uint32 takes two inputs: first list of type: []uint16, second list of type: []uint32.
// Then it merges two list and returns a new map of type: map[uint16]uint32
func ZipUint16Uint32(list1 []uint16, list2 []uint32) map[uint16]uint32 {
	newMap := make(map[uint16]uint32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint16 takes two inputs: first list of type: []uint16, second list of type: []uint16.
// Then it merges two list and returns a new map of type: map[uint16]uint16
func ZipUint16(list1 []uint16, list2 []uint16) map[uint16]uint16 {
	newMap := make(map[uint16]uint16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint16Uint8 takes two inputs: first list of type: []uint16, second list of type: []uint8.
// Then it merges two list and returns a new map of type: map[uint16]uint8
func ZipUint16Uint8(list1 []uint16, list2 []uint8) map[uint16]uint8 {
	newMap := make(map[uint16]uint8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint16Str takes two inputs: first list of type: []uint16, second list of type: []string.
// Then it merges two list and returns a new map of type: map[uint16]string
func ZipUint16Str(list1 []uint16, list2 []string) map[uint16]string {
	newMap := make(map[uint16]string)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint16Bool takes two inputs: first list of type: []uint16, second list of type: []bool.
// Then it merges two list and returns a new map of type: map[uint16]bool
func ZipUint16Bool(list1 []uint16, list2 []bool) map[uint16]bool {
	newMap := make(map[uint16]bool)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint16Float32 takes two inputs: first list of type: []uint16, second list of type: []float32.
// Then it merges two list and returns a new map of type: map[uint16]float32
func ZipUint16Float32(list1 []uint16, list2 []float32) map[uint16]float32 {
	newMap := make(map[uint16]float32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint16Float64 takes two inputs: first list of type: []uint16, second list of type: []float64.
// Then it merges two list and returns a new map of type: map[uint16]float64
func ZipUint16Float64(list1 []uint16, list2 []float64) map[uint16]float64 {
	newMap := make(map[uint16]float64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint8Int takes two inputs: first list of type: []uint8, second list of type: []int.
// Then it merges two list and returns a new map of type: map[uint8]int
func ZipUint8Int(list1 []uint8, list2 []int) map[uint8]int {
	newMap := make(map[uint8]int)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint8Int64 takes two inputs: first list of type: []uint8, second list of type: []int64.
// Then it merges two list and returns a new map of type: map[uint8]int64
func ZipUint8Int64(list1 []uint8, list2 []int64) map[uint8]int64 {
	newMap := make(map[uint8]int64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint8Int32 takes two inputs: first list of type: []uint8, second list of type: []int32.
// Then it merges two list and returns a new map of type: map[uint8]int32
func ZipUint8Int32(list1 []uint8, list2 []int32) map[uint8]int32 {
	newMap := make(map[uint8]int32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint8Int16 takes two inputs: first list of type: []uint8, second list of type: []int16.
// Then it merges two list and returns a new map of type: map[uint8]int16
func ZipUint8Int16(list1 []uint8, list2 []int16) map[uint8]int16 {
	newMap := make(map[uint8]int16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint8Int8 takes two inputs: first list of type: []uint8, second list of type: []int8.
// Then it merges two list and returns a new map of type: map[uint8]int8
func ZipUint8Int8(list1 []uint8, list2 []int8) map[uint8]int8 {
	newMap := make(map[uint8]int8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint8Uint takes two inputs: first list of type: []uint8, second list of type: []uint.
// Then it merges two list and returns a new map of type: map[uint8]uint
func ZipUint8Uint(list1 []uint8, list2 []uint) map[uint8]uint {
	newMap := make(map[uint8]uint)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint8Uint64 takes two inputs: first list of type: []uint8, second list of type: []uint64.
// Then it merges two list and returns a new map of type: map[uint8]uint64
func ZipUint8Uint64(list1 []uint8, list2 []uint64) map[uint8]uint64 {
	newMap := make(map[uint8]uint64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint8Uint32 takes two inputs: first list of type: []uint8, second list of type: []uint32.
// Then it merges two list and returns a new map of type: map[uint8]uint32
func ZipUint8Uint32(list1 []uint8, list2 []uint32) map[uint8]uint32 {
	newMap := make(map[uint8]uint32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint8Uint16 takes two inputs: first list of type: []uint8, second list of type: []uint16.
// Then it merges two list and returns a new map of type: map[uint8]uint16
func ZipUint8Uint16(list1 []uint8, list2 []uint16) map[uint8]uint16 {
	newMap := make(map[uint8]uint16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint8 takes two inputs: first list of type: []uint8, second list of type: []uint8.
// Then it merges two list and returns a new map of type: map[uint8]uint8
func ZipUint8(list1 []uint8, list2 []uint8) map[uint8]uint8 {
	newMap := make(map[uint8]uint8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint8Str takes two inputs: first list of type: []uint8, second list of type: []string.
// Then it merges two list and returns a new map of type: map[uint8]string
func ZipUint8Str(list1 []uint8, list2 []string) map[uint8]string {
	newMap := make(map[uint8]string)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint8Bool takes two inputs: first list of type: []uint8, second list of type: []bool.
// Then it merges two list and returns a new map of type: map[uint8]bool
func ZipUint8Bool(list1 []uint8, list2 []bool) map[uint8]bool {
	newMap := make(map[uint8]bool)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint8Float32 takes two inputs: first list of type: []uint8, second list of type: []float32.
// Then it merges two list and returns a new map of type: map[uint8]float32
func ZipUint8Float32(list1 []uint8, list2 []float32) map[uint8]float32 {
	newMap := make(map[uint8]float32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipUint8Float64 takes two inputs: first list of type: []uint8, second list of type: []float64.
// Then it merges two list and returns a new map of type: map[uint8]float64
func ZipUint8Float64(list1 []uint8, list2 []float64) map[uint8]float64 {
	newMap := make(map[uint8]float64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipStrInt takes two inputs: first list of type: []string, second list of type: []int.
// Then it merges two list and returns a new map of type: map[string]int
func ZipStrInt(list1 []string, list2 []int) map[string]int {
	newMap := make(map[string]int)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipStrInt64 takes two inputs: first list of type: []string, second list of type: []int64.
// Then it merges two list and returns a new map of type: map[string]int64
func ZipStrInt64(list1 []string, list2 []int64) map[string]int64 {
	newMap := make(map[string]int64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipStrInt32 takes two inputs: first list of type: []string, second list of type: []int32.
// Then it merges two list and returns a new map of type: map[string]int32
func ZipStrInt32(list1 []string, list2 []int32) map[string]int32 {
	newMap := make(map[string]int32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipStrInt16 takes two inputs: first list of type: []string, second list of type: []int16.
// Then it merges two list and returns a new map of type: map[string]int16
func ZipStrInt16(list1 []string, list2 []int16) map[string]int16 {
	newMap := make(map[string]int16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipStrInt8 takes two inputs: first list of type: []string, second list of type: []int8.
// Then it merges two list and returns a new map of type: map[string]int8
func ZipStrInt8(list1 []string, list2 []int8) map[string]int8 {
	newMap := make(map[string]int8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipStrUint takes two inputs: first list of type: []string, second list of type: []uint.
// Then it merges two list and returns a new map of type: map[string]uint
func ZipStrUint(list1 []string, list2 []uint) map[string]uint {
	newMap := make(map[string]uint)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipStrUint64 takes two inputs: first list of type: []string, second list of type: []uint64.
// Then it merges two list and returns a new map of type: map[string]uint64
func ZipStrUint64(list1 []string, list2 []uint64) map[string]uint64 {
	newMap := make(map[string]uint64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipStrUint32 takes two inputs: first list of type: []string, second list of type: []uint32.
// Then it merges two list and returns a new map of type: map[string]uint32
func ZipStrUint32(list1 []string, list2 []uint32) map[string]uint32 {
	newMap := make(map[string]uint32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipStrUint16 takes two inputs: first list of type: []string, second list of type: []uint16.
// Then it merges two list and returns a new map of type: map[string]uint16
func ZipStrUint16(list1 []string, list2 []uint16) map[string]uint16 {
	newMap := make(map[string]uint16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipStrUint8 takes two inputs: first list of type: []string, second list of type: []uint8.
// Then it merges two list and returns a new map of type: map[string]uint8
func ZipStrUint8(list1 []string, list2 []uint8) map[string]uint8 {
	newMap := make(map[string]uint8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipStr takes two inputs: first list of type: []string, second list of type: []string.
// Then it merges two list and returns a new map of type: map[string]string
func ZipStr(list1 []string, list2 []string) map[string]string {
	newMap := make(map[string]string)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipStrBool takes two inputs: first list of type: []string, second list of type: []bool.
// Then it merges two list and returns a new map of type: map[string]bool
func ZipStrBool(list1 []string, list2 []bool) map[string]bool {
	newMap := make(map[string]bool)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipStrFloat32 takes two inputs: first list of type: []string, second list of type: []float32.
// Then it merges two list and returns a new map of type: map[string]float32
func ZipStrFloat32(list1 []string, list2 []float32) map[string]float32 {
	newMap := make(map[string]float32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipStrFloat64 takes two inputs: first list of type: []string, second list of type: []float64.
// Then it merges two list and returns a new map of type: map[string]float64
func ZipStrFloat64(list1 []string, list2 []float64) map[string]float64 {
	newMap := make(map[string]float64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipBoolInt takes two inputs: first list of type: []bool, second list of type: []int.
// Then it merges two list and returns a new map of type: map[bool]int
func ZipBoolInt(list1 []bool, list2 []int) map[bool]int {
	newMap := make(map[bool]int)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipBoolInt64 takes two inputs: first list of type: []bool, second list of type: []int64.
// Then it merges two list and returns a new map of type: map[bool]int64
func ZipBoolInt64(list1 []bool, list2 []int64) map[bool]int64 {
	newMap := make(map[bool]int64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipBoolInt32 takes two inputs: first list of type: []bool, second list of type: []int32.
// Then it merges two list and returns a new map of type: map[bool]int32
func ZipBoolInt32(list1 []bool, list2 []int32) map[bool]int32 {
	newMap := make(map[bool]int32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipBoolInt16 takes two inputs: first list of type: []bool, second list of type: []int16.
// Then it merges two list and returns a new map of type: map[bool]int16
func ZipBoolInt16(list1 []bool, list2 []int16) map[bool]int16 {
	newMap := make(map[bool]int16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipBoolInt8 takes two inputs: first list of type: []bool, second list of type: []int8.
// Then it merges two list and returns a new map of type: map[bool]int8
func ZipBoolInt8(list1 []bool, list2 []int8) map[bool]int8 {
	newMap := make(map[bool]int8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipBoolUint takes two inputs: first list of type: []bool, second list of type: []uint.
// Then it merges two list and returns a new map of type: map[bool]uint
func ZipBoolUint(list1 []bool, list2 []uint) map[bool]uint {
	newMap := make(map[bool]uint)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipBoolUint64 takes two inputs: first list of type: []bool, second list of type: []uint64.
// Then it merges two list and returns a new map of type: map[bool]uint64
func ZipBoolUint64(list1 []bool, list2 []uint64) map[bool]uint64 {
	newMap := make(map[bool]uint64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipBoolUint32 takes two inputs: first list of type: []bool, second list of type: []uint32.
// Then it merges two list and returns a new map of type: map[bool]uint32
func ZipBoolUint32(list1 []bool, list2 []uint32) map[bool]uint32 {
	newMap := make(map[bool]uint32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipBoolUint16 takes two inputs: first list of type: []bool, second list of type: []uint16.
// Then it merges two list and returns a new map of type: map[bool]uint16
func ZipBoolUint16(list1 []bool, list2 []uint16) map[bool]uint16 {
	newMap := make(map[bool]uint16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipBoolUint8 takes two inputs: first list of type: []bool, second list of type: []uint8.
// Then it merges two list and returns a new map of type: map[bool]uint8
func ZipBoolUint8(list1 []bool, list2 []uint8) map[bool]uint8 {
	newMap := make(map[bool]uint8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipBoolStr takes two inputs: first list of type: []bool, second list of type: []string.
// Then it merges two list and returns a new map of type: map[bool]string
func ZipBoolStr(list1 []bool, list2 []string) map[bool]string {
	newMap := make(map[bool]string)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipBool takes two inputs: first list of type: []bool, second list of type: []bool.
// Then it merges two list and returns a new map of type: map[bool]bool
func ZipBool(list1 []bool, list2 []bool) map[bool]bool {
	newMap := make(map[bool]bool)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipBoolFloat32 takes two inputs: first list of type: []bool, second list of type: []float32.
// Then it merges two list and returns a new map of type: map[bool]float32
func ZipBoolFloat32(list1 []bool, list2 []float32) map[bool]float32 {
	newMap := make(map[bool]float32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipBoolFloat64 takes two inputs: first list of type: []bool, second list of type: []float64.
// Then it merges two list and returns a new map of type: map[bool]float64
func ZipBoolFloat64(list1 []bool, list2 []float64) map[bool]float64 {
	newMap := make(map[bool]float64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat32Int takes two inputs: first list of type: []float32, second list of type: []int.
// Then it merges two list and returns a new map of type: map[float32]int
func ZipFloat32Int(list1 []float32, list2 []int) map[float32]int {
	newMap := make(map[float32]int)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat32Int64 takes two inputs: first list of type: []float32, second list of type: []int64.
// Then it merges two list and returns a new map of type: map[float32]int64
func ZipFloat32Int64(list1 []float32, list2 []int64) map[float32]int64 {
	newMap := make(map[float32]int64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat32Int32 takes two inputs: first list of type: []float32, second list of type: []int32.
// Then it merges two list and returns a new map of type: map[float32]int32
func ZipFloat32Int32(list1 []float32, list2 []int32) map[float32]int32 {
	newMap := make(map[float32]int32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat32Int16 takes two inputs: first list of type: []float32, second list of type: []int16.
// Then it merges two list and returns a new map of type: map[float32]int16
func ZipFloat32Int16(list1 []float32, list2 []int16) map[float32]int16 {
	newMap := make(map[float32]int16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat32Int8 takes two inputs: first list of type: []float32, second list of type: []int8.
// Then it merges two list and returns a new map of type: map[float32]int8
func ZipFloat32Int8(list1 []float32, list2 []int8) map[float32]int8 {
	newMap := make(map[float32]int8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat32Uint takes two inputs: first list of type: []float32, second list of type: []uint.
// Then it merges two list and returns a new map of type: map[float32]uint
func ZipFloat32Uint(list1 []float32, list2 []uint) map[float32]uint {
	newMap := make(map[float32]uint)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat32Uint64 takes two inputs: first list of type: []float32, second list of type: []uint64.
// Then it merges two list and returns a new map of type: map[float32]uint64
func ZipFloat32Uint64(list1 []float32, list2 []uint64) map[float32]uint64 {
	newMap := make(map[float32]uint64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat32Uint32 takes two inputs: first list of type: []float32, second list of type: []uint32.
// Then it merges two list and returns a new map of type: map[float32]uint32
func ZipFloat32Uint32(list1 []float32, list2 []uint32) map[float32]uint32 {
	newMap := make(map[float32]uint32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat32Uint16 takes two inputs: first list of type: []float32, second list of type: []uint16.
// Then it merges two list and returns a new map of type: map[float32]uint16
func ZipFloat32Uint16(list1 []float32, list2 []uint16) map[float32]uint16 {
	newMap := make(map[float32]uint16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat32Uint8 takes two inputs: first list of type: []float32, second list of type: []uint8.
// Then it merges two list and returns a new map of type: map[float32]uint8
func ZipFloat32Uint8(list1 []float32, list2 []uint8) map[float32]uint8 {
	newMap := make(map[float32]uint8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat32Str takes two inputs: first list of type: []float32, second list of type: []string.
// Then it merges two list and returns a new map of type: map[float32]string
func ZipFloat32Str(list1 []float32, list2 []string) map[float32]string {
	newMap := make(map[float32]string)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat32Bool takes two inputs: first list of type: []float32, second list of type: []bool.
// Then it merges two list and returns a new map of type: map[float32]bool
func ZipFloat32Bool(list1 []float32, list2 []bool) map[float32]bool {
	newMap := make(map[float32]bool)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat32 takes two inputs: first list of type: []float32, second list of type: []float32.
// Then it merges two list and returns a new map of type: map[float32]float32
func ZipFloat32(list1 []float32, list2 []float32) map[float32]float32 {
	newMap := make(map[float32]float32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat32Float64 takes two inputs: first list of type: []float32, second list of type: []float64.
// Then it merges two list and returns a new map of type: map[float32]float64
func ZipFloat32Float64(list1 []float32, list2 []float64) map[float32]float64 {
	newMap := make(map[float32]float64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat64Int takes two inputs: first list of type: []float64, second list of type: []int.
// Then it merges two list and returns a new map of type: map[float64]int
func ZipFloat64Int(list1 []float64, list2 []int) map[float64]int {
	newMap := make(map[float64]int)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat64Int64 takes two inputs: first list of type: []float64, second list of type: []int64.
// Then it merges two list and returns a new map of type: map[float64]int64
func ZipFloat64Int64(list1 []float64, list2 []int64) map[float64]int64 {
	newMap := make(map[float64]int64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat64Int32 takes two inputs: first list of type: []float64, second list of type: []int32.
// Then it merges two list and returns a new map of type: map[float64]int32
func ZipFloat64Int32(list1 []float64, list2 []int32) map[float64]int32 {
	newMap := make(map[float64]int32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat64Int16 takes two inputs: first list of type: []float64, second list of type: []int16.
// Then it merges two list and returns a new map of type: map[float64]int16
func ZipFloat64Int16(list1 []float64, list2 []int16) map[float64]int16 {
	newMap := make(map[float64]int16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat64Int8 takes two inputs: first list of type: []float64, second list of type: []int8.
// Then it merges two list and returns a new map of type: map[float64]int8
func ZipFloat64Int8(list1 []float64, list2 []int8) map[float64]int8 {
	newMap := make(map[float64]int8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat64Uint takes two inputs: first list of type: []float64, second list of type: []uint.
// Then it merges two list and returns a new map of type: map[float64]uint
func ZipFloat64Uint(list1 []float64, list2 []uint) map[float64]uint {
	newMap := make(map[float64]uint)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat64Uint64 takes two inputs: first list of type: []float64, second list of type: []uint64.
// Then it merges two list and returns a new map of type: map[float64]uint64
func ZipFloat64Uint64(list1 []float64, list2 []uint64) map[float64]uint64 {
	newMap := make(map[float64]uint64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat64Uint32 takes two inputs: first list of type: []float64, second list of type: []uint32.
// Then it merges two list and returns a new map of type: map[float64]uint32
func ZipFloat64Uint32(list1 []float64, list2 []uint32) map[float64]uint32 {
	newMap := make(map[float64]uint32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat64Uint16 takes two inputs: first list of type: []float64, second list of type: []uint16.
// Then it merges two list and returns a new map of type: map[float64]uint16
func ZipFloat64Uint16(list1 []float64, list2 []uint16) map[float64]uint16 {
	newMap := make(map[float64]uint16)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat64Uint8 takes two inputs: first list of type: []float64, second list of type: []uint8.
// Then it merges two list and returns a new map of type: map[float64]uint8
func ZipFloat64Uint8(list1 []float64, list2 []uint8) map[float64]uint8 {
	newMap := make(map[float64]uint8)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat64Str takes two inputs: first list of type: []float64, second list of type: []string.
// Then it merges two list and returns a new map of type: map[float64]string
func ZipFloat64Str(list1 []float64, list2 []string) map[float64]string {
	newMap := make(map[float64]string)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat64Bool takes two inputs: first list of type: []float64, second list of type: []bool.
// Then it merges two list and returns a new map of type: map[float64]bool
func ZipFloat64Bool(list1 []float64, list2 []bool) map[float64]bool {
	newMap := make(map[float64]bool)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat64Float32 takes two inputs: first list of type: []float64, second list of type: []float32.
// Then it merges two list and returns a new map of type: map[float64]float32
func ZipFloat64Float32(list1 []float64, list2 []float32) map[float64]float32 {
	newMap := make(map[float64]float32)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}

// ZipFloat64 takes two inputs: first list of type: []float64, second list of type: []float64.
// Then it merges two list and returns a new map of type: map[float64]float64
func ZipFloat64(list1 []float64, list2 []float64) map[float64]float64 {
	newMap := make(map[float64]float64)

	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 {
		return newMap
	}

	minLen := len1
	if len2 < minLen {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		newMap[list1[i]] = list2[i]
	}

	return newMap
}
