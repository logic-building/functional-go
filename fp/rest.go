/*
	Input: list
    output: New list includes all the items except 1st one.
            Returns new empty list if the input is either nil or empty or length 1 list
    Ex.
		list := []int{1, 2, 3, 4, 5}
		RestInt(list) // returns [2, 3, 4, 5]
*/

package fp

// RestInt removes 1st item of the list and return new list having rest of the items
func RestInt(l []int) []int {
	if l == nil {
		return []int{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []int{}
	}

	newList := make([]int, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestInt64 removes 1st item of the list and return new list having rest of the items
func RestInt64(l []int64) []int64 {
	if l == nil {
		return []int64{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []int64{}
	}

	newList := make([]int64, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestInt32 removes 1st item of the list and return new list having rest of the items
func RestInt32(l []int32) []int32 {
	if l == nil {
		return []int32{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []int32{}
	}

	newList := make([]int32, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestInt16 removes 1st item of the list and return new list having rest of the items
func RestInt16(l []int16) []int16 {
	if l == nil {
		return []int16{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []int16{}
	}

	newList := make([]int16, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestInt8 removes 1st item of the list and return new list having rest of the items
func RestInt8(l []int8) []int8 {
	if l == nil {
		return []int8{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []int8{}
	}

	newList := make([]int8, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestUint removes 1st item of the list and return new list having rest of the items
func RestUint(l []uint) []uint {
	if l == nil {
		return []uint{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []uint{}
	}

	newList := make([]uint, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestUint64 removes 1st item of the list and return new list having rest of the items
func RestUint64(l []uint64) []uint64 {
	if l == nil {
		return []uint64{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []uint64{}
	}

	newList := make([]uint64, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestUint32 removes 1st item of the list and return new list having rest of the items
func RestUint32(l []uint32) []uint32 {
	if l == nil {
		return []uint32{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []uint32{}
	}

	newList := make([]uint32, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestUint16 removes 1st item of the list and return new list having rest of the items
func RestUint16(l []uint16) []uint16 {
	if l == nil {
		return []uint16{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []uint16{}
	}

	newList := make([]uint16, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestUint8 removes 1st item of the list and return new list having rest of the items
func RestUint8(l []uint8) []uint8 {
	if l == nil {
		return []uint8{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []uint8{}
	}

	newList := make([]uint8, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestFloat64 removes 1st item of the list and return new list having rest of the items
func RestFloat64(l []float64) []float64 {
	if l == nil {
		return []float64{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []float64{}
	}

	newList := make([]float64, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestFloat32 removes 1st item of the list and return new list having rest of the items
func RestFloat32(l []float32) []float32 {
	if l == nil {
		return []float32{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []float32{}
	}

	newList := make([]float32, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestStr removes 1st item of the list and return new list having rest of the items
func RestStr(l []string) []string {
	if l == nil {
		return []string{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []string{}
	}

	newList := make([]string, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}
