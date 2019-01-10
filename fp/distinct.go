package fp

func DistinctInt(list []int) []int {
	s := make(map[int]struct{}, len(list))
	i := 0
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		list[i] = v
		i++
	}
	return list[:i]
}

func DistinctInt64(list []int64) []int64 {
	s := make(map[int64]struct{}, len(list))
	i := 0
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		list[i] = v
		i++
	}
	return list[:i]
}

func DistinctInt32(list []int32) []int32 {
	s := make(map[int32]struct{}, len(list))
	i := 0
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		list[i] = v
		i++
	}
	return list[:i]
}

func DistinctInt16(list []int16) []int16 {
	s := make(map[int16]struct{}, len(list))
	i := 0
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		list[i] = v
		i++
	}
	return list[:i]
}

func DistinctInt8(list []int8) []int8 {
	s := make(map[int8]struct{}, len(list))
	i := 0
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		list[i] = v
		i++
	}
	return list[:i]
}

func DistinctUint(list []uint) []uint {
	s := make(map[uint]struct{}, len(list))
	i := 0
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		list[i] = v
		i++
	}
	return list[:i]
}

func DistinctUint64(list []uint64) []uint64 {
	s := make(map[uint64]struct{}, len(list))
	i := 0
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		list[i] = v
		i++
	}
	return list[:i]
}

func DistinctUint32(list []uint32) []uint32 {
	s := make(map[uint32]struct{}, len(list))
	i := 0
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		list[i] = v
		i++
	}
	return list[:i]
}

func DistinctUint16(list []uint16) []uint16 {
	s := make(map[uint16]struct{}, len(list))
	i := 0
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		list[i] = v
		i++
	}
	return list[:i]
}

func DistinctUint8(list []uint8) []uint8 {
	s := make(map[uint8]struct{}, len(list))
	i := 0
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		list[i] = v
		i++
	}
	return list[:i]
}

func DistinctFloat64(list []float64) []float64 {
	s := make(map[float64]struct{}, len(list))
	i := 0
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		list[i] = v
		i++
	}
	return list[:i]
}

func DistinctFloat32(list []float32) []float32 {
	s := make(map[float32]struct{}, len(list))
	i := 0
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		list[i] = v
		i++
	}
	return list[:i]
}

func DistinctStr(list []string) []string {
	s := make(map[string]struct{}, len(list))
	i := 0
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		list[i] = v
		i++
	}
	return list[:i]
}
