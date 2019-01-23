package fp

import "strings"

func DropInt(num int, list []int) []int {
	var newList []int
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropInts(nums []int, list []int) []int {
	var newList []int
	for _, v := range list {
		if !ExistsInt(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropInt64(num int64, list []int64) []int64 {
	var newList []int64
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropInts64(nums []int64, list []int64) []int64 {
	var newList []int64
	for _, v := range list {
		if !ExistsInt64(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropInt32(num int32, list []int32) []int32 {
	var newList []int32
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropInts32(nums []int32, list []int32) []int32 {
	var newList []int32
	for _, v := range list {
		if !ExistsInt32(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropInt16(num int16, list []int16) []int16 {
	var newList []int16
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropInts16(nums []int16, list []int16) []int16 {
	var newList []int16
	for _, v := range list {
		if !ExistsInt16(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropInt8(num int8, list []int8) []int8 {
	var newList []int8
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropInts8(nums []int8, list []int8) []int8 {
	var newList []int8
	for _, v := range list {
		if !ExistsInt8(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropUint(num uint, list []uint) []uint {
	var newList []uint
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropUints(nums []uint, list []uint) []uint {
	var newList []uint
	for _, v := range list {
		if !ExistsUint(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropUint64(num uint64, list []uint64) []uint64 {
	var newList []uint64
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropUints64(nums []uint64, list []uint64) []uint64 {
	var newList []uint64
	for _, v := range list {
		if !ExistsUint64(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropUint32(num uint32, list []uint32) []uint32 {
	var newList []uint32
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropUints32(nums []uint32, list []uint32) []uint32 {
	var newList []uint32
	for _, v := range list {
		if !ExistsUint32(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropUint16(num uint16, list []uint16) []uint16 {
	var newList []uint16
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropUints16(nums []uint16, list []uint16) []uint16 {
	var newList []uint16
	for _, v := range list {
		if !ExistsUint16(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropUint8(num uint8, list []uint8) []uint8 {
	var newList []uint8
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropUints8(nums []uint8, list []uint8) []uint8 {
	var newList []uint8
	for _, v := range list {
		if !ExistsUint8(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropFloat64(num float64, list []float64) []float64 {
	var newList []float64
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropFloats64(nums []float64, list []float64) []float64 {
	var newList []float64
	for _, v := range list {
		if !ExistsFloat64(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropFloat32(num float32, list []float32) []float32 {
	var newList []float32
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropFloats32(nums []float32, list []float32) []float32 {
	var newList []float32
	for _, v := range list {
		if !ExistsFloat32(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropStr(str string, list []string) []string {
	var newList []string
	for _, v := range list {
		if v != str {
			newList = append(newList, v)
		}
	}
	return newList
}

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

func DropStrs(strs []string, list []string) []string {
	var newList []string
	for _, v := range list {
		if !ExistsStr(v, strs) {
			newList = append(newList, v)
		}
	}
	return newList
}

func DropStrsIgnoreCase(strs []string, list []string) []string {
	var newList []string

	for _, v := range list {
		if !ExistsStrIgnoreCase(strings.ToLower(v), strs) {
			newList = append(newList, v)
		}
	}
	return newList
}
