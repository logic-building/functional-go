package fp

// SubsetInt returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt(list1, list2 []int) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetIntPtr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetIntPtr(list1, list2 []*int) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetInt64 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt64(list1, list2 []int64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int64]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetInt64Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt64Ptr(list1, list2 []*int64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int64]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetInt32 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt32(list1, list2 []int32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int32]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetInt32Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt32Ptr(list1, list2 []*int32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int32]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetInt16 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt16(list1, list2 []int16) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int16]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetInt16Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt16Ptr(list1, list2 []*int16) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int16]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetInt8 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt8(list1, list2 []int8) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int8]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetInt8Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetInt8Ptr(list1, list2 []*int8) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[int8]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetUint returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint(list1, list2 []uint) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetUintPtr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUintPtr(list1, list2 []*uint) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetUint64 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint64(list1, list2 []uint64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint64]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetUint64Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint64Ptr(list1, list2 []*uint64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint64]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetUint32 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint32(list1, list2 []uint32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint32]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetUint32Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint32Ptr(list1, list2 []*uint32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint32]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetUint16 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint16(list1, list2 []uint16) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint16]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetUint16Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint16Ptr(list1, list2 []*uint16) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint16]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetUint8 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint8(list1, list2 []uint8) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint8]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetUint8Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetUint8Ptr(list1, list2 []*uint8) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[uint8]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetStr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetStr(list1, list2 []string) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[string]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetStrPtr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetStrPtr(list1, list2 []*string) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[string]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetBool returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetBool(list1, list2 []bool) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[bool]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetBoolPtr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetBoolPtr(list1, list2 []*bool) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[bool]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetFloat32 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetFloat32(list1, list2 []float32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[float32]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetFloat32Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetFloat32Ptr(list1, list2 []*float32) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[float32]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetFloat64 returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetFloat64(list1, list2 []float64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[float64]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[list1[i]]
		if !ok {
			resultMap[list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}

// SubsetFloat64Ptr returns true or false by checking if set1 is a subset of set2
// repeated value within list parameter will be ignored
func SubsetFloat64Ptr(list1, list2 []*float64) bool {
	if list1 == nil || len(list1) == 0 || list2 == nil || len(list2) == 0 {
		return false
	}

	resultMap := make(map[float64]bool)
	foundCounter := 0
	for i := 0; i < len(list1); i++ {
		_, ok := resultMap[*list1[i]]
		if !ok {
			resultMap[*list1[i]] = true
			for j := 0; j < len(list2); j++ {
				if list1[i] == list2[j] {
					foundCounter++
					break
				}
			}
		}
	}

	if len(resultMap) == foundCounter {
		return true
	}
	return false
}
