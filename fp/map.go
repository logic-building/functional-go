package fp

// MapInt applies the function(1st argument) on each item of the list and returns new list
//
// Takes 2 inputs
//	1. Function - takes 1 input
//	2. List
//
// Returns
//	New List.
//	Empty list if all arguments are nil or either one is nil
//
// Example: Square each item in the list
//	MapInt(squareInt, []int{1, 2, 3}) // Returns [1, 4, 9]
//
//	func squareInt(num int) int {
//		return num * num
//	}
func MapInt(f func(int) int, list []int) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt64 applies the function(1st argument) on each item of the list and returns new list
//
// Takes 2 inputs
//	1. Function - takes 1 input
//	2. List
//
// Returns
//	New List.
//	Empty list if all arguments are nil or either one is nil
//
// Example: Square each item in the list
//	MapInt64(squareInt, []int64{1, 2, 3}) // Returns [1, 4, 9]
//
//	func squareInt(num int64) int64 {
//		return num * num
//	}
func MapInt64(f func(int64) int64, list []int64) []int64 {
	if f == nil {
		return []int64{}
	}
	newList := make([]int64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt32 applies the function(1st argument) on each item of the list and returns new list
//
// Takes 2 inputs
//	1. Function - takes 1 input
//	2. List
//
// Returns
//	New List.
//	Empty list if all arguments are nil or either one is nil
//
// Example: Square each item in the list
//	MapInt32(squareInt, []int32{1, 2, 3}) // Returns [1, 4, 9]
//
//	func squareInt(num int32) int32 {
//		return num * num
//	}
func MapInt32(f func(int32) int32, list []int32) []int32 {
	if f == nil {
		return []int32{}
	}
	newList := make([]int32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt8 applies the function(1st argument) on each item of the list and returns new list
//
// Takes 2 inputs
//	1. Function - takes 1 input
//	2. List
//
// Returns
//	New List.
//	Empty list if all arguments are nil or either one is nil
//
// Example: Square each item in the list
//	MapInt8(squareInt, []int8{1, 2, 3}) // Returns [1, 4, 9]
//
//	func squareInt(num int8) int8 {
//		return num * num
//	}
func MapInt8(f func(int8) int8, list []int8) []int8 {
	if f == nil {
		return []int8{}
	}
	newList := make([]int8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapInt16 applies the function(1st argument) on each item of the list and returns new list
//
// Takes 2 inputs
//	1. Function - takes 1 input
//	2. List
//
// Returns
//	New List.
//	Empty list if all arguments are nil or either one is nil
//
// Example: Square each item in the list
//	MapInt16(squareInt, []int16{1, 2, 3}) // Returns [1, 4, 9]
//
//	func squareInt(num int16) int16 {
//		return num * num
//	}
func MapInt16(f func(int16) int16, list []int16) []int16 {
	if f == nil {
		return []int16{}
	}
	newList := make([]int16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint64 applies the function(1st argument) on each item of the list and returns new list
//
// Takes 2 inputs
//	1. Function - takes 1 input
//	2. List
//
// Returns
//	New List.
//	Empty list if all arguments are nil or either one is nil
//
// Example: Square each item in the list
//	MapUint64(squareInt, []uint64{1, 2, 3}) // Returns [1, 4, 9]
//
//	func squareInt(num uint64) uint64 {
//		return num * num
//	}
func MapUint64(f func(uint64) uint64, list []uint64) []uint64 {
	if f == nil {
		return []uint64{}
	}
	newList := make([]uint64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint32 applies the function(1st argument) on each item of the list and returns new list
//
// Takes 2 inputs
//	1. Function - takes 1 input
//	2. List
//
// Returns
//	New List.
//	Empty list if all arguments are nil or either one is nil
//
// Example: Square each item in the list
//	MapUint32(squareInt, []uint32{1, 2, 3}) // Returns [1, 4, 9]
//
//	func squareInt(num uint32) uint32 {
//		return num * num
//	}
func MapUint32(f func(uint32) uint32, list []uint32) []uint32 {
	if f == nil {
		return []uint32{}
	}
	newList := make([]uint32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint16 applies the function(1st argument) on each item of the list and returns new list
//
// Takes 2 inputs
//	1. Function - takes 1 input
//	2. List
//
// Returns
//	New List.
//	Empty list if all arguments are nil or either one is nil
//
// Example: Square each item in the list
//	MapUint16(squareInt, []uint16{1, 2, 3}) // Returns [1, 4, 9]
//
//	func squareInt(num uint16) uint16 {
//		return num * num
//	}
func MapUint16(f func(uint16) uint16, list []uint16) []uint16 {
	if f == nil {
		return []uint16{}
	}
	newList := make([]uint16, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint8 applies the function(1st argument) on each item of the list and returns new list
//
// Takes 2 inputs
//	1. Function - takes 1 input
//	2. List
//
// Returns
//	New List.
//	Empty list if all arguments are nil or either one is nil
//
// Example: Square each item in the list
//	MapUint8(squareInt, []uint8{1, 2, 3}) // Returns [1, 4, 9]
//
//	func squareInt(num uint8) uint8 {
//		return num * num
//	}
func MapUint8(f func(uint8) uint8, list []uint8) []uint8 {
	if f == nil {
		return []uint8{}
	}
	newList := make([]uint8, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapUint applies the function(1st argument) on each item of the list and returns new list
//
// Takes 2 inputs
//	1. Function - takes 1 input
//	2. List
//
// Returns
//	New List.
//	Empty list if all arguments are nil or either one is nil
//
// Example: Square each item in the list
//	MapUint(squareInt, []uint{1, 2, 3}) // Returns [1, 4, 9]
//
//	func squareInt(num uint) uint {
//		return num * num
//	}
func MapUint(f func(uint) uint, list []uint) []uint {
	if f == nil {
		return []uint{}
	}
	newList := make([]uint, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat64 applies the function(1st argument) on each item of the list and returns new list
//
// Takes 2 inputs
//	1. Function - takes 1 input
//	2. List
//
// Returns
//	New List.
//	Empty list if all arguments are nil or either one is nil
//
// Example: add 1 to each item in the list
//	MapFloat64(addOne, []float64{1, 2, 3}) // Returns [2, 3, 4]
//
//	func addOne(num float64) float64 {
//		return num + 1
//	}
func MapFloat64(f func(float64) float64, list []float64) []float64 {
	if f == nil {
		return []float64{}
	}
	newList := make([]float64, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapFloat32 applies the function(1st argument) on each item of the list and returns new list
//
// Takes 2 inputs
//	1. Function - takes 1 input
//	2. List
//
// Returns
//	New List.
//	Empty list if all arguments are nil or either one is nil
//
// Example: add 1 to each item in the list
//	MapFloat32(addOne, []float32{1, 2, 3}) // Returns [2, 3, 4]
//
//	func addOne(num float32) float32 {
//		return num + 1
//	}
func MapFloat32(f func(float32) float32, list []float32) []float32 {
	if f == nil {
		return []float32{}
	}
	newList := make([]float32, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapStr applies the function(1st argument) on each item of the list and returns new list
//
// Takes 2 inputs
//	1. Function - takes 1 input
//	2. List
//
// Returns
//	New List.
//	Empty list if all arguments are nil or either one is nil
//
// Example: change case to upper for each item in the list
//	MapStr(strings.ToUpper, []string{"govinda", "gopal", "shyam"}) // Returns ["GOVINDA", "GOPAL", "SHYAM"]
func MapStr(f func(string) string, list []string) []string {
	if f == nil {
		return []string{}
	}
	newList := make([]string, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapBool applies the function(1st argument) on each item of the list and returns new list
//
// Takes 2 inputs
//	1. Function - takes 1 input
//	2. List
//
// Returns
//	New List.
//	Empty list if all arguments are nil or either one is nil
func MapBool(f func(bool) bool, list []bool) []bool {
	if f == nil {
		return []bool{}
	}
	newList := make([]bool, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}
