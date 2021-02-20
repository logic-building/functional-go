package fp

import "sort"

type int641 []int64

func (a int641) Len() int           { return len(a) }
func (a int641) Less(i, j int) bool { return a[i] < a[j] }
func (a int641) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortInts64 return new sorted list
func SortInts64(list []int64) []int64 {
	if len(list) == 0 {
		return []int64{}
	}
	var newList int641
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortInts64Ptr return new sorted list
func SortInts64Ptr(list []*int64) []*int64 {
	if len(list) == 0 {
		return []*int64{}
	}
	var newList int641
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*int64, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type int642 []int64

func (a int642) Len() int           { return len(a) }
func (a int642) Less(i, j int) bool { return a[i] > a[j] }
func (a int642) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortInts64Desc return new sorted list
func SortInts64Desc(list []int64) []int64 {
	if len(list) == 0 {
		return []int64{}
	}
	var newList int642
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortInts64DescPtr return new sorted list
func SortInts64DescPtr(list []*int64) []*int64 {
	if len(list) == 0 {
		return []*int64{}
	}
	var newList int642
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*int64, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type int321 []int32

func (a int321) Len() int           { return len(a) }
func (a int321) Less(i, j int) bool { return a[i] < a[j] }
func (a int321) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortInts32 return new sorted list
func SortInts32(list []int32) []int32 {
	if len(list) == 0 {
		return []int32{}
	}
	var newList int321
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortInts32Ptr return new sorted list
func SortInts32Ptr(list []*int32) []*int32 {
	if len(list) == 0 {
		return []*int32{}
	}
	var newList int321
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*int32, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type int322 []int32

func (a int322) Len() int           { return len(a) }
func (a int322) Less(i, j int) bool { return a[i] > a[j] }
func (a int322) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortInts32Desc return new sorted list
func SortInts32Desc(list []int32) []int32 {
	if len(list) == 0 {
		return []int32{}
	}
	var newList int322
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortInts32DescPtr return new sorted list
func SortInts32DescPtr(list []*int32) []*int32 {
	if len(list) == 0 {
		return []*int32{}
	}
	var newList int322
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*int32, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type int161 []int16

func (a int161) Len() int           { return len(a) }
func (a int161) Less(i, j int) bool { return a[i] < a[j] }
func (a int161) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortInts16 return new sorted list
func SortInts16(list []int16) []int16 {
	if len(list) == 0 {
		return []int16{}
	}
	var newList int161
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortInts16Ptr return new sorted list
func SortInts16Ptr(list []*int16) []*int16 {
	if len(list) == 0 {
		return []*int16{}
	}
	var newList int161
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*int16, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type int162 []int16

func (a int162) Len() int           { return len(a) }
func (a int162) Less(i, j int) bool { return a[i] > a[j] }
func (a int162) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortInts16Desc return new sorted list
func SortInts16Desc(list []int16) []int16 {
	if len(list) == 0 {
		return []int16{}
	}
	var newList int162
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortInts16DescPtr return new sorted list
func SortInts16DescPtr(list []*int16) []*int16 {
	if len(list) == 0 {
		return []*int16{}
	}
	var newList int162
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*int16, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type int81 []int8

func (a int81) Len() int           { return len(a) }
func (a int81) Less(i, j int) bool { return a[i] < a[j] }
func (a int81) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortInts8 return new sorted list
func SortInts8(list []int8) []int8 {
	if len(list) == 0 {
		return []int8{}
	}
	var newList int81
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortInts8Ptr return new sorted list
func SortInts8Ptr(list []*int8) []*int8 {
	if len(list) == 0 {
		return []*int8{}
	}
	var newList int81
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*int8, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type int82 []int8

func (a int82) Len() int           { return len(a) }
func (a int82) Less(i, j int) bool { return a[i] > a[j] }
func (a int82) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortInts8Desc return new sorted list
func SortInts8Desc(list []int8) []int8 {
	if len(list) == 0 {
		return []int8{}
	}
	var newList int82
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortInts8DescPtr return new sorted list
func SortInts8DescPtr(list []*int8) []*int8 {
	if len(list) == 0 {
		return []*int8{}
	}
	var newList int82
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*int8, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type uint1 []uint

func (a uint1) Len() int           { return len(a) }
func (a uint1) Less(i, j int) bool { return a[i] < a[j] }
func (a uint1) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortUints return new sorted list
func SortUints(list []uint) []uint {
	if len(list) == 0 {
		return []uint{}
	}
	var newList uint1
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortUintsPtr return new sorted list
func SortUintsPtr(list []*uint) []*uint {
	if len(list) == 0 {
		return []*uint{}
	}
	var newList uint1
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*uint, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type uint2 []uint

func (a uint2) Len() int           { return len(a) }
func (a uint2) Less(i, j int) bool { return a[i] > a[j] }
func (a uint2) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortUintsDesc return new sorted list
func SortUintsDesc(list []uint) []uint {
	if len(list) == 0 {
		return []uint{}
	}
	var newList uint2
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortUintsDescPtr return new sorted list
func SortUintsDescPtr(list []*uint) []*uint {
	if len(list) == 0 {
		return []*uint{}
	}
	var newList uint2
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*uint, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type uint641 []uint64

func (a uint641) Len() int           { return len(a) }
func (a uint641) Less(i, j int) bool { return a[i] < a[j] }
func (a uint641) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortUint64s return new sorted list
func SortUint64s(list []uint64) []uint64 {
	if len(list) == 0 {
		return []uint64{}
	}
	var newList uint641
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortUint64sPtr return new sorted list
func SortUint64sPtr(list []*uint64) []*uint64 {
	if len(list) == 0 {
		return []*uint64{}
	}
	var newList uint641
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*uint64, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type uint642 []uint64

func (a uint642) Len() int           { return len(a) }
func (a uint642) Less(i, j int) bool { return a[i] > a[j] }
func (a uint642) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortUint64sDesc return new sorted list
func SortUint64sDesc(list []uint64) []uint64 {
	if len(list) == 0 {
		return []uint64{}
	}
	var newList uint642
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortUint64sDescPtr return new sorted list
func SortUint64sDescPtr(list []*uint64) []*uint64 {
	if len(list) == 0 {
		return []*uint64{}
	}
	var newList uint642
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*uint64, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type uint321 []uint32

func (a uint321) Len() int           { return len(a) }
func (a uint321) Less(i, j int) bool { return a[i] < a[j] }
func (a uint321) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortUints32 return new sorted list
func SortUints32(list []uint32) []uint32 {
	if len(list) == 0 {
		return []uint32{}
	}
	var newList uint321
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortUints32Ptr return new sorted list
func SortUints32Ptr(list []*uint32) []*uint32 {
	if len(list) == 0 {
		return []*uint32{}
	}
	var newList uint321
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*uint32, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type uint322 []uint32

func (a uint322) Len() int           { return len(a) }
func (a uint322) Less(i, j int) bool { return a[i] > a[j] }
func (a uint322) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortUints32Desc return new sorted list
func SortUints32Desc(list []uint32) []uint32 {
	if len(list) == 0 {
		return []uint32{}
	}
	var newList uint322
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortUints32DescPtr return new sorted list
func SortUints32DescPtr(list []*uint32) []*uint32 {
	if len(list) == 0 {
		return []*uint32{}
	}
	var newList uint322
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*uint32, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type uint161 []uint16

func (a uint161) Len() int           { return len(a) }
func (a uint161) Less(i, j int) bool { return a[i] < a[j] }
func (a uint161) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortUints16 return new sorted list
func SortUints16(list []uint16) []uint16 {
	if len(list) == 0 {
		return []uint16{}
	}
	var newList uint161
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortUints16Ptr return new sorted list
func SortUints16Ptr(list []*uint16) []*uint16 {
	if len(list) == 0 {
		return []*uint16{}
	}
	var newList uint161
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*uint16, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type uint162 []uint16

func (a uint162) Len() int           { return len(a) }
func (a uint162) Less(i, j int) bool { return a[i] > a[j] }
func (a uint162) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortUints16Desc return new sorted list
func SortUints16Desc(list []uint16) []uint16 {
	if len(list) == 0 {
		return []uint16{}
	}
	var newList uint162
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortUints16DescPtr return new sorted list
func SortUints16DescPtr(list []*uint16) []*uint16 {
	if len(list) == 0 {
		return []*uint16{}
	}
	var newList uint162
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*uint16, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type uint81 []uint8

func (a uint81) Len() int           { return len(a) }
func (a uint81) Less(i, j int) bool { return a[i] < a[j] }
func (a uint81) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortUints8 return new sorted list
func SortUints8(list []uint8) []uint8 {
	if len(list) == 0 {
		return []uint8{}
	}
	var newList uint81
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortUints8Ptr return new sorted list
func SortUints8Ptr(list []*uint8) []*uint8 {
	if len(list) == 0 {
		return []*uint8{}
	}
	var newList uint81
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*uint8, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type uint82 []uint8

func (a uint82) Len() int           { return len(a) }
func (a uint82) Less(i, j int) bool { return a[i] > a[j] }
func (a uint82) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortUints8Desc return new sorted list
func SortUints8Desc(list []uint8) []uint8 {
	if len(list) == 0 {
		return []uint8{}
	}
	var newList uint82
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortUints8DescPtr return new sorted list
func SortUints8DescPtr(list []*uint8) []*uint8 {
	if len(list) == 0 {
		return []*uint8{}
	}
	var newList uint82
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*uint8, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type float321 []float32

func (a float321) Len() int           { return len(a) }
func (a float321) Less(i, j int) bool { return a[i] < a[j] }
func (a float321) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortFloat32s return new sorted list
func SortFloat32s(list []float32) []float32 {
	if len(list) == 0 {
		return []float32{}
	}
	var newList float321
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortFloat32sPtr return new sorted list
func SortFloat32sPtr(list []*float32) []*float32 {
	if len(list) == 0 {
		return []*float32{}
	}
	var newList float321
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*float32, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}

type float322 []float32

func (a float322) Len() int           { return len(a) }
func (a float322) Less(i, j int) bool { return a[i] > a[j] }
func (a float322) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SortFloat32sDesc return new sorted list
func SortFloat32sDesc(list []float32) []float32 {
	if len(list) == 0 {
		return []float32{}
	}
	var newList float322
	for _, item := range list {
		newList = append(newList, item)
	}
	sort.Sort(newList)
	return newList
}

// SortFloat32sDescPtr return new sorted list
func SortFloat32sDescPtr(list []*float32) []*float32 {
	if len(list) == 0 {
		return []*float32{}
	}
	var newList float322
	for _, item := range list {
		newList = append(newList, *item)
	}
	sort.Sort(newList)

	newListPtr := make([]*float32, len(list))
	for i := 0; i < len(newList); i++ {
		newListPtr[i] = &newList[i]
	}

	return newListPtr
}
