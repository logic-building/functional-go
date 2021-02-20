package fp

import "strings"

// DistinctInt removes duplicates.
//
// Example
// 	list := []int{8, 2, 8, 0, 2, 0}
// 	fp.DistinctInt(list) // returns [8, 2, 0]
func DistinctInt(list []int) []int {
	var newList []int
	s := make(map[int]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctInt64 removes duplicates.
//
// Example
// 	list := []int64{8, 2, 8, 0, 2, 0}
// 	fp.DistinctInt64(list) // returns [8, 2, 0]
func DistinctInt64(list []int64) []int64 {
	var newList []int64
	s := make(map[int64]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctInt32 removes duplicates.
//
// Example
// 	list := []int32{8, 2, 8, 0, 2, 0}
// 	fp.DistinctInt32(list) // returns [8, 2, 0]
func DistinctInt32(list []int32) []int32 {
	var newList []int32
	s := make(map[int32]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctInt16 removes duplicates.
//
// Example
// 	list := []int16{8, 2, 8, 0, 2, 0}
// 	fp.DistinctInt16(list) // returns [8, 2, 0]
func DistinctInt16(list []int16) []int16 {
	var newList []int16
	s := make(map[int16]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctInt8 removes duplicates.
//
// Example
// 	list := []int8{8, 2, 8, 0, 2, 0}
// 	fp.DistinctInt8(list) // returns [8, 2, 0]
func DistinctInt8(list []int8) []int8 {
	var newList []int8
	s := make(map[int8]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctUint removes duplicates.
//
// Example
// 	list := []uint{8, 2, 8, 0, 2, 0}
// 	fp.DistinctUint(list) // returns [8, 2, 0]
func DistinctUint(list []uint) []uint {
	var newList []uint
	s := make(map[uint]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctUint64 removes duplicates.
//
// Example
// 	list := []uint64{8, 2, 8, 0, 2, 0}
// 	fp.DistinctUint64(list) // returns [8, 2, 0]
func DistinctUint64(list []uint64) []uint64 {
	var newList []uint64
	s := make(map[uint64]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctUint32 removes duplicates.
//
// Example
// 	list := []uint32{8, 2, 8, 0, 2, 0}
// 	fp.DistinctUint32(list) // returns [8, 2, 0]
func DistinctUint32(list []uint32) []uint32 {
	var newList []uint32
	s := make(map[uint32]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctUint16 removes duplicates.
//
// Example
// 	list := []uint16{8, 2, 8, 0, 2, 0}
// 	fp.DistinctUint16(list) // returns [8, 2, 0]
func DistinctUint16(list []uint16) []uint16 {
	var newList []uint16
	s := make(map[uint16]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctUint8 removes duplicates.
//
// Example
// 	list := []uint8{8, 2, 8, 0, 2, 0}
// 	fp.DistinctUint8(list) // returns [8, 2, 0]
func DistinctUint8(list []uint8) []uint8 {
	var newList []uint8
	s := make(map[uint8]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctFloat64 removes duplicates.
//
// Example
// 	list := []float64{8.1, 2.1, 8.1, 0, 2.1, 0}
// 	fp.DistinctFloat64(list) // returns [8.1, 2.1, 0]
func DistinctFloat64(list []float64) []float64 {
	var newList []float64
	s := make(map[float64]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctFloat32 removes duplicates.
//
// Example
// 	list := []float32{8.1, 2.1, 8.1, 0, 2.1, 0}
// 	fp.DistinctFloat32(list) // returns [8.1, 2.1, 0]
func DistinctFloat32(list []float32) []float32 {
	var newList []float32
	s := make(map[float32]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctStr removes duplicates.
//
// Example
// 	list := []string{"Bharat", "Hanuman", "Bharat", "Sita", "Hanuman", "Sita"}
// 	fp.DistinctStr(list) // returns ["Bharat", "Hanuman", "Sita"]
func DistinctStr(list []string) []string {
	var newList []string
	s := make(map[string]struct{}, len(list))
	for _, v := range list {
		if _, ok := s[v]; ok {
			continue
		}
		s[v] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}

// DistinctStrIgnoreCase ignores case and removes duplicates.
//
// Example
// 	list := []string{"Bharat", "HanumaN", "BharaT", "SiTa", "Hanuman", "Sita"}
// 	fp.DistinctStrIgnoreCase(list) // returns ["Bharat", "HanumaN", "SiTa"]
func DistinctStrIgnoreCase(list []string) []string {
	var newList []string
	s := make(map[string]struct{}, len(list))
	for _, v := range list {
		lowerV := strings.ToLower(v)
		if _, ok := s[lowerV]; ok {
			continue
		}
		s[lowerV] = struct{}{}
		newList = append(newList, v)
	}
	return newList
}
