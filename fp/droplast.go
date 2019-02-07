package fp

func DropLastInt(list []int) []int {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []int{}
	}

	newList := make([]int, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

func DropLastInt64(list []int64) []int64 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []int64{}
	}

	newList := make([]int64, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

func DropLastInt32(list []int32) []int32 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []int32{}
	}

	newList := make([]int32, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

func DropLastInt16(list []int16) []int16 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []int16{}
	}

	newList := make([]int16, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

func DropLastInt8(list []int8) []int8 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []int8{}
	}

	newList := make([]int8, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

func DropLastUint(list []uint) []uint {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []uint{}
	}

	newList := make([]uint, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

func DropLastUint64(list []uint64) []uint64 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []uint64{}
	}

	newList := make([]uint64, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

func DropLastUint32(list []uint32) []uint32 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []uint32{}
	}

	newList := make([]uint32, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

func DropLastUint16(list []uint16) []uint16 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []uint16{}
	}

	newList := make([]uint16, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

func DropLastUint8(list []uint8) []uint8 {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []uint8{}
	}

	newList := make([]uint8, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

func DropLastStr(list []string) []string {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []string{}
	}

	newList := make([]string, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}
