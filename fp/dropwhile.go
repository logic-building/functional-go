package fp

func DropWhileInt(f func(int) bool, list []int) []int {
	if f == nil {
		return []int{}
	}
	var newList []int
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]int, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func DropWhileInt64(f func(int64) bool, list []int64) []int64 {
	if f == nil {
		return []int64{}
	}
	var newList []int64
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]int64, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func DropWhileInt32(f func(int32) bool, list []int32) []int32 {
	if f == nil {
		return []int32{}
	}
	var newList []int32
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]int32, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func DropWhileInt16(f func(int16) bool, list []int16) []int16 {
	if f == nil {
		return []int16{}
	}
	var newList []int16
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]int16, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func DropWhileInt8(f func(int8) bool, list []int8) []int8 {
	if f == nil {
		return []int8{}
	}
	var newList []int8
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]int8, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func DropWhileUint(f func(uint) bool, list []uint) []uint {
	if f == nil {
		return []uint{}
	}
	var newList []uint
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]uint, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func DropWhileUint64(f func(uint64) bool, list []uint64) []uint64 {
	if f == nil {
		return []uint64{}
	}
	var newList []uint64
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]uint64, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func DropWhileUint32(f func(uint32) bool, list []uint32) []uint32 {
	if f == nil {
		return []uint32{}
	}
	var newList []uint32
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]uint32, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func DropWhileUint16(f func(uint16) bool, list []uint16) []uint16 {
	if f == nil {
		return []uint16{}
	}
	var newList []uint16
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]uint16, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func DropWhileUint8(f func(uint8) bool, list []uint8) []uint8 {
	if f == nil {
		return []uint8{}
	}
	var newList []uint8
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]uint8, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func DropWhileFloat64(f func(float64) bool, list []float64) []float64 {
	if f == nil {
		return []float64{}
	}
	var newList []float64
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]float64, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func DropWhileFloat32(f func(float32) bool, list []float32) []float32 {
	if f == nil {
		return []float32{}
	}
	var newList []float32
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]float32, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func DropWhileStr(f func(string) bool, list []string) []string {
	if f == nil {
		return []string{}
	}
	var newList []string
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]string, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}
