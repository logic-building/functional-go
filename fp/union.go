package fp

// UnionInt return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionInt(arrList ...[]int) []int {
	resultMap := make(map[int]bool)
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	resultArr := make([]int, len(resultMap))
	i := 0
	for k := range resultMap {
		resultArr[i] = k
		i++
	}
	return resultArr
}

// UnionIntPtr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionIntPtr(arrList ...[]*int) []*int {
	resultMap := make(map[int]bool)
	var resultArr []*int
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[*v]
			if !ok {
				resultMap[*v] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}

// UnionInt64 return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionInt64(arrList ...[]int64) []int64 {
	resultMap := make(map[int64]bool)
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	resultArr := make([]int64, len(resultMap))
	i := 0
	for k := range resultMap {
		resultArr[i] = k
		i++
	}
	return resultArr
}

// UnionInt64Ptr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionInt64Ptr(arrList ...[]*int64) []*int64 {
	resultMap := make(map[int64]bool)
	var resultArr []*int64
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[*v]
			if !ok {
				resultMap[*v] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}

// UnionInt32 return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionInt32(arrList ...[]int32) []int32 {
	resultMap := make(map[int32]bool)
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	resultArr := make([]int32, len(resultMap))
	i := 0
	for k := range resultMap {
		resultArr[i] = k
		i++
	}
	return resultArr
}

// UnionInt32Ptr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionInt32Ptr(arrList ...[]*int32) []*int32 {
	resultMap := make(map[int32]bool)
	var resultArr []*int32
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[*v]
			if !ok {
				resultMap[*v] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}

// UnionInt16 return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionInt16(arrList ...[]int16) []int16 {
	resultMap := make(map[int16]bool)
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	resultArr := make([]int16, len(resultMap))
	i := 0
	for k := range resultMap {
		resultArr[i] = k
		i++
	}
	return resultArr
}

// UnionInt16Ptr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionInt16Ptr(arrList ...[]*int16) []*int16 {
	resultMap := make(map[int16]bool)
	var resultArr []*int16
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[*v]
			if !ok {
				resultMap[*v] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}

// UnionInt8 return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionInt8(arrList ...[]int8) []int8 {
	resultMap := make(map[int8]bool)
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	resultArr := make([]int8, len(resultMap))
	i := 0
	for k := range resultMap {
		resultArr[i] = k
		i++
	}
	return resultArr
}

// UnionInt8Ptr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionInt8Ptr(arrList ...[]*int8) []*int8 {
	resultMap := make(map[int8]bool)
	var resultArr []*int8
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[*v]
			if !ok {
				resultMap[*v] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}

// UnionUint return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionUint(arrList ...[]uint) []uint {
	resultMap := make(map[uint]bool)
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	resultArr := make([]uint, len(resultMap))
	i := 0
	for k := range resultMap {
		resultArr[i] = k
		i++
	}
	return resultArr
}

// UnionUintPtr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionUintPtr(arrList ...[]*uint) []*uint {
	resultMap := make(map[uint]bool)
	var resultArr []*uint
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[*v]
			if !ok {
				resultMap[*v] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}

// UnionUint64 return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionUint64(arrList ...[]uint64) []uint64 {
	resultMap := make(map[uint64]bool)
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	resultArr := make([]uint64, len(resultMap))
	i := 0
	for k := range resultMap {
		resultArr[i] = k
		i++
	}
	return resultArr
}

// UnionUint64Ptr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionUint64Ptr(arrList ...[]*uint64) []*uint64 {
	resultMap := make(map[uint64]bool)
	var resultArr []*uint64
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[*v]
			if !ok {
				resultMap[*v] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}

// UnionUint32 return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionUint32(arrList ...[]uint32) []uint32 {
	resultMap := make(map[uint32]bool)
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	resultArr := make([]uint32, len(resultMap))
	i := 0
	for k := range resultMap {
		resultArr[i] = k
		i++
	}
	return resultArr
}

// UnionUint32Ptr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionUint32Ptr(arrList ...[]*uint32) []*uint32 {
	resultMap := make(map[uint32]bool)
	var resultArr []*uint32
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[*v]
			if !ok {
				resultMap[*v] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}

// UnionUint16 return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionUint16(arrList ...[]uint16) []uint16 {
	resultMap := make(map[uint16]bool)
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	resultArr := make([]uint16, len(resultMap))
	i := 0
	for k := range resultMap {
		resultArr[i] = k
		i++
	}
	return resultArr
}

// UnionUint16Ptr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionUint16Ptr(arrList ...[]*uint16) []*uint16 {
	resultMap := make(map[uint16]bool)
	var resultArr []*uint16
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[*v]
			if !ok {
				resultMap[*v] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}

// UnionUint8 return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionUint8(arrList ...[]uint8) []uint8 {
	resultMap := make(map[uint8]bool)
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	resultArr := make([]uint8, len(resultMap))
	i := 0
	for k := range resultMap {
		resultArr[i] = k
		i++
	}
	return resultArr
}

// UnionUint8Ptr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionUint8Ptr(arrList ...[]*uint8) []*uint8 {
	resultMap := make(map[uint8]bool)
	var resultArr []*uint8
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[*v]
			if !ok {
				resultMap[*v] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}

// UnionStr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionStr(arrList ...[]string) []string {
	resultMap := make(map[string]bool)
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	resultArr := make([]string, len(resultMap))
	i := 0
	for k := range resultMap {
		resultArr[i] = k
		i++
	}
	return resultArr
}

// UnionStrPtr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionStrPtr(arrList ...[]*string) []*string {
	resultMap := make(map[string]bool)
	var resultArr []*string
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[*v]
			if !ok {
				resultMap[*v] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}

// UnionBool return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionBool(arrList ...[]bool) []bool {
	resultMap := make(map[bool]bool)
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	resultArr := make([]bool, len(resultMap))
	i := 0
	for k := range resultMap {
		resultArr[i] = k
		i++
	}
	return resultArr
}

// UnionBoolPtr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionBoolPtr(arrList ...[]*bool) []*bool {
	resultMap := make(map[bool]bool)
	var resultArr []*bool
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[*v]
			if !ok {
				resultMap[*v] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}

// UnionFloat32 return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionFloat32(arrList ...[]float32) []float32 {
	resultMap := make(map[float32]bool)
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	resultArr := make([]float32, len(resultMap))
	i := 0
	for k := range resultMap {
		resultArr[i] = k
		i++
	}
	return resultArr
}

// UnionFloat32Ptr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionFloat32Ptr(arrList ...[]*float32) []*float32 {
	resultMap := make(map[float32]bool)
	var resultArr []*float32
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[*v]
			if !ok {
				resultMap[*v] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}

// UnionFloat64 return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionFloat64(arrList ...[]float64) []float64 {
	resultMap := make(map[float64]bool)
	for _, arr := range arrList {
		for _, v := range arr {
			resultMap[v] = true
		}
	}

	resultArr := make([]float64, len(resultMap))
	i := 0
	for k := range resultMap {
		resultArr[i] = k
		i++
	}
	return resultArr
}

// UnionFloat64Ptr return a set that is the union of the input sets
// repeated value within list parameter will be ignored
func UnionFloat64Ptr(arrList ...[]*float64) []*float64 {
	resultMap := make(map[float64]bool)
	var resultArr []*float64
	for _, arr := range arrList {
		for _, v := range arr {
			_, ok := resultMap[*v]
			if !ok {
				resultMap[*v] = true
				resultArr = append(resultArr, v)
			}
		}
	}
	return resultArr
}
