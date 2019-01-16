package fp

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
