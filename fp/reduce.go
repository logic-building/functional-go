package fp

// Takes three inputs
//    A. function - takes two arguments
//    B. list
//    C. initializer (optional)
// Returns:
//    single value
//
// Example
//   list := []int{1, 2, 3, 4, 5}
//   ReduceInt(plusInt, list) // returns: 15
//
//   func plusInt(num1, num2 int) int {
//	    return num1 + num2
//   }
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

func ReduceInt32(f func(int32, int32) int32, list []int32, initializer ...int32) int32 {
	var init int32 = 0
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

func ReduceInt16(f func(int16, int16) int16, list []int16, initializer ...int16) int16 {
	var init int16 = 0
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

func ReduceInt8(f func(int8, int8) int8, list []int8, initializer ...int8) int8 {
	var init int8 = 0
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

func ReduceUint(f func(uint, uint) uint, list []uint, initializer ...uint) uint {
	var init uint = 0
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

func ReduceUint64(f func(uint64, uint64) uint64, list []uint64, initializer ...uint64) uint64 {
	var init uint64 = 0
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

func ReduceUint32(f func(uint32, uint32) uint32, list []uint32, initializer ...uint32) uint32 {
	var init uint32 = 0
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

func ReduceUint16(f func(uint16, uint16) uint16, list []uint16, initializer ...uint16) uint16 {
	var init uint16 = 0
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

func ReduceUint8(f func(uint8, uint8) uint8, list []uint8, initializer ...uint8) uint8 {
	var init uint8 = 0
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

func ReduceFloat64(f func(float64, float64) float64, list []float64, initializer ...float64) float64 {
	var init float64 = 0
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

func ReduceFloat32(f func(float32, float32) float32, list []float32, initializer ...float32) float32 {
	var init float32 = 0
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

func ReduceStr(f func(string, string) string, list []string, initializer ...string) string {
	var init string = ""
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
