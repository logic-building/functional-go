package fp

// EqualIntsP Returns true if both list are equal else returns false
func EqualIntsP(list1, list2 []int) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// EqualIntsPPtr Returns true if both list are equal else returns false
func EqualIntsPPtr(list1, list2 []*int) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if *list1[i] != *list2[i] {
			return false
		}
	}
	return true
}

// EqualInts64P Returns true if both list are equal else returns false
func EqualInts64P(list1, list2 []int64) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// EqualInts64PPtr Returns true if both list are equal else returns false
func EqualInts64PPtr(list1, list2 []*int64) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if *list1[i] != *list2[i] {
			return false
		}
	}
	return true
}

// EqualInts32P Returns true if both list are equal else returns false
func EqualInts32P(list1, list2 []int32) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// EqualInts32PPtr Returns true if both list are equal else returns false
func EqualInts32PPtr(list1, list2 []*int32) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if *list1[i] != *list2[i] {
			return false
		}
	}
	return true
}

// EqualInts16P Returns true if both list are equal else returns false
func EqualInts16P(list1, list2 []int16) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// EqualInts16PPtr Returns true if both list are equal else returns false
func EqualInts16PPtr(list1, list2 []*int16) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if *list1[i] != *list2[i] {
			return false
		}
	}
	return true
}

// EqualInts8P Returns true if both list are equal else returns false
func EqualInts8P(list1, list2 []int8) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// EqualInts8PPtr Returns true if both list are equal else returns false
func EqualInts8PPtr(list1, list2 []*int8) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if *list1[i] != *list2[i] {
			return false
		}
	}
	return true
}

// EqualUintsP Returns true if both list are equal else returns false
func EqualUintsP(list1, list2 []uint) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// EqualUintsPPtr Returns true if both list are equal else returns false
func EqualUintsPPtr(list1, list2 []*uint) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if *list1[i] != *list2[i] {
			return false
		}
	}
	return true
}

// EqualUint64sP Returns true if both list are equal else returns false
func EqualUint64sP(list1, list2 []uint64) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// EqualUint64sPPtr Returns true if both list are equal else returns false
func EqualUint64sPPtr(list1, list2 []*uint64) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if *list1[i] != *list2[i] {
			return false
		}
	}
	return true
}

// EqualUints32P Returns true if both list are equal else returns false
func EqualUints32P(list1, list2 []uint32) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// EqualUints32PPtr Returns true if both list are equal else returns false
func EqualUints32PPtr(list1, list2 []*uint32) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if *list1[i] != *list2[i] {
			return false
		}
	}
	return true
}

// EqualUints16P Returns true if both list are equal else returns false
func EqualUints16P(list1, list2 []uint16) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// EqualUints16PPtr Returns true if both list are equal else returns false
func EqualUints16PPtr(list1, list2 []*uint16) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if *list1[i] != *list2[i] {
			return false
		}
	}
	return true
}

// EqualUints8P Returns true if both list are equal else returns false
func EqualUints8P(list1, list2 []uint8) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// EqualUints8PPtr Returns true if both list are equal else returns false
func EqualUints8PPtr(list1, list2 []*uint8) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if *list1[i] != *list2[i] {
			return false
		}
	}
	return true
}

// EqualStrsP Returns true if both list are equal else returns false
func EqualStrsP(list1, list2 []string) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// EqualStrsPPtr Returns true if both list are equal else returns false
func EqualStrsPPtr(list1, list2 []*string) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if *list1[i] != *list2[i] {
			return false
		}
	}
	return true
}

// EqualBoolsP Returns true if both list are equal else returns false
func EqualBoolsP(list1, list2 []bool) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// EqualBoolsPPtr Returns true if both list are equal else returns false
func EqualBoolsPPtr(list1, list2 []*bool) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if *list1[i] != *list2[i] {
			return false
		}
	}
	return true
}

// EqualFloat32sP Returns true if both list are equal else returns false
func EqualFloat32sP(list1, list2 []float32) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// EqualFloat32sPPtr Returns true if both list are equal else returns false
func EqualFloat32sPPtr(list1, list2 []*float32) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if *list1[i] != *list2[i] {
			return false
		}
	}
	return true
}

// EqualFloat64sP Returns true if both list are equal else returns false
func EqualFloat64sP(list1, list2 []float64) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if list1[i] != list2[i] {
			return false
		}
	}
	return true
}

// EqualFloat64sPPtr Returns true if both list are equal else returns false
func EqualFloat64sPPtr(list1, list2 []*float64) bool {
	len1 := len(list1)
	len2 := len(list2)

	if len1 == 0 || len2 == 0 || len1 != len2 {
		return false
	}

	for i := 0; i < len1; i++ {
		if *list1[i] != *list2[i] {
			return false
		}
	}
	return true
}
