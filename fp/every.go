package fp

// EveryBool returns true if supplied function returns logical true for every item in the list
//
// Example:
//	fp.EveryBool(fp.True, []bool{true, true, true, true}) // Returns true
//	fp.EveryBool(fp.True, []bool{true, false, true, true}) // Returns false
//	fp.EveryBool(fp.False, []bool{false, false, false, false}) // Returns true
//	fp.EveryBool(fp.False, []bool{}) // Returns false
//	fp.EveryBool(fp.True, []bool{}) // Returns false
//	fp.EveryBool(fp.True, nil) // Returns false
func EveryBool(f func() bool, list []bool) bool {
	if f == nil || len(list) == 0 {
		return false
	}

	b := f()
	for _, v := range list {
		if !(b == v) {
			return false
		}
	}
	return true
}

// EveryInt returns true if supplied function returns logical true for every item in the list
//
// Example:
//	fp.EveryInt(even, []int{8, 2, 10, 4}) // Returns true
//
//	func isEven(num int) bool {
//		return num%2 == 0
//	}
//
// fp.EveryInt(even, []int{}) // Returns false
// fp.EveryInt(nil, []int{}) // Returns false
// fp.EveryInt(nil, nil) // Returns false
func EveryInt(f func(int) bool, list []int) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryInt64 returns true if supplied function returns logical true for every item in the list
//
// Example:
//	fp.EveryInt64(even, []int64{8, 2, 10, 4}) // Returns true
//
//	func isEven(num int64) bool {
//		return num%2 == 0
//	}
//
// fp.EveryInt64(even, []int64{}) // Returns false
// fp.EveryInt64(nil, []int64{}) // Returns false
// fp.EveryInt64(nil, nil) // Returns false
func EveryInt64(f func(int64) bool, list []int64) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryInt32 returns true if supplied function returns logical true for every item in the list
//
// Example:
//	fp.EveryInt32(even, []int32{8, 2, 10, 4}) // Returns true
//
//	func isEven(num int32) bool {
//		return num%2 == 0
//	}
//
// fp.EveryInt32(even, []int32{}) // Returns false
// fp.EveryInt32(nil, []int32{}) // Returns false
// fp.EveryInt32(nil, nil) // Returns false
func EveryInt32(f func(int32) bool, list []int32) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryInt16 returns true if supplied function returns logical true for every item in the list
//
// Example:
//	fp.EveryInt16(even, []int16{8, 2, 10, 4}) // Returns true
//
//	func isEven(num int16) bool {
//		return num%2 == 0
//	}
//
// fp.EveryInt16(even, []int16{}) // Returns false
// fp.EveryInt16(nil, []int16{}) // Returns false
// fp.EveryInt16(nil, nil) // Returns false
func EveryInt16(f func(int16) bool, list []int16) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryInt8 returns true if supplied function returns logical true for every item in the list
//
// Example:
//	fp.EveryInt8(even, []int8{8, 2, 10, 4}) // Returns true
//
//	func isEven(num int8) bool {
//		return num%2 == 0
//	}
//
// fp.EveryInt8(even, []int8{}) // Returns false
// fp.EveryInt8(nil, []int8{}) // Returns false
// fp.EveryInt8(nil, nil) // Returns false
func EveryInt8(f func(int8) bool, list []int8) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryUint64 returns true if supplied function returns logical true for every item in the list
//
// Example:
//	fp.EveryUint64(even, []uint64{8, 2, 10, 4}) // Returns true
//
//	func isEven(num uint64) bool {
//		return num%2 == 0
//	}
//
// fp.EveryUint64(even, []uint64{}) // Returns false
// fp.EveryUint64(nil, []uint64{}) // Returns false
// fp.EveryUint64(nil, nil) // Returns false
func EveryUint64(f func(uint64) bool, list []uint64) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryUint32 returns true if supplied function returns logical true for every item in the list
//
// Example:
//	fp.EveryUint32(even, []uint32{8, 2, 10, 4}) // Returns true
//
//	func isEven(num uint32) bool {
//		return num%2 == 0
//	}
//
// fp.EveryUint32(even, []uint32{}) // Returns false
// fp.EveryUint32(nil, []uint32{}) // Returns false
// fp.EveryUint32(nil, nil) // Returns false
func EveryUint32(f func(uint32) bool, list []uint32) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryUint16 returns true if supplied function returns logical true for every item in the list
//
// Example:
//	fp.EveryUint16(even, []uint16{8, 2, 10, 4}) // Returns true
//
//	func isEven(num uint16) bool {
//		return num%2 == 0
//	}
//
// fp.EveryUint16(even, []uint16{}) // Returns false
// fp.EveryUint16(nil, []uint16{}) // Returns false
// fp.EveryUint16(nil, nil) // Returns false
func EveryUint16(f func(uint16) bool, list []uint16) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryUint8 returns true if supplied function returns logical true for every item in the list
//
// Example:
//	fp.EveryUint8(even, []uint8{8, 2, 10, 4}) // Returns true
//
//	func isEven(num uint8) bool {
//		return num%2 == 0
//	}
//
// fp.EveryUint8(even, []uint8{}) // Returns false
// fp.EveryUint8(nil, []uint8{}) // Returns false
// fp.EveryUint8(nil, nil) // Returns false
func EveryUint8(f func(uint8) bool, list []uint8) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryUint returns true if supplied function returns logical true for every item in the list
//
// Example:
//	fp.EveryUint(even, []uint{8, 2, 10, 4}) // Returns true
//
//	func isEven(num uint) bool {
//		return num%2 == 0
//	}
//
// fp.EveryUint(even, []uint{}) // Returns false
// fp.EveryUint(nil, []uint{}) // Returns false
// fp.EveryUint(nil, nil) // Returns false
func EveryUint(f func(uint) bool, list []uint) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryFloat64 returns true if supplied function returns logical true for every item in the list
//
// Example:
//	fp.EveryFloat64(isPositive, []float64{8.2, 2.3, 10.4, 4}) // Returns true
//
//	func isPositive(num float64) bool {
//		return num > 0
//	}
//
// fp.EveryFloat64(even, []float64{}) // Returns false
// fp.EveryFloat64(nil, []float64{}) // Returns false
// fp.EveryFloat64(nil, nil) // Returns false
func EveryFloat64(f func(float64) bool, list []float64) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryFloat32 returns true if supplied function returns logical true for every item in the list
//
// Example:
//	fp.EveryFloat32(isPositive, []float32{8.2, 2.3, 10.4, 4}) // Returns true
//
//	func isPositive(num float32) bool {
//		return num > 0
//	}
//
// fp.EveryFloat32(even, []float32{}) // Returns false
// fp.EveryFloat32(nil, []float32{}) // Returns false
// fp.EveryFloat32(nil, nil) // Returns false
func EveryFloat32(f func(float32) bool, list []float32) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

// EveryStr returns true if supplied function returns logical true for every item in the list
//
// Example:
//	fp.EveryStr(isStrLen3, []string{"Ram", "Raj", "Sai"}) // Returns true
//
//	func isStrLen3(str string) bool {
//		return len(str) == 3
//	}
//
// fp.EveryStr(even, []string{}) // Returns false
// fp.EveryStr(nil, []string{}) // Returns false
// fp.EveryStr(nil, nil) // Returns false
func EveryStr(f func(string) bool, list []string) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}
