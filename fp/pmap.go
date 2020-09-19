package fp

import "sync"

// PMapInt applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt(f func(int) int, list []int, optional ...Optional) []int {
	if f == nil {
		return []int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntNoOrder(f, list, worker)
		}
	}

	return pMapIntPreserveOrder(f, list, worker)
}

func pMapIntPreserveOrder(f func(int) int, list []int, worker int) []int {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]int) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]int{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]int, len(list))
	newList := make([]int, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapIntNoOrder(f func(int) int, list []int, worker int) []int {
	chJobs := make(chan int, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan int, chJobs chan int) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]int, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64(f func(int64) int64, list []int64, optional ...Optional) []int64 {
	if f == nil {
		return []int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64NoOrder(f, list, worker)
		}
	}

	return pMapInt64PreserveOrder(f, list, worker)
}

func pMapInt64PreserveOrder(f func(int64) int64, list []int64, worker int) []int64 {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]int64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]int64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]int64, len(list))
	newList := make([]int64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt64NoOrder(f func(int64) int64, list []int64, worker int) []int64 {
	chJobs := make(chan int64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan int64, chJobs chan int64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]int64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32(f func(int32) int32, list []int32, optional ...Optional) []int32 {
	if f == nil {
		return []int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32NoOrder(f, list, worker)
		}
	}

	return pMapInt32PreserveOrder(f, list, worker)
}

func pMapInt32PreserveOrder(f func(int32) int32, list []int32, worker int) []int32 {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]int32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]int32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]int32, len(list))
	newList := make([]int32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt32NoOrder(f func(int32) int32, list []int32, worker int) []int32 {
	chJobs := make(chan int32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan int32, chJobs chan int32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]int32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16(f func(int16) int16, list []int16, optional ...Optional) []int16 {
	if f == nil {
		return []int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16NoOrder(f, list, worker)
		}
	}

	return pMapInt16PreserveOrder(f, list, worker)
}

func pMapInt16PreserveOrder(f func(int16) int16, list []int16, worker int) []int16 {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]int16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]int16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]int16, len(list))
	newList := make([]int16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt16NoOrder(f func(int16) int16, list []int16, worker int) []int16 {
	chJobs := make(chan int16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan int16, chJobs chan int16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]int16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapInt8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8(f func(int8) int8, list []int8, optional ...Optional) []int8 {
	if f == nil {
		return []int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8NoOrder(f, list, worker)
		}
	}

	return pMapInt8PreserveOrder(f, list, worker)
}

func pMapInt8PreserveOrder(f func(int8) int8, list []int8, worker int) []int8 {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]int8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]int8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]int8, len(list))
	newList := make([]int8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapInt8NoOrder(f func(int8) int8, list []int8, worker int) []int8 {
	chJobs := make(chan int8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan int8, chJobs chan int8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]int8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint(f func(uint) uint, list []uint, optional ...Optional) []uint {
	if f == nil {
		return []uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintNoOrder(f, list, worker)
		}
	}

	return pMapUintPreserveOrder(f, list, worker)
}

func pMapUintPreserveOrder(f func(uint) uint, list []uint, worker int) []uint {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]uint) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]uint{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]uint, len(list))
	newList := make([]uint, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUintNoOrder(f func(uint) uint, list []uint, worker int) []uint {
	chJobs := make(chan uint, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan uint, chJobs chan uint) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]uint, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64(f func(uint64) uint64, list []uint64, optional ...Optional) []uint64 {
	if f == nil {
		return []uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64NoOrder(f, list, worker)
		}
	}

	return pMapUint64PreserveOrder(f, list, worker)
}

