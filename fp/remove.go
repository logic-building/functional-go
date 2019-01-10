package fp

import "strings"

func RemoveInt(num int, list []int) []int {
	var newList []int
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveInts(nums []int, list []int) []int {
	var newList []int
	for _, v := range list {
		if !SomeInt(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveInt64(num int64, list []int64) []int64 {
	var newList []int64
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveInts64(nums []int64, list []int64) []int64 {
	var newList []int64
	for _, v := range list {
		if !SomeInt64(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveInt32(num int32, list []int32) []int32 {
	var newList []int32
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveInts32(nums []int32, list []int32) []int32 {
	var newList []int32
	for _, v := range list {
		if !SomeInt32(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveInt16(num int16, list []int16) []int16 {
	var newList []int16
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveInts16(nums []int16, list []int16) []int16 {
	var newList []int16
	for _, v := range list {
		if !SomeInt16(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveInt8(num int8, list []int8) []int8 {
	var newList []int8
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveInts8(nums []int8, list []int8) []int8 {
	var newList []int8
	for _, v := range list {
		if !SomeInt8(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveUint(num uint, list []uint) []uint {
	var newList []uint
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveUints(nums []uint, list []uint) []uint {
	var newList []uint
	for _, v := range list {
		if !SomeUint(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveUint64(num uint64, list []uint64) []uint64 {
	var newList []uint64
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveUints64(nums []uint64, list []uint64) []uint64 {
	var newList []uint64
	for _, v := range list {
		if !SomeUint64(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveUint32(num uint32, list []uint32) []uint32 {
	var newList []uint32
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveUints32(nums []uint32, list []uint32) []uint32 {
	var newList []uint32
	for _, v := range list {
		if !SomeUint32(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveUint16(num uint16, list []uint16) []uint16 {
	var newList []uint16
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveUints16(nums []uint16, list []uint16) []uint16 {
	var newList []uint16
	for _, v := range list {
		if !SomeUint16(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveUint8(num uint8, list []uint8) []uint8 {
	var newList []uint8
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveUints8(nums []uint8, list []uint8) []uint8 {
	var newList []uint8
	for _, v := range list {
		if !SomeUint8(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveFloat64(num float64, list []float64) []float64 {
	var newList []float64
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveFloats64(nums []float64, list []float64) []float64 {
	var newList []float64
	for _, v := range list {
		if !SomeFloat64(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveFloat32(num float32, list []float32) []float32 {
	var newList []float32
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveFloats32(nums []float32, list []float32) []float32 {
	var newList []float32
	for _, v := range list {
		if !SomeFloat32(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveStr(str string, list []string) []string {
	var newList []string
	for _, v := range list {
		if v != str {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveStrIgnoreCase(str string, list []string) []string {
	strLower := strings.ToLower(str)
	var newList []string
	for _, v := range list {
		if strLower != strings.ToLower(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveStrs(strs []string, list []string) []string {
	var newList []string
	for _, v := range list {
		if !SomeStr(v, strs) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveStrsIgnoreCase(strs []string, list []string) []string {
	var newList []string

	for _, v := range list {
		if !SomeStrIgnoreCase(strings.ToLower(v), strs) {
			newList = append(newList, v)
		}
	}
	return newList
}
