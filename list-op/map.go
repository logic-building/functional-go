package list_op

func MapInt(f func(int) int, list []int) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt64(f func(int64) int64, list []int64) []int64 {
	if f == nil {
		return []int64{}
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt32(f func(int32) int32, list []int32) []int32 {
	if f == nil {
		return []int32{}
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt8(f func(int8) int8, list []int8) []int8 {
	if f == nil {
		return []int8{}
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapInt16(f func(int16) int16, list []int16) []int16 {
	if f == nil {
		return []int16{}
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint64(f func(uint64) uint64, list []uint64) []uint64 {
	if f == nil {
		return []uint64{}
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint32(f func(uint32) uint32, list []uint32) []uint32 {
	if f == nil {
		return []uint32{}
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint16(f func(uint16) uint16, list []uint16) []uint16 {
	if f == nil {
		return []uint16{}
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint8(f func(uint8) uint8, list []uint8) []uint8 {
	if f == nil {
		return []uint8{}
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapUint(f func(uint) uint, list []uint) []uint {
	if f == nil {
		return []uint{}
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapFloat64(f func(float64) float64, list []float64) []float64 {
	if f == nil {
		return []float64{}
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapFloat32(f func(float32) float32, list []float32) []float32 {
	if f == nil {
		return []float32{}
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapStr(f func(string) string, list []string) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
