package fp 

import "sync"

// PMapIntInt64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: int64
//	2. List
//
// Returns
//	New List of type int64
//	Empty list if all arguments are nil or either one is nil
func PMapIntInt64(f func(int) int64, list []int) []int64 {
	if f == nil {
		return []int64{}
	}

	ch := make(chan map[int]int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v int) {
			defer wg.Done()
			ch <- map[int]int64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapIntInt32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: int32
//	2. List
//
// Returns
//	New List of type int32
//	Empty list if all arguments are nil or either one is nil
func PMapIntInt32(f func(int) int32, list []int) []int32 {
	if f == nil {
		return []int32{}
	}

	ch := make(chan map[int]int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v int) {
			defer wg.Done()
			ch <- map[int]int32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapIntInt16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: int16
//	2. List
//
// Returns
//	New List of type int16
//	Empty list if all arguments are nil or either one is nil
func PMapIntInt16(f func(int) int16, list []int) []int16 {
	if f == nil {
		return []int16{}
	}

	ch := make(chan map[int]int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v int) {
			defer wg.Done()
			ch <- map[int]int16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapIntInt8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: int8
//	2. List
//
// Returns
//	New List of type int8
//	Empty list if all arguments are nil or either one is nil
func PMapIntInt8(f func(int) int8, list []int) []int8 {
	if f == nil {
		return []int8{}
	}

	ch := make(chan map[int]int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v int) {
			defer wg.Done()
			ch <- map[int]int8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapIntUint applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: uint
//	2. List
//
// Returns
//	New List of type uint
//	Empty list if all arguments are nil or either one is nil
func PMapIntUint(f func(int) uint, list []int) []uint {
	if f == nil {
		return []uint{}
	}

	ch := make(chan map[int]uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v int) {
			defer wg.Done()
			ch <- map[int]uint{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapIntUint64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: uint64
//	2. List
//
// Returns
//	New List of type uint64
//	Empty list if all arguments are nil or either one is nil
func PMapIntUint64(f func(int) uint64, list []int) []uint64 {
	if f == nil {
		return []uint64{}
	}

	ch := make(chan map[int]uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v int) {
			defer wg.Done()
			ch <- map[int]uint64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapIntUint32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: uint32
//	2. List
//
// Returns
//	New List of type uint32
//	Empty list if all arguments are nil or either one is nil
func PMapIntUint32(f func(int) uint32, list []int) []uint32 {
	if f == nil {
		return []uint32{}
	}

	ch := make(chan map[int]uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v int) {
			defer wg.Done()
			ch <- map[int]uint32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapIntUint16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: uint16
//	2. List
//
// Returns
//	New List of type uint16
//	Empty list if all arguments are nil or either one is nil
func PMapIntUint16(f func(int) uint16, list []int) []uint16 {
	if f == nil {
		return []uint16{}
	}

	ch := make(chan map[int]uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v int) {
			defer wg.Done()
			ch <- map[int]uint16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapIntUint8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: uint8
//	2. List
//
// Returns
//	New List of type uint8
//	Empty list if all arguments are nil or either one is nil
func PMapIntUint8(f func(int) uint8, list []int) []uint8 {
	if f == nil {
		return []uint8{}
	}

	ch := make(chan map[int]uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v int) {
			defer wg.Done()
			ch <- map[int]uint8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapIntStr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: string
//	2. List
//
// Returns
//	New List of type string
//	Empty list if all arguments are nil or either one is nil
func PMapIntStr(f func(int) string, list []int) []string {
	if f == nil {
		return []string{}
	}

	ch := make(chan map[int]string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v int) {
			defer wg.Done()
			ch <- map[int]string{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapIntBool applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: bool
//	2. List
//
// Returns
//	New List of type bool
//	Empty list if all arguments are nil or either one is nil
func PMapIntBool(f func(int) bool, list []int) []bool {
	if f == nil {
		return []bool{}
	}

	ch := make(chan map[int]bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v int) {
			defer wg.Done()
			ch <- map[int]bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
// PMapInt64Int applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: int
//	2. List
//
// Returns
//	New List of type int
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Int(f func(int64) int, list []int64) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v int64) {
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
// PMapInt64Int32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: int32
//	2. List
//
// Returns
//	New List of type int32
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Int32(f func(int64) int32, list []int64) []int32 {
	if f == nil {
		return []int32{}
	}

	ch := make(chan map[int]int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v int64) {
			defer wg.Done()
			ch <- map[int]int32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt64Int16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: int16
//	2. List
//
// Returns
//	New List of type int16
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Int16(f func(int64) int16, list []int64) []int16 {
	if f == nil {
		return []int16{}
	}

	ch := make(chan map[int]int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v int64) {
			defer wg.Done()
			ch <- map[int]int16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt64Int8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: int8
//	2. List
//
// Returns
//	New List of type int8
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Int8(f func(int64) int8, list []int64) []int8 {
	if f == nil {
		return []int8{}
	}

	ch := make(chan map[int]int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v int64) {
			defer wg.Done()
			ch <- map[int]int8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt64Uint applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: uint
//	2. List
//
// Returns
//	New List of type uint
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Uint(f func(int64) uint, list []int64) []uint {
	if f == nil {
		return []uint{}
	}

	ch := make(chan map[int]uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v int64) {
			defer wg.Done()
			ch <- map[int]uint{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt64Uint64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: uint64
//	2. List
//
// Returns
//	New List of type uint64
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Uint64(f func(int64) uint64, list []int64) []uint64 {
	if f == nil {
		return []uint64{}
	}

	ch := make(chan map[int]uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v int64) {
			defer wg.Done()
			ch <- map[int]uint64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt64Uint32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: uint32
//	2. List
//
// Returns
//	New List of type uint32
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Uint32(f func(int64) uint32, list []int64) []uint32 {
	if f == nil {
		return []uint32{}
	}

	ch := make(chan map[int]uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v int64) {
			defer wg.Done()
			ch <- map[int]uint32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt64Uint16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: uint16
//	2. List
//
// Returns
//	New List of type uint16
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Uint16(f func(int64) uint16, list []int64) []uint16 {
	if f == nil {
		return []uint16{}
	}

	ch := make(chan map[int]uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v int64) {
			defer wg.Done()
			ch <- map[int]uint16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt64Uint8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: uint8
//	2. List
//
// Returns
//	New List of type uint8
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Uint8(f func(int64) uint8, list []int64) []uint8 {
	if f == nil {
		return []uint8{}
	}

	ch := make(chan map[int]uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v int64) {
			defer wg.Done()
			ch <- map[int]uint8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt64Str applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: string
//	2. List
//
// Returns
//	New List of type string
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Str(f func(int64) string, list []int64) []string {
	if f == nil {
		return []string{}
	}

	ch := make(chan map[int]string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v int64) {
			defer wg.Done()
			ch <- map[int]string{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt64Bool applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: bool
//	2. List
//
// Returns
//	New List of type bool
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Bool(f func(int64) bool, list []int64) []bool {
	if f == nil {
		return []bool{}
	}

	ch := make(chan map[int]bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v int64) {
			defer wg.Done()
			ch <- map[int]bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
// PMapInt32Int applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: int
//	2. List
//
// Returns
//	New List of type int
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Int(f func(int32) int, list []int32) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v int32) {
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
// PMapInt32Int64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: int64
//	2. List
//
// Returns
//	New List of type int64
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Int64(f func(int32) int64, list []int32) []int64 {
	if f == nil {
		return []int64{}
	}

	ch := make(chan map[int]int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v int32) {
			defer wg.Done()
			ch <- map[int]int64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt32Int16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: int16
//	2. List
//
// Returns
//	New List of type int16
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Int16(f func(int32) int16, list []int32) []int16 {
	if f == nil {
		return []int16{}
	}

	ch := make(chan map[int]int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v int32) {
			defer wg.Done()
			ch <- map[int]int16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt32Int8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: int8
//	2. List
//
// Returns
//	New List of type int8
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Int8(f func(int32) int8, list []int32) []int8 {
	if f == nil {
		return []int8{}
	}

	ch := make(chan map[int]int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v int32) {
			defer wg.Done()
			ch <- map[int]int8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt32Uint applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: uint
//	2. List
//
// Returns
//	New List of type uint
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Uint(f func(int32) uint, list []int32) []uint {
	if f == nil {
		return []uint{}
	}

	ch := make(chan map[int]uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v int32) {
			defer wg.Done()
			ch <- map[int]uint{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt32Uint64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: uint64
//	2. List
//
// Returns
//	New List of type uint64
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Uint64(f func(int32) uint64, list []int32) []uint64 {
	if f == nil {
		return []uint64{}
	}

	ch := make(chan map[int]uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v int32) {
			defer wg.Done()
			ch <- map[int]uint64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt32Uint32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: uint32
//	2. List
//
// Returns
//	New List of type uint32
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Uint32(f func(int32) uint32, list []int32) []uint32 {
	if f == nil {
		return []uint32{}
	}

	ch := make(chan map[int]uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v int32) {
			defer wg.Done()
			ch <- map[int]uint32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt32Uint16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: uint16
//	2. List
//
// Returns
//	New List of type uint16
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Uint16(f func(int32) uint16, list []int32) []uint16 {
	if f == nil {
		return []uint16{}
	}

	ch := make(chan map[int]uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v int32) {
			defer wg.Done()
			ch <- map[int]uint16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt32Uint8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: uint8
//	2. List
//
// Returns
//	New List of type uint8
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Uint8(f func(int32) uint8, list []int32) []uint8 {
	if f == nil {
		return []uint8{}
	}

	ch := make(chan map[int]uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v int32) {
			defer wg.Done()
			ch <- map[int]uint8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt32Str applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: string
//	2. List
//
// Returns
//	New List of type string
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Str(f func(int32) string, list []int32) []string {
	if f == nil {
		return []string{}
	}

	ch := make(chan map[int]string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v int32) {
			defer wg.Done()
			ch <- map[int]string{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt32Bool applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: bool
//	2. List
//
// Returns
//	New List of type bool
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Bool(f func(int32) bool, list []int32) []bool {
	if f == nil {
		return []bool{}
	}

	ch := make(chan map[int]bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v int32) {
			defer wg.Done()
			ch <- map[int]bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
// PMapInt16Int applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: int
//	2. List
//
// Returns
//	New List of type int
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Int(f func(int16) int, list []int16) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v int16) {
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
// PMapInt16Int64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: int64
//	2. List
//
// Returns
//	New List of type int64
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Int64(f func(int16) int64, list []int16) []int64 {
	if f == nil {
		return []int64{}
	}

	ch := make(chan map[int]int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v int16) {
			defer wg.Done()
			ch <- map[int]int64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt16Int32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: int32
//	2. List
//
// Returns
//	New List of type int32
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Int32(f func(int16) int32, list []int16) []int32 {
	if f == nil {
		return []int32{}
	}

	ch := make(chan map[int]int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v int16) {
			defer wg.Done()
			ch <- map[int]int32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt16Int8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: int8
//	2. List
//
// Returns
//	New List of type int8
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Int8(f func(int16) int8, list []int16) []int8 {
	if f == nil {
		return []int8{}
	}

	ch := make(chan map[int]int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v int16) {
			defer wg.Done()
			ch <- map[int]int8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt16Uint applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: uint
//	2. List
//
// Returns
//	New List of type uint
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Uint(f func(int16) uint, list []int16) []uint {
	if f == nil {
		return []uint{}
	}

	ch := make(chan map[int]uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v int16) {
			defer wg.Done()
			ch <- map[int]uint{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt16Uint64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: uint64
//	2. List
//
// Returns
//	New List of type uint64
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Uint64(f func(int16) uint64, list []int16) []uint64 {
	if f == nil {
		return []uint64{}
	}

	ch := make(chan map[int]uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v int16) {
			defer wg.Done()
			ch <- map[int]uint64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt16Uint32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: uint32
//	2. List
//
// Returns
//	New List of type uint32
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Uint32(f func(int16) uint32, list []int16) []uint32 {
	if f == nil {
		return []uint32{}
	}

	ch := make(chan map[int]uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v int16) {
			defer wg.Done()
			ch <- map[int]uint32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt16Uint16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: uint16
//	2. List
//
// Returns
//	New List of type uint16
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Uint16(f func(int16) uint16, list []int16) []uint16 {
	if f == nil {
		return []uint16{}
	}

	ch := make(chan map[int]uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v int16) {
			defer wg.Done()
			ch <- map[int]uint16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt16Uint8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: uint8
//	2. List
//
// Returns
//	New List of type uint8
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Uint8(f func(int16) uint8, list []int16) []uint8 {
	if f == nil {
		return []uint8{}
	}

	ch := make(chan map[int]uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v int16) {
			defer wg.Done()
			ch <- map[int]uint8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt16Str applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: string
//	2. List
//
// Returns
//	New List of type string
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Str(f func(int16) string, list []int16) []string {
	if f == nil {
		return []string{}
	}

	ch := make(chan map[int]string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v int16) {
			defer wg.Done()
			ch <- map[int]string{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt16Bool applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: bool
//	2. List
//
// Returns
//	New List of type bool
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Bool(f func(int16) bool, list []int16) []bool {
	if f == nil {
		return []bool{}
	}

	ch := make(chan map[int]bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v int16) {
			defer wg.Done()
			ch <- map[int]bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
// PMapInt8Int applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: int
//	2. List
//
// Returns
//	New List of type int
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Int(f func(int8) int, list []int8) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v int8) {
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
// PMapInt8Int64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: int64
//	2. List
//
// Returns
//	New List of type int64
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Int64(f func(int8) int64, list []int8) []int64 {
	if f == nil {
		return []int64{}
	}

	ch := make(chan map[int]int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v int8) {
			defer wg.Done()
			ch <- map[int]int64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt8Int32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: int32
//	2. List
//
// Returns
//	New List of type int32
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Int32(f func(int8) int32, list []int8) []int32 {
	if f == nil {
		return []int32{}
	}

	ch := make(chan map[int]int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v int8) {
			defer wg.Done()
			ch <- map[int]int32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt8Int16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: int16
//	2. List
//
// Returns
//	New List of type int16
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Int16(f func(int8) int16, list []int8) []int16 {
	if f == nil {
		return []int16{}
	}

	ch := make(chan map[int]int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v int8) {
			defer wg.Done()
			ch <- map[int]int16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt8Uint applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: uint
//	2. List
//
// Returns
//	New List of type uint
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Uint(f func(int8) uint, list []int8) []uint {
	if f == nil {
		return []uint{}
	}

	ch := make(chan map[int]uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v int8) {
			defer wg.Done()
			ch <- map[int]uint{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt8Uint64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: uint64
//	2. List
//
// Returns
//	New List of type uint64
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Uint64(f func(int8) uint64, list []int8) []uint64 {
	if f == nil {
		return []uint64{}
	}

	ch := make(chan map[int]uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v int8) {
			defer wg.Done()
			ch <- map[int]uint64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt8Uint32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: uint32
//	2. List
//
// Returns
//	New List of type uint32
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Uint32(f func(int8) uint32, list []int8) []uint32 {
	if f == nil {
		return []uint32{}
	}

	ch := make(chan map[int]uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v int8) {
			defer wg.Done()
			ch <- map[int]uint32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt8Uint16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: uint16
//	2. List
//
// Returns
//	New List of type uint16
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Uint16(f func(int8) uint16, list []int8) []uint16 {
	if f == nil {
		return []uint16{}
	}

	ch := make(chan map[int]uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v int8) {
			defer wg.Done()
			ch <- map[int]uint16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt8Uint8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: uint8
//	2. List
//
// Returns
//	New List of type uint8
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Uint8(f func(int8) uint8, list []int8) []uint8 {
	if f == nil {
		return []uint8{}
	}

	ch := make(chan map[int]uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v int8) {
			defer wg.Done()
			ch <- map[int]uint8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt8Str applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: string
//	2. List
//
// Returns
//	New List of type string
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Str(f func(int8) string, list []int8) []string {
	if f == nil {
		return []string{}
	}

	ch := make(chan map[int]string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v int8) {
			defer wg.Done()
			ch <- map[int]string{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapInt8Bool applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: bool
//	2. List
//
// Returns
//	New List of type bool
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Bool(f func(int8) bool, list []int8) []bool {
	if f == nil {
		return []bool{}
	}

	ch := make(chan map[int]bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v int8) {
			defer wg.Done()
			ch <- map[int]bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
// PMapUintInt applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: int
//	2. List
//
// Returns
//	New List of type int
//	Empty list if all arguments are nil or either one is nil
func PMapUintInt(f func(uint) int, list []uint) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v uint) {
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
// PMapUintInt64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: int64
//	2. List
//
// Returns
//	New List of type int64
//	Empty list if all arguments are nil or either one is nil
func PMapUintInt64(f func(uint) int64, list []uint) []int64 {
	if f == nil {
		return []int64{}
	}

	ch := make(chan map[int]int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v uint) {
			defer wg.Done()
			ch <- map[int]int64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUintInt32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: int32
//	2. List
//
// Returns
//	New List of type int32
//	Empty list if all arguments are nil or either one is nil
func PMapUintInt32(f func(uint) int32, list []uint) []int32 {
	if f == nil {
		return []int32{}
	}

	ch := make(chan map[int]int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v uint) {
			defer wg.Done()
			ch <- map[int]int32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUintInt16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: int16
//	2. List
//
// Returns
//	New List of type int16
//	Empty list if all arguments are nil or either one is nil
func PMapUintInt16(f func(uint) int16, list []uint) []int16 {
	if f == nil {
		return []int16{}
	}

	ch := make(chan map[int]int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v uint) {
			defer wg.Done()
			ch <- map[int]int16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUintInt8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: int8
//	2. List
//
// Returns
//	New List of type int8
//	Empty list if all arguments are nil or either one is nil
func PMapUintInt8(f func(uint) int8, list []uint) []int8 {
	if f == nil {
		return []int8{}
	}

	ch := make(chan map[int]int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v uint) {
			defer wg.Done()
			ch <- map[int]int8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUintUint64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: uint64
//	2. List
//
// Returns
//	New List of type uint64
//	Empty list if all arguments are nil or either one is nil
func PMapUintUint64(f func(uint) uint64, list []uint) []uint64 {
	if f == nil {
		return []uint64{}
	}

	ch := make(chan map[int]uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v uint) {
			defer wg.Done()
			ch <- map[int]uint64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUintUint32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: uint32
//	2. List
//
// Returns
//	New List of type uint32
//	Empty list if all arguments are nil or either one is nil
func PMapUintUint32(f func(uint) uint32, list []uint) []uint32 {
	if f == nil {
		return []uint32{}
	}

	ch := make(chan map[int]uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v uint) {
			defer wg.Done()
			ch <- map[int]uint32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUintUint16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: uint16
//	2. List
//
// Returns
//	New List of type uint16
//	Empty list if all arguments are nil or either one is nil
func PMapUintUint16(f func(uint) uint16, list []uint) []uint16 {
	if f == nil {
		return []uint16{}
	}

	ch := make(chan map[int]uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v uint) {
			defer wg.Done()
			ch <- map[int]uint16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUintUint8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: uint8
//	2. List
//
// Returns
//	New List of type uint8
//	Empty list if all arguments are nil or either one is nil
func PMapUintUint8(f func(uint) uint8, list []uint) []uint8 {
	if f == nil {
		return []uint8{}
	}

	ch := make(chan map[int]uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v uint) {
			defer wg.Done()
			ch <- map[int]uint8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUintStr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: string
//	2. List
//
// Returns
//	New List of type string
//	Empty list if all arguments are nil or either one is nil
func PMapUintStr(f func(uint) string, list []uint) []string {
	if f == nil {
		return []string{}
	}

	ch := make(chan map[int]string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v uint) {
			defer wg.Done()
			ch <- map[int]string{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUintBool applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: bool
//	2. List
//
// Returns
//	New List of type bool
//	Empty list if all arguments are nil or either one is nil
func PMapUintBool(f func(uint) bool, list []uint) []bool {
	if f == nil {
		return []bool{}
	}

	ch := make(chan map[int]bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v uint) {
			defer wg.Done()
			ch <- map[int]bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
// PMapUint64Int applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: int
//	2. List
//
// Returns
//	New List of type int
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Int(f func(uint64) int, list []uint64) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v uint64) {
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
// PMapUint64Int64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: int64
//	2. List
//
// Returns
//	New List of type int64
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Int64(f func(uint64) int64, list []uint64) []int64 {
	if f == nil {
		return []int64{}
	}

	ch := make(chan map[int]int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v uint64) {
			defer wg.Done()
			ch <- map[int]int64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint64Int32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: int32
//	2. List
//
// Returns
//	New List of type int32
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Int32(f func(uint64) int32, list []uint64) []int32 {
	if f == nil {
		return []int32{}
	}

	ch := make(chan map[int]int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v uint64) {
			defer wg.Done()
			ch <- map[int]int32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint64Int16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: int16
//	2. List
//
// Returns
//	New List of type int16
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Int16(f func(uint64) int16, list []uint64) []int16 {
	if f == nil {
		return []int16{}
	}

	ch := make(chan map[int]int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v uint64) {
			defer wg.Done()
			ch <- map[int]int16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint64Int8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: int8
//	2. List
//
// Returns
//	New List of type int8
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Int8(f func(uint64) int8, list []uint64) []int8 {
	if f == nil {
		return []int8{}
	}

	ch := make(chan map[int]int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v uint64) {
			defer wg.Done()
			ch <- map[int]int8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint64Uint applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: uint
//	2. List
//
// Returns
//	New List of type uint
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Uint(f func(uint64) uint, list []uint64) []uint {
	if f == nil {
		return []uint{}
	}

	ch := make(chan map[int]uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v uint64) {
			defer wg.Done()
			ch <- map[int]uint{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint64Uint32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: uint32
//	2. List
//
// Returns
//	New List of type uint32
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Uint32(f func(uint64) uint32, list []uint64) []uint32 {
	if f == nil {
		return []uint32{}
	}

	ch := make(chan map[int]uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v uint64) {
			defer wg.Done()
			ch <- map[int]uint32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint64Uint16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: uint16
//	2. List
//
// Returns
//	New List of type uint16
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Uint16(f func(uint64) uint16, list []uint64) []uint16 {
	if f == nil {
		return []uint16{}
	}

	ch := make(chan map[int]uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v uint64) {
			defer wg.Done()
			ch <- map[int]uint16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint64Uint8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: uint8
//	2. List
//
// Returns
//	New List of type uint8
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Uint8(f func(uint64) uint8, list []uint64) []uint8 {
	if f == nil {
		return []uint8{}
	}

	ch := make(chan map[int]uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v uint64) {
			defer wg.Done()
			ch <- map[int]uint8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint64Str applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: string
//	2. List
//
// Returns
//	New List of type string
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Str(f func(uint64) string, list []uint64) []string {
	if f == nil {
		return []string{}
	}

	ch := make(chan map[int]string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v uint64) {
			defer wg.Done()
			ch <- map[int]string{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint64Bool applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: bool
//	2. List
//
// Returns
//	New List of type bool
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Bool(f func(uint64) bool, list []uint64) []bool {
	if f == nil {
		return []bool{}
	}

	ch := make(chan map[int]bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v uint64) {
			defer wg.Done()
			ch <- map[int]bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
// PMapUint32Int applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: int
//	2. List
//
// Returns
//	New List of type int
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Int(f func(uint32) int, list []uint32) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v uint32) {
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
// PMapUint32Int64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: int64
//	2. List
//
// Returns
//	New List of type int64
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Int64(f func(uint32) int64, list []uint32) []int64 {
	if f == nil {
		return []int64{}
	}

	ch := make(chan map[int]int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v uint32) {
			defer wg.Done()
			ch <- map[int]int64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint32Int32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: int32
//	2. List
//
// Returns
//	New List of type int32
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Int32(f func(uint32) int32, list []uint32) []int32 {
	if f == nil {
		return []int32{}
	}

	ch := make(chan map[int]int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v uint32) {
			defer wg.Done()
			ch <- map[int]int32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint32Int16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: int16
//	2. List
//
// Returns
//	New List of type int16
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Int16(f func(uint32) int16, list []uint32) []int16 {
	if f == nil {
		return []int16{}
	}

	ch := make(chan map[int]int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v uint32) {
			defer wg.Done()
			ch <- map[int]int16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint32Int8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: int8
//	2. List
//
// Returns
//	New List of type int8
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Int8(f func(uint32) int8, list []uint32) []int8 {
	if f == nil {
		return []int8{}
	}

	ch := make(chan map[int]int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v uint32) {
			defer wg.Done()
			ch <- map[int]int8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint32Uint applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: uint
//	2. List
//
// Returns
//	New List of type uint
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Uint(f func(uint32) uint, list []uint32) []uint {
	if f == nil {
		return []uint{}
	}

	ch := make(chan map[int]uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v uint32) {
			defer wg.Done()
			ch <- map[int]uint{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint32Uint64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: uint64
//	2. List
//
// Returns
//	New List of type uint64
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Uint64(f func(uint32) uint64, list []uint32) []uint64 {
	if f == nil {
		return []uint64{}
	}

	ch := make(chan map[int]uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v uint32) {
			defer wg.Done()
			ch <- map[int]uint64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint32Uint16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: uint16
//	2. List
//
// Returns
//	New List of type uint16
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Uint16(f func(uint32) uint16, list []uint32) []uint16 {
	if f == nil {
		return []uint16{}
	}

	ch := make(chan map[int]uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v uint32) {
			defer wg.Done()
			ch <- map[int]uint16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint32Uint8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: uint8
//	2. List
//
// Returns
//	New List of type uint8
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Uint8(f func(uint32) uint8, list []uint32) []uint8 {
	if f == nil {
		return []uint8{}
	}

	ch := make(chan map[int]uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v uint32) {
			defer wg.Done()
			ch <- map[int]uint8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint32Str applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: string
//	2. List
//
// Returns
//	New List of type string
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Str(f func(uint32) string, list []uint32) []string {
	if f == nil {
		return []string{}
	}

	ch := make(chan map[int]string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v uint32) {
			defer wg.Done()
			ch <- map[int]string{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint32Bool applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: bool
//	2. List
//
// Returns
//	New List of type bool
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Bool(f func(uint32) bool, list []uint32) []bool {
	if f == nil {
		return []bool{}
	}

	ch := make(chan map[int]bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v uint32) {
			defer wg.Done()
			ch <- map[int]bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
// PMapUint16Int applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: int
//	2. List
//
// Returns
//	New List of type int
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Int(f func(uint16) int, list []uint16) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v uint16) {
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
// PMapUint16Int64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: int64
//	2. List
//
// Returns
//	New List of type int64
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Int64(f func(uint16) int64, list []uint16) []int64 {
	if f == nil {
		return []int64{}
	}

	ch := make(chan map[int]int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v uint16) {
			defer wg.Done()
			ch <- map[int]int64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint16Int32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: int32
//	2. List
//
// Returns
//	New List of type int32
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Int32(f func(uint16) int32, list []uint16) []int32 {
	if f == nil {
		return []int32{}
	}

	ch := make(chan map[int]int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v uint16) {
			defer wg.Done()
			ch <- map[int]int32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint16Int16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: int16
//	2. List
//
// Returns
//	New List of type int16
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Int16(f func(uint16) int16, list []uint16) []int16 {
	if f == nil {
		return []int16{}
	}

	ch := make(chan map[int]int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v uint16) {
			defer wg.Done()
			ch <- map[int]int16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint16Int8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: int8
//	2. List
//
// Returns
//	New List of type int8
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Int8(f func(uint16) int8, list []uint16) []int8 {
	if f == nil {
		return []int8{}
	}

	ch := make(chan map[int]int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v uint16) {
			defer wg.Done()
			ch <- map[int]int8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint16Uint applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: uint
//	2. List
//
// Returns
//	New List of type uint
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Uint(f func(uint16) uint, list []uint16) []uint {
	if f == nil {
		return []uint{}
	}

	ch := make(chan map[int]uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v uint16) {
			defer wg.Done()
			ch <- map[int]uint{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint16Uint64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: uint64
//	2. List
//
// Returns
//	New List of type uint64
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Uint64(f func(uint16) uint64, list []uint16) []uint64 {
	if f == nil {
		return []uint64{}
	}

	ch := make(chan map[int]uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v uint16) {
			defer wg.Done()
			ch <- map[int]uint64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint16Uint32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: uint32
//	2. List
//
// Returns
//	New List of type uint32
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Uint32(f func(uint16) uint32, list []uint16) []uint32 {
	if f == nil {
		return []uint32{}
	}

	ch := make(chan map[int]uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v uint16) {
			defer wg.Done()
			ch <- map[int]uint32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint16Uint8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: uint8
//	2. List
//
// Returns
//	New List of type uint8
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Uint8(f func(uint16) uint8, list []uint16) []uint8 {
	if f == nil {
		return []uint8{}
	}

	ch := make(chan map[int]uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v uint16) {
			defer wg.Done()
			ch <- map[int]uint8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint16Str applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: string
//	2. List
//
// Returns
//	New List of type string
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Str(f func(uint16) string, list []uint16) []string {
	if f == nil {
		return []string{}
	}

	ch := make(chan map[int]string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v uint16) {
			defer wg.Done()
			ch <- map[int]string{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint16Bool applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: bool
//	2. List
//
// Returns
//	New List of type bool
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Bool(f func(uint16) bool, list []uint16) []bool {
	if f == nil {
		return []bool{}
	}

	ch := make(chan map[int]bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v uint16) {
			defer wg.Done()
			ch <- map[int]bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
// PMapUint8Int applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: int
//	2. List
//
// Returns
//	New List of type int
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Int(f func(uint8) int, list []uint8) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v uint8) {
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
// PMapUint8Int64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: int64
//	2. List
//
// Returns
//	New List of type int64
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Int64(f func(uint8) int64, list []uint8) []int64 {
	if f == nil {
		return []int64{}
	}

	ch := make(chan map[int]int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v uint8) {
			defer wg.Done()
			ch <- map[int]int64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint8Int32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: int32
//	2. List
//
// Returns
//	New List of type int32
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Int32(f func(uint8) int32, list []uint8) []int32 {
	if f == nil {
		return []int32{}
	}

	ch := make(chan map[int]int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v uint8) {
			defer wg.Done()
			ch <- map[int]int32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint8Int16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: int16
//	2. List
//
// Returns
//	New List of type int16
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Int16(f func(uint8) int16, list []uint8) []int16 {
	if f == nil {
		return []int16{}
	}

	ch := make(chan map[int]int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v uint8) {
			defer wg.Done()
			ch <- map[int]int16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint8Int8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: int8
//	2. List
//
// Returns
//	New List of type int8
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Int8(f func(uint8) int8, list []uint8) []int8 {
	if f == nil {
		return []int8{}
	}

	ch := make(chan map[int]int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v uint8) {
			defer wg.Done()
			ch <- map[int]int8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint8Uint applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: uint
//	2. List
//
// Returns
//	New List of type uint
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Uint(f func(uint8) uint, list []uint8) []uint {
	if f == nil {
		return []uint{}
	}

	ch := make(chan map[int]uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v uint8) {
			defer wg.Done()
			ch <- map[int]uint{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint8Uint64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: uint64
//	2. List
//
// Returns
//	New List of type uint64
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Uint64(f func(uint8) uint64, list []uint8) []uint64 {
	if f == nil {
		return []uint64{}
	}

	ch := make(chan map[int]uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v uint8) {
			defer wg.Done()
			ch <- map[int]uint64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint8Uint32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: uint32
//	2. List
//
// Returns
//	New List of type uint32
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Uint32(f func(uint8) uint32, list []uint8) []uint32 {
	if f == nil {
		return []uint32{}
	}

	ch := make(chan map[int]uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v uint8) {
			defer wg.Done()
			ch <- map[int]uint32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint8Uint16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: uint16
//	2. List
//
// Returns
//	New List of type uint16
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Uint16(f func(uint8) uint16, list []uint8) []uint16 {
	if f == nil {
		return []uint16{}
	}

	ch := make(chan map[int]uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v uint8) {
			defer wg.Done()
			ch <- map[int]uint16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint8Str applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: string
//	2. List
//
// Returns
//	New List of type string
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Str(f func(uint8) string, list []uint8) []string {
	if f == nil {
		return []string{}
	}

	ch := make(chan map[int]string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v uint8) {
			defer wg.Done()
			ch <- map[int]string{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapUint8Bool applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: bool
//	2. List
//
// Returns
//	New List of type bool
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Bool(f func(uint8) bool, list []uint8) []bool {
	if f == nil {
		return []bool{}
	}

	ch := make(chan map[int]bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v uint8) {
			defer wg.Done()
			ch <- map[int]bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
// PMapStrInt applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: int
//	2. List
//
// Returns
//	New List of type int
//	Empty list if all arguments are nil or either one is nil
func PMapStrInt(f func(string) int, list []string) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v string) {
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
// PMapStrInt64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: int64
//	2. List
//
// Returns
//	New List of type int64
//	Empty list if all arguments are nil or either one is nil
func PMapStrInt64(f func(string) int64, list []string) []int64 {
	if f == nil {
		return []int64{}
	}

	ch := make(chan map[int]int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v string) {
			defer wg.Done()
			ch <- map[int]int64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapStrInt32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: int32
//	2. List
//
// Returns
//	New List of type int32
//	Empty list if all arguments are nil or either one is nil
func PMapStrInt32(f func(string) int32, list []string) []int32 {
	if f == nil {
		return []int32{}
	}

	ch := make(chan map[int]int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v string) {
			defer wg.Done()
			ch <- map[int]int32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapStrInt16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: int16
//	2. List
//
// Returns
//	New List of type int16
//	Empty list if all arguments are nil or either one is nil
func PMapStrInt16(f func(string) int16, list []string) []int16 {
	if f == nil {
		return []int16{}
	}

	ch := make(chan map[int]int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v string) {
			defer wg.Done()
			ch <- map[int]int16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapStrInt8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: int8
//	2. List
//
// Returns
//	New List of type int8
//	Empty list if all arguments are nil or either one is nil
func PMapStrInt8(f func(string) int8, list []string) []int8 {
	if f == nil {
		return []int8{}
	}

	ch := make(chan map[int]int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v string) {
			defer wg.Done()
			ch <- map[int]int8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapStrUint applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: uint
//	2. List
//
// Returns
//	New List of type uint
//	Empty list if all arguments are nil or either one is nil
func PMapStrUint(f func(string) uint, list []string) []uint {
	if f == nil {
		return []uint{}
	}

	ch := make(chan map[int]uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v string) {
			defer wg.Done()
			ch <- map[int]uint{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapStrUint64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: uint64
//	2. List
//
// Returns
//	New List of type uint64
//	Empty list if all arguments are nil or either one is nil
func PMapStrUint64(f func(string) uint64, list []string) []uint64 {
	if f == nil {
		return []uint64{}
	}

	ch := make(chan map[int]uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v string) {
			defer wg.Done()
			ch <- map[int]uint64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapStrUint32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: uint32
//	2. List
//
// Returns
//	New List of type uint32
//	Empty list if all arguments are nil or either one is nil
func PMapStrUint32(f func(string) uint32, list []string) []uint32 {
	if f == nil {
		return []uint32{}
	}

	ch := make(chan map[int]uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v string) {
			defer wg.Done()
			ch <- map[int]uint32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapStrUint16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: uint16
//	2. List
//
// Returns
//	New List of type uint16
//	Empty list if all arguments are nil or either one is nil
func PMapStrUint16(f func(string) uint16, list []string) []uint16 {
	if f == nil {
		return []uint16{}
	}

	ch := make(chan map[int]uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v string) {
			defer wg.Done()
			ch <- map[int]uint16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapStrUint8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: uint8
//	2. List
//
// Returns
//	New List of type uint8
//	Empty list if all arguments are nil or either one is nil
func PMapStrUint8(f func(string) uint8, list []string) []uint8 {
	if f == nil {
		return []uint8{}
	}

	ch := make(chan map[int]uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v string) {
			defer wg.Done()
			ch <- map[int]uint8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapStrBool applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: bool
//	2. List
//
// Returns
//	New List of type bool
//	Empty list if all arguments are nil or either one is nil
func PMapStrBool(f func(string) bool, list []string) []bool {
	if f == nil {
		return []bool{}
	}

	ch := make(chan map[int]bool)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v string) {
			defer wg.Done()
			ch <- map[int]bool{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
// PMapBoolInt applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: int
//	2. List
//
// Returns
//	New List of type int
//	Empty list if all arguments are nil or either one is nil
func PMapBoolInt(f func(bool) int, list []bool) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v bool) {
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
// PMapBoolInt64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: int64
//	2. List
//
// Returns
//	New List of type int64
//	Empty list if all arguments are nil or either one is nil
func PMapBoolInt64(f func(bool) int64, list []bool) []int64 {
	if f == nil {
		return []int64{}
	}

	ch := make(chan map[int]int64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v bool) {
			defer wg.Done()
			ch <- map[int]int64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapBoolInt32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: int32
//	2. List
//
// Returns
//	New List of type int32
//	Empty list if all arguments are nil or either one is nil
func PMapBoolInt32(f func(bool) int32, list []bool) []int32 {
	if f == nil {
		return []int32{}
	}

	ch := make(chan map[int]int32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v bool) {
			defer wg.Done()
			ch <- map[int]int32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapBoolInt16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: int16
//	2. List
//
// Returns
//	New List of type int16
//	Empty list if all arguments are nil or either one is nil
func PMapBoolInt16(f func(bool) int16, list []bool) []int16 {
	if f == nil {
		return []int16{}
	}

	ch := make(chan map[int]int16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v bool) {
			defer wg.Done()
			ch <- map[int]int16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapBoolInt8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: int8
//	2. List
//
// Returns
//	New List of type int8
//	Empty list if all arguments are nil or either one is nil
func PMapBoolInt8(f func(bool) int8, list []bool) []int8 {
	if f == nil {
		return []int8{}
	}

	ch := make(chan map[int]int8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v bool) {
			defer wg.Done()
			ch <- map[int]int8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapBoolUint applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: uint
//	2. List
//
// Returns
//	New List of type uint
//	Empty list if all arguments are nil or either one is nil
func PMapBoolUint(f func(bool) uint, list []bool) []uint {
	if f == nil {
		return []uint{}
	}

	ch := make(chan map[int]uint)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v bool) {
			defer wg.Done()
			ch <- map[int]uint{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapBoolUint64 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: uint64
//	2. List
//
// Returns
//	New List of type uint64
//	Empty list if all arguments are nil or either one is nil
func PMapBoolUint64(f func(bool) uint64, list []bool) []uint64 {
	if f == nil {
		return []uint64{}
	}

	ch := make(chan map[int]uint64)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v bool) {
			defer wg.Done()
			ch <- map[int]uint64{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapBoolUint32 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: uint32
//	2. List
//
// Returns
//	New List of type uint32
//	Empty list if all arguments are nil or either one is nil
func PMapBoolUint32(f func(bool) uint32, list []bool) []uint32 {
	if f == nil {
		return []uint32{}
	}

	ch := make(chan map[int]uint32)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v bool) {
			defer wg.Done()
			ch <- map[int]uint32{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapBoolUint16 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: uint16
//	2. List
//
// Returns
//	New List of type uint16
//	Empty list if all arguments are nil or either one is nil
func PMapBoolUint16(f func(bool) uint16, list []bool) []uint16 {
	if f == nil {
		return []uint16{}
	}

	ch := make(chan map[int]uint16)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v bool) {
			defer wg.Done()
			ch <- map[int]uint16{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapBoolUint8 applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: uint8
//	2. List
//
// Returns
//	New List of type uint8
//	Empty list if all arguments are nil or either one is nil
func PMapBoolUint8(f func(bool) uint8, list []bool) []uint8 {
	if f == nil {
		return []uint8{}
	}

	ch := make(chan map[int]uint8)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v bool) {
			defer wg.Done()
			ch <- map[int]uint8{i: f(v)}
		}(&wg, ch, i, v)
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
// PMapBoolStr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: string
//	2. List
//
// Returns
//	New List of type string
//	Empty list if all arguments are nil or either one is nil
func PMapBoolStr(f func(bool) string, list []bool) []string {
	if f == nil {
		return []string{}
	}

	ch := make(chan map[int]string)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v bool) {
			defer wg.Done()
			ch <- map[int]string{i: f(v)}
		}(&wg, ch, i, v)
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