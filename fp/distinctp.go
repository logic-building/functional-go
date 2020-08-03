package fp

// DistinctIntP returns true if no two of the arguments are =
func DistinctIntP(list []int) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[int]bool)
	for _, v := range list {
		if _, ok := s[v]; ok {
			return false
		}
		s[v] = true
	}
	return true
}

// DistinctIntPPtr returns true if no two of the arguments are =
func DistinctIntPPtr(list []*int) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[int]bool)
	for _, v := range list {
		if _, ok := s[*v]; ok {
			return false
		}
		s[*v] = true
	}
	return true
}

// DistinctInt64P returns true if no two of the arguments are =
func DistinctInt64P(list []int64) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[int64]bool)
	for _, v := range list {
		if _, ok := s[v]; ok {
			return false
		}
		s[v] = true
	}
	return true
}

// DistinctInt64PPtr returns true if no two of the arguments are =
func DistinctInt64PPtr(list []*int64) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[int64]bool)
	for _, v := range list {
		if _, ok := s[*v]; ok {
			return false
		}
		s[*v] = true
	}
	return true
}

// DistinctInt32P returns true if no two of the arguments are =
func DistinctInt32P(list []int32) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[int32]bool)
	for _, v := range list {
		if _, ok := s[v]; ok {
			return false
		}
		s[v] = true
	}
	return true
}

// DistinctInt32PPtr returns true if no two of the arguments are =
func DistinctInt32PPtr(list []*int32) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[int32]bool)
	for _, v := range list {
		if _, ok := s[*v]; ok {
			return false
		}
		s[*v] = true
	}
	return true
}

// DistinctInt16P returns true if no two of the arguments are =
func DistinctInt16P(list []int16) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[int16]bool)
	for _, v := range list {
		if _, ok := s[v]; ok {
			return false
		}
		s[v] = true
	}
	return true
}

// DistinctInt16PPtr returns true if no two of the arguments are =
func DistinctInt16PPtr(list []*int16) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[int16]bool)
	for _, v := range list {
		if _, ok := s[*v]; ok {
			return false
		}
		s[*v] = true
	}
	return true
}

// DistinctInt8P returns true if no two of the arguments are =
func DistinctInt8P(list []int8) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[int8]bool)
	for _, v := range list {
		if _, ok := s[v]; ok {
			return false
		}
		s[v] = true
	}
	return true
}

// DistinctInt8PPtr returns true if no two of the arguments are =
func DistinctInt8PPtr(list []*int8) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[int8]bool)
	for _, v := range list {
		if _, ok := s[*v]; ok {
			return false
		}
		s[*v] = true
	}
	return true
}

// DistinctUintP returns true if no two of the arguments are =
func DistinctUintP(list []uint) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[uint]bool)
	for _, v := range list {
		if _, ok := s[v]; ok {
			return false
		}
		s[v] = true
	}
	return true
}

// DistinctUintPPtr returns true if no two of the arguments are =
func DistinctUintPPtr(list []*uint) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[uint]bool)
	for _, v := range list {
		if _, ok := s[*v]; ok {
			return false
		}
		s[*v] = true
	}
	return true
}

// DistinctUint64P returns true if no two of the arguments are =
func DistinctUint64P(list []uint64) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[uint64]bool)
	for _, v := range list {
		if _, ok := s[v]; ok {
			return false
		}
		s[v] = true
	}
	return true
}

// DistinctUint64PPtr returns true if no two of the arguments are =
func DistinctUint64PPtr(list []*uint64) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[uint64]bool)
	for _, v := range list {
		if _, ok := s[*v]; ok {
			return false
		}
		s[*v] = true
	}
	return true
}

// DistinctUint32P returns true if no two of the arguments are =
func DistinctUint32P(list []uint32) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[uint32]bool)
	for _, v := range list {
		if _, ok := s[v]; ok {
			return false
		}
		s[v] = true
	}
	return true
}

// DistinctUint32PPtr returns true if no two of the arguments are =
func DistinctUint32PPtr(list []*uint32) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[uint32]bool)
	for _, v := range list {
		if _, ok := s[*v]; ok {
			return false
		}
		s[*v] = true
	}
	return true
}

// DistinctUint16P returns true if no two of the arguments are =
func DistinctUint16P(list []uint16) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[uint16]bool)
	for _, v := range list {
		if _, ok := s[v]; ok {
			return false
		}
		s[v] = true
	}
	return true
}

// DistinctUint16PPtr returns true if no two of the arguments are =
func DistinctUint16PPtr(list []*uint16) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[uint16]bool)
	for _, v := range list {
		if _, ok := s[*v]; ok {
			return false
		}
		s[*v] = true
	}
	return true
}

// DistinctUint8P returns true if no two of the arguments are =
func DistinctUint8P(list []uint8) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[uint8]bool)
	for _, v := range list {
		if _, ok := s[v]; ok {
			return false
		}
		s[v] = true
	}
	return true
}

// DistinctUint8PPtr returns true if no two of the arguments are =
func DistinctUint8PPtr(list []*uint8) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[uint8]bool)
	for _, v := range list {
		if _, ok := s[*v]; ok {
			return false
		}
		s[*v] = true
	}
	return true
}

// DistinctStrP returns true if no two of the arguments are =
func DistinctStrP(list []string) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[string]bool)
	for _, v := range list {
		if _, ok := s[v]; ok {
			return false
		}
		s[v] = true
	}
	return true
}

// DistinctStrPPtr returns true if no two of the arguments are =
func DistinctStrPPtr(list []*string) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[string]bool)
	for _, v := range list {
		if _, ok := s[*v]; ok {
			return false
		}
		s[*v] = true
	}
	return true
}

// DistinctBoolP returns true if no two of the arguments are =
func DistinctBoolP(list []bool) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[bool]bool)
	for _, v := range list {
		if _, ok := s[v]; ok {
			return false
		}
		s[v] = true
	}
	return true
}

// DistinctBoolPPtr returns true if no two of the arguments are =
func DistinctBoolPPtr(list []*bool) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[bool]bool)
	for _, v := range list {
		if _, ok := s[*v]; ok {
			return false
		}
		s[*v] = true
	}
	return true
}

// DistinctFloat32P returns true if no two of the arguments are =
func DistinctFloat32P(list []float32) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[float32]bool)
	for _, v := range list {
		if _, ok := s[v]; ok {
			return false
		}
		s[v] = true
	}
	return true
}

// DistinctFloat32PPtr returns true if no two of the arguments are =
func DistinctFloat32PPtr(list []*float32) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[float32]bool)
	for _, v := range list {
		if _, ok := s[*v]; ok {
			return false
		}
		s[*v] = true
	}
	return true
}

// DistinctFloat64P returns true if no two of the arguments are =
func DistinctFloat64P(list []float64) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[float64]bool)
	for _, v := range list {
		if _, ok := s[v]; ok {
			return false
		}
		s[v] = true
	}
	return true
}

// DistinctFloat64PPtr returns true if no two of the arguments are =
func DistinctFloat64PPtr(list []*float64) bool {
	if len(list) == 0 {
		return false
	}

	s := make(map[float64]bool)
	for _, v := range list {
		if _, ok := s[*v]; ok {
			return false
		}
		s[*v] = true
	}
	return true
}
