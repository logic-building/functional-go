package fp 

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
