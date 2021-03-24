package fp

// DistinctIntPtr removes duplicates.
func DistinctIntPtr(list []*int) []*int {
	var newList []*int
	s := make(map[int]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[*v]; ok {
			continue
		}
		s[*v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctInt64Ptr removes duplicates.
func DistinctInt64Ptr(list []*int64) []*int64 {
	var newList []*int64
	s := make(map[int64]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[*v]; ok {
			continue
		}
		s[*v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctInt32Ptr removes duplicates.
func DistinctInt32Ptr(list []*int32) []*int32 {
	var newList []*int32
	s := make(map[int32]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[*v]; ok {
			continue
		}
		s[*v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctInt16Ptr removes duplicates.
func DistinctInt16Ptr(list []*int16) []*int16 {
	var newList []*int16
	s := make(map[int16]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[*v]; ok {
			continue
		}
		s[*v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctInt8Ptr removes duplicates.
func DistinctInt8Ptr(list []*int8) []*int8 {
	var newList []*int8
	s := make(map[int8]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[*v]; ok {
			continue
		}
		s[*v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctUintPtr removes duplicates.
func DistinctUintPtr(list []*uint) []*uint {
	var newList []*uint
	s := make(map[uint]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[*v]; ok {
			continue
		}
		s[*v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctUint64Ptr removes duplicates.
func DistinctUint64Ptr(list []*uint64) []*uint64 {
	var newList []*uint64
	s := make(map[uint64]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[*v]; ok {
			continue
		}
		s[*v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctUint32Ptr removes duplicates.
func DistinctUint32Ptr(list []*uint32) []*uint32 {
	var newList []*uint32
	s := make(map[uint32]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[*v]; ok {
			continue
		}
		s[*v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctUint16Ptr removes duplicates.
func DistinctUint16Ptr(list []*uint16) []*uint16 {
	var newList []*uint16
	s := make(map[uint16]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[*v]; ok {
			continue
		}
		s[*v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctUint8Ptr removes duplicates.
func DistinctUint8Ptr(list []*uint8) []*uint8 {
	var newList []*uint8
	s := make(map[uint8]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[*v]; ok {
			continue
		}
		s[*v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctStrPtr removes duplicates.
func DistinctStrPtr(list []*string) []*string {
	var newList []*string
	s := make(map[string]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[*v]; ok {
			continue
		}
		s[*v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctBoolPtr removes duplicates.
func DistinctBoolPtr(list []*bool) []*bool {
	var newList []*bool
	s := make(map[bool]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[*v]; ok {
			continue
		}
		s[*v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctFloat32Ptr removes duplicates.
func DistinctFloat32Ptr(list []*float32) []*float32 {
	var newList []*float32
	s := make(map[float32]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[*v]; ok {
			continue
		}
		s[*v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctFloat64Ptr removes duplicates.
func DistinctFloat64Ptr(list []*float64) []*float64 {
	var newList []*float64
	s := make(map[float64]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[*v]; ok {
			continue
		}
		s[*v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}
