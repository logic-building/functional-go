package fp

// SubsetInt returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt(list1, list2 []int) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			found := false
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetIntPtr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetIntPtr(list1, list2 []*int) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			found := false
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetInt64 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt64(list1, list2 []int64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int64]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			found := false
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetInt64Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt64Ptr(list1, list2 []*int64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int64]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			found := false
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetInt32 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt32(list1, list2 []int32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int32]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			found := false
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetInt32Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt32Ptr(list1, list2 []*int32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int32]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			found := false
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetInt16 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt16(list1, list2 []int16) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int16]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			found := false
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetInt16Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt16Ptr(list1, list2 []*int16) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int16]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			found := false
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetInt8 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt8(list1, list2 []int8) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int8]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			found := false
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetInt8Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt8Ptr(list1, list2 []*int8) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int8]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			found := false
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetUint returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint(list1, list2 []uint) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			found := false
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetUintPtr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUintPtr(list1, list2 []*uint) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			found := false
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetUint64 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint64(list1, list2 []uint64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint64]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			found := false
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetUint64Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint64Ptr(list1, list2 []*uint64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint64]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			found := false
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetUint32 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint32(list1, list2 []uint32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint32]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			found := false
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetUint32Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint32Ptr(list1, list2 []*uint32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint32]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			found := false
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetUint16 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint16(list1, list2 []uint16) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint16]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			found := false
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetUint16Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint16Ptr(list1, list2 []*uint16) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint16]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			found := false
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetUint8 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint8(list1, list2 []uint8) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint8]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			found := false
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetUint8Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint8Ptr(list1, list2 []*uint8) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint8]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			found := false
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetStr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetStr(list1, list2 []string) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[string]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			found := false
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetStrPtr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetStrPtr(list1, list2 []*string) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[string]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			found := false
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetBool returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetBool(list1, list2 []bool) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[bool]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			found := false
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetBoolPtr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetBoolPtr(list1, list2 []*bool) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[bool]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			found := false
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetFloat32 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetFloat32(list1, list2 []float32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[float32]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			found := false
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetFloat32Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetFloat32Ptr(list1, list2 []*float32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[float32]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			found := false
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetFloat64 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetFloat64(list1, list2 []float64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[float64]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			found := false
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}

// SubsetFloat64Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetFloat64Ptr(list1, list2 []*float64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[float64]bool)
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			found := false
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					found = true
					break
				}
			}
			if !found {
				return false
			}
		}
	}
	return true
}
