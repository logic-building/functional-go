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
