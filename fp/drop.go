package fp

import "strings"

// DropInt returns a new list after dropping the given item
//
// Example:
// 	fp.DropInt(1, []int{1, 2, 3, 1}) // returns [2, 3]
func DropInt(num int, list []int) []int {
	var newList []int
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts returns a new list after dropping the given items
//
// Example:
// 	fp.DropInts([]int{1, 4}, []int{1, 2, 3, 1, 4}) // returns [2, 3]
func DropInts(nums []int, list []int) []int {
	var newList []int
	for _, v := range list {
		if !ExistsInt(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInt64 returns a new list after dropping the given item
//
// Example:
// 	fp.DropInt64(1, []int64{1, 2, 3, 1}) // returns [2, 3]
func DropInt64(num int64, list []int64) []int64 {
	var newList []int64
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts64 returns a new list after dropping the given items
//
// Example:
// 	fp.DropInts64([]int64{1, 4}, []int64{1, 2, 3, 1, 4}) // returns [2, 3]
func DropInts64(nums []int64, list []int64) []int64 {
	var newList []int64
	for _, v := range list {
		if !ExistsInt64(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInt32 returns a new list after dropping the given item
//
// Example:
// 	fp.DropInt32(1, []int32{1, 2, 3, 1}) // returns [2, 3]
func DropInt32(num int32, list []int32) []int32 {
	var newList []int32
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts32 returns a new list after dropping the given items
//
// Example:
// 	fp.DropInts32([]int32{1, 4}, []int32{1, 2, 3, 1, 4}) // returns [2, 3]
func DropInts32(nums []int32, list []int32) []int32 {
	var newList []int32
	for _, v := range list {
		if !ExistsInt32(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInt16 returns a new list after dropping the given item
//
// Example:
// 	fp.DropInt16(1, []int16{1, 2, 3, 1}) // returns [2, 3]
func DropInt16(num int16, list []int16) []int16 {
	var newList []int16
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts16 returns a new list after dropping the given items
//
// Example:
// 	fp.DropInts16([]int16{1, 4}, []int16{1, 2, 3, 1, 4}) // returns [2, 3]
func DropInts16(nums []int16, list []int16) []int16 {
	var newList []int16
	for _, v := range list {
		if !ExistsInt16(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInt8 returns a new list after dropping the given item
//
// Example:
// 	fp.DropInt8(1, []int8{1, 2, 3, 1}) // returns [2, 3]
func DropInt8(num int8, list []int8) []int8 {
	var newList []int8
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts8 returns a new list after dropping the given items
//
// Example:
// 	fp.DropInts8([]int8{1, 4}, []int8{1, 2, 3, 1, 4}) // returns [2, 3]
func DropInts8(nums []int8, list []int8) []int8 {
	var newList []int8
	for _, v := range list {
		if !ExistsInt8(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropUint returns a new list after dropping the given item
//
// Example:
// 	fp.DropUint(1, []uint{1, 2, 3, 1}) // returns [2, 3]
func DropUint(num uint, list []uint) []uint {
	var newList []uint
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropUints returns a new list after dropping the given items
//
// Example:
// 	fp.DropUints([]uint{1, 4}, []uint{1, 2, 3, 1, 4}) // returns [2, 3]
func DropUints(nums []uint, list []uint) []uint {
	var newList []uint
	for _, v := range list {
		if !ExistsUint(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropUint64 returns a new list after dropping the given item
//
// Example:
// 	fp.DropUint64(1, []uint64{1, 2, 3, 1}) // returns [2, 3]
func DropUint64(num uint64, list []uint64) []uint64 {
	var newList []uint64
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropUints64 returns a new list after dropping the given items
//
// Example:
// 	fp.DropUints64([]uint64{1, 4}, []uint64{1, 2, 3, 1, 4}) // returns [2, 3]
func DropUints64(nums []uint64, list []uint64) []uint64 {
	var newList []uint64
	for _, v := range list {
		if !ExistsUint64(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropUint32 returns a new list after dropping the given item
//
// Example:
// 	fp.DropUint32(1, []uint32{1, 2, 3, 1}) // returns [2, 3]
func DropUint32(num uint32, list []uint32) []uint32 {
	var newList []uint32
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropUints32 returns a new list after dropping the given items
//
// Example:
// 	fp.DropUints32([]uint32{1, 4}, []uint32{1, 2, 3, 1, 4}) // returns [2, 3]
func DropUints32(nums []uint32, list []uint32) []uint32 {
	var newList []uint32
	for _, v := range list {
		if !ExistsUint32(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropUint16 returns a new list after dropping the given item
//
// Example:
// 	fp.DropUint16(1, []uint16{1, 2, 3, 1}) // returns [2, 3]
func DropUint16(num uint16, list []uint16) []uint16 {
	var newList []uint16
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropUints16 returns a new list after dropping the given items
//
// Example:
// 	fp.DropUints16([]uint16{1, 4}, []uint16{1, 2, 3, 1, 4}) // returns [2, 3]
func DropUints16(nums []uint16, list []uint16) []uint16 {
	var newList []uint16
	for _, v := range list {
		if !ExistsUint16(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropUint8 returns a new list after dropping the given item
//
// Example:
// 	fp.DropUint8(1, []uint8{1, 2, 3, 1}) // returns [2, 3]
func DropUint8(num uint8, list []uint8) []uint8 {
	var newList []uint8
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropUints8 returns a new list after dropping the given items
//
// Example:
// 	fp.DropUints8([]uint8{1, 4}, []uint8{1, 2, 3, 1, 4}) // returns [2, 3]
func DropUints8(nums []uint8, list []uint8) []uint8 {
	var newList []uint8
	for _, v := range list {
		if !ExistsUint8(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropFloat64 returns a new list after dropping the given item
//
// Example:
// 	fp.DropFloat64(1.1, []float64{1.1, 2.1, 3.1, 1.1}) // returns [2.1, 3.1]
func DropFloat64(num float64, list []float64) []float64 {
	var newList []float64
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropFloats64 returns a new list after dropping the given items
//
// Example:
// 	fp.DropFloats64([]float64{1.1, 4.1}, []float64{1.1, 2.1, 3.1, 1.1, 4.1}) // returns [2.1, 3.1]
func DropFloats64(nums []float64, list []float64) []float64 {
	var newList []float64
	for _, v := range list {
		if !ExistsFloat64(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropFloat32 returns a new list after dropping the given item
//
// Example:
// 	fp.DropFloat32(1.1, []float32{1.1, 2.1, 3.1, 1.1}) // returns [2.1, 3.1]
func DropFloat32(num float32, list []float32) []float32 {
	var newList []float32
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropFloats32 returns a new list after dropping the given items
//
// Example:
// 	fp.DropFloats32([]float32{1.1, 4.1}, []float32{1.1, 2.1, 3.1, 1.1, 4.1}) // returns [2.1, 3.1]
func DropFloats32(nums []float32, list []float32) []float32 {
	var newList []float32
	for _, v := range list {
		if !ExistsFloat32(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropStr returns a new list after dropping the given item
//
// Example:
// 	fp.DropStr("Ravan", []string{"Ravan", "Ram", "Shyam", "Ravan"} // returns ["Ram", "Shyam"]
func DropStr(str string, list []string) []string {
	var newList []string
	for _, v := range list {
		if v != str {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropStrIgnoreCase returns a new list after dropping the given item. It ignores case.
//
// Example:
// 	fp.DropStrIgnoreCase("ravan", []string{"Ravan", "Ram", "Shyam", "Ravan"} // returns ["Ram", "Shyam"]
func DropStrIgnoreCase(str string, list []string) []string {
	strLower := strings.ToLower(str)
	var newList []string
	for _, v := range list {
		if strLower != strings.ToLower(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropStrs returns a new list after dropping the given items
//
// Example:
// 	fp.DropStrs([]string{"nks", "bharat"}, []string{"nks", "Ram", "Shyam", "Nks", "bharat"}) // returns ["Ram", "Shyam"]
func DropStrs(strs []string, list []string) []string {
	var newList []string
	for _, v := range list {
		if !ExistsStr(v, strs) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropStrsIgnoreCase returns a new list after dropping the given items. It ignores the case.
//
// Example:
// 	fp.DropStrs([]string{"nks", "bharat"}, []string{"Nks", "Ram", "Shyam", "Nks", "Bharat"}) // returns ["Ram", "Shyam"]
func DropStrsIgnoreCase(strs []string, list []string) []string {
	var newList []string

	for _, v := range list {
		if !ExistsStrIgnoreCase(strings.ToLower(v), strs) {
			newList = append(newList, v)
		}
	}
	return newList
}
