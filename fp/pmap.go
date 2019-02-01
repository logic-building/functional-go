package fp

import (
	"sync"
)

// Apply the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
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
func PmapInt(f func(int) int, list []int) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i, v int) {
			defer wg.Done()
			ch <- map[int]int{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// Apply the function(1st argument) on each item of the list and returns new list
// Run in parallel. no_of_goroutines = no_of_items_in_list
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
func PmapInt64(f func(int64) int64, list []int64) []int64 {
	if f == nil {
		return []int64{}
	}

	ch := make(chan map[int64]int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int64]int64, i, v int64) {
			defer wg.Done()
			ch <- map[int64]int64{i: f(v)}
		}(&wg, ch, int64(i), v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// Apply the function(1st argument) on each item of the list and returns new list
// Run in parallel. no_of_goroutines = no_of_items_in_list
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
func PmapInt32(f func(int32) int32, list []int32) []int32 {
	if f == nil {
		return []int32{}
	}

	ch := make(chan map[int32]int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int32]int32, i, v int32) {
			defer wg.Done()
			ch <- map[int32]int32{i: f(v)}
		}(&wg, ch, int32(i), v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// Apply the function(1st argument) on each item of the list and returns new list
// Run in parallel. no_of_goroutines = no_of_items_in_list
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
func PmapInt16(f func(int16) int16, list []int16) []int16 {
	if f == nil {
		return []int16{}
	}

	ch := make(chan map[int16]int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int16]int16, i, v int16) {
			defer wg.Done()
			ch <- map[int16]int16{i: f(v)}
		}(&wg, ch, int16(i), v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// Apply the function(1st argument) on each item of the list and returns new list
// Run in parallel. no_of_goroutines = no_of_items_in_list
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
func PmapInt8(f func(int8) int8, list []int8) []int8 {
	if f == nil {
		return []int8{}
	}

	ch := make(chan map[int8]int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int8]int8, i, v int8) {
			defer wg.Done()
			ch <- map[int8]int8{i: f(v)}
		}(&wg, ch, int8(i), v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// Apply the function(1st argument) on each item of the list and returns new list
// Run in parallel. no_of_goroutines = no_of_items_in_list
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
func PmapUint(f func(uint) uint, list []uint) []uint {
	if f == nil {
		return []uint{}
	}

	ch := make(chan map[uint]uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[uint]uint, i, v uint) {
			defer wg.Done()
			ch <- map[uint]uint{i: f(v)}
		}(&wg, ch, uint(i), v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// Apply the function(1st argument) on each item of the list and returns new list
// Run in parallel. no_of_goroutines = no_of_items_in_list
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
func PmapUint64(f func(uint64) uint64, list []uint64) []uint64 {
	if f == nil {
		return []uint64{}
	}

	ch := make(chan map[uint64]uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[uint64]uint64, i, v uint64) {
			defer wg.Done()
			ch <- map[uint64]uint64{i: f(v)}
		}(&wg, ch, uint64(i), v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// Apply the function(1st argument) on each item of the list and returns new list
// Run in parallel. no_of_goroutines = no_of_items_in_list
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
func PmapUint32(f func(uint32) uint32, list []uint32) []uint32 {
	if f == nil {
		return []uint32{}
	}

	ch := make(chan map[uint32]uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[uint32]uint32, i, v uint32) {
			defer wg.Done()
			ch <- map[uint32]uint32{i: f(v)}
		}(&wg, ch, uint32(i), v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// Apply the function(1st argument) on each item of the list and returns new list
// Run in parallel. no_of_goroutines = no_of_items_in_list
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
func PmapUint16(f func(uint16) uint16, list []uint16) []uint16 {
	if f == nil {
		return []uint16{}
	}

	ch := make(chan map[uint16]uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[uint16]uint16, i, v uint16) {
			defer wg.Done()
			ch <- map[uint16]uint16{i: f(v)}
		}(&wg, ch, uint16(i), v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// Apply the function(1st argument) on each item of the list and returns new list
// Run in parallel. no_of_goroutines = no_of_items_in_list
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
func PmapUint8(f func(uint8) uint8, list []uint8) []uint8 {
	if f == nil {
		return []uint8{}
	}

	ch := make(chan map[uint8]uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[uint8]uint8, i, v uint8) {
			defer wg.Done()
			ch <- map[uint8]uint8{i: f(v)}
		}(&wg, ch, uint8(i), v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// Apply the function(1st argument) on each item of the list and returns new list
// Run in parallel. no_of_goroutines = no_of_items_in_list
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
func PmapFloat64(f func(float64) float64, list []float64) []float64 {
	if f == nil {
		return []float64{}
	}

	ch := make(chan map[int64]float64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int64]float64, i int64, v float64) {
			defer wg.Done()
			ch <- map[int64]float64{i: f(v)}
		}(&wg, ch, int64(i), v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// Apply the function(1st argument) on each item of the list and returns new list
// Run in parallel. no_of_goroutines = no_of_items_in_list
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
func PmapFloat32(f func(float32) float32, list []float32) []float32 {
	if f == nil {
		return []float32{}
	}

	ch := make(chan map[int32]float32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int32]float32, i int32, v float32) {
			defer wg.Done()
			ch <- map[int32]float32{i: f(v)}
		}(&wg, ch, int32(i), v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// Apply the function(1st argument) on each item of the list and returns new list
// Run in parallel. no_of_goroutines = no_of_items_in_list
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
func PmapStr(f func(string) string, list []string) []string {
	if f == nil {
		return []string{}
	}

	ch := make(chan map[int64]string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int64]string, i int64, v string) {
			defer wg.Done()
			ch <- map[int64]string{i: f(v)}
		}(&wg, ch, int64(i), v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
