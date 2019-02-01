package fp

// Returns a list of range between lower and upper value
//
// Takes 3 inputs
//	1. lower limit
//	2. Upper limit
//	3. Hops (optional)
//
// Returns
//	List of range between lower and upper value
//	Empty list if 3rd argument is either 0 or negative number
//
// Example:
//	RangeInt(-2, 2) // Returns: [-2, -1, 0, 1]
//	RangeInt(0, 2) // Returns: [0, 1]
//	RangeInt(3, 7, 2) // Returns: [3, 5]
func RangeInt(lower, higher int, hops ...int) []int {
	hop := 1
	if len(hops) >= 1 {
		if hops[0] <= 0 {
			return []int{}
		}
		hop = hops[0]
	}

	if lower >= higher {
		return []int{}
	}

	l := make([]int, higher-lower)
	for i, v := 0, lower; v < higher; v += hop {
		l[i] = v
		i++
	}
	return l
}

// Returns a list of range between lower and upper value
//
// Takes 3 inputs
//	1. lower limit
//	2. Upper limit
//	3. Hops (optional)
//
// Returns
//	List of range between lower and upper value
//	Empty list if 3rd argument is either 0 or negative number
//
// Example:
//	RangeInt64(-2, 2) // Returns: [-2, -1, 0, 1]
//	RangeInt64(0, 2) // Returns: [0, 1]
//	RangeInt64(3, 7, 2) // Returns: [3, 5]
func RangeInt64(lower, higher int64, hops ...int64) []int64 {
	var hop int64 = 1
	if len(hops) >= 1 {
		if hops[0] <= 0 {
			return []int64{}
		}
		hop = hops[0]
	}

	if lower >= higher {
		return []int64{}
	}

	l := make([]int64, higher-lower)
	for i, v := 0, lower; v < higher; v += hop {
		l[i] = v
		i++
	}
	return l
}

// Returns a list of range between lower and upper value
//
// Takes 3 inputs
//	1. lower limit
//	2. Upper limit
//	3. Hops (optional)
//
// Returns
//	List of range between lower and upper value
//	Empty list if 3rd argument is either 0 or negative number
//
// Example:
//	RangeInt32(-2, 2) // Returns: [-2, -1, 0, 1]
//	RangeInt32(0, 2) // Returns: [0, 1]
//	RangeInt32(3, 7, 2) // Returns: [3, 5]
func RangeInt32(lower, higher int32, hops ...int32) []int32 {
	var hop int32 = 1
	if len(hops) >= 1 {
		if hops[0] <= 0 {
			return []int32{}
		}
		hop = hops[0]
	}

	if lower >= higher {
		return []int32{}
	}

	l := make([]int32, higher-lower)
	for i, v := 0, lower; v < higher; v += hop {
		l[i] = v
		i++
	}
	return l
}

// Returns a list of range between lower and upper value
//
// Takes 3 inputs
//	1. lower limit
//	2. Upper limit
//	3. Hops (optional)
//
// Returns
//	List of range between lower and upper value
//	Empty list if 3rd argument is either 0 or negative number
//
// Example:
//	RangeInt16(-2, 2) // Returns: [-2, -1, 0, 1]
//	RangeInt16(0, 2) // Returns: [0, 1]
//	RangeInt16(3, 7, 2) // Returns: [3, 5]
func RangeInt16(lower, higher int16, hops ...int16) []int16 {
	var hop int16 = 1
	if len(hops) >= 1 {
		if hops[0] <= 0 {
			return []int16{}
		}
		hop = hops[0]
	}

	if lower >= higher {
		return []int16{}
	}

	l := make([]int16, higher-lower)
	for i, v := 0, lower; v < higher; v += hop {
		l[i] = v
		i++
	}
	return l
}

// Returns a list of range between lower and upper value
//
// Takes 3 inputs
//	1. lower limit
//	2. Upper limit
//	3. Hops (optional)
//
// Returns
//	List of range between lower and upper value
//	Empty list if 3rd argument is either 0 or negative number
//
// Example:
//	RangeInt8(-2, 2) // Returns: [-2, -1, 0, 1]
//	RangeInt8(0, 2) // Returns: [0, 1]
//	RangeInt8(3, 7, 2) // Returns: [3, 5]
func RangeInt8(lower, higher int8, hops ...int8) []int8 {
	var hop int8 = 1
	if len(hops) >= 1 {
		if hops[0] <= 0 {
			return []int8{}
		}
		hop = hops[0]
	}

	if lower >= higher {
		return []int8{}
	}

	l := make([]int8, higher-lower)
	for i, v := 0, lower; v < higher; v += hop {
		l[i] = v
		i++
	}
	return l
}

