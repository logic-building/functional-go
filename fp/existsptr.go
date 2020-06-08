package fp

// ExistsIntPtr checks if given item exists in the list
func ExistsIntPtr(num *int, list []*int) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsInt64Ptr checks if given item exists in the list
func ExistsInt64Ptr(num *int64, list []*int64) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsInt32Ptr checks if given item exists in the list
func ExistsInt32Ptr(num *int32, list []*int32) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsInt16Ptr checks if given item exists in the list
func ExistsInt16Ptr(num *int16, list []*int16) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsInt8Ptr checks if given item exists in the list
func ExistsInt8Ptr(num *int8, list []*int8) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsUintPtr checks if given item exists in the list
func ExistsUintPtr(num *uint, list []*uint) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsUint64Ptr checks if given item exists in the list
func ExistsUint64Ptr(num *uint64, list []*uint64) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsUint32Ptr checks if given item exists in the list
func ExistsUint32Ptr(num *uint32, list []*uint32) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsUint16Ptr checks if given item exists in the list
func ExistsUint16Ptr(num *uint16, list []*uint16) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsUint8Ptr checks if given item exists in the list
func ExistsUint8Ptr(num *uint8, list []*uint8) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsStrPtr checks if given item exists in the list
func ExistsStrPtr(num *string, list []*string) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsBoolPtr checks if given item exists in the list
func ExistsBoolPtr(num *bool, list []*bool) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsFloat32Ptr checks if given item exists in the list
func ExistsFloat32Ptr(num *float32, list []*float32) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsFloat64Ptr checks if given item exists in the list
func ExistsFloat64Ptr(num *float64, list []*float64) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}
