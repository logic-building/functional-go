package fp

func RangeInt(lower, higher int) []int {
	if lower >= higher {
		return []int{}
	}

	l := make([]int, higher-lower)
	for i, v := 0, lower; v < higher; v++ {
		l[i] = v
		i++
	}
	return l
}

func RangeInt64(lower, higher int64) []int64 {
	if lower >= higher {
		return []int64{}
	}

	l := make([]int64, higher-lower)
	for i, v := 0, lower; v < higher; v++ {
		l[i] = v
		i++
	}
	return l
}

func RangeInt32(lower, higher int32) []int32 {
	if lower >= higher {
		return []int32{}
	}

	l := make([]int32, higher-lower)
	for i, v := 0, lower; v < higher; v++ {
		l[i] = v
		i++
	}
	return l
}

func RangeInt16(lower, higher int16) []int16 {
	if lower >= higher {
		return []int16{}
	}

	l := make([]int16, higher-lower)
	for i, v := 0, lower; v < higher; v++ {
		l[i] = v
		i++
	}
	return l
}

func RangeInt8(lower, higher int8) []int8 {
	if lower >= higher {
		return []int8{}
	}

	l := make([]int8, higher-lower)
	for i, v := 0, lower; v < higher; v++ {
		l[i] = v
		i++
	}
	return l
}

func RangeUint(lower, higher uint) []uint {
	if lower >= higher {
		return []uint{}
	}

	l := make([]uint, higher-lower)
	for i, v := 0, lower; v < higher; v++ {
		l[i] = v
		i++
	}
	return l
}

func RangeUint64(lower, higher uint64) []uint64 {
	if lower >= higher {
		return []uint64{}
	}

	l := make([]uint64, higher-lower)
	for i, v := 0, lower; v < higher; v++ {
		l[i] = v
		i++
	}
	return l
}

func RangeUint32(lower, higher uint32) []uint32 {
	if lower >= higher {
		return []uint32{}
	}

	l := make([]uint32, higher-lower)
	for i, v := 0, lower; v < higher; v++ {
		l[i] = v
		i++
	}
	return l
}

func RangeUint16(lower, higher uint16) []uint16 {
	if lower >= higher {
		return []uint16{}
	}

	l := make([]uint16, higher-lower)
	for i, v := 0, lower; v < higher; v++ {
		l[i] = v
		i++
	}
	return l
}

func RangeUint8(lower, higher uint8) []uint8 {
	if lower >= higher {
		return []uint8{}
	}

	l := make([]uint8, higher-lower)
	for i, v := 0, lower; v < higher; v++ {
		l[i] = v
		i++
	}
	return l
}
