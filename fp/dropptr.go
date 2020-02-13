package fp

// DropInt returns a new list after dropping the given item
func DropIntPtr(num *int, list []*int) []*int {
	var newList []*int
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts returns a new list after dropping the given items
func DropIntsPtr(nums []*int, list []*int) []*int {
	var newList []*int
	for _, v := range list {
		if !ExistsIntPtr(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInt64 returns a new list after dropping the given item
func DropInt64Ptr(num *int64, list []*int64) []*int64 {
	var newList []*int64
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts returns a new list after dropping the given items
func DropInts64Ptr(nums []*int64, list []*int64) []*int64 {
	var newList []*int64
	for _, v := range list {
		if !ExistsInt64Ptr(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInt32 returns a new list after dropping the given item
func DropInt32Ptr(num *int32, list []*int32) []*int32 {
	var newList []*int32
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts returns a new list after dropping the given items
func DropInts32Ptr(nums []*int32, list []*int32) []*int32 {
	var newList []*int32
	for _, v := range list {
		if !ExistsInt32Ptr(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInt16 returns a new list after dropping the given item
func DropInt16Ptr(num *int16, list []*int16) []*int16 {
	var newList []*int16
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts returns a new list after dropping the given items
func DropInts16Ptr(nums []*int16, list []*int16) []*int16 {
	var newList []*int16
	for _, v := range list {
		if !ExistsInt16Ptr(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInt8 returns a new list after dropping the given item
func DropInt8Ptr(num *int8, list []*int8) []*int8 {
	var newList []*int8
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts returns a new list after dropping the given items
func DropInts8Ptr(nums []*int8, list []*int8) []*int8 {
	var newList []*int8
	for _, v := range list {
		if !ExistsInt8Ptr(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropUint returns a new list after dropping the given item
func DropUintPtr(num *uint, list []*uint) []*uint {
	var newList []*uint
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts returns a new list after dropping the given items
func DropUintsPtr(nums []*uint, list []*uint) []*uint {
	var newList []*uint
	for _, v := range list {
		if !ExistsUintPtr(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropUint64 returns a new list after dropping the given item
func DropUint64Ptr(num *uint64, list []*uint64) []*uint64 {
	var newList []*uint64
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts returns a new list after dropping the given items
func DropUint64sPtr(nums []*uint64, list []*uint64) []*uint64 {
	var newList []*uint64
	for _, v := range list {
		if !ExistsUint64Ptr(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropUint32 returns a new list after dropping the given item
func DropUint32Ptr(num *uint32, list []*uint32) []*uint32 {
	var newList []*uint32
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts returns a new list after dropping the given items
func DropUints32Ptr(nums []*uint32, list []*uint32) []*uint32 {
	var newList []*uint32
	for _, v := range list {
		if !ExistsUint32Ptr(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropUint16 returns a new list after dropping the given item
func DropUint16Ptr(num *uint16, list []*uint16) []*uint16 {
	var newList []*uint16
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts returns a new list after dropping the given items
func DropUints16Ptr(nums []*uint16, list []*uint16) []*uint16 {
	var newList []*uint16
	for _, v := range list {
		if !ExistsUint16Ptr(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropUint8 returns a new list after dropping the given item
func DropUint8Ptr(num *uint8, list []*uint8) []*uint8 {
	var newList []*uint8
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts returns a new list after dropping the given items
func DropUints8Ptr(nums []*uint8, list []*uint8) []*uint8 {
	var newList []*uint8
	for _, v := range list {
		if !ExistsUint8Ptr(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropStr returns a new list after dropping the given item
func DropStrPtr(num *string, list []*string) []*string {
	var newList []*string
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts returns a new list after dropping the given items
func DropStrsPtr(nums []*string, list []*string) []*string {
	var newList []*string
	for _, v := range list {
		if !ExistsStrPtr(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropBool returns a new list after dropping the given item
func DropBoolPtr(num *bool, list []*bool) []*bool {
	var newList []*bool
	for _, v := range list {
		if v != num {
			newList = append(newList, v)
		}
	}
	return newList
}

// DropInts returns a new list after dropping the given items
func DropBoolsPtr(nums []*bool, list []*bool) []*bool {
	var newList []*bool
	for _, v := range list {
		if !ExistsBoolPtr(v, nums) {
			newList = append(newList, v)
		}
	}
	return newList
}
