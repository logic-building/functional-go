package fp

// EqualMapIntP Returns true if both maps are equal else returns false
func EqualMapIntP(map1, map2 map[int]int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntPPtr Returns true if both maps are equal else returns false
func EqualMapIntPPtr(map1, map2 map[*int]*int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntInt64P Returns true if both maps are equal else returns false
func EqualMapIntInt64P(map1, map2 map[int]int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntInt64PPtr Returns true if both maps are equal else returns false
func EqualMapIntInt64PPtr(map1, map2 map[*int]*int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntInt32P Returns true if both maps are equal else returns false
func EqualMapIntInt32P(map1, map2 map[int]int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntInt32PPtr Returns true if both maps are equal else returns false
func EqualMapIntInt32PPtr(map1, map2 map[*int]*int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntInt16P Returns true if both maps are equal else returns false
func EqualMapIntInt16P(map1, map2 map[int]int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntInt16PPtr Returns true if both maps are equal else returns false
func EqualMapIntInt16PPtr(map1, map2 map[*int]*int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntInt8P Returns true if both maps are equal else returns false
func EqualMapIntInt8P(map1, map2 map[int]int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntInt8PPtr Returns true if both maps are equal else returns false
func EqualMapIntInt8PPtr(map1, map2 map[*int]*int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntUintP Returns true if both maps are equal else returns false
func EqualMapIntUintP(map1, map2 map[int]uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntUintPPtr Returns true if both maps are equal else returns false
func EqualMapIntUintPPtr(map1, map2 map[*int]*uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntUint64P Returns true if both maps are equal else returns false
func EqualMapIntUint64P(map1, map2 map[int]uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntUint64PPtr Returns true if both maps are equal else returns false
func EqualMapIntUint64PPtr(map1, map2 map[*int]*uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntUint32P Returns true if both maps are equal else returns false
func EqualMapIntUint32P(map1, map2 map[int]uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntUint32PPtr Returns true if both maps are equal else returns false
func EqualMapIntUint32PPtr(map1, map2 map[*int]*uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntUint16P Returns true if both maps are equal else returns false
func EqualMapIntUint16P(map1, map2 map[int]uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntUint16PPtr Returns true if both maps are equal else returns false
func EqualMapIntUint16PPtr(map1, map2 map[*int]*uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntUint8P Returns true if both maps are equal else returns false
func EqualMapIntUint8P(map1, map2 map[int]uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntUint8PPtr Returns true if both maps are equal else returns false
func EqualMapIntUint8PPtr(map1, map2 map[*int]*uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntStrP Returns true if both maps are equal else returns false
func EqualMapIntStrP(map1, map2 map[int]string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntStrPPtr Returns true if both maps are equal else returns false
func EqualMapIntStrPPtr(map1, map2 map[*int]*string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntBoolP Returns true if both maps are equal else returns false
func EqualMapIntBoolP(map1, map2 map[int]bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntBoolPPtr Returns true if both maps are equal else returns false
func EqualMapIntBoolPPtr(map1, map2 map[*int]*bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntFloat32P Returns true if both maps are equal else returns false
func EqualMapIntFloat32P(map1, map2 map[int]float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntFloat32PPtr Returns true if both maps are equal else returns false
func EqualMapIntFloat32PPtr(map1, map2 map[*int]*float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntFloat64P Returns true if both maps are equal else returns false
func EqualMapIntFloat64P(map1, map2 map[int]float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapIntFloat64PPtr Returns true if both maps are equal else returns false
func EqualMapIntFloat64PPtr(map1, map2 map[*int]*float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64IntP Returns true if both maps are equal else returns false
func EqualMapInt64IntP(map1, map2 map[int64]int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64IntPPtr Returns true if both maps are equal else returns false
func EqualMapInt64IntPPtr(map1, map2 map[*int64]*int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64P Returns true if both maps are equal else returns false
func EqualMapInt64P(map1, map2 map[int64]int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64PPtr Returns true if both maps are equal else returns false
func EqualMapInt64PPtr(map1, map2 map[*int64]*int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Int32P Returns true if both maps are equal else returns false
func EqualMapInt64Int32P(map1, map2 map[int64]int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Int32PPtr Returns true if both maps are equal else returns false
func EqualMapInt64Int32PPtr(map1, map2 map[*int64]*int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Int16P Returns true if both maps are equal else returns false
func EqualMapInt64Int16P(map1, map2 map[int64]int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Int16PPtr Returns true if both maps are equal else returns false
func EqualMapInt64Int16PPtr(map1, map2 map[*int64]*int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Int8P Returns true if both maps are equal else returns false
func EqualMapInt64Int8P(map1, map2 map[int64]int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Int8PPtr Returns true if both maps are equal else returns false
func EqualMapInt64Int8PPtr(map1, map2 map[*int64]*int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64UintP Returns true if both maps are equal else returns false
func EqualMapInt64UintP(map1, map2 map[int64]uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64UintPPtr Returns true if both maps are equal else returns false
func EqualMapInt64UintPPtr(map1, map2 map[*int64]*uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Uint64P Returns true if both maps are equal else returns false
func EqualMapInt64Uint64P(map1, map2 map[int64]uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Uint64PPtr Returns true if both maps are equal else returns false
func EqualMapInt64Uint64PPtr(map1, map2 map[*int64]*uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Uint32P Returns true if both maps are equal else returns false
func EqualMapInt64Uint32P(map1, map2 map[int64]uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Uint32PPtr Returns true if both maps are equal else returns false
func EqualMapInt64Uint32PPtr(map1, map2 map[*int64]*uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Uint16P Returns true if both maps are equal else returns false
func EqualMapInt64Uint16P(map1, map2 map[int64]uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Uint16PPtr Returns true if both maps are equal else returns false
func EqualMapInt64Uint16PPtr(map1, map2 map[*int64]*uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Uint8P Returns true if both maps are equal else returns false
func EqualMapInt64Uint8P(map1, map2 map[int64]uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Uint8PPtr Returns true if both maps are equal else returns false
func EqualMapInt64Uint8PPtr(map1, map2 map[*int64]*uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64StrP Returns true if both maps are equal else returns false
func EqualMapInt64StrP(map1, map2 map[int64]string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64StrPPtr Returns true if both maps are equal else returns false
func EqualMapInt64StrPPtr(map1, map2 map[*int64]*string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64BoolP Returns true if both maps are equal else returns false
func EqualMapInt64BoolP(map1, map2 map[int64]bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64BoolPPtr Returns true if both maps are equal else returns false
func EqualMapInt64BoolPPtr(map1, map2 map[*int64]*bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Float32P Returns true if both maps are equal else returns false
func EqualMapInt64Float32P(map1, map2 map[int64]float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Float32PPtr Returns true if both maps are equal else returns false
func EqualMapInt64Float32PPtr(map1, map2 map[*int64]*float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Float64P Returns true if both maps are equal else returns false
func EqualMapInt64Float64P(map1, map2 map[int64]float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt64Float64PPtr Returns true if both maps are equal else returns false
func EqualMapInt64Float64PPtr(map1, map2 map[*int64]*float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32IntP Returns true if both maps are equal else returns false
func EqualMapInt32IntP(map1, map2 map[int32]int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32IntPPtr Returns true if both maps are equal else returns false
func EqualMapInt32IntPPtr(map1, map2 map[*int32]*int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Int64P Returns true if both maps are equal else returns false
func EqualMapInt32Int64P(map1, map2 map[int32]int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Int64PPtr Returns true if both maps are equal else returns false
func EqualMapInt32Int64PPtr(map1, map2 map[*int32]*int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32P Returns true if both maps are equal else returns false
func EqualMapInt32P(map1, map2 map[int32]int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32PPtr Returns true if both maps are equal else returns false
func EqualMapInt32PPtr(map1, map2 map[*int32]*int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Int16P Returns true if both maps are equal else returns false
func EqualMapInt32Int16P(map1, map2 map[int32]int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Int16PPtr Returns true if both maps are equal else returns false
func EqualMapInt32Int16PPtr(map1, map2 map[*int32]*int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Int8P Returns true if both maps are equal else returns false
func EqualMapInt32Int8P(map1, map2 map[int32]int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Int8PPtr Returns true if both maps are equal else returns false
func EqualMapInt32Int8PPtr(map1, map2 map[*int32]*int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32UintP Returns true if both maps are equal else returns false
func EqualMapInt32UintP(map1, map2 map[int32]uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32UintPPtr Returns true if both maps are equal else returns false
func EqualMapInt32UintPPtr(map1, map2 map[*int32]*uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Uint64P Returns true if both maps are equal else returns false
func EqualMapInt32Uint64P(map1, map2 map[int32]uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Uint64PPtr Returns true if both maps are equal else returns false
func EqualMapInt32Uint64PPtr(map1, map2 map[*int32]*uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Uint32P Returns true if both maps are equal else returns false
func EqualMapInt32Uint32P(map1, map2 map[int32]uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Uint32PPtr Returns true if both maps are equal else returns false
func EqualMapInt32Uint32PPtr(map1, map2 map[*int32]*uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Uint16P Returns true if both maps are equal else returns false
func EqualMapInt32Uint16P(map1, map2 map[int32]uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Uint16PPtr Returns true if both maps are equal else returns false
func EqualMapInt32Uint16PPtr(map1, map2 map[*int32]*uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Uint8P Returns true if both maps are equal else returns false
func EqualMapInt32Uint8P(map1, map2 map[int32]uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Uint8PPtr Returns true if both maps are equal else returns false
func EqualMapInt32Uint8PPtr(map1, map2 map[*int32]*uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32StrP Returns true if both maps are equal else returns false
func EqualMapInt32StrP(map1, map2 map[int32]string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32StrPPtr Returns true if both maps are equal else returns false
func EqualMapInt32StrPPtr(map1, map2 map[*int32]*string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32BoolP Returns true if both maps are equal else returns false
func EqualMapInt32BoolP(map1, map2 map[int32]bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32BoolPPtr Returns true if both maps are equal else returns false
func EqualMapInt32BoolPPtr(map1, map2 map[*int32]*bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Float32P Returns true if both maps are equal else returns false
func EqualMapInt32Float32P(map1, map2 map[int32]float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Float32PPtr Returns true if both maps are equal else returns false
func EqualMapInt32Float32PPtr(map1, map2 map[*int32]*float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Float64P Returns true if both maps are equal else returns false
func EqualMapInt32Float64P(map1, map2 map[int32]float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt32Float64PPtr Returns true if both maps are equal else returns false
func EqualMapInt32Float64PPtr(map1, map2 map[*int32]*float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16IntP Returns true if both maps are equal else returns false
func EqualMapInt16IntP(map1, map2 map[int16]int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16IntPPtr Returns true if both maps are equal else returns false
func EqualMapInt16IntPPtr(map1, map2 map[*int16]*int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Int64P Returns true if both maps are equal else returns false
func EqualMapInt16Int64P(map1, map2 map[int16]int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Int64PPtr Returns true if both maps are equal else returns false
func EqualMapInt16Int64PPtr(map1, map2 map[*int16]*int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Int32P Returns true if both maps are equal else returns false
func EqualMapInt16Int32P(map1, map2 map[int16]int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Int32PPtr Returns true if both maps are equal else returns false
func EqualMapInt16Int32PPtr(map1, map2 map[*int16]*int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16P Returns true if both maps are equal else returns false
func EqualMapInt16P(map1, map2 map[int16]int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16PPtr Returns true if both maps are equal else returns false
func EqualMapInt16PPtr(map1, map2 map[*int16]*int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Int8P Returns true if both maps are equal else returns false
func EqualMapInt16Int8P(map1, map2 map[int16]int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Int8PPtr Returns true if both maps are equal else returns false
func EqualMapInt16Int8PPtr(map1, map2 map[*int16]*int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16UintP Returns true if both maps are equal else returns false
func EqualMapInt16UintP(map1, map2 map[int16]uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16UintPPtr Returns true if both maps are equal else returns false
func EqualMapInt16UintPPtr(map1, map2 map[*int16]*uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Uint64P Returns true if both maps are equal else returns false
func EqualMapInt16Uint64P(map1, map2 map[int16]uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Uint64PPtr Returns true if both maps are equal else returns false
func EqualMapInt16Uint64PPtr(map1, map2 map[*int16]*uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Uint32P Returns true if both maps are equal else returns false
func EqualMapInt16Uint32P(map1, map2 map[int16]uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Uint32PPtr Returns true if both maps are equal else returns false
func EqualMapInt16Uint32PPtr(map1, map2 map[*int16]*uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Uint16P Returns true if both maps are equal else returns false
func EqualMapInt16Uint16P(map1, map2 map[int16]uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Uint16PPtr Returns true if both maps are equal else returns false
func EqualMapInt16Uint16PPtr(map1, map2 map[*int16]*uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Uint8P Returns true if both maps are equal else returns false
func EqualMapInt16Uint8P(map1, map2 map[int16]uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Uint8PPtr Returns true if both maps are equal else returns false
func EqualMapInt16Uint8PPtr(map1, map2 map[*int16]*uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16StrP Returns true if both maps are equal else returns false
func EqualMapInt16StrP(map1, map2 map[int16]string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16StrPPtr Returns true if both maps are equal else returns false
func EqualMapInt16StrPPtr(map1, map2 map[*int16]*string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16BoolP Returns true if both maps are equal else returns false
func EqualMapInt16BoolP(map1, map2 map[int16]bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16BoolPPtr Returns true if both maps are equal else returns false
func EqualMapInt16BoolPPtr(map1, map2 map[*int16]*bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Float32P Returns true if both maps are equal else returns false
func EqualMapInt16Float32P(map1, map2 map[int16]float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Float32PPtr Returns true if both maps are equal else returns false
func EqualMapInt16Float32PPtr(map1, map2 map[*int16]*float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Float64P Returns true if both maps are equal else returns false
func EqualMapInt16Float64P(map1, map2 map[int16]float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt16Float64PPtr Returns true if both maps are equal else returns false
func EqualMapInt16Float64PPtr(map1, map2 map[*int16]*float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8IntP Returns true if both maps are equal else returns false
func EqualMapInt8IntP(map1, map2 map[int8]int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8IntPPtr Returns true if both maps are equal else returns false
func EqualMapInt8IntPPtr(map1, map2 map[*int8]*int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Int64P Returns true if both maps are equal else returns false
func EqualMapInt8Int64P(map1, map2 map[int8]int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Int64PPtr Returns true if both maps are equal else returns false
func EqualMapInt8Int64PPtr(map1, map2 map[*int8]*int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Int32P Returns true if both maps are equal else returns false
func EqualMapInt8Int32P(map1, map2 map[int8]int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Int32PPtr Returns true if both maps are equal else returns false
func EqualMapInt8Int32PPtr(map1, map2 map[*int8]*int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Int16P Returns true if both maps are equal else returns false
func EqualMapInt8Int16P(map1, map2 map[int8]int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Int16PPtr Returns true if both maps are equal else returns false
func EqualMapInt8Int16PPtr(map1, map2 map[*int8]*int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8P Returns true if both maps are equal else returns false
func EqualMapInt8P(map1, map2 map[int8]int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8PPtr Returns true if both maps are equal else returns false
func EqualMapInt8PPtr(map1, map2 map[*int8]*int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8UintP Returns true if both maps are equal else returns false
func EqualMapInt8UintP(map1, map2 map[int8]uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8UintPPtr Returns true if both maps are equal else returns false
func EqualMapInt8UintPPtr(map1, map2 map[*int8]*uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Uint64P Returns true if both maps are equal else returns false
func EqualMapInt8Uint64P(map1, map2 map[int8]uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Uint64PPtr Returns true if both maps are equal else returns false
func EqualMapInt8Uint64PPtr(map1, map2 map[*int8]*uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Uint32P Returns true if both maps are equal else returns false
func EqualMapInt8Uint32P(map1, map2 map[int8]uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Uint32PPtr Returns true if both maps are equal else returns false
func EqualMapInt8Uint32PPtr(map1, map2 map[*int8]*uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Uint16P Returns true if both maps are equal else returns false
func EqualMapInt8Uint16P(map1, map2 map[int8]uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Uint16PPtr Returns true if both maps are equal else returns false
func EqualMapInt8Uint16PPtr(map1, map2 map[*int8]*uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Uint8P Returns true if both maps are equal else returns false
func EqualMapInt8Uint8P(map1, map2 map[int8]uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Uint8PPtr Returns true if both maps are equal else returns false
func EqualMapInt8Uint8PPtr(map1, map2 map[*int8]*uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8StrP Returns true if both maps are equal else returns false
func EqualMapInt8StrP(map1, map2 map[int8]string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8StrPPtr Returns true if both maps are equal else returns false
func EqualMapInt8StrPPtr(map1, map2 map[*int8]*string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8BoolP Returns true if both maps are equal else returns false
func EqualMapInt8BoolP(map1, map2 map[int8]bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8BoolPPtr Returns true if both maps are equal else returns false
func EqualMapInt8BoolPPtr(map1, map2 map[*int8]*bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Float32P Returns true if both maps are equal else returns false
func EqualMapInt8Float32P(map1, map2 map[int8]float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Float32PPtr Returns true if both maps are equal else returns false
func EqualMapInt8Float32PPtr(map1, map2 map[*int8]*float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Float64P Returns true if both maps are equal else returns false
func EqualMapInt8Float64P(map1, map2 map[int8]float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapInt8Float64PPtr Returns true if both maps are equal else returns false
func EqualMapInt8Float64PPtr(map1, map2 map[*int8]*float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintIntP Returns true if both maps are equal else returns false
func EqualMapUintIntP(map1, map2 map[uint]int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintIntPPtr Returns true if both maps are equal else returns false
func EqualMapUintIntPPtr(map1, map2 map[*uint]*int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintInt64P Returns true if both maps are equal else returns false
func EqualMapUintInt64P(map1, map2 map[uint]int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintInt64PPtr Returns true if both maps are equal else returns false
func EqualMapUintInt64PPtr(map1, map2 map[*uint]*int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintInt32P Returns true if both maps are equal else returns false
func EqualMapUintInt32P(map1, map2 map[uint]int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintInt32PPtr Returns true if both maps are equal else returns false
func EqualMapUintInt32PPtr(map1, map2 map[*uint]*int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintInt16P Returns true if both maps are equal else returns false
func EqualMapUintInt16P(map1, map2 map[uint]int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintInt16PPtr Returns true if both maps are equal else returns false
func EqualMapUintInt16PPtr(map1, map2 map[*uint]*int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintInt8P Returns true if both maps are equal else returns false
func EqualMapUintInt8P(map1, map2 map[uint]int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintInt8PPtr Returns true if both maps are equal else returns false
func EqualMapUintInt8PPtr(map1, map2 map[*uint]*int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintP Returns true if both maps are equal else returns false
func EqualMapUintP(map1, map2 map[uint]uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintPPtr Returns true if both maps are equal else returns false
func EqualMapUintPPtr(map1, map2 map[*uint]*uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintUint64P Returns true if both maps are equal else returns false
func EqualMapUintUint64P(map1, map2 map[uint]uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintUint64PPtr Returns true if both maps are equal else returns false
func EqualMapUintUint64PPtr(map1, map2 map[*uint]*uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintUint32P Returns true if both maps are equal else returns false
func EqualMapUintUint32P(map1, map2 map[uint]uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintUint32PPtr Returns true if both maps are equal else returns false
func EqualMapUintUint32PPtr(map1, map2 map[*uint]*uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintUint16P Returns true if both maps are equal else returns false
func EqualMapUintUint16P(map1, map2 map[uint]uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintUint16PPtr Returns true if both maps are equal else returns false
func EqualMapUintUint16PPtr(map1, map2 map[*uint]*uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintUint8P Returns true if both maps are equal else returns false
func EqualMapUintUint8P(map1, map2 map[uint]uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintUint8PPtr Returns true if both maps are equal else returns false
func EqualMapUintUint8PPtr(map1, map2 map[*uint]*uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintStrP Returns true if both maps are equal else returns false
func EqualMapUintStrP(map1, map2 map[uint]string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintStrPPtr Returns true if both maps are equal else returns false
func EqualMapUintStrPPtr(map1, map2 map[*uint]*string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintBoolP Returns true if both maps are equal else returns false
func EqualMapUintBoolP(map1, map2 map[uint]bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintBoolPPtr Returns true if both maps are equal else returns false
func EqualMapUintBoolPPtr(map1, map2 map[*uint]*bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintFloat32P Returns true if both maps are equal else returns false
func EqualMapUintFloat32P(map1, map2 map[uint]float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintFloat32PPtr Returns true if both maps are equal else returns false
func EqualMapUintFloat32PPtr(map1, map2 map[*uint]*float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintFloat64P Returns true if both maps are equal else returns false
func EqualMapUintFloat64P(map1, map2 map[uint]float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUintFloat64PPtr Returns true if both maps are equal else returns false
func EqualMapUintFloat64PPtr(map1, map2 map[*uint]*float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64IntP Returns true if both maps are equal else returns false
func EqualMapUint64IntP(map1, map2 map[uint64]int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64IntPPtr Returns true if both maps are equal else returns false
func EqualMapUint64IntPPtr(map1, map2 map[*uint64]*int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Int64P Returns true if both maps are equal else returns false
func EqualMapUint64Int64P(map1, map2 map[uint64]int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Int64PPtr Returns true if both maps are equal else returns false
func EqualMapUint64Int64PPtr(map1, map2 map[*uint64]*int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Int32P Returns true if both maps are equal else returns false
func EqualMapUint64Int32P(map1, map2 map[uint64]int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Int32PPtr Returns true if both maps are equal else returns false
func EqualMapUint64Int32PPtr(map1, map2 map[*uint64]*int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Int16P Returns true if both maps are equal else returns false
func EqualMapUint64Int16P(map1, map2 map[uint64]int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Int16PPtr Returns true if both maps are equal else returns false
func EqualMapUint64Int16PPtr(map1, map2 map[*uint64]*int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Int8P Returns true if both maps are equal else returns false
func EqualMapUint64Int8P(map1, map2 map[uint64]int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Int8PPtr Returns true if both maps are equal else returns false
func EqualMapUint64Int8PPtr(map1, map2 map[*uint64]*int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64UintP Returns true if both maps are equal else returns false
func EqualMapUint64UintP(map1, map2 map[uint64]uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64UintPPtr Returns true if both maps are equal else returns false
func EqualMapUint64UintPPtr(map1, map2 map[*uint64]*uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64P Returns true if both maps are equal else returns false
func EqualMapUint64P(map1, map2 map[uint64]uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64PPtr Returns true if both maps are equal else returns false
func EqualMapUint64PPtr(map1, map2 map[*uint64]*uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Uint32P Returns true if both maps are equal else returns false
func EqualMapUint64Uint32P(map1, map2 map[uint64]uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Uint32PPtr Returns true if both maps are equal else returns false
func EqualMapUint64Uint32PPtr(map1, map2 map[*uint64]*uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Uint16P Returns true if both maps are equal else returns false
func EqualMapUint64Uint16P(map1, map2 map[uint64]uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Uint16PPtr Returns true if both maps are equal else returns false
func EqualMapUint64Uint16PPtr(map1, map2 map[*uint64]*uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Uint8P Returns true if both maps are equal else returns false
func EqualMapUint64Uint8P(map1, map2 map[uint64]uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Uint8PPtr Returns true if both maps are equal else returns false
func EqualMapUint64Uint8PPtr(map1, map2 map[*uint64]*uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64StrP Returns true if both maps are equal else returns false
func EqualMapUint64StrP(map1, map2 map[uint64]string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64StrPPtr Returns true if both maps are equal else returns false
func EqualMapUint64StrPPtr(map1, map2 map[*uint64]*string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64BoolP Returns true if both maps are equal else returns false
func EqualMapUint64BoolP(map1, map2 map[uint64]bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64BoolPPtr Returns true if both maps are equal else returns false
func EqualMapUint64BoolPPtr(map1, map2 map[*uint64]*bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Float32P Returns true if both maps are equal else returns false
func EqualMapUint64Float32P(map1, map2 map[uint64]float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Float32PPtr Returns true if both maps are equal else returns false
func EqualMapUint64Float32PPtr(map1, map2 map[*uint64]*float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Float64P Returns true if both maps are equal else returns false
func EqualMapUint64Float64P(map1, map2 map[uint64]float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint64Float64PPtr Returns true if both maps are equal else returns false
func EqualMapUint64Float64PPtr(map1, map2 map[*uint64]*float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32IntP Returns true if both maps are equal else returns false
func EqualMapUint32IntP(map1, map2 map[uint32]int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32IntPPtr Returns true if both maps are equal else returns false
func EqualMapUint32IntPPtr(map1, map2 map[*uint32]*int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Int64P Returns true if both maps are equal else returns false
func EqualMapUint32Int64P(map1, map2 map[uint32]int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Int64PPtr Returns true if both maps are equal else returns false
func EqualMapUint32Int64PPtr(map1, map2 map[*uint32]*int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Int32P Returns true if both maps are equal else returns false
func EqualMapUint32Int32P(map1, map2 map[uint32]int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Int32PPtr Returns true if both maps are equal else returns false
func EqualMapUint32Int32PPtr(map1, map2 map[*uint32]*int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Int16P Returns true if both maps are equal else returns false
func EqualMapUint32Int16P(map1, map2 map[uint32]int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Int16PPtr Returns true if both maps are equal else returns false
func EqualMapUint32Int16PPtr(map1, map2 map[*uint32]*int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Int8P Returns true if both maps are equal else returns false
func EqualMapUint32Int8P(map1, map2 map[uint32]int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Int8PPtr Returns true if both maps are equal else returns false
func EqualMapUint32Int8PPtr(map1, map2 map[*uint32]*int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32UintP Returns true if both maps are equal else returns false
func EqualMapUint32UintP(map1, map2 map[uint32]uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32UintPPtr Returns true if both maps are equal else returns false
func EqualMapUint32UintPPtr(map1, map2 map[*uint32]*uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Uint64P Returns true if both maps are equal else returns false
func EqualMapUint32Uint64P(map1, map2 map[uint32]uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Uint64PPtr Returns true if both maps are equal else returns false
func EqualMapUint32Uint64PPtr(map1, map2 map[*uint32]*uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32P Returns true if both maps are equal else returns false
func EqualMapUint32P(map1, map2 map[uint32]uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32PPtr Returns true if both maps are equal else returns false
func EqualMapUint32PPtr(map1, map2 map[*uint32]*uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Uint16P Returns true if both maps are equal else returns false
func EqualMapUint32Uint16P(map1, map2 map[uint32]uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Uint16PPtr Returns true if both maps are equal else returns false
func EqualMapUint32Uint16PPtr(map1, map2 map[*uint32]*uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Uint8P Returns true if both maps are equal else returns false
func EqualMapUint32Uint8P(map1, map2 map[uint32]uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Uint8PPtr Returns true if both maps are equal else returns false
func EqualMapUint32Uint8PPtr(map1, map2 map[*uint32]*uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32StrP Returns true if both maps are equal else returns false
func EqualMapUint32StrP(map1, map2 map[uint32]string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32StrPPtr Returns true if both maps are equal else returns false
func EqualMapUint32StrPPtr(map1, map2 map[*uint32]*string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32BoolP Returns true if both maps are equal else returns false
func EqualMapUint32BoolP(map1, map2 map[uint32]bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32BoolPPtr Returns true if both maps are equal else returns false
func EqualMapUint32BoolPPtr(map1, map2 map[*uint32]*bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Float32P Returns true if both maps are equal else returns false
func EqualMapUint32Float32P(map1, map2 map[uint32]float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Float32PPtr Returns true if both maps are equal else returns false
func EqualMapUint32Float32PPtr(map1, map2 map[*uint32]*float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Float64P Returns true if both maps are equal else returns false
func EqualMapUint32Float64P(map1, map2 map[uint32]float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint32Float64PPtr Returns true if both maps are equal else returns false
func EqualMapUint32Float64PPtr(map1, map2 map[*uint32]*float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16IntP Returns true if both maps are equal else returns false
func EqualMapUint16IntP(map1, map2 map[uint16]int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16IntPPtr Returns true if both maps are equal else returns false
func EqualMapUint16IntPPtr(map1, map2 map[*uint16]*int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Int64P Returns true if both maps are equal else returns false
func EqualMapUint16Int64P(map1, map2 map[uint16]int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Int64PPtr Returns true if both maps are equal else returns false
func EqualMapUint16Int64PPtr(map1, map2 map[*uint16]*int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Int32P Returns true if both maps are equal else returns false
func EqualMapUint16Int32P(map1, map2 map[uint16]int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Int32PPtr Returns true if both maps are equal else returns false
func EqualMapUint16Int32PPtr(map1, map2 map[*uint16]*int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Int16P Returns true if both maps are equal else returns false
func EqualMapUint16Int16P(map1, map2 map[uint16]int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Int16PPtr Returns true if both maps are equal else returns false
func EqualMapUint16Int16PPtr(map1, map2 map[*uint16]*int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Int8P Returns true if both maps are equal else returns false
func EqualMapUint16Int8P(map1, map2 map[uint16]int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Int8PPtr Returns true if both maps are equal else returns false
func EqualMapUint16Int8PPtr(map1, map2 map[*uint16]*int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16UintP Returns true if both maps are equal else returns false
func EqualMapUint16UintP(map1, map2 map[uint16]uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16UintPPtr Returns true if both maps are equal else returns false
func EqualMapUint16UintPPtr(map1, map2 map[*uint16]*uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Uint64P Returns true if both maps are equal else returns false
func EqualMapUint16Uint64P(map1, map2 map[uint16]uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Uint64PPtr Returns true if both maps are equal else returns false
func EqualMapUint16Uint64PPtr(map1, map2 map[*uint16]*uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Uint32P Returns true if both maps are equal else returns false
func EqualMapUint16Uint32P(map1, map2 map[uint16]uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Uint32PPtr Returns true if both maps are equal else returns false
func EqualMapUint16Uint32PPtr(map1, map2 map[*uint16]*uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16P Returns true if both maps are equal else returns false
func EqualMapUint16P(map1, map2 map[uint16]uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16PPtr Returns true if both maps are equal else returns false
func EqualMapUint16PPtr(map1, map2 map[*uint16]*uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Uint8P Returns true if both maps are equal else returns false
func EqualMapUint16Uint8P(map1, map2 map[uint16]uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Uint8PPtr Returns true if both maps are equal else returns false
func EqualMapUint16Uint8PPtr(map1, map2 map[*uint16]*uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16StrP Returns true if both maps are equal else returns false
func EqualMapUint16StrP(map1, map2 map[uint16]string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16StrPPtr Returns true if both maps are equal else returns false
func EqualMapUint16StrPPtr(map1, map2 map[*uint16]*string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16BoolP Returns true if both maps are equal else returns false
func EqualMapUint16BoolP(map1, map2 map[uint16]bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16BoolPPtr Returns true if both maps are equal else returns false
func EqualMapUint16BoolPPtr(map1, map2 map[*uint16]*bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Float32P Returns true if both maps are equal else returns false
func EqualMapUint16Float32P(map1, map2 map[uint16]float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Float32PPtr Returns true if both maps are equal else returns false
func EqualMapUint16Float32PPtr(map1, map2 map[*uint16]*float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Float64P Returns true if both maps are equal else returns false
func EqualMapUint16Float64P(map1, map2 map[uint16]float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint16Float64PPtr Returns true if both maps are equal else returns false
func EqualMapUint16Float64PPtr(map1, map2 map[*uint16]*float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8IntP Returns true if both maps are equal else returns false
func EqualMapUint8IntP(map1, map2 map[uint8]int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8IntPPtr Returns true if both maps are equal else returns false
func EqualMapUint8IntPPtr(map1, map2 map[*uint8]*int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Int64P Returns true if both maps are equal else returns false
func EqualMapUint8Int64P(map1, map2 map[uint8]int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Int64PPtr Returns true if both maps are equal else returns false
func EqualMapUint8Int64PPtr(map1, map2 map[*uint8]*int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Int32P Returns true if both maps are equal else returns false
func EqualMapUint8Int32P(map1, map2 map[uint8]int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Int32PPtr Returns true if both maps are equal else returns false
func EqualMapUint8Int32PPtr(map1, map2 map[*uint8]*int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Int16P Returns true if both maps are equal else returns false
func EqualMapUint8Int16P(map1, map2 map[uint8]int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Int16PPtr Returns true if both maps are equal else returns false
func EqualMapUint8Int16PPtr(map1, map2 map[*uint8]*int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Int8P Returns true if both maps are equal else returns false
func EqualMapUint8Int8P(map1, map2 map[uint8]int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Int8PPtr Returns true if both maps are equal else returns false
func EqualMapUint8Int8PPtr(map1, map2 map[*uint8]*int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8UintP Returns true if both maps are equal else returns false
func EqualMapUint8UintP(map1, map2 map[uint8]uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8UintPPtr Returns true if both maps are equal else returns false
func EqualMapUint8UintPPtr(map1, map2 map[*uint8]*uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Uint64P Returns true if both maps are equal else returns false
func EqualMapUint8Uint64P(map1, map2 map[uint8]uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Uint64PPtr Returns true if both maps are equal else returns false
func EqualMapUint8Uint64PPtr(map1, map2 map[*uint8]*uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Uint32P Returns true if both maps are equal else returns false
func EqualMapUint8Uint32P(map1, map2 map[uint8]uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Uint32PPtr Returns true if both maps are equal else returns false
func EqualMapUint8Uint32PPtr(map1, map2 map[*uint8]*uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Uint16P Returns true if both maps are equal else returns false
func EqualMapUint8Uint16P(map1, map2 map[uint8]uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Uint16PPtr Returns true if both maps are equal else returns false
func EqualMapUint8Uint16PPtr(map1, map2 map[*uint8]*uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8P Returns true if both maps are equal else returns false
func EqualMapUint8P(map1, map2 map[uint8]uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8PPtr Returns true if both maps are equal else returns false
func EqualMapUint8PPtr(map1, map2 map[*uint8]*uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8StrP Returns true if both maps are equal else returns false
func EqualMapUint8StrP(map1, map2 map[uint8]string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8StrPPtr Returns true if both maps are equal else returns false
func EqualMapUint8StrPPtr(map1, map2 map[*uint8]*string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8BoolP Returns true if both maps are equal else returns false
func EqualMapUint8BoolP(map1, map2 map[uint8]bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8BoolPPtr Returns true if both maps are equal else returns false
func EqualMapUint8BoolPPtr(map1, map2 map[*uint8]*bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Float32P Returns true if both maps are equal else returns false
func EqualMapUint8Float32P(map1, map2 map[uint8]float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Float32PPtr Returns true if both maps are equal else returns false
func EqualMapUint8Float32PPtr(map1, map2 map[*uint8]*float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Float64P Returns true if both maps are equal else returns false
func EqualMapUint8Float64P(map1, map2 map[uint8]float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapUint8Float64PPtr Returns true if both maps are equal else returns false
func EqualMapUint8Float64PPtr(map1, map2 map[*uint8]*float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrIntP Returns true if both maps are equal else returns false
func EqualMapStrIntP(map1, map2 map[string]int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrIntPPtr Returns true if both maps are equal else returns false
func EqualMapStrIntPPtr(map1, map2 map[*string]*int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrInt64P Returns true if both maps are equal else returns false
func EqualMapStrInt64P(map1, map2 map[string]int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrInt64PPtr Returns true if both maps are equal else returns false
func EqualMapStrInt64PPtr(map1, map2 map[*string]*int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrInt32P Returns true if both maps are equal else returns false
func EqualMapStrInt32P(map1, map2 map[string]int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrInt32PPtr Returns true if both maps are equal else returns false
func EqualMapStrInt32PPtr(map1, map2 map[*string]*int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrInt16P Returns true if both maps are equal else returns false
func EqualMapStrInt16P(map1, map2 map[string]int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrInt16PPtr Returns true if both maps are equal else returns false
func EqualMapStrInt16PPtr(map1, map2 map[*string]*int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrInt8P Returns true if both maps are equal else returns false
func EqualMapStrInt8P(map1, map2 map[string]int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrInt8PPtr Returns true if both maps are equal else returns false
func EqualMapStrInt8PPtr(map1, map2 map[*string]*int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrUintP Returns true if both maps are equal else returns false
func EqualMapStrUintP(map1, map2 map[string]uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrUintPPtr Returns true if both maps are equal else returns false
func EqualMapStrUintPPtr(map1, map2 map[*string]*uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrUint64P Returns true if both maps are equal else returns false
func EqualMapStrUint64P(map1, map2 map[string]uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrUint64PPtr Returns true if both maps are equal else returns false
func EqualMapStrUint64PPtr(map1, map2 map[*string]*uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrUint32P Returns true if both maps are equal else returns false
func EqualMapStrUint32P(map1, map2 map[string]uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrUint32PPtr Returns true if both maps are equal else returns false
func EqualMapStrUint32PPtr(map1, map2 map[*string]*uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrUint16P Returns true if both maps are equal else returns false
func EqualMapStrUint16P(map1, map2 map[string]uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrUint16PPtr Returns true if both maps are equal else returns false
func EqualMapStrUint16PPtr(map1, map2 map[*string]*uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrUint8P Returns true if both maps are equal else returns false
func EqualMapStrUint8P(map1, map2 map[string]uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrUint8PPtr Returns true if both maps are equal else returns false
func EqualMapStrUint8PPtr(map1, map2 map[*string]*uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrP Returns true if both maps are equal else returns false
func EqualMapStrP(map1, map2 map[string]string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrPPtr Returns true if both maps are equal else returns false
func EqualMapStrPPtr(map1, map2 map[*string]*string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrBoolP Returns true if both maps are equal else returns false
func EqualMapStrBoolP(map1, map2 map[string]bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrBoolPPtr Returns true if both maps are equal else returns false
func EqualMapStrBoolPPtr(map1, map2 map[*string]*bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrFloat32P Returns true if both maps are equal else returns false
func EqualMapStrFloat32P(map1, map2 map[string]float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrFloat32PPtr Returns true if both maps are equal else returns false
func EqualMapStrFloat32PPtr(map1, map2 map[*string]*float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrFloat64P Returns true if both maps are equal else returns false
func EqualMapStrFloat64P(map1, map2 map[string]float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapStrFloat64PPtr Returns true if both maps are equal else returns false
func EqualMapStrFloat64PPtr(map1, map2 map[*string]*float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolIntP Returns true if both maps are equal else returns false
func EqualMapBoolIntP(map1, map2 map[bool]int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolIntPPtr Returns true if both maps are equal else returns false
func EqualMapBoolIntPPtr(map1, map2 map[*bool]*int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolInt64P Returns true if both maps are equal else returns false
func EqualMapBoolInt64P(map1, map2 map[bool]int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolInt64PPtr Returns true if both maps are equal else returns false
func EqualMapBoolInt64PPtr(map1, map2 map[*bool]*int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolInt32P Returns true if both maps are equal else returns false
func EqualMapBoolInt32P(map1, map2 map[bool]int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolInt32PPtr Returns true if both maps are equal else returns false
func EqualMapBoolInt32PPtr(map1, map2 map[*bool]*int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolInt16P Returns true if both maps are equal else returns false
func EqualMapBoolInt16P(map1, map2 map[bool]int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolInt16PPtr Returns true if both maps are equal else returns false
func EqualMapBoolInt16PPtr(map1, map2 map[*bool]*int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolInt8P Returns true if both maps are equal else returns false
func EqualMapBoolInt8P(map1, map2 map[bool]int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolInt8PPtr Returns true if both maps are equal else returns false
func EqualMapBoolInt8PPtr(map1, map2 map[*bool]*int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolUintP Returns true if both maps are equal else returns false
func EqualMapBoolUintP(map1, map2 map[bool]uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolUintPPtr Returns true if both maps are equal else returns false
func EqualMapBoolUintPPtr(map1, map2 map[*bool]*uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolUint64P Returns true if both maps are equal else returns false
func EqualMapBoolUint64P(map1, map2 map[bool]uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolUint64PPtr Returns true if both maps are equal else returns false
func EqualMapBoolUint64PPtr(map1, map2 map[*bool]*uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolUint32P Returns true if both maps are equal else returns false
func EqualMapBoolUint32P(map1, map2 map[bool]uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolUint32PPtr Returns true if both maps are equal else returns false
func EqualMapBoolUint32PPtr(map1, map2 map[*bool]*uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolUint16P Returns true if both maps are equal else returns false
func EqualMapBoolUint16P(map1, map2 map[bool]uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolUint16PPtr Returns true if both maps are equal else returns false
func EqualMapBoolUint16PPtr(map1, map2 map[*bool]*uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolUint8P Returns true if both maps are equal else returns false
func EqualMapBoolUint8P(map1, map2 map[bool]uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolUint8PPtr Returns true if both maps are equal else returns false
func EqualMapBoolUint8PPtr(map1, map2 map[*bool]*uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolStrP Returns true if both maps are equal else returns false
func EqualMapBoolStrP(map1, map2 map[bool]string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolStrPPtr Returns true if both maps are equal else returns false
func EqualMapBoolStrPPtr(map1, map2 map[*bool]*string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolP Returns true if both maps are equal else returns false
func EqualMapBoolP(map1, map2 map[bool]bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolPPtr Returns true if both maps are equal else returns false
func EqualMapBoolPPtr(map1, map2 map[*bool]*bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolFloat32P Returns true if both maps are equal else returns false
func EqualMapBoolFloat32P(map1, map2 map[bool]float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolFloat32PPtr Returns true if both maps are equal else returns false
func EqualMapBoolFloat32PPtr(map1, map2 map[*bool]*float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolFloat64P Returns true if both maps are equal else returns false
func EqualMapBoolFloat64P(map1, map2 map[bool]float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapBoolFloat64PPtr Returns true if both maps are equal else returns false
func EqualMapBoolFloat64PPtr(map1, map2 map[*bool]*float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32IntP Returns true if both maps are equal else returns false
func EqualMapFloat32IntP(map1, map2 map[float32]int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32IntPPtr Returns true if both maps are equal else returns false
func EqualMapFloat32IntPPtr(map1, map2 map[*float32]*int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Int64P Returns true if both maps are equal else returns false
func EqualMapFloat32Int64P(map1, map2 map[float32]int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Int64PPtr Returns true if both maps are equal else returns false
func EqualMapFloat32Int64PPtr(map1, map2 map[*float32]*int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Int32P Returns true if both maps are equal else returns false
func EqualMapFloat32Int32P(map1, map2 map[float32]int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Int32PPtr Returns true if both maps are equal else returns false
func EqualMapFloat32Int32PPtr(map1, map2 map[*float32]*int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Int16P Returns true if both maps are equal else returns false
func EqualMapFloat32Int16P(map1, map2 map[float32]int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Int16PPtr Returns true if both maps are equal else returns false
func EqualMapFloat32Int16PPtr(map1, map2 map[*float32]*int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Int8P Returns true if both maps are equal else returns false
func EqualMapFloat32Int8P(map1, map2 map[float32]int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Int8PPtr Returns true if both maps are equal else returns false
func EqualMapFloat32Int8PPtr(map1, map2 map[*float32]*int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32UintP Returns true if both maps are equal else returns false
func EqualMapFloat32UintP(map1, map2 map[float32]uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32UintPPtr Returns true if both maps are equal else returns false
func EqualMapFloat32UintPPtr(map1, map2 map[*float32]*uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Uint64P Returns true if both maps are equal else returns false
func EqualMapFloat32Uint64P(map1, map2 map[float32]uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Uint64PPtr Returns true if both maps are equal else returns false
func EqualMapFloat32Uint64PPtr(map1, map2 map[*float32]*uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Uint32P Returns true if both maps are equal else returns false
func EqualMapFloat32Uint32P(map1, map2 map[float32]uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Uint32PPtr Returns true if both maps are equal else returns false
func EqualMapFloat32Uint32PPtr(map1, map2 map[*float32]*uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Uint16P Returns true if both maps are equal else returns false
func EqualMapFloat32Uint16P(map1, map2 map[float32]uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Uint16PPtr Returns true if both maps are equal else returns false
func EqualMapFloat32Uint16PPtr(map1, map2 map[*float32]*uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Uint8P Returns true if both maps are equal else returns false
func EqualMapFloat32Uint8P(map1, map2 map[float32]uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Uint8PPtr Returns true if both maps are equal else returns false
func EqualMapFloat32Uint8PPtr(map1, map2 map[*float32]*uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32StrP Returns true if both maps are equal else returns false
func EqualMapFloat32StrP(map1, map2 map[float32]string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32StrPPtr Returns true if both maps are equal else returns false
func EqualMapFloat32StrPPtr(map1, map2 map[*float32]*string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32BoolP Returns true if both maps are equal else returns false
func EqualMapFloat32BoolP(map1, map2 map[float32]bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32BoolPPtr Returns true if both maps are equal else returns false
func EqualMapFloat32BoolPPtr(map1, map2 map[*float32]*bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32P Returns true if both maps are equal else returns false
func EqualMapFloat32P(map1, map2 map[float32]float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32PPtr Returns true if both maps are equal else returns false
func EqualMapFloat32PPtr(map1, map2 map[*float32]*float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Float64P Returns true if both maps are equal else returns false
func EqualMapFloat32Float64P(map1, map2 map[float32]float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat32Float64PPtr Returns true if both maps are equal else returns false
func EqualMapFloat32Float64PPtr(map1, map2 map[*float32]*float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64IntP Returns true if both maps are equal else returns false
func EqualMapFloat64IntP(map1, map2 map[float64]int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64IntPPtr Returns true if both maps are equal else returns false
func EqualMapFloat64IntPPtr(map1, map2 map[*float64]*int) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Int64P Returns true if both maps are equal else returns false
func EqualMapFloat64Int64P(map1, map2 map[float64]int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Int64PPtr Returns true if both maps are equal else returns false
func EqualMapFloat64Int64PPtr(map1, map2 map[*float64]*int64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Int32P Returns true if both maps are equal else returns false
func EqualMapFloat64Int32P(map1, map2 map[float64]int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Int32PPtr Returns true if both maps are equal else returns false
func EqualMapFloat64Int32PPtr(map1, map2 map[*float64]*int32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Int16P Returns true if both maps are equal else returns false
func EqualMapFloat64Int16P(map1, map2 map[float64]int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Int16PPtr Returns true if both maps are equal else returns false
func EqualMapFloat64Int16PPtr(map1, map2 map[*float64]*int16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Int8P Returns true if both maps are equal else returns false
func EqualMapFloat64Int8P(map1, map2 map[float64]int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Int8PPtr Returns true if both maps are equal else returns false
func EqualMapFloat64Int8PPtr(map1, map2 map[*float64]*int8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64UintP Returns true if both maps are equal else returns false
func EqualMapFloat64UintP(map1, map2 map[float64]uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64UintPPtr Returns true if both maps are equal else returns false
func EqualMapFloat64UintPPtr(map1, map2 map[*float64]*uint) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Uint64P Returns true if both maps are equal else returns false
func EqualMapFloat64Uint64P(map1, map2 map[float64]uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Uint64PPtr Returns true if both maps are equal else returns false
func EqualMapFloat64Uint64PPtr(map1, map2 map[*float64]*uint64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Uint32P Returns true if both maps are equal else returns false
func EqualMapFloat64Uint32P(map1, map2 map[float64]uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Uint32PPtr Returns true if both maps are equal else returns false
func EqualMapFloat64Uint32PPtr(map1, map2 map[*float64]*uint32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Uint16P Returns true if both maps are equal else returns false
func EqualMapFloat64Uint16P(map1, map2 map[float64]uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Uint16PPtr Returns true if both maps are equal else returns false
func EqualMapFloat64Uint16PPtr(map1, map2 map[*float64]*uint16) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Uint8P Returns true if both maps are equal else returns false
func EqualMapFloat64Uint8P(map1, map2 map[float64]uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Uint8PPtr Returns true if both maps are equal else returns false
func EqualMapFloat64Uint8PPtr(map1, map2 map[*float64]*uint8) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64StrP Returns true if both maps are equal else returns false
func EqualMapFloat64StrP(map1, map2 map[float64]string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64StrPPtr Returns true if both maps are equal else returns false
func EqualMapFloat64StrPPtr(map1, map2 map[*float64]*string) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64BoolP Returns true if both maps are equal else returns false
func EqualMapFloat64BoolP(map1, map2 map[float64]bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64BoolPPtr Returns true if both maps are equal else returns false
func EqualMapFloat64BoolPPtr(map1, map2 map[*float64]*bool) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Float32P Returns true if both maps are equal else returns false
func EqualMapFloat64Float32P(map1, map2 map[float64]float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64Float32PPtr Returns true if both maps are equal else returns false
func EqualMapFloat64Float32PPtr(map1, map2 map[*float64]*float32) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64P Returns true if both maps are equal else returns false
func EqualMapFloat64P(map1, map2 map[float64]float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if k1 == k2 && v1 == v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}

// EqualMapFloat64PPtr Returns true if both maps are equal else returns false
func EqualMapFloat64PPtr(map1, map2 map[*float64]*float64) bool {
	len1 := len(map1)
	len2 := len(map2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for k1, v1 := range map1 {
		found := false
		for k2, v2 := range map2 {
			if *k1 == *k2 && *v1 == *v2 {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}
	return true
}