// Returns a list of range between lower and upper value
//
// Takes 3 inputs
//	1. lower limit
//	2. Upper limit
//	3. Hops (optional)
//
// Returns
//	List of range between lower and upper value
//	Empty list if 3rd argument is either 0 or negative number
//
// Example:
//	RangeUint(-2, 2) // Returns: [-2, -1, 0, 1]
//	RangeUint(0, 2) // Returns: [0, 1]
//	RangeUint(3, 7, 2) // Returns: [3, 5]
func RangeUint(lower, higher uint, hops ...uint) []uint {
	var hop uint = 1
	if len(hops) >= 1 {
		if hops[0] == 0 {
			return []uint{}
		}
		hop = hops[0]
	}

	if lower >= higher {
		return []uint{}
	}

	l := make([]uint, higher-lower)
	for i, v := 0, lower; v < higher; v += hop {
		l[i] = v
		i++
	}
	return l
}

// Returns a list of range between lower and upper value
//
// Takes 3 inputs
//	1. lower limit
//	2. Upper limit
//	3. Hops (optional)
//
// Returns
//	List of range between lower and upper value
//	Empty list if 3rd argument is either 0 or negative number
//
// Example:
//	RangeUint64(-2, 2) // Returns: [-2, -1, 0, 1]
//	RangeUint64(0, 2) // Returns: [0, 1]
//	RangeUint64(3, 7, 2) // Returns: [3, 5]
func RangeUint64(lower, higher uint64, hops ...uint64) []uint64 {
	var hop uint64 = 1
	if len(hops) >= 1 {
		if hops[0] == 0 {
			return []uint64{}
		}
		hop = hops[0]
	}

	if lower >= higher {
		return []uint64{}
	}

	l := make([]uint64, higher-lower)
	for i, v := 0, lower; v < higher; v += hop {
		l[i] = v
		i++
	}
	return l
}

// Returns a list of range between lower and upper value
//
// Takes 3 inputs
//	1. lower limit
//	2. Upper limit
//	3. Hops (optional)
//
// Returns
//	List of range between lower and upper value
//	Empty list if 3rd argument is either 0 or negative number
//
// Example:
//	RangeUint32(-2, 2) // Returns: [-2, -1, 0, 1]
//	RangeUint32(0, 2) // Returns: [0, 1]
//	RangeUint32(3, 7, 2) // Returns: [3, 5]
func RangeUint32(lower, higher uint32, hops ...uint32) []uint32 {
	var hop uint32 = 1
	if len(hops) >= 1 {
		if hops[0] == 0 {
			return []uint32{}
		}
		hop = hops[0]
	}

	if lower >= higher {
		return []uint32{}
	}

	l := make([]uint32, higher-lower)
	for i, v := 0, lower; v < higher; v += hop {
		l[i] = v
		i++
	}
	return l
}

// Returns a list of range between lower and upper value
//
// Takes 3 inputs
//	1. lower limit
//	2. Upper limit
//	3. Hops (optional)
//
// Returns
//	List of range between lower and upper value
//	Empty list if 3rd argument is either 0 or negative number
//
// Example:
//	RangeUint16(-2, 2) // Returns: [-2, -1, 0, 1]
//	RangeUint16(0, 2) // Returns: [0, 1]
//	RangeUint16(3, 7, 2) // Returns: [3, 5]
func RangeUint16(lower, higher uint16, hops ...uint16) []uint16 {
	var hop uint16 = 1
	if len(hops) >= 1 {
		if hops[0] == 0 {
			return []uint16{}
		}
		hop = hops[0]
	}
	if lower >= higher {
		return []uint16{}
	}

	l := make([]uint16, higher-lower)
	for i, v := 0, lower; v < higher; v += hop {
		l[i] = v
		i++
	}
	return l
}

// Returns a list of range between lower and upper value
//
// Takes 3 inputs
//	1. lower limit
//	2. Upper limit
//	3. Hops (optional)
//
// Returns
//	List of range between lower and upper value
//	Empty list if 3rd argument is either 0 or negative number
//
// Example:
//	RangeUint8(-2, 2) // Returns: [-2, -1, 0, 1]
//	RangeUint8(0, 2) // Returns: [0, 1]
//	RangeUint8(3, 7, 2) // Returns: [3, 5]
func RangeUint8(lower, higher uint8, hops ...uint8) []uint8 {
	var hop uint8 = 1
	if len(hops) >= 1 {
		if hops[0] == 0 {
			return []uint8{}
		}
		hop = hops[0]
	}

	if lower >= higher {
		return []uint8{}
	}

	l := make([]uint8, higher-lower)
	for i, v := 0, lower; v < higher; v += hop {
		l[i] = v
		i++
	}
	return l
}
