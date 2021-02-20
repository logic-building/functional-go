package fp

// FilterMapInt filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List.
//  Empty list if all there parameters are nil or either of parameter is nil
//
// Example: Multiply all positive numbers in the list by 2
//	fp.FilterMapInt(isPositive, multiplyBy2, []int{-1, 0, 2, 4}) // Returns [4, 8]
//
//	func isPositive(num int) bool {
//		return num > 0
//	}
//
//	func multiplyBy2(num int) int {
//		return num * 2
//	}
func FilterMapInt(fFilter func(int) bool, fMap func(int) int, list []int) []int {
	if fFilter == nil || fMap == nil {
		return []int{}
	}
	var newList []int
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapInt64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List.
//  Empty list if all there parameters are nil or either of parameter is nil
//
// Example: Multiply all positive numbers in the list by 2
//	fp.FilterMapInt64(isPositive, multiplyBy2, []int64{-1, 0, 2, 4}) // Returns [4, 8]
//
//	func isPositive(num int64) bool {
//		return num > 0
//	}
//
//	func multiplyBy2(num int64) int64 {
//		return num * 2
//	}
func FilterMapInt64(fFilter func(int64) bool, fMap func(int64) int64, list []int64) []int64 {
	if fFilter == nil || fMap == nil {
		return []int64{}
	}

	var newList []int64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapInt32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List.
//  Empty list if all there parameters are nil or either of parameter is nil
//
// Example: Multiply all positive numbers in the list by 2
//	fp.FilterMapInt32(isPositive, multiplyBy2, []int32{-1, 0, 2, 4}) // Returns [4, 8]
//
//	func isPositive(num int32) bool {
//		return num > 0
//	}
//
//	func multiplyBy2(num int32) int32 {
//		return num * 2
//	}
func FilterMapInt32(fFilter func(int32) bool, fMap func(int32) int32, list []int32) []int32 {
	if fFilter == nil || fMap == nil {
		return []int32{}
	}
	var newList []int32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapInt16 filter the list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List.
//  Empty list if all there parameters are nil or either of parameter is nil
//
// Example: Multiply all positive numbers in the list by 2
//	fp.FilterMapInt16(isPositive, multiplyBy2, []int16{-1, 0, 2, 4}) // Returns [4, 8]
//
//	func isPositive(num int16) bool {
//		return num > 0
//	}
//
//	func multiplyBy2(num int16) int16 {
//		return num * 2
//	}
func FilterMapInt16(fFilter func(int16) bool, fMap func(int16) int16, list []int16) []int16 {
	if fFilter == nil || fMap == nil {
		return []int16{}
	}
	var newList []int16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapInt8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List.
//  Empty list if all there parameters are nil or either of parameter is nil
//
// Example: Multiply all positive numbers in the list by 2
//	fp.FilterMapInt8(isPositive, multiplyBy2, []int8{-1, 0, 2, 4}) // Returns [4, 8]
//
//	func isPositive(num int8) bool {
//		return num > 0
//	}
//
//	func multiplyBy2(num int8) int8 {
//		return num * 2
//	}
func FilterMapInt8(fFilter func(int8) bool, fMap func(int8) int8, list []int8) []int8 {
	if fFilter == nil || fMap == nil {
		return []int8{}
	}
	var newList []int8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapUint64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List.
//  Empty list if all there parameters are nil or either of parameter is nil
//
// Example: Multiply all positive numbers in the list by 2
//	fp.FilterMapUint64(isPositive, multiplyBy2, []uint64{-1, 0, 2, 4}) // Returns [4, 8]
//
//	func isPositive(num uint64) bool {
//		return num > 0
//	}
//
//	func multiplyBy2(num uint64) uint64 {
//		return num * 2
//	}
func FilterMapUint64(fFilter func(uint64) bool, fMap func(uint64) uint64, list []uint64) []uint64 {
	if fFilter == nil || fMap == nil {
		return []uint64{}
	}
	var newList []uint64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapUint32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List.
//  Empty list if all there parameters are nil or either of parameter is nil
//
// Example: Multiply all positive numbers in the list by 2
//	fp.FilterMapUint32(isPositive, multiplyBy2, []uint64{-1, 0, 2, 4}) // Returns [4, 8]
//
//	func isPositive(num uint32) bool {
//		return num > 0
//	}
//
//	func multiplyBy2(num uint32) uint32 {
//		return num * 2
//	}
func FilterMapUint32(fFilter func(uint32) bool, fMap func(uint32) uint32, list []uint32) []uint32 {
	if fFilter == nil || fMap == nil {
		return []uint32{}
	}
	var newList []uint32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapUint16 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List.
//  Empty list if all there parameters are nil or either of parameter is nil
//
// Example: Multiply all positive numbers in the list by 2
//	fp.FilterMapUint16(isPositive, multiplyBy2, []uint16{-1, 0, 2, 4}) // Returns [4, 8]
//
//	func isPositive(num uint16) bool {
//		return num > 0
//	}
//
//	func multiplyBy2(num uint16) uint16 {
//		return num * 2
//	}
func FilterMapUint16(fFilter func(uint16) bool, fMap func(uint16) uint16, list []uint16) []uint16 {
	if fFilter == nil || fMap == nil {
		return []uint16{}
	}
	var newList []uint16
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapUint8 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List.
//  Empty list if all there parameters are nil or either of parameter is nil
//
// Example: Multiply all positive numbers in the list by 2
//	fp.FilterMapUint8(isPositive, multiplyBy2, []uint8{-1, 0, 2, 4}) // Returns [4, 8]
//
//	func isPositive(num uint8) bool {
//		return num > 0
//	}
//
//	func multiplyBy2(num uint8) uint8 {
//		return num * 2
//	}
func FilterMapUint8(fFilter func(uint8) bool, fMap func(uint8) uint8, list []uint8) []uint8 {
	if fFilter == nil || fMap == nil {
		return []uint8{}
	}
	var newList []uint8
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapUint filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List.
//  Empty list if all there parameters are nil or either of parameter is nil
//
// Example: Multiply all positive numbers in the list by 2
//	fp.FilterMapUint(isPositive, multiplyBy2, []uint{-1, 0, 2, 4}) // Returns [4, 8]
//
//	func isPositive(num uint) bool {
//		return num > 0
//	}
//
//	func multiplyBy2(num uint) uint {
//		return num * 2
//	}
func FilterMapUint(fFilter func(uint) bool, fMap func(uint) uint, list []uint) []uint {
	if fFilter == nil || fMap == nil {
		return []uint{}
	}
	var newList []uint
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapFloat64 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List.
//  Empty list if all there parameters are nil or either of parameter is nil
//
// Example: Multiply all positive numbers in the list by 2
//	fp.FilterMapFloat64(isPositive, multiplyBy2, []float64{-1, 0, 2, 4}) // Returns [4, 8]
//
//	func isPositive(num float64) bool {
//		return num > 0
//	}
//
//	func multiplyBy2(num float64) float64 {
//		return num * 2
//	}
func FilterMapFloat64(fFilter func(float64) bool, fMap func(float64) float64, list []float64) []float64 {
	if fFilter == nil || fMap == nil {
		return []float64{}
	}
	var newList []float64
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapFloat32 filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List.
//  Empty list if all there parameters are nil or either of parameter is nil
//
// Example: Multiply all positive numbers in the list by 2
//	fp.FilterMapFloat32(isPositive, multiplyBy2, []float32{-1, 0, 2, 4}) // Returns [4, 8]
//
//	func isPositive(num float32) bool {
//		return num > 0
//	}
//
//	func multiplyBy2(num float32) float32 {
//		return num * 2
//	}
func FilterMapFloat32(fFilter func(float32) bool, fMap func(float32) float32, list []float32) []float32 {
	if fFilter == nil || fMap == nil {
		return []float32{}
	}
	var newList []float32
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// FilterMapStr filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input and returns true/false.
//	2. Function: takes int as argument and returns int
// 	3. List
//
// Returns:
//	New List.
//  Empty list if all there parameters are nil or either of parameter is nil
//
// Example: filter all the names in the list which length is <10 and change them to upper case
//	fp.FilterMapStr(isStringLengthLessThan10, strings.ToUpper, []string{"gopal", "govinda", "nandeshwar", "Nandeshwar Sah"}) // Returns [{"GOPAL", "GOVINDA"]
//
//	func isStringLengthLessThan10(str string) bool {
//		return len(str) < 10
//	}
func FilterMapStr(fFilter func(string) bool, fMap func(string) string, list []string) []string {
	if fFilter == nil || fMap == nil {
		return []string{}
	}
	var newList []string
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}
