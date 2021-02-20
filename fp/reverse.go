package fp

// ReverseInts reverse the list
func ReverseInts(list []int) []int {
	newList := make([]int, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseIntsPtr reverse the list
func ReverseIntsPtr(list []*int) []*int {
	newList := make([]*int, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseInts64 reverse the list
func ReverseInts64(list []int64) []int64 {
	newList := make([]int64, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseInts64Ptr reverse the list
func ReverseInts64Ptr(list []*int64) []*int64 {
	newList := make([]*int64, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseInts32 reverse the list
func ReverseInts32(list []int32) []int32 {
	newList := make([]int32, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseInts32Ptr reverse the list
func ReverseInts32Ptr(list []*int32) []*int32 {
	newList := make([]*int32, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseInts16 reverse the list
func ReverseInts16(list []int16) []int16 {
	newList := make([]int16, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseInts16Ptr reverse the list
func ReverseInts16Ptr(list []*int16) []*int16 {
	newList := make([]*int16, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseInts8 reverse the list
func ReverseInts8(list []int8) []int8 {
	newList := make([]int8, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseInts8Ptr reverse the list
func ReverseInts8Ptr(list []*int8) []*int8 {
	newList := make([]*int8, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseUints reverse the list
func ReverseUints(list []uint) []uint {
	newList := make([]uint, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseUintsPtr reverse the list
func ReverseUintsPtr(list []*uint) []*uint {
	newList := make([]*uint, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseUint64s reverse the list
func ReverseUint64s(list []uint64) []uint64 {
	newList := make([]uint64, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseUint64sPtr reverse the list
func ReverseUint64sPtr(list []*uint64) []*uint64 {
	newList := make([]*uint64, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseUints32 reverse the list
func ReverseUints32(list []uint32) []uint32 {
	newList := make([]uint32, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseUints32Ptr reverse the list
func ReverseUints32Ptr(list []*uint32) []*uint32 {
	newList := make([]*uint32, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseUints16 reverse the list
func ReverseUints16(list []uint16) []uint16 {
	newList := make([]uint16, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseUints16Ptr reverse the list
func ReverseUints16Ptr(list []*uint16) []*uint16 {
	newList := make([]*uint16, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseUints8 reverse the list
func ReverseUints8(list []uint8) []uint8 {
	newList := make([]uint8, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseUints8Ptr reverse the list
func ReverseUints8Ptr(list []*uint8) []*uint8 {
	newList := make([]*uint8, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseStrs reverse the list
func ReverseStrs(list []string) []string {
	newList := make([]string, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseStrsPtr reverse the list
func ReverseStrsPtr(list []*string) []*string {
	newList := make([]*string, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseBools reverse the list
func ReverseBools(list []bool) []bool {
	newList := make([]bool, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseBoolsPtr reverse the list
func ReverseBoolsPtr(list []*bool) []*bool {
	newList := make([]*bool, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseFloat32s reverse the list
func ReverseFloat32s(list []float32) []float32 {
	newList := make([]float32, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseFloat32sPtr reverse the list
func ReverseFloat32sPtr(list []*float32) []*float32 {
	newList := make([]*float32, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseFloat64s reverse the list
func ReverseFloat64s(list []float64) []float64 {
	newList := make([]float64, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}

// ReverseFloat64sPtr reverse the list
func ReverseFloat64sPtr(list []*float64) []*float64 {
	newList := make([]*float64, len(list))
	for i := 0; i < len(list); i++ {
		newList[i] = list[len(list)-(i+1)]
	}
	return newList
}
