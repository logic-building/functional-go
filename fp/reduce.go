package fp

// ReduceInt reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments
//	B. list
// 	C. initializer (optional)
//
// Returns:
//	single value.
//
// Example
//	list := []int{1, 2, 3, 4, 5}
//	ReduceInt(plusInt, list) // returns: 15
//	ReduceInt(plusInt, list, 3) // returns: 18
//
//	func plusInt(num1, num2 int) int {
//		return num1 + num2
//	}
func ReduceInt(f func(int, int) int, list []int, initializer ...int) int {
	var init int
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init
	}

	r := f(init, list[0])
	return ReduceInt(f, list[1:], r)
}

// ReduceInt64 reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments
//	B. list
// 	C. initializer (optional)
//
// Returns:
//	single value.
//
// Example
//	list := []int64{1, 2, 3, 4, 5}
//	ReduceInt64(plusInt, list) // returns: 15
//	ReduceInt64(plusInt, list, 3) // returns: 18
//
//	func plusInt(num1, num2 int64) int64 {
//		return num1 + num2
//	}
func ReduceInt64(f func(int64, int64) int64, list []int64, initializer ...int64) int64 {
	var init int64
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return ReduceInt64(f, list[1:], r)
}

// ReduceInt32 reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments
//	B. list
// 	C. initializer (optional)
//
// Returns:
//	single value.
//
// Example
//	list := []int32{1, 2, 3, 4, 5}
//	ReduceInt32(plusInt, list) // returns: 15
//	ReduceInt32(plusInt, list, 3) // returns: 18
//
//	func plusInt(num1, num2 int32) int32 {
//		return num1 + num2
//	}
func ReduceInt32(f func(int32, int32) int32, list []int32, initializer ...int32) int32 {
	var init int32
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return ReduceInt32(f, list[1:], r)
}

// ReduceInt16 reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments
//	B. list
// 	C. initializer (optional)
//
// Returns:
//	single value.
//
// Example
//	list := []int16{1, 2, 3, 4, 5}
//	ReduceInt16(plusInt, list) // returns: 15
//	ReduceInt16(plusInt, list, 3) // returns: 18
//
//	func plusInt(num1, num2 int16) int16 {
//		return num1 + num2
//	}
func ReduceInt16(f func(int16, int16) int16, list []int16, initializer ...int16) int16 {
	var init int16
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return ReduceInt16(f, list[1:], r)
}

// ReduceInt8 reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments
//	B. list
// 	C. initializer (optional)
//
// Returns:
//	single value.
//
// Example
//	list := []int8{1, 2, 3, 4, 5}
//	ReduceInt8(plusInt, list) // returns: 15
//	ReduceInt8(plusInt, list, 3) // returns: 18
//
//	func plusInt(num1, num2 int8) int8 {
//		return num1 + num2
//	}
func ReduceInt8(f func(int8, int8) int8, list []int8, initializer ...int8) int8 {
	var init int8
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return ReduceInt8(f, list[1:], r)
}

// ReduceUint reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments
//	B. list
// 	C. initializer (optional)
//
// Returns:
//	single value.
//
// Example
//	list := []int8{1, 2, 3, 4, 5}
//	ReduceUint(plusInt, list) // returns: 15
//	ReduceUint(plusInt, list, 3) // returns: 18
//
//	func plusInt(num1, num2 int8) int8 {
//		return num1 + num2
//	}
func ReduceUint(f func(uint, uint) uint, list []uint, initializer ...uint) uint {
	var init uint
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]

		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return ReduceUint(f, list[1:], r)
}

// ReduceUint64 reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments
//	B. list
// 	C. initializer (optional)
//
// Returns:
//	single value.
//
// Example
//	list := []uint64{1, 2, 3, 4, 5}
//	ReduceUint64(plusInt, list) // returns: 15
//	ReduceUint64(plusInt, list, 3) // returns: 18
//
//	func plusInt(num1, num2 uint64) uint64 {
//		return num1 + num2
//	}
func ReduceUint64(f func(uint64, uint64) uint64, list []uint64, initializer ...uint64) uint64 {
	var init uint64
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]

		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return ReduceUint64(f, list[1:], r)
}

