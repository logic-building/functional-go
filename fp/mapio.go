package fp 

func MapIntInt64(f func(int) int64, list []int) []int64 {
	if f == nil {
		return []int64{}
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapIntInt32(f func(int) int32, list []int) []int32 {
	if f == nil {
		return []int32{}
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapIntInt16(f func(int) int16, list []int) []int16 {
	if f == nil {
		return []int16{}
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapIntInt8(f func(int) int8, list []int) []int8 {
	if f == nil {
		return []int8{}
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapIntUint(f func(int) uint, list []int) []uint {
	if f == nil {
		return []uint{}
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapIntUint64(f func(int) uint64, list []int) []uint64 {
	if f == nil {
		return []uint64{}
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapIntUint32(f func(int) uint32, list []int) []uint32 {
	if f == nil {
		return []uint32{}
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapIntUint16(f func(int) uint16, list []int) []uint16 {
	if f == nil {
		return []uint16{}
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapIntUint8(f func(int) uint8, list []int) []uint8 {
	if f == nil {
		return []uint8{}
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapIntStr(f func(int) string, list []int) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapIntBool(f func(int) bool, list []int) []bool {
	if f == nil {
		return []bool{}
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt64Int(f func(int64) int, list []int64) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt64Int32(f func(int64) int32, list []int64) []int32 {
	if f == nil {
		return []int32{}
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt64Int16(f func(int64) int16, list []int64) []int16 {
	if f == nil {
		return []int16{}
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt64Int8(f func(int64) int8, list []int64) []int8 {
	if f == nil {
		return []int8{}
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt64Uint(f func(int64) uint, list []int64) []uint {
	if f == nil {
		return []uint{}
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt64Uint64(f func(int64) uint64, list []int64) []uint64 {
	if f == nil {
		return []uint64{}
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt64Uint32(f func(int64) uint32, list []int64) []uint32 {
	if f == nil {
		return []uint32{}
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt64Uint16(f func(int64) uint16, list []int64) []uint16 {
	if f == nil {
		return []uint16{}
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt64Uint8(f func(int64) uint8, list []int64) []uint8 {
	if f == nil {
		return []uint8{}
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt64Str(f func(int64) string, list []int64) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt64Bool(f func(int64) bool, list []int64) []bool {
	if f == nil {
		return []bool{}
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt32Int(f func(int32) int, list []int32) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt32Int64(f func(int32) int64, list []int32) []int64 {
	if f == nil {
		return []int64{}
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt32Int16(f func(int32) int16, list []int32) []int16 {
	if f == nil {
		return []int16{}
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt32Int8(f func(int32) int8, list []int32) []int8 {
	if f == nil {
		return []int8{}
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt32Uint(f func(int32) uint, list []int32) []uint {
	if f == nil {
		return []uint{}
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt32Uint64(f func(int32) uint64, list []int32) []uint64 {
	if f == nil {
		return []uint64{}
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt32Uint32(f func(int32) uint32, list []int32) []uint32 {
	if f == nil {
		return []uint32{}
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt32Uint16(f func(int32) uint16, list []int32) []uint16 {
	if f == nil {
		return []uint16{}
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt32Uint8(f func(int32) uint8, list []int32) []uint8 {
	if f == nil {
		return []uint8{}
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt32Str(f func(int32) string, list []int32) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt32Bool(f func(int32) bool, list []int32) []bool {
	if f == nil {
		return []bool{}
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt16Int(f func(int16) int, list []int16) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt16Int64(f func(int16) int64, list []int16) []int64 {
	if f == nil {
		return []int64{}
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt16Int32(f func(int16) int32, list []int16) []int32 {
	if f == nil {
		return []int32{}
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt16Int8(f func(int16) int8, list []int16) []int8 {
	if f == nil {
		return []int8{}
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt16Uint(f func(int16) uint, list []int16) []uint {
	if f == nil {
		return []uint{}
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt16Uint64(f func(int16) uint64, list []int16) []uint64 {
	if f == nil {
		return []uint64{}
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt16Uint32(f func(int16) uint32, list []int16) []uint32 {
	if f == nil {
		return []uint32{}
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt16Uint16(f func(int16) uint16, list []int16) []uint16 {
	if f == nil {
		return []uint16{}
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt16Uint8(f func(int16) uint8, list []int16) []uint8 {
	if f == nil {
		return []uint8{}
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt16Str(f func(int16) string, list []int16) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt16Bool(f func(int16) bool, list []int16) []bool {
	if f == nil {
		return []bool{}
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt8Int(f func(int8) int, list []int8) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt8Int64(f func(int8) int64, list []int8) []int64 {
	if f == nil {
		return []int64{}
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt8Int32(f func(int8) int32, list []int8) []int32 {
	if f == nil {
		return []int32{}
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt8Int16(f func(int8) int16, list []int8) []int16 {
	if f == nil {
		return []int16{}
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt8Uint(f func(int8) uint, list []int8) []uint {
	if f == nil {
		return []uint{}
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt8Uint64(f func(int8) uint64, list []int8) []uint64 {
	if f == nil {
		return []uint64{}
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt8Uint32(f func(int8) uint32, list []int8) []uint32 {
	if f == nil {
		return []uint32{}
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt8Uint16(f func(int8) uint16, list []int8) []uint16 {
	if f == nil {
		return []uint16{}
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt8Uint8(f func(int8) uint8, list []int8) []uint8 {
	if f == nil {
		return []uint8{}
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt8Str(f func(int8) string, list []int8) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt8Bool(f func(int8) bool, list []int8) []bool {
	if f == nil {
		return []bool{}
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUintInt(f func(uint) int, list []uint) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUintInt64(f func(uint) int64, list []uint) []int64 {
	if f == nil {
		return []int64{}
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUintInt32(f func(uint) int32, list []uint) []int32 {
	if f == nil {
		return []int32{}
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUintInt16(f func(uint) int16, list []uint) []int16 {
	if f == nil {
		return []int16{}
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUintInt8(f func(uint) int8, list []uint) []int8 {
	if f == nil {
		return []int8{}
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUintUint64(f func(uint) uint64, list []uint) []uint64 {
	if f == nil {
		return []uint64{}
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUintUint32(f func(uint) uint32, list []uint) []uint32 {
	if f == nil {
		return []uint32{}
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUintUint16(f func(uint) uint16, list []uint) []uint16 {
	if f == nil {
		return []uint16{}
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUintUint8(f func(uint) uint8, list []uint) []uint8 {
	if f == nil {
		return []uint8{}
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUintStr(f func(uint) string, list []uint) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUintBool(f func(uint) bool, list []uint) []bool {
	if f == nil {
		return []bool{}
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint64Int(f func(uint64) int, list []uint64) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint64Int64(f func(uint64) int64, list []uint64) []int64 {
	if f == nil {
		return []int64{}
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint64Int32(f func(uint64) int32, list []uint64) []int32 {
	if f == nil {
		return []int32{}
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint64Int16(f func(uint64) int16, list []uint64) []int16 {
	if f == nil {
		return []int16{}
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint64Int8(f func(uint64) int8, list []uint64) []int8 {
	if f == nil {
		return []int8{}
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint64Uint(f func(uint64) uint, list []uint64) []uint {
	if f == nil {
		return []uint{}
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint64Uint32(f func(uint64) uint32, list []uint64) []uint32 {
	if f == nil {
		return []uint32{}
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint64Uint16(f func(uint64) uint16, list []uint64) []uint16 {
	if f == nil {
		return []uint16{}
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint64Uint8(f func(uint64) uint8, list []uint64) []uint8 {
	if f == nil {
		return []uint8{}
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint64Str(f func(uint64) string, list []uint64) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint64Bool(f func(uint64) bool, list []uint64) []bool {
	if f == nil {
		return []bool{}
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint32Int(f func(uint32) int, list []uint32) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint32Int64(f func(uint32) int64, list []uint32) []int64 {
	if f == nil {
		return []int64{}
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint32Int32(f func(uint32) int32, list []uint32) []int32 {
	if f == nil {
		return []int32{}
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint32Int16(f func(uint32) int16, list []uint32) []int16 {
	if f == nil {
		return []int16{}
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint32Int8(f func(uint32) int8, list []uint32) []int8 {
	if f == nil {
		return []int8{}
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint32Uint(f func(uint32) uint, list []uint32) []uint {
	if f == nil {
		return []uint{}
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint32Uint64(f func(uint32) uint64, list []uint32) []uint64 {
	if f == nil {
		return []uint64{}
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint32Uint16(f func(uint32) uint16, list []uint32) []uint16 {
	if f == nil {
		return []uint16{}
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint32Uint8(f func(uint32) uint8, list []uint32) []uint8 {
	if f == nil {
		return []uint8{}
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint32Str(f func(uint32) string, list []uint32) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint32Bool(f func(uint32) bool, list []uint32) []bool {
	if f == nil {
		return []bool{}
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint16Int(f func(uint16) int, list []uint16) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint16Int64(f func(uint16) int64, list []uint16) []int64 {
	if f == nil {
		return []int64{}
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint16Int32(f func(uint16) int32, list []uint16) []int32 {
	if f == nil {
		return []int32{}
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint16Int16(f func(uint16) int16, list []uint16) []int16 {
	if f == nil {
		return []int16{}
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint16Int8(f func(uint16) int8, list []uint16) []int8 {
	if f == nil {
		return []int8{}
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint16Uint(f func(uint16) uint, list []uint16) []uint {
	if f == nil {
		return []uint{}
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint16Uint64(f func(uint16) uint64, list []uint16) []uint64 {
	if f == nil {
		return []uint64{}
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint16Uint32(f func(uint16) uint32, list []uint16) []uint32 {
	if f == nil {
		return []uint32{}
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint16Uint8(f func(uint16) uint8, list []uint16) []uint8 {
	if f == nil {
		return []uint8{}
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint16Str(f func(uint16) string, list []uint16) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint16Bool(f func(uint16) bool, list []uint16) []bool {
	if f == nil {
		return []bool{}
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint8Int(f func(uint8) int, list []uint8) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint8Int64(f func(uint8) int64, list []uint8) []int64 {
	if f == nil {
		return []int64{}
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint8Int32(f func(uint8) int32, list []uint8) []int32 {
	if f == nil {
		return []int32{}
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint8Int16(f func(uint8) int16, list []uint8) []int16 {
	if f == nil {
		return []int16{}
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint8Int8(f func(uint8) int8, list []uint8) []int8 {
	if f == nil {
		return []int8{}
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint8Uint(f func(uint8) uint, list []uint8) []uint {
	if f == nil {
		return []uint{}
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint8Uint64(f func(uint8) uint64, list []uint8) []uint64 {
	if f == nil {
		return []uint64{}
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint8Uint32(f func(uint8) uint32, list []uint8) []uint32 {
	if f == nil {
		return []uint32{}
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint8Uint16(f func(uint8) uint16, list []uint8) []uint16 {
	if f == nil {
		return []uint16{}
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint8Str(f func(uint8) string, list []uint8) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint8Bool(f func(uint8) bool, list []uint8) []bool {
	if f == nil {
		return []bool{}
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapStrInt(f func(string) int, list []string) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapStrInt64(f func(string) int64, list []string) []int64 {
	if f == nil {
		return []int64{}
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapStrInt32(f func(string) int32, list []string) []int32 {
	if f == nil {
		return []int32{}
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapStrInt16(f func(string) int16, list []string) []int16 {
	if f == nil {
		return []int16{}
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapStrInt8(f func(string) int8, list []string) []int8 {
	if f == nil {
		return []int8{}
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapStrUint(f func(string) uint, list []string) []uint {
	if f == nil {
		return []uint{}
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapStrUint64(f func(string) uint64, list []string) []uint64 {
	if f == nil {
		return []uint64{}
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapStrUint32(f func(string) uint32, list []string) []uint32 {
	if f == nil {
		return []uint32{}
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapStrUint16(f func(string) uint16, list []string) []uint16 {
	if f == nil {
		return []uint16{}
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapStrUint8(f func(string) uint8, list []string) []uint8 {
	if f == nil {
		return []uint8{}
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapStrBool(f func(string) bool, list []string) []bool {
	if f == nil {
		return []bool{}
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapBoolInt(f func(bool) int, list []bool) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapBoolInt64(f func(bool) int64, list []bool) []int64 {
	if f == nil {
		return []int64{}
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapBoolInt32(f func(bool) int32, list []bool) []int32 {
	if f == nil {
		return []int32{}
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapBoolInt16(f func(bool) int16, list []bool) []int16 {
	if f == nil {
		return []int16{}
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapBoolInt8(f func(bool) int8, list []bool) []int8 {
	if f == nil {
		return []int8{}
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapBoolUint(f func(bool) uint, list []bool) []uint {
	if f == nil {
		return []uint{}
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapBoolUint64(f func(bool) uint64, list []bool) []uint64 {
	if f == nil {
		return []uint64{}
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapBoolUint32(f func(bool) uint32, list []bool) []uint32 {
	if f == nil {
		return []uint32{}
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapBoolUint16(f func(bool) uint16, list []bool) []uint16 {
	if f == nil {
		return []uint16{}
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapBoolUint8(f func(bool) uint8, list []bool) []uint8 {
	if f == nil {
		return []uint8{}
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapBoolStr(f func(bool) string, list []bool) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
