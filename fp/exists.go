package fp

import "strings"

// ExistsInt checks if given item exists in the list
//
// Example:
//	ExistsInt(8, []int{8, 2, 10, 4}) // Returns true
//	ExistsInt(8, []int{}) // Returns false
//	ExistsInt(8, nil) // Returns false
func ExistsInt(num int, list []int) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsInt64 checks if given item exists in the list
//
// Example:
//	ExistsInt64(8, []int64{8, 2, 10, 4}) // Returns true
//	ExistsInt(8, []int64{}) // Returns false
//	ExistsInt(8, nil) // Returns false
func ExistsInt64(num int64, list []int64) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsInt32 checks if given item exists in the list
//
// Example:
//	ExistsInt32(8, []int32{8, 2, 10, 4}) // Returns true
//	ExistsInt32(8, []int32{}) // Returns false
//	ExistsInt32(8, nil) // Returns false
func ExistsInt32(num int32, list []int32) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsInt16 checks if given item exists in the list
//
// Example:
//	ExistsInt16(8, []int16{8, 2, 10, 4}) // Returns true
//	ExistsInt16(8, []int16{}) // Returns false
//	ExistsInt16(8, nil) // Returns false
func ExistsInt16(num int16, list []int16) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsInt8 checks if given item exists in the list
//
// Example:
//	ExistsInt8(8, []int8{8, 2, 10, 4}) // Returns true
//	ExistsInt8(8, []int8{}) // Returns false
//	ExistsInt8(8, nil) // Returns false
func ExistsInt8(num int8, list []int8) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsUint64 checks if given item exists in the list
//
// Example:
//	ExistsUint64(8, []uint64{8, 2, 10, 4}) // Returns true
//	ExistsUint64(8, []uint64{}) // Returns false
//	ExistsUint64(8, nil) // Returns false
func ExistsUint64(num uint64, list []uint64) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsUint32 checks if given item exists in the list
//
// Example:
//	ExistsUint32(8, []uint32{8, 2, 10, 4}) // Returns true
//	ExistsUint32(8, []uint32{}) // Returns false
//	ExistsUint32(8, nil) // Returns false
func ExistsUint32(num uint32, list []uint32) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsUint16 checks if given item exists in the list
//
// Example:
//	ExistsUint16(8, []uint16{8, 2, 10, 4}) // Returns true
//	ExistsUint16(8, []uint16{}) // Returns false
//	ExistsUint16(8, nil) // Returns false
func ExistsUint16(num uint16, list []uint16) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsUint8 checks if given item exists in the list
//
// Example:
//	ExistsUint8(8, []uint8{8, 2, 10, 4}) // Returns true
//	ExistsUint8(8, []uint8{}) // Returns false
//	ExistsUint8(8, nil) // Returns false
func ExistsUint8(num uint8, list []uint8) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsUint checks if given item exists in the list
//
// Example:
//	ExistsUint(8, []uint{8, 2, 10, 4}) // Returns true
//	ExistsUint(8, []uint{}) // Returns false
//	ExistsUint(8, nil) // Returns false
func ExistsUint(num uint, list []uint) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsFloat64 checks if given item exists in the list
//
// Example:
//	ExistsFloat64(8, []float64{8, 2, 10, 4}) // Returns true
//	ExistsFloat64(8, []float64{}) // Returns false
//	ExistsFloat64(8, nil) // Returns false
func ExistsFloat64(num float64, list []float64) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsFloat32 checks if given item exists in the list
//
// Example:
//	ExistsFloat32(8, []float32{8, 2, 10, 4}) // Returns true
//	ExistsFloat32(8, []float32{}) // Returns false
//	ExistsFloat32(8, nil) // Returns false
func ExistsFloat32(num float32, list []float32) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsStr checks if given item exists in the list
//
// Example:
//	ExistsStr("ram", []string{"shyam", "ram", "Hanuman"}) // Returns true
//	ExistsStr("ram", []string{}) // Returns false
//	ExistsStr("ram", nil) // Returns false
func ExistsStr(num string, list []string) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}

// ExistsStrIgnoreCase checks if given item exists in the list
//
// Example:
//	ExistsStrIgnoreCase("Ram", []string{"shyam", "ram", "Hanuman"}) // Returns true
//	ExistsStrIgnoreCase("ram", []string{}) // Returns false
//	ExistsStrIgnoreCase("ram", nil) // Returns false
func ExistsStrIgnoreCase(str string, list []string) bool {
	strLowerCase := strings.ToLower(str)
	for _, v := range list {
		if strings.ToLower(v) == strLowerCase {
			return true
		}
	}
	return false
}

// ExistsBool checks if given item exists in the list
func ExistsBool(num bool, list []bool) bool {
	for _, v := range list {
		if v == num {
			return true
		}
	}
	return false
}
