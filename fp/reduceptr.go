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

func ReduceIntPtr(f func(*int, *int) *int, list []*int, initializer ...int) *int {
	var init *int
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
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
	return ReduceIntPtr(f, list[1:], *r)
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

func ReduceInt64Ptr(f func(*int64, *int64) *int64, list []*int64, initializer ...int64) *int64 {
	var init *int64
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
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
	return ReduceInt64Ptr(f, list[1:], *r)
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

func ReduceInt32Ptr(f func(*int32, *int32) *int32, list []*int32, initializer ...int32) *int32 {
	var init *int32
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
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
	return ReduceInt32Ptr(f, list[1:], *r)
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

func ReduceInt16Ptr(f func(*int16, *int16) *int16, list []*int16, initializer ...int16) *int16 {
	var init *int16
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
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
	return ReduceInt16Ptr(f, list[1:], *r)
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

func ReduceInt8Ptr(f func(*int8, *int8) *int8, list []*int8, initializer ...int8) *int8 {
	var init *int8
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
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
	return ReduceInt8Ptr(f, list[1:], *r)
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

func ReduceUintPtr(f func(*uint, *uint) *uint, list []*uint, initializer ...uint) *uint {
	var init *uint
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
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
	return ReduceUintPtr(f, list[1:], *r)
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

func ReduceUint64Ptr(f func(*uint64, *uint64) *uint64, list []*uint64, initializer ...uint64) *uint64 {
	var init *uint64
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
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
	return ReduceUint64Ptr(f, list[1:], *r)
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

func ReduceUint32Ptr(f func(*uint32, *uint32) *uint32, list []*uint32, initializer ...uint32) *uint32 {
	var init *uint32
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
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
	return ReduceUint32Ptr(f, list[1:], *r)
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

func ReduceUint16Ptr(f func(*uint16, *uint16) *uint16, list []*uint16, initializer ...uint16) *uint16 {
	var init *uint16
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
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
	return ReduceUint16Ptr(f, list[1:], *r)
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

func ReduceUint8Ptr(f func(*uint8, *uint8) *uint8, list []*uint8, initializer ...uint8) *uint8 {
	var init *uint8
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
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
	return ReduceUint8Ptr(f, list[1:], *r)
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

func ReduceStrPtr(f func(*string, *string) *string, list []*string, initializer ...string) *string {
	var init *string
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
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
	return ReduceStrPtr(f, list[1:], *r)
}

// ReduceBool reduces a list to a single value by combining elements via a supplied function
//
// Takes three inputs
//	A. function - takes two arguments
//	B. list
// 	C. initializer (optional)
//
// Returns:
//	single value.

func ReduceBoolPtr(f func(*bool, *bool) *bool, list []*bool, initializer ...bool) *bool {
	var init *bool
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
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
	return ReduceBoolPtr(f, list[1:], *r)
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

func ReduceFloat32Ptr(f func(*float32, *float32) *float32, list []*float32, initializer ...float32) *float32 {
	var init *float32
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
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
	return ReduceFloat32Ptr(f, list[1:], *r)
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

func ReduceFloat64Ptr(f func(*float64, *float64) *float64, list []*float64, initializer ...float64) *float64 {
	var init *float64
	lenList := len(list)

	if initializer != nil {
		init = &initializer[0]
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
	return ReduceFloat64Ptr(f, list[1:], *r)
}
