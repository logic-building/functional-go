package fp

// SupersetInt returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetInt(list1, list2 []int) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetIntPtr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetIntPtr(list1, list2 []*int) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetInt64 returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetInt64(list1, list2 []int64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int64]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetInt64Ptr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetInt64Ptr(list1, list2 []*int64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int64]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetInt32 returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetInt32(list1, list2 []int32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int32]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetInt32Ptr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetInt32Ptr(list1, list2 []*int32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int32]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetInt16 returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetInt16(list1, list2 []int16) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int16]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetInt16Ptr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetInt16Ptr(list1, list2 []*int16) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int16]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetInt8 returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetInt8(list1, list2 []int8) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int8]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetInt8Ptr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetInt8Ptr(list1, list2 []*int8) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int8]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetUint returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetUint(list1, list2 []uint) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetUintPtr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetUintPtr(list1, list2 []*uint) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetUint64 returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetUint64(list1, list2 []uint64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint64]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetUint64Ptr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetUint64Ptr(list1, list2 []*uint64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint64]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetUint32 returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetUint32(list1, list2 []uint32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint32]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetUint32Ptr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetUint32Ptr(list1, list2 []*uint32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint32]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetUint16 returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetUint16(list1, list2 []uint16) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint16]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetUint16Ptr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetUint16Ptr(list1, list2 []*uint16) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint16]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetUint8 returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetUint8(list1, list2 []uint8) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint8]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetUint8Ptr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetUint8Ptr(list1, list2 []*uint8) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint8]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetStr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetStr(list1, list2 []string) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[string]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetStrPtr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetStrPtr(list1, list2 []*string) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[string]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetBool returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetBool(list1, list2 []bool) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[bool]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetBoolPtr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetBoolPtr(list1, list2 []*bool) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[bool]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetFloat32 returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetFloat32(list1, list2 []float32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[float32]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetFloat32Ptr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetFloat32Ptr(list1, list2 []*float32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[float32]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetFloat64 returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetFloat64(list1, list2 []float64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[float64]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[list2[i]]
		if !ok {
			found := false
			resultMap[list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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

// SupersetFloat64Ptr returns true or false by checking if set1 is a superset of set2
// repeated value within list parameter will be ignored
func SupersetFloat64Ptr(list1, list2 []*float64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[float64]bool)

	for i := 0; i < len(list2); i++ {
		_, ok := resultMap[*list2[i]]
		if !ok {
			found := false
			resultMap[*list2[i]] = true
			for j := 0; j < len(list1); j++ {
				if list2[i] == list1[j] {
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
