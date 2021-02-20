package fp

// MinMaxIntPtr returns min and max items from the list.
// Return 0,0 if the list is either empty or nil
func MinMaxIntPtr(list []*int) (*int, *int) {
	var zero int

	if list == nil || len(list) == 0 {
		return &zero, &zero
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if *v < *min {
			min = v
		} else if *v > *max {
			max = v
		}
	}
	return min, max
}

// MinMaxInt64Ptr returns min and max items from the list.
// Return 0,0 if the list is either empty or nil
func MinMaxInt64Ptr(list []*int64) (*int64, *int64) {
	var zero int64

	if list == nil || len(list) == 0 {
		return &zero, &zero
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if *v < *min {
			min = v
		} else if *v > *max {
			max = v
		}
	}
	return min, max
}

// MinMaxInt32Ptr returns min and max items from the list.
// Return 0,0 if the list is either empty or nil
func MinMaxInt32Ptr(list []*int32) (*int32, *int32) {
	var zero int32

	if list == nil || len(list) == 0 {
		return &zero, &zero
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if *v < *min {
			min = v
		} else if *v > *max {
			max = v
		}
	}
	return min, max
}

// MinMaxInt16Ptr returns min and max items from the list.
// Return 0,0 if the list is either empty or nil
func MinMaxInt16Ptr(list []*int16) (*int16, *int16) {
	var zero int16

	if list == nil || len(list) == 0 {
		return &zero, &zero
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if *v < *min {
			min = v
		} else if *v > *max {
			max = v
		}
	}
	return min, max
}

// MinMaxInt8Ptr returns min and max items from the list.
// Return 0,0 if the list is either empty or nil
func MinMaxInt8Ptr(list []*int8) (*int8, *int8) {
	var zero int8

	if list == nil || len(list) == 0 {
		return &zero, &zero
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if *v < *min {
			min = v
		} else if *v > *max {
			max = v
		}
	}
	return min, max
}

// MinMaxUintPtr returns min and max items from the list.
// Return 0,0 if the list is either empty or nil
func MinMaxUintPtr(list []*uint) (*uint, *uint) {
	var zero uint

	if list == nil || len(list) == 0 {
		return &zero, &zero
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if *v < *min {
			min = v
		} else if *v > *max {
			max = v
		}
	}
	return min, max
}

// MinMaxUint64Ptr returns min and max items from the list.
// Return 0,0 if the list is either empty or nil
func MinMaxUint64Ptr(list []*uint64) (*uint64, *uint64) {
	var zero uint64

	if list == nil || len(list) == 0 {
		return &zero, &zero
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if *v < *min {
			min = v
		} else if *v > *max {
			max = v
		}
	}
	return min, max
}

// MinMaxUint32Ptr returns min and max items from the list.
// Return 0,0 if the list is either empty or nil
func MinMaxUint32Ptr(list []*uint32) (*uint32, *uint32) {
	var zero uint32

	if list == nil || len(list) == 0 {
		return &zero, &zero
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if *v < *min {
			min = v
		} else if *v > *max {
			max = v
		}
	}
	return min, max
}

// MinMaxUint16Ptr returns min and max items from the list.
// Return 0,0 if the list is either empty or nil
func MinMaxUint16Ptr(list []*uint16) (*uint16, *uint16) {
	var zero uint16

	if list == nil || len(list) == 0 {
		return &zero, &zero
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if *v < *min {
			min = v
		} else if *v > *max {
			max = v
		}
	}
	return min, max
}

// MinMaxUint8Ptr returns min and max items from the list.
// Return 0,0 if the list is either empty or nil
func MinMaxUint8Ptr(list []*uint8) (*uint8, *uint8) {
	var zero uint8

	if list == nil || len(list) == 0 {
		return &zero, &zero
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if *v < *min {
			min = v
		} else if *v > *max {
			max = v
		}
	}
	return min, max
}

// MinMaxFloat32Ptr returns min and max items from the list.
// Return 0,0 if the list is either empty or nil
func MinMaxFloat32Ptr(list []*float32) (*float32, *float32) {
	var zero float32

	if list == nil || len(list) == 0 {
		return &zero, &zero
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if *v < *min {
			min = v
		} else if *v > *max {
			max = v
		}
	}
	return min, max
}

// MinMaxFloat64Ptr returns min and max items from the list.
// Return 0,0 if the list is either empty or nil
func MinMaxFloat64Ptr(list []*float64) (*float64, *float64) {
	var zero float64

	if list == nil || len(list) == 0 {
		return &zero, &zero
	}
	min := list[0]
	max := list[0]

	for _, v := range list {
		if *v < *min {
			min = v
		} else if *v > *max {
			max = v
		}
	}
	return min, max
}
