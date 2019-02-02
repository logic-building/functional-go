package fp

// DropWhileInt drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
//
// Example: Drops even number. Returns the remaining items once odd number is found in the list.
//	fp.DropWhileInt(isEven, []int{4, 2, 3, 4, 5}) // Returns [3, 4, 5]
//
//	func isEven(num int) bool {
//		return num%2 == 0
//	}
func DropWhileInt(f func(int) bool, list []int) []int {
	if f == nil {
		return []int{}
	}
	var newList []int
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]int, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileInt64 drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
//
// Example: Drops even number. Returns the remaining items once odd number is found in the list.
//	fp.DropWhileInt64(isEven, []int64{4, 2, 3, 4, 5}) // Returns [3, 4, 5]
//
//	func isEven(num int64) bool {
//		return num%2 == 0
//	}
func DropWhileInt64(f func(int64) bool, list []int64) []int64 {
	if f == nil {
		return []int64{}
	}
	var newList []int64
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]int64, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileInt32 drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
//	New List.
//  Empty list if either one of arguments or both of them are nil
//
// Example: Drops even number. Returns the remaining items once odd number is found in the list.
//	fp.DropWhileInt32(isEven, []int32{4, 2, 3, 4, 5}) // Returns [3, 4, 5]
//
//	func isEven(num int32) bool {
//		return num%2 == 0
//	}
func DropWhileInt32(f func(int32) bool, list []int32) []int32 {
	if f == nil {
		return []int32{}
	}
	var newList []int32
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]int32, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileInt16 drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
//
// Example: Drops even number. Returns the remaining items once odd number is found in the list.
//	fp.DropWhileInt16(isEven, []int16{4, 2, 3, 4, 5}) // Returns [3, 4, 5]
//
//	func isEven(num int16) bool {
//		return num%2 == 0
//	}
func DropWhileInt16(f func(int16) bool, list []int16) []int16 {
	if f == nil {
		return []int16{}
	}
	var newList []int16
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]int16, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileInt8 drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
// Returns:
//
// 	New List.
//  Empty list if either one of arguments or both of them are nil
//
// Example: Drops even number. Returns the remaining items once odd number is found in the list.
//	fp.DropWhileInt8(isEven, []int8{4, 2, 3, 4, 5}) // Returns [3, 4, 5]
//
//	func isEven(num int8) bool {
//		return num%2 == 0
//	}
func DropWhileInt8(f func(int8) bool, list []int8) []int8 {
	if f == nil {
		return []int8{}
	}
	var newList []int8
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]int8, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileUint drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
//
// Example: Drops even number. Returns the remaining items once odd number is found in the list.
//	fp.DropWhileUint(isEven, []uint{4, 2, 3, 4, 5}) // Returns [3, 4, 5]
//
//	func isEven(num uint) bool {
//		return num%2 == 0
//	}
func DropWhileUint(f func(uint) bool, list []uint) []uint {
	if f == nil {
		return []uint{}
	}
	var newList []uint
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]uint, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileUint64 drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
//
// Example: Drops even number. Returns the remaining items once odd number is found in the list.
//	fp.DropWhileUint64(isEven, []uint64{4, 2, 3, 4, 5}) // Returns [3, 4, 5]
//
//	func isEven(num uint64) bool {
//		return num%2 == 0
//	}
func DropWhileUint64(f func(uint64) bool, list []uint64) []uint64 {
	if f == nil {
		return []uint64{}
	}
	var newList []uint64
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]uint64, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileUint32 drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
//
// Example: Drops even number. Returns the remaining items once odd number is found in the list.
//	fp.DropWhileUint32(isEven, []uint32{4, 2, 3, 4, 5}) // Returns [3, 4, 5]
//
//	func isEven(num uint32) bool {
//		return num%2 == 0
//	}
func DropWhileUint32(f func(uint32) bool, list []uint32) []uint32 {
	if f == nil {
		return []uint32{}
	}
	var newList []uint32
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]uint32, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileUint16 drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
//
// Example: Drops even number. Returns the remaining items once odd number is found in the list.
//	fp.DropWhileUint16(isEven, []uint16{4, 2, 3, 4, 5}) // Returns [3, 4, 5]
//
//	func isEven(num uint16) bool {
//		return num%2 == 0
//	}
func DropWhileUint16(f func(uint16) bool, list []uint16) []uint16 {
	if f == nil {
		return []uint16{}
	}
	var newList []uint16
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]uint16, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileUint8 drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
//
// Example: Drops even number. Returns the remaining items once odd number is found in the list.
//	fp.DropWhileUint8(isEven, []uint8{4, 2, 3, 4, 5}) // Returns [3, 4, 5]
//
//	func isEven(num uint8) bool {
//		return num%2 == 0
//	}
func DropWhileUint8(f func(uint8) bool, list []uint8) []uint8 {
	if f == nil {
		return []uint8{}
	}
	var newList []uint8
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]uint8, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileFloat64 drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
//
// Example: Drops positive numbers. Returns the remaining items once negative number is found in the list.
//	fp. DropWhileFloat64(isPositive, []float64{4, 2, 3, -1, 4, 5}) // Returns [-1, 4, 5]
//
//	func isPositiveFloat64(num float64) bool {
//		return num > 0
//	}
func DropWhileFloat64(f func(float64) bool, list []float64) []float64 {
	if f == nil {
		return []float64{}
	}
	var newList []float64
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]float64, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileFloat32 drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
//
// Example: Drops positive numbers. Returns the remaining items once negative number is found in the list.
//	fp. DropWhileFloat32(isPositive, []float32{4, 2, 3, -1, 4, 5}) // Returns [-1, 4, 5]
//
//	func isPositiveFloat64(num float32) bool {
//		return num > 0
//	}
func DropWhileFloat32(f func(float32) bool, list []float32) []float32 {
	if f == nil {
		return []float32{}
	}
	var newList []float32
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]float32, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

// DropWhileStr drops the items from the list as long as condition satisfies.
//
// Takes two inputs
//	1. Function: takes one input and returns boolean
//	2. list
//
// Returns:
// 	New List.
//  Empty list if either one of arguments or both of them are nil
//
// Example: Drops names as long as it's length is <10.
//	fp.DropWhileStr(
//	func(str string) bool {
//		return len(str) < 10
//	},
//	[]string{"Ram", "Shyam", "Nandeshwar", "ShyamSundar", "Hari Shankar"}) // Returns ["Nandeshwar", "ShyamSundar", "Hari Shankar"]
func DropWhileStr(f func(string) bool, list []string) []string {
	if f == nil {
		return []string{}
	}
	var newList []string
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]string, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}
