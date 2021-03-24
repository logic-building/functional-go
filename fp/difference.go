package fp

// DifferenceInt returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceInt(arrList ...[]int) []int {
	if arrList == nil {
		return []int{}
	}

	resultMap := make(map[int]struct{})
	if len(arrList) == 1 {
		var newList []int
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
		return newList
	}

	var newList []int
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceIntPtr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceIntPtr(arrList ...[]*int) []*int {
	if arrList == nil {
		return []*int{}
	}

	resultMap := make(map[int]struct{})
	if len(arrList) == 1 {
		var newList []*int
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = struct{}{}
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*int
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceInt64 returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceInt64(arrList ...[]int64) []int64 {
	if arrList == nil {
		return []int64{}
	}

	resultMap := make(map[int64]struct{})
	if len(arrList) == 1 {
		var newList []int64
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
		return newList
	}

	var newList []int64
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceInt64Ptr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceInt64Ptr(arrList ...[]*int64) []*int64 {
	if arrList == nil {
		return []*int64{}
	}

	resultMap := make(map[int64]struct{})
	if len(arrList) == 1 {
		var newList []*int64
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = struct{}{}
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*int64
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceInt32 returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceInt32(arrList ...[]int32) []int32 {
	if arrList == nil {
		return []int32{}
	}

	resultMap := make(map[int32]struct{})
	if len(arrList) == 1 {
		var newList []int32
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
		return newList
	}

	var newList []int32
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceInt32Ptr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceInt32Ptr(arrList ...[]*int32) []*int32 {
	if arrList == nil {
		return []*int32{}
	}

	resultMap := make(map[int32]struct{})
	if len(arrList) == 1 {
		var newList []*int32
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = struct{}{}
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*int32
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceInt16 returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceInt16(arrList ...[]int16) []int16 {
	if arrList == nil {
		return []int16{}
	}

	resultMap := make(map[int16]struct{})
	if len(arrList) == 1 {
		var newList []int16
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
		return newList
	}

	var newList []int16
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceInt16Ptr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceInt16Ptr(arrList ...[]*int16) []*int16 {
	if arrList == nil {
		return []*int16{}
	}

	resultMap := make(map[int16]struct{})
	if len(arrList) == 1 {
		var newList []*int16
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = struct{}{}
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*int16
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceInt8 returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceInt8(arrList ...[]int8) []int8 {
	if arrList == nil {
		return []int8{}
	}

	resultMap := make(map[int8]struct{})
	if len(arrList) == 1 {
		var newList []int8
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
		return newList
	}

	var newList []int8
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceInt8Ptr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceInt8Ptr(arrList ...[]*int8) []*int8 {
	if arrList == nil {
		return []*int8{}
	}

	resultMap := make(map[int8]struct{})
	if len(arrList) == 1 {
		var newList []*int8
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = struct{}{}
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*int8
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceUint returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceUint(arrList ...[]uint) []uint {
	if arrList == nil {
		return []uint{}
	}

	resultMap := make(map[uint]struct{})
	if len(arrList) == 1 {
		var newList []uint
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
		return newList
	}

	var newList []uint
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceUintPtr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceUintPtr(arrList ...[]*uint) []*uint {
	if arrList == nil {
		return []*uint{}
	}

	resultMap := make(map[uint]struct{})
	if len(arrList) == 1 {
		var newList []*uint
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = struct{}{}
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*uint
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceUint64 returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceUint64(arrList ...[]uint64) []uint64 {
	if arrList == nil {
		return []uint64{}
	}

	resultMap := make(map[uint64]struct{})
	if len(arrList) == 1 {
		var newList []uint64
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
		return newList
	}

	var newList []uint64
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceUint64Ptr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceUint64Ptr(arrList ...[]*uint64) []*uint64 {
	if arrList == nil {
		return []*uint64{}
	}

	resultMap := make(map[uint64]struct{})
	if len(arrList) == 1 {
		var newList []*uint64
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = struct{}{}
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*uint64
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceUint32 returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceUint32(arrList ...[]uint32) []uint32 {
	if arrList == nil {
		return []uint32{}
	}

	resultMap := make(map[uint32]struct{})
	if len(arrList) == 1 {
		var newList []uint32
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
		return newList
	}

	var newList []uint32
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceUint32Ptr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceUint32Ptr(arrList ...[]*uint32) []*uint32 {
	if arrList == nil {
		return []*uint32{}
	}

	resultMap := make(map[uint32]struct{})
	if len(arrList) == 1 {
		var newList []*uint32
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = struct{}{}
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*uint32
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceUint16 returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceUint16(arrList ...[]uint16) []uint16 {
	if arrList == nil {
		return []uint16{}
	}

	resultMap := make(map[uint16]struct{})
	if len(arrList) == 1 {
		var newList []uint16
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
		return newList
	}

	var newList []uint16
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceUint16Ptr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceUint16Ptr(arrList ...[]*uint16) []*uint16 {
	if arrList == nil {
		return []*uint16{}
	}

	resultMap := make(map[uint16]struct{})
	if len(arrList) == 1 {
		var newList []*uint16
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = struct{}{}
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*uint16
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceUint8 returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceUint8(arrList ...[]uint8) []uint8 {
	if arrList == nil {
		return []uint8{}
	}

	resultMap := make(map[uint8]struct{})
	if len(arrList) == 1 {
		var newList []uint8
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
		return newList
	}

	var newList []uint8
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceUint8Ptr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceUint8Ptr(arrList ...[]*uint8) []*uint8 {
	if arrList == nil {
		return []*uint8{}
	}

	resultMap := make(map[uint8]struct{})
	if len(arrList) == 1 {
		var newList []*uint8
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = struct{}{}
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*uint8
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceStr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceStr(arrList ...[]string) []string {
	if arrList == nil {
		return []string{}
	}

	resultMap := make(map[string]struct{})
	if len(arrList) == 1 {
		var newList []string
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
		return newList
	}

	var newList []string
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceStrPtr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceStrPtr(arrList ...[]*string) []*string {
	if arrList == nil {
		return []*string{}
	}

	resultMap := make(map[string]struct{})
	if len(arrList) == 1 {
		var newList []*string
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = struct{}{}
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*string
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceBool returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceBool(arrList ...[]bool) []bool {
	if arrList == nil {
		return []bool{}
	}

	resultMap := make(map[bool]struct{})
	if len(arrList) == 1 {
		var newList []bool
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
		return newList
	}

	var newList []bool
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceBoolPtr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceBoolPtr(arrList ...[]*bool) []*bool {
	if arrList == nil {
		return []*bool{}
	}

	resultMap := make(map[bool]struct{})
	if len(arrList) == 1 {
		var newList []*bool
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = struct{}{}
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*bool
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceFloat32 returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceFloat32(arrList ...[]float32) []float32 {
	if arrList == nil {
		return []float32{}
	}

	resultMap := make(map[float32]struct{})
	if len(arrList) == 1 {
		var newList []float32
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
		return newList
	}

	var newList []float32
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceFloat32Ptr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceFloat32Ptr(arrList ...[]*float32) []*float32 {
	if arrList == nil {
		return []*float32{}
	}

	resultMap := make(map[float32]struct{})
	if len(arrList) == 1 {
		var newList []*float32
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = struct{}{}
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*float32
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceFloat64 returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceFloat64(arrList ...[]float64) []float64 {
	if arrList == nil {
		return []float64{}
	}

	resultMap := make(map[float64]struct{})
	if len(arrList) == 1 {
		var newList []float64
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
		return newList
	}

	var newList []float64
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if arrList[0][i] == v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}

// DifferenceFloat64Ptr returns a set that is the first set without elements of the remaining sets
// repeated value within list parameter will be ignored
func DifferenceFloat64Ptr(arrList ...[]*float64) []*float64 {
	if arrList == nil {
		return []*float64{}
	}

	resultMap := make(map[float64]struct{})
	if len(arrList) == 1 {
		var newList []*float64
		for i := 0; i < len(arrList[0]); i++ {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				resultMap[*arrList[0][i]] = struct{}{}
				newList = append(newList, arrList[0][i])
			}
		}
		return newList
	}

	var newList []*float64
	// 1st loop iterates items in 1st array
	// 2nd loop iterates all the rest of the arrays
	// 3rd loop iterates items in the rest of the arrays
	for i := 0; i < len(arrList[0]); i++ {

		matchCount := 0
		for j := 1; j < len(arrList); j++ {
			for _, v := range arrList[j] {
				// compare every items in 1st array to every items in the rest of the arrays
				if *arrList[0][i] == *v {
					matchCount++
					break
				}
			}
		}
		if matchCount == 0 {
			_, ok := resultMap[*arrList[0][i]]
			if !ok {
				newList = append(newList, arrList[0][i])
				resultMap[*arrList[0][i]] = struct{}{}
			}
		}
	}
	return newList
}
