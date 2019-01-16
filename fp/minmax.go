package fp

func MinMaxInt(list []int) (int, int) {
	if list == nil || len(list) == 0 {
		return 0, 0
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}

func MinMaxInt64(list []int64) (int64, int64) {
	if list == nil || len(list) == 0 {
		return 0, 0
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}

func MinMaxInt32(list []int32) (int32, int32) {
	if list == nil || len(list) == 0 {
		return 0, 0
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}

func MinMaxInt16(list []int16) (int16, int16) {
	if list == nil || len(list) == 0 {
		return 0, 0
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}

func MinMaxInt8(list []int8) (int8, int8) {
	if list == nil || len(list) == 0 {
		return 0, 0
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}

func MinMaxUint(list []uint) (uint, uint) {
	if list == nil || len(list) == 0 {
		return 0, 0
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}

func MinMaxUint64(list []uint64) (uint64, uint64) {
	if list == nil || len(list) == 0 {
		return 0, 0
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}

func MinMaxUint32(list []uint32) (uint32, uint32) {
	if list == nil || len(list) == 0 {
		return 0, 0
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}

func MinMaxUint16(list []uint16) (uint16, uint16) {
	if list == nil || len(list) == 0 {
		return 0, 0
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}

func MinMaxUint8(list []uint8) (uint8, uint8) {
	if list == nil || len(list) == 0 {
		return 0, 0
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}

func MinMaxFloat64(list []float64) (float64, float64) {
	if list == nil || len(list) == 0 {
		return 0, 0
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}

func MinMaxFloat32(list []float32) (float32, float32) {
	if list == nil || len(list) == 0 {
		return 0, 0
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}
	return min, max
}
