package fp

// RestIntPtr removes 1st item of the list and return new list having rest of the items
func RestIntPtr(l []*int) []*int {
	if l == nil {
		return []*int{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []*int{}
	}

	newList := make([]*int, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestInt64Ptr removes 1st item of the list and return new list having rest of the items
func RestInt64Ptr(l []*int64) []*int64 {
	if l == nil {
		return []*int64{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []*int64{}
	}

	newList := make([]*int64, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestInt32Ptr removes 1st item of the list and return new list having rest of the items
func RestInt32Ptr(l []*int32) []*int32 {
	if l == nil {
		return []*int32{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []*int32{}
	}

	newList := make([]*int32, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestInt16Ptr removes 1st item of the list and return new list having rest of the items
func RestInt16Ptr(l []*int16) []*int16 {
	if l == nil {
		return []*int16{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []*int16{}
	}

	newList := make([]*int16, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestInt8Ptr removes 1st item of the list and return new list having rest of the items
func RestInt8Ptr(l []*int8) []*int8 {
	if l == nil {
		return []*int8{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []*int8{}
	}

	newList := make([]*int8, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestUintPtr removes 1st item of the list and return new list having rest of the items
func RestUintPtr(l []*uint) []*uint {
	if l == nil {
		return []*uint{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []*uint{}
	}

	newList := make([]*uint, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestUint64Ptr removes 1st item of the list and return new list having rest of the items
func RestUint64Ptr(l []*uint64) []*uint64 {
	if l == nil {
		return []*uint64{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []*uint64{}
	}

	newList := make([]*uint64, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestUint32Ptr removes 1st item of the list and return new list having rest of the items
func RestUint32Ptr(l []*uint32) []*uint32 {
	if l == nil {
		return []*uint32{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []*uint32{}
	}

	newList := make([]*uint32, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestUint16Ptr removes 1st item of the list and return new list having rest of the items
func RestUint16Ptr(l []*uint16) []*uint16 {
	if l == nil {
		return []*uint16{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []*uint16{}
	}

	newList := make([]*uint16, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestUint8Ptr removes 1st item of the list and return new list having rest of the items
func RestUint8Ptr(l []*uint8) []*uint8 {
	if l == nil {
		return []*uint8{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []*uint8{}
	}

	newList := make([]*uint8, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestStrPtr removes 1st item of the list and return new list having rest of the items
func RestStrPtr(l []*string) []*string {
	if l == nil {
		return []*string{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []*string{}
	}

	newList := make([]*string, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestBoolPtr removes 1st item of the list and return new list having rest of the items
func RestBoolPtr(l []*bool) []*bool {
	if l == nil {
		return []*bool{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []*bool{}
	}

	newList := make([]*bool, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestFloat32Ptr removes 1st item of the list and return new list having rest of the items
func RestFloat32Ptr(l []*float32) []*float32 {
	if l == nil {
		return []*float32{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []*float32{}
	}

	newList := make([]*float32, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

// RestFloat64Ptr removes 1st item of the list and return new list having rest of the items
func RestFloat64Ptr(l []*float64) []*float64 {
	if l == nil {
		return []*float64{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []*float64{}
	}

	newList := make([]*float64, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}
