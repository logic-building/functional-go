package fp

// MinIntPtr returns min item from the list.
// Return 0 if the list is either empty or nil
func MinIntPtr(list []*int) *int {
	if list == nil || len(list) == 0 {
		var zero int
		return &zero
	}
	min := list[0]
	for _, v := range list {
		if *v < *min {
			min = v
		}
	}
	return min
}

// MinInt64Ptr returns min item from the list.
// Return 0 if the list is either empty or nil
func MinInt64Ptr(list []*int64) *int64 {
	if list == nil || len(list) == 0 {
		var zero int64
		return &zero
	}
	min := list[0]
	for _, v := range list {
		if *v < *min {
			min = v
		}
	}
	return min
}

// MinInt32Ptr returns min item from the list.
// Return 0 if the list is either empty or nil
func MinInt32Ptr(list []*int32) *int32 {
	if list == nil || len(list) == 0 {
		var zero int32
		return &zero
	}
	min := list[0]
	for _, v := range list {
		if *v < *min {
			min = v
		}
	}
	return min
}

// MinInt16Ptr returns min item from the list.
// Return 0 if the list is either empty or nil
func MinInt16Ptr(list []*int16) *int16 {
	if list == nil || len(list) == 0 {
		var zero int16
		return &zero
	}
	min := list[0]
	for _, v := range list {
		if *v < *min {
			min = v
		}
	}
	return min
}

// MinInt8Ptr returns min item from the list.
// Return 0 if the list is either empty or nil
func MinInt8Ptr(list []*int8) *int8 {
	if list == nil || len(list) == 0 {
		var zero int8
		return &zero
	}
	min := list[0]
	for _, v := range list {
		if *v < *min {
			min = v
		}
	}
	return min
}

// MinUintPtr returns min item from the list.
// Return 0 if the list is either empty or nil
func MinUintPtr(list []*uint) *uint {
	if list == nil || len(list) == 0 {
		var zero uint
		return &zero
	}
	min := list[0]
	for _, v := range list {
		if *v < *min {
			min = v
		}
	}
	return min
}

// MinUint64Ptr returns min item from the list.
// Return 0 if the list is either empty or nil
func MinUint64Ptr(list []*uint64) *uint64 {
	if list == nil || len(list) == 0 {
		var zero uint64
		return &zero
	}
	min := list[0]
	for _, v := range list {
		if *v < *min {
			min = v
		}
	}
	return min
}

// MinUint32Ptr returns min item from the list.
// Return 0 if the list is either empty or nil
func MinUint32Ptr(list []*uint32) *uint32 {
	if list == nil || len(list) == 0 {
		var zero uint32
		return &zero
	}
	min := list[0]
	for _, v := range list {
		if *v < *min {
			min = v
		}
	}
	return min
}

// MinUint16Ptr returns min item from the list.
// Return 0 if the list is either empty or nil
func MinUint16Ptr(list []*uint16) *uint16 {
	if list == nil || len(list) == 0 {
		var zero uint16
		return &zero
	}
	min := list[0]
	for _, v := range list {
		if *v < *min {
			min = v
		}
	}
	return min
}

// MinUint8Ptr returns min item from the list.
// Return 0 if the list is either empty or nil
func MinUint8Ptr(list []*uint8) *uint8 {
	if list == nil || len(list) == 0 {
		var zero uint8
		return &zero
	}
	min := list[0]
	for _, v := range list {
		if *v < *min {
			min = v
		}
	}
	return min
}

// MinFloat32Ptr returns min item from the list.
// Return 0 if the list is either empty or nil
func MinFloat32Ptr(list []*float32) *float32 {
	if list == nil || len(list) == 0 {
		var zero float32
		return &zero
	}
	min := list[0]
	for _, v := range list {
		if *v < *min {
			min = v
		}
	}
	return min
}

// MinFloat64Ptr returns min item from the list.
// Return 0 if the list is either empty or nil
func MinFloat64Ptr(list []*float64) *float64 {
	if list == nil || len(list) == 0 {
		var zero float64
		return &zero
	}
	min := list[0]
	for _, v := range list {
		if *v < *min {
			min = v
		}
	}
	return min
}