// ReduceUint32 reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments
//	B. list
// 	C. initializer (optional)
//
// Returns:
//	single value.
//
// Example
//	list := []uint32{1, 2, 3, 4, 5}
//	ReduceUint32(plusInt, list) // returns: 15
//	ReduceUint32(plusInt, list, 3) // returns: 18
//
//	func plusInt(num1, num2 uint32) uint32 {
//		return num1 + num2
//	}
func ReduceUint32(f func(uint32, uint32) uint32, list []uint32, initializer ...uint32) uint32 {
	var init uint32
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]

		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return ReduceUint32(f, list[1:], r)
}

// ReduceUint16 reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments
//	B. list
// 	C. initializer (optional)
//
// Returns:
//	single value.
//
// Example
//	list := []uint16{1, 2, 3, 4, 5}
//	ReduceUint16(plusInt, list) // returns: 15
//	ReduceUint16(plusInt, list, 3) // returns: 18
//
//	func plusInt(num1, num2 uint16) uint16 {
//		return num1 + num2
//	}
func ReduceUint16(f func(uint16, uint16) uint16, list []uint16, initializer ...uint16) uint16 {
	var init uint16
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]

		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return ReduceUint16(f, list[1:], r)
}

// ReduceUint8 reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments
//	B. list
// 	C. initializer (optional)
//
// Returns:
//	single value.
//
// Example
//	list := []uint8{1, 2, 3, 4, 5}
//	ReduceUint8(plusInt, list) // returns: 15
//	ReduceUint8(plusInt, list, 3) // returns: 18
//
//	func plusInt(num1, num2 uint8) uint8 {
//		return num1 + num2
//	}
func ReduceUint8(f func(uint8, uint8) uint8, list []uint8, initializer ...uint8) uint8 {
	var init uint8
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]

		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return ReduceUint8(f, list[1:], r)
}

// ReduceFloat64 reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments
//	B. list
// 	C. initializer (optional)
//
// Returns:
//	single value.
//
// Example
//	list := []float64{1, 2, 3, 4, 5}
//	ReduceFloat64(plusInt, list) // returns: 15
//	ReduceFloat64(plusInt, list, 3) // returns: 18
//
//	func plusInt(num1, num2 float64) float64 {
//		return num1 + num2
//	}
func ReduceFloat64(f func(float64, float64) float64, list []float64, initializer ...float64) float64 {
	var init float64
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]

		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return ReduceFloat64(f, list[1:], r)
}

// ReduceFloat32 reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments
//	B. list
// 	C. initializer (optional)
//
// Returns:
//	single value.
//
// Example
//	list := []float32{1, 2, 3, 4, 5}
//	ReduceFloat32(plusInt, list) // returns: 15
//	ReduceFloat32(plusInt, list, 3) // returns: 18
//
//	func plusInt(num1, num2 float32) float32 {
//		return num1 + num2
//	}
func ReduceFloat32(f func(float32, float32) float32, list []float32, initializer ...float32) float32 {
	var init float32
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]

		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return ReduceFloat32(f, list[1:], r)
}

// ReduceStr reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments
//	B. list
// 	C. initializer (optional)
//
// Returns:
//	single value.
//
// Example
//	list := []string{"Nandeshwar", "Kumar", "Sah"}
//	ReduceStr(plusInt, list) // returns: "NandeshwarKumarSah"
//
//	func plusStr(num1, num2 string) string {
//		return num1 + num2
//	}
func ReduceStr(f func(string, string) string, list []string, initializer ...string) string {
	var init string
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]

		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}

	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return ReduceStr(f, list[1:], r)
}