func pMapUint64PreserveOrder(f func(uint64) uint64, list []uint64, worker int) []uint64 {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]uint64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]uint64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]uint64, len(list))
	newList := make([]uint64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint64NoOrder(f func(uint64) uint64, list []uint64, worker int) []uint64 {
	chJobs := make(chan uint64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan uint64, chJobs chan uint64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]uint64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32(f func(uint32) uint32, list []uint32, optional ...Optional) []uint32 {
	if f == nil {
		return []uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32NoOrder(f, list, worker)
		}
	}

	return pMapUint32PreserveOrder(f, list, worker)
}

func pMapUint32PreserveOrder(f func(uint32) uint32, list []uint32, worker int) []uint32 {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]uint32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]uint32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]uint32, len(list))
	newList := make([]uint32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint32NoOrder(f func(uint32) uint32, list []uint32, worker int) []uint32 {
	chJobs := make(chan uint32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan uint32, chJobs chan uint32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]uint32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16(f func(uint16) uint16, list []uint16, optional ...Optional) []uint16 {
	if f == nil {
		return []uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16NoOrder(f, list, worker)
		}
	}

	return pMapUint16PreserveOrder(f, list, worker)
}

func pMapUint16PreserveOrder(f func(uint16) uint16, list []uint16, worker int) []uint16 {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]uint16) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]uint16{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]uint16, len(list))
	newList := make([]uint16, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint16NoOrder(f func(uint16) uint16, list []uint16, worker int) []uint16 {
	chJobs := make(chan uint16, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan uint16, chJobs chan uint16) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]uint16, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapUint8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8(f func(uint8) uint8, list []uint8, optional ...Optional) []uint8 {
	if f == nil {
		return []uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8NoOrder(f, list, worker)
		}
	}

	return pMapUint8PreserveOrder(f, list, worker)
}

func pMapUint8PreserveOrder(f func(uint8) uint8, list []uint8, worker int) []uint8 {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]uint8) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]uint8{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]uint8, len(list))
	newList := make([]uint8, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapUint8NoOrder(f func(uint8) uint8, list []uint8, worker int) []uint8 {
	chJobs := make(chan uint8, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan uint8, chJobs chan uint8) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]uint8, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapStr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStr(f func(string) string, list []string, optional ...Optional) []string {
	if f == nil {
		return []string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrNoOrder(f, list, worker)
		}
	}

	return pMapStrPreserveOrder(f, list, worker)
}

func pMapStrPreserveOrder(f func(string) string, list []string, worker int) []string {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]string) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]string{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]string, len(list))
	newList := make([]string, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapStrNoOrder(f func(string) string, list []string, worker int) []string {
	chJobs := make(chan string, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan string, chJobs chan string) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]string, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapBool applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBool(f func(bool) bool, list []bool, optional ...Optional) []bool {
	if f == nil {
		return []bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolNoOrder(f, list, worker)
		}
	}

	return pMapBoolPreserveOrder(f, list, worker)
}

func pMapBoolPreserveOrder(f func(bool) bool, list []bool, worker int) []bool {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]bool) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]bool{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]bool, len(list))
	newList := make([]bool, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapBoolNoOrder(f func(bool) bool, list []bool, worker int) []bool {
	chJobs := make(chan bool, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan bool, chJobs chan bool) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]bool, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32(f func(float32) float32, list []float32, optional ...Optional) []float32 {
	if f == nil {
		return []float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32NoOrder(f, list, worker)
		}
	}

	return pMapFloat32PreserveOrder(f, list, worker)
}

func pMapFloat32PreserveOrder(f func(float32) float32, list []float32, worker int) []float32 {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]float32) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]float32{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]float32, len(list))
	newList := make([]float32, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat32NoOrder(f func(float32) float32, list []float32, worker int) []float32 {
	chJobs := make(chan float32, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan float32, chJobs chan float32) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]float32, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}

// PMapFloat64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64(f func(float64) float64, list []float64, optional ...Optional) []float64 {
	if f == nil {
		return []float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64NoOrder(f, list, worker)
		}
	}

	return pMapFloat64PreserveOrder(f, list, worker)
}

func pMapFloat64PreserveOrder(f func(float64) float64, list []float64, worker int) []float64 {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]float64) {
			defer wg.Done()

			for m := range chJobs {
				for k, v := range m {
					chResult <- map[int]float64{k: f(v)}
				}
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newListMap := make(map[int]float64, len(list))
	newList := make([]float64, len(list))

	for m := range chResult {
		for k, v := range m {
			newListMap[k] = v
		}
	}

	for i := 0; i < len(list); i++ {
		newList[i] = newListMap[i]
	}

	return newList
}

func pMapFloat64NoOrder(f func(float64) float64, list []float64, worker int) []float64 {
	chJobs := make(chan float64, len(list))
	go func() {
		for _, v := range list {
			chJobs <- v
		}
		close(chJobs)
	}()

	chResult := make(chan float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan float64, chJobs chan float64) {
			defer wg.Done()

			for v := range chJobs {
				chResult <- f(v)
			}
		}(chResult, chJobs)
	}

	// This will wait for the workers to complete their job and then close the channel
	go func() {
		wg.Wait()
		close(chResult)
	}()

	newList := make([]float64, len(list))
	i := 0

	for v := range chResult {
		newList[i] = v
		i++
	}

	return newList
}
