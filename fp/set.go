package fp

// SetInt returns a set of the distinct elements of coll.
func SetInt(list []int) []int {
	if list == nil || len(list) == 0 {
		return []int{}
	}

	resultMap := make(map[int]struct{})
	newList := []int{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i]]
		if !ok {
			resultMap[list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetIntPtr returns a set of the distinct elements of coll.
func SetIntPtr(list []*int) []*int {
	if list == nil || len(list) == 0 {
		return []*int{}
	}

	resultMap := make(map[int]struct{})
	newList := []*int{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[*list[i]]
		if !ok {
			resultMap[*list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetInt64 returns a set of the distinct elements of coll.
func SetInt64(list []int64) []int64 {
	if list == nil || len(list) == 0 {
		return []int64{}
	}

	resultMap := make(map[int64]struct{})
	newList := []int64{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i]]
		if !ok {
			resultMap[list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetInt64Ptr returns a set of the distinct elements of coll.
func SetInt64Ptr(list []*int64) []*int64 {
	if list == nil || len(list) == 0 {
		return []*int64{}
	}

	resultMap := make(map[int64]struct{})
	newList := []*int64{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[*list[i]]
		if !ok {
			resultMap[*list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetInt32 returns a set of the distinct elements of coll.
func SetInt32(list []int32) []int32 {
	if list == nil || len(list) == 0 {
		return []int32{}
	}

	resultMap := make(map[int32]struct{})
	newList := []int32{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i]]
		if !ok {
			resultMap[list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetInt32Ptr returns a set of the distinct elements of coll.
func SetInt32Ptr(list []*int32) []*int32 {
	if list == nil || len(list) == 0 {
		return []*int32{}
	}

	resultMap := make(map[int32]struct{})
	newList := []*int32{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[*list[i]]
		if !ok {
			resultMap[*list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetInt16 returns a set of the distinct elements of coll.
func SetInt16(list []int16) []int16 {
	if list == nil || len(list) == 0 {
		return []int16{}
	}

	resultMap := make(map[int16]struct{})
	newList := []int16{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i]]
		if !ok {
			resultMap[list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetInt16Ptr returns a set of the distinct elements of coll.
func SetInt16Ptr(list []*int16) []*int16 {
	if list == nil || len(list) == 0 {
		return []*int16{}
	}

	resultMap := make(map[int16]struct{})
	newList := []*int16{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[*list[i]]
		if !ok {
			resultMap[*list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetInt8 returns a set of the distinct elements of coll.
func SetInt8(list []int8) []int8 {
	if list == nil || len(list) == 0 {
		return []int8{}
	}

	resultMap := make(map[int8]struct{})
	newList := []int8{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i]]
		if !ok {
			resultMap[list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetInt8Ptr returns a set of the distinct elements of coll.
func SetInt8Ptr(list []*int8) []*int8 {
	if list == nil || len(list) == 0 {
		return []*int8{}
	}

	resultMap := make(map[int8]struct{})
	newList := []*int8{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[*list[i]]
		if !ok {
			resultMap[*list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetUint returns a set of the distinct elements of coll.
func SetUint(list []uint) []uint {
	if list == nil || len(list) == 0 {
		return []uint{}
	}

	resultMap := make(map[uint]struct{})
	newList := []uint{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i]]
		if !ok {
			resultMap[list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetUintPtr returns a set of the distinct elements of coll.
func SetUintPtr(list []*uint) []*uint {
	if list == nil || len(list) == 0 {
		return []*uint{}
	}

	resultMap := make(map[uint]struct{})
	newList := []*uint{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[*list[i]]
		if !ok {
			resultMap[*list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetUint64 returns a set of the distinct elements of coll.
func SetUint64(list []uint64) []uint64 {
	if list == nil || len(list) == 0 {
		return []uint64{}
	}

	resultMap := make(map[uint64]struct{})
	newList := []uint64{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i]]
		if !ok {
			resultMap[list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetUint64Ptr returns a set of the distinct elements of coll.
func SetUint64Ptr(list []*uint64) []*uint64 {
	if list == nil || len(list) == 0 {
		return []*uint64{}
	}

	resultMap := make(map[uint64]struct{})
	newList := []*uint64{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[*list[i]]
		if !ok {
			resultMap[*list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetUint32 returns a set of the distinct elements of coll.
func SetUint32(list []uint32) []uint32 {
	if list == nil || len(list) == 0 {
		return []uint32{}
	}

	resultMap := make(map[uint32]struct{})
	newList := []uint32{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i]]
		if !ok {
			resultMap[list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetUint32Ptr returns a set of the distinct elements of coll.
func SetUint32Ptr(list []*uint32) []*uint32 {
	if list == nil || len(list) == 0 {
		return []*uint32{}
	}

	resultMap := make(map[uint32]struct{})
	newList := []*uint32{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[*list[i]]
		if !ok {
			resultMap[*list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetUint16 returns a set of the distinct elements of coll.
func SetUint16(list []uint16) []uint16 {
	if list == nil || len(list) == 0 {
		return []uint16{}
	}

	resultMap := make(map[uint16]struct{})
	newList := []uint16{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i]]
		if !ok {
			resultMap[list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetUint16Ptr returns a set of the distinct elements of coll.
func SetUint16Ptr(list []*uint16) []*uint16 {
	if list == nil || len(list) == 0 {
		return []*uint16{}
	}

	resultMap := make(map[uint16]struct{})
	newList := []*uint16{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[*list[i]]
		if !ok {
			resultMap[*list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetUint8 returns a set of the distinct elements of coll.
func SetUint8(list []uint8) []uint8 {
	if list == nil || len(list) == 0 {
		return []uint8{}
	}

	resultMap := make(map[uint8]struct{})
	newList := []uint8{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i]]
		if !ok {
			resultMap[list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetUint8Ptr returns a set of the distinct elements of coll.
func SetUint8Ptr(list []*uint8) []*uint8 {
	if list == nil || len(list) == 0 {
		return []*uint8{}
	}

	resultMap := make(map[uint8]struct{})
	newList := []*uint8{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[*list[i]]
		if !ok {
			resultMap[*list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetStr returns a set of the distinct elements of coll.
func SetStr(list []string) []string {
	if list == nil || len(list) == 0 {
		return []string{}
	}

	resultMap := make(map[string]struct{})
	newList := []string{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i]]
		if !ok {
			resultMap[list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetStrPtr returns a set of the distinct elements of coll.
func SetStrPtr(list []*string) []*string {
	if list == nil || len(list) == 0 {
		return []*string{}
	}

	resultMap := make(map[string]struct{})
	newList := []*string{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[*list[i]]
		if !ok {
			resultMap[*list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetBool returns a set of the distinct elements of coll.
func SetBool(list []bool) []bool {
	if list == nil || len(list) == 0 {
		return []bool{}
	}

	resultMap := make(map[bool]struct{})
	newList := []bool{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i]]
		if !ok {
			resultMap[list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetBoolPtr returns a set of the distinct elements of coll.
func SetBoolPtr(list []*bool) []*bool {
	if list == nil || len(list) == 0 {
		return []*bool{}
	}

	resultMap := make(map[bool]struct{})
	newList := []*bool{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[*list[i]]
		if !ok {
			resultMap[*list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetFloat32 returns a set of the distinct elements of coll.
func SetFloat32(list []float32) []float32 {
	if list == nil || len(list) == 0 {
		return []float32{}
	}

	resultMap := make(map[float32]struct{})
	newList := []float32{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i]]
		if !ok {
			resultMap[list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetFloat32Ptr returns a set of the distinct elements of coll.
func SetFloat32Ptr(list []*float32) []*float32 {
	if list == nil || len(list) == 0 {
		return []*float32{}
	}

	resultMap := make(map[float32]struct{})
	newList := []*float32{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[*list[i]]
		if !ok {
			resultMap[*list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetFloat64 returns a set of the distinct elements of coll.
func SetFloat64(list []float64) []float64 {
	if list == nil || len(list) == 0 {
		return []float64{}
	}

	resultMap := make(map[float64]struct{})
	newList := []float64{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[list[i]]
		if !ok {
			resultMap[list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}

// SetFloat64Ptr returns a set of the distinct elements of coll.
func SetFloat64Ptr(list []*float64) []*float64 {
	if list == nil || len(list) == 0 {
		return []*float64{}
	}

	resultMap := make(map[float64]struct{})
	newList := []*float64{}
	for i := 0; i < len(list); i++ {
		_, ok := resultMap[*list[i]]
		if !ok {
			resultMap[*list[i]] = struct{}{}
			newList = append(newList, list[i])
		}
	}
	return newList
}
