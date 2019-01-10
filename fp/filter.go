package fp

func FilterInt(f func(int) bool, list []int) []int {
	if f == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterInt64(f func(int64) bool, list []int64) []int64 {
	if f == nil {
		return []int64{}
	}
	var newList []int64
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterInt32(f func(int32) bool, list []int32) []int32 {
	if f == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterInt16(f func(int16) bool, list []int16) []int16 {
	if f == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterInt8(f func(int8) bool, list []int8) []int8 {
	if f == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterUint64(f func(uint64) bool, list []uint64) []uint64 {
	if f == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterUint32(f func(uint32) bool, list []uint32) []uint32 {
	if f == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterUint16(f func(uint16) bool, list []uint16) []uint16 {
	if f == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterUint8(f func(uint8) bool, list []uint8) []uint8 {
	if f == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterUint(f func(uint) bool, list []uint) []uint {
	if f == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterFloat64(f func(float64) bool, list []float64) []float64 {
	if f == nil {
		return []float64{}
	}
	var newList []float64
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterFloat32(f func(float32) bool, list []float32) []float32 {
	if f == nil {
		return []float32{}
	}
	var newList []float32
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterStr(f func(string) bool, list []string) []string {
	if f == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}
