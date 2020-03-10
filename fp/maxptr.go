package fp

   func MaxIntPtr(list []*int) *int {
	var initialVal int = 0
	var max *int = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

   func MaxInt64Ptr(list []*int64) *int64 {
	var initialVal int64 = 0
	var max *int64 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

   func MaxInt32Ptr(list []*int32) *int32 {
	var initialVal int32 = 0
	var max *int32 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

   func MaxInt16Ptr(list []*int16) *int16 {
	var initialVal int16 = 0
	var max *int16 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

   func MaxInt8Ptr(list []*int8) *int8 {
	var initialVal int8 = 0
	var max *int8 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

   func MaxUintPtr(list []*uint) *uint {
	var initialVal uint = 0
	var max *uint = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

   func MaxUint64Ptr(list []*uint64) *uint64 {
	var initialVal uint64 = 0
	var max *uint64 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

   func MaxUint32Ptr(list []*uint32) *uint32 {
	var initialVal uint32 = 0
	var max *uint32 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

   func MaxUint16Ptr(list []*uint16) *uint16 {
	var initialVal uint16 = 0
	var max *uint16 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

   func MaxUint8Ptr(list []*uint8) *uint8 {
	var initialVal uint8 = 0
	var max *uint8 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

   func MaxFloat32Ptr(list []*float32) *float32 {
	var initialVal float32 = 0
	var max *float32 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}

   func MaxFloat64Ptr(list []*float64) *float64 {
	var initialVal float64 = 0
	var max *float64 = &initialVal
	for _, v := range list {
		if *v > *max {
			max = v
		}
	}
	return max
}
