package fp

// MaxIntPtr Returns the greatest of the nums
func MaxIntPtr(list []*int) *int {
	var initialVal int
	var max *int = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

// MaxInt64Ptr Returns the greatest of the nums
func MaxInt64Ptr(list []*int64) *int64 {
	var initialVal int64
	var max *int64 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

// MaxInt32Ptr Returns the greatest of the nums
func MaxInt32Ptr(list []*int32) *int32 {
	var initialVal int32
	var max *int32 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

// MaxInt16Ptr Returns the greatest of the nums
func MaxInt16Ptr(list []*int16) *int16 {
	var initialVal int16
	var max *int16 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

// MaxInt8Ptr Returns the greatest of the nums
func MaxInt8Ptr(list []*int8) *int8 {
	var initialVal int8
	var max *int8 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

// MaxUintPtr Returns the greatest of the nums
func MaxUintPtr(list []*uint) *uint {
	var initialVal uint
	var max *uint = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

// MaxUint64Ptr Returns the greatest of the nums
func MaxUint64Ptr(list []*uint64) *uint64 {
	var initialVal uint64
	var max *uint64 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

// MaxUint32Ptr Returns the greatest of the nums
func MaxUint32Ptr(list []*uint32) *uint32 {
	var initialVal uint32
	var max *uint32 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

// MaxUint16Ptr Returns the greatest of the nums
func MaxUint16Ptr(list []*uint16) *uint16 {
	var initialVal uint16
	var max *uint16 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

// MaxUint8Ptr Returns the greatest of the nums
func MaxUint8Ptr(list []*uint8) *uint8 {
	var initialVal uint8
	var max *uint8 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

// MaxFloat32Ptr Returns the greatest of the nums
func MaxFloat32Ptr(list []*float32) *float32 {
	var initialVal float32
	var max *float32 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

// MaxFloat64Ptr Returns the greatest of the nums
func MaxFloat64Ptr(list []*float64) *float64 {
	var initialVal float64
	var max *float64 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}
