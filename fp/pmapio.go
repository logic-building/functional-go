package fp

import "sync"

// PMapIntInt64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt64(f func(int) int64, list []int, optional ...Optional) []int64 {
	if f == nil {
		return []int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt64NoOrder(f, list, worker)
		}
	}

	return pMapIntInt64PreserveOrder(f, list, worker)
}

func pMapIntInt64PreserveOrder(f func(int) int64, list []int, worker int) []int64 {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]int) {
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

func pMapIntInt64NoOrder(f func(int) int64, list []int, worker int) []int64 {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan int64, chJobs chan int) {
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

// PMapIntInt32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt32(f func(int) int32, list []int, optional ...Optional) []int32 {
	if f == nil {
		return []int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt32NoOrder(f, list, worker)
		}
	}

	return pMapIntInt32PreserveOrder(f, list, worker)
}

func pMapIntInt32PreserveOrder(f func(int) int32, list []int, worker int) []int32 {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]int) {
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

func pMapIntInt32NoOrder(f func(int) int32, list []int, worker int) []int32 {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan int32, chJobs chan int) {
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

// PMapIntInt16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt16(f func(int) int16, list []int, optional ...Optional) []int16 {
	if f == nil {
		return []int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt16NoOrder(f, list, worker)
		}
	}

	return pMapIntInt16PreserveOrder(f, list, worker)
}

func pMapIntInt16PreserveOrder(f func(int) int16, list []int, worker int) []int16 {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]int) {
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

func pMapIntInt16NoOrder(f func(int) int16, list []int, worker int) []int16 {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan int16, chJobs chan int) {
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

// PMapIntInt8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntInt8(f func(int) int8, list []int, optional ...Optional) []int8 {
	if f == nil {
		return []int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntInt8NoOrder(f, list, worker)
		}
	}

	return pMapIntInt8PreserveOrder(f, list, worker)
}

func pMapIntInt8PreserveOrder(f func(int) int8, list []int, worker int) []int8 {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]int) {
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

func pMapIntInt8NoOrder(f func(int) int8, list []int, worker int) []int8 {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan int8, chJobs chan int) {
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

// PMapIntUint applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint(f func(int) uint, list []int, optional ...Optional) []uint {
	if f == nil {
		return []uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUintNoOrder(f, list, worker)
		}
	}

	return pMapIntUintPreserveOrder(f, list, worker)
}

func pMapIntUintPreserveOrder(f func(int) uint, list []int, worker int) []uint {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]int) {
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

func pMapIntUintNoOrder(f func(int) uint, list []int, worker int) []uint {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan uint, chJobs chan int) {
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

// PMapIntUint64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint64(f func(int) uint64, list []int, optional ...Optional) []uint64 {
	if f == nil {
		return []uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint64NoOrder(f, list, worker)
		}
	}

	return pMapIntUint64PreserveOrder(f, list, worker)
}

func pMapIntUint64PreserveOrder(f func(int) uint64, list []int, worker int) []uint64 {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]int) {
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

func pMapIntUint64NoOrder(f func(int) uint64, list []int, worker int) []uint64 {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan uint64, chJobs chan int) {
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

// PMapIntUint32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint32(f func(int) uint32, list []int, optional ...Optional) []uint32 {
	if f == nil {
		return []uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint32NoOrder(f, list, worker)
		}
	}

	return pMapIntUint32PreserveOrder(f, list, worker)
}

func pMapIntUint32PreserveOrder(f func(int) uint32, list []int, worker int) []uint32 {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]int) {
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

func pMapIntUint32NoOrder(f func(int) uint32, list []int, worker int) []uint32 {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan uint32, chJobs chan int) {
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

// PMapIntUint16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint16(f func(int) uint16, list []int, optional ...Optional) []uint16 {
	if f == nil {
		return []uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint16NoOrder(f, list, worker)
		}
	}

	return pMapIntUint16PreserveOrder(f, list, worker)
}

func pMapIntUint16PreserveOrder(f func(int) uint16, list []int, worker int) []uint16 {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]int) {
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

func pMapIntUint16NoOrder(f func(int) uint16, list []int, worker int) []uint16 {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan uint16, chJobs chan int) {
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

// PMapIntUint8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntUint8(f func(int) uint8, list []int, optional ...Optional) []uint8 {
	if f == nil {
		return []uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntUint8NoOrder(f, list, worker)
		}
	}

	return pMapIntUint8PreserveOrder(f, list, worker)
}

func pMapIntUint8PreserveOrder(f func(int) uint8, list []int, worker int) []uint8 {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]int) {
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

func pMapIntUint8NoOrder(f func(int) uint8, list []int, worker int) []uint8 {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan uint8, chJobs chan int) {
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

// PMapIntStr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntStr(f func(int) string, list []int, optional ...Optional) []string {
	if f == nil {
		return []string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntStrNoOrder(f, list, worker)
		}
	}

	return pMapIntStrPreserveOrder(f, list, worker)
}

func pMapIntStrPreserveOrder(f func(int) string, list []int, worker int) []string {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]int) {
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

func pMapIntStrNoOrder(f func(int) string, list []int, worker int) []string {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan string, chJobs chan int) {
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

// PMapIntBool applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntBool(f func(int) bool, list []int, optional ...Optional) []bool {
	if f == nil {
		return []bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntBoolNoOrder(f, list, worker)
		}
	}

	return pMapIntBoolPreserveOrder(f, list, worker)
}

func pMapIntBoolPreserveOrder(f func(int) bool, list []int, worker int) []bool {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]int) {
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

func pMapIntBoolNoOrder(f func(int) bool, list []int, worker int) []bool {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan bool, chJobs chan int) {
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

// PMapIntFloat32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntFloat32(f func(int) float32, list []int, optional ...Optional) []float32 {
	if f == nil {
		return []float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntFloat32NoOrder(f, list, worker)
		}
	}

	return pMapIntFloat32PreserveOrder(f, list, worker)
}

func pMapIntFloat32PreserveOrder(f func(int) float32, list []int, worker int) []float32 {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]int) {
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

func pMapIntFloat32NoOrder(f func(int) float32, list []int, worker int) []float32 {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan float32, chJobs chan int) {
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

// PMapIntFloat64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapIntFloat64(f func(int) float64, list []int, optional ...Optional) []float64 {
	if f == nil {
		return []float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapIntFloat64NoOrder(f, list, worker)
		}
	}

	return pMapIntFloat64PreserveOrder(f, list, worker)
}

func pMapIntFloat64PreserveOrder(f func(int) float64, list []int, worker int) []float64 {
	chJobs := make(chan map[int]int, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]int) {
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

func pMapIntFloat64NoOrder(f func(int) float64, list []int, worker int) []float64 {
	chJobs := make(chan int, len(list))
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

		go func(chResult chan float64, chJobs chan int) {
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

// PMapInt64Int applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Int(f func(int64) int, list []int64, optional ...Optional) []int {
	if f == nil {
		return []int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64IntNoOrder(f, list, worker)
		}
	}

	return pMapInt64IntPreserveOrder(f, list, worker)
}

func pMapInt64IntPreserveOrder(f func(int64) int, list []int64, worker int) []int {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]int64) {
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

func pMapInt64IntNoOrder(f func(int64) int, list []int64, worker int) []int {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan int, chJobs chan int64) {
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

// PMapInt64Int32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Int32(f func(int64) int32, list []int64, optional ...Optional) []int32 {
	if f == nil {
		return []int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Int32NoOrder(f, list, worker)
		}
	}

	return pMapInt64Int32PreserveOrder(f, list, worker)
}

func pMapInt64Int32PreserveOrder(f func(int64) int32, list []int64, worker int) []int32 {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]int64) {
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

func pMapInt64Int32NoOrder(f func(int64) int32, list []int64, worker int) []int32 {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan int32, chJobs chan int64) {
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

// PMapInt64Int16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Int16(f func(int64) int16, list []int64, optional ...Optional) []int16 {
	if f == nil {
		return []int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Int16NoOrder(f, list, worker)
		}
	}

	return pMapInt64Int16PreserveOrder(f, list, worker)
}

func pMapInt64Int16PreserveOrder(f func(int64) int16, list []int64, worker int) []int16 {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]int64) {
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

func pMapInt64Int16NoOrder(f func(int64) int16, list []int64, worker int) []int16 {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan int16, chJobs chan int64) {
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

// PMapInt64Int8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Int8(f func(int64) int8, list []int64, optional ...Optional) []int8 {
	if f == nil {
		return []int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Int8NoOrder(f, list, worker)
		}
	}

	return pMapInt64Int8PreserveOrder(f, list, worker)
}

func pMapInt64Int8PreserveOrder(f func(int64) int8, list []int64, worker int) []int8 {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]int64) {
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

func pMapInt64Int8NoOrder(f func(int64) int8, list []int64, worker int) []int8 {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan int8, chJobs chan int64) {
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

// PMapInt64Uint applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint(f func(int64) uint, list []int64, optional ...Optional) []uint {
	if f == nil {
		return []uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64UintNoOrder(f, list, worker)
		}
	}

	return pMapInt64UintPreserveOrder(f, list, worker)
}

func pMapInt64UintPreserveOrder(f func(int64) uint, list []int64, worker int) []uint {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]int64) {
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

func pMapInt64UintNoOrder(f func(int64) uint, list []int64, worker int) []uint {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan uint, chJobs chan int64) {
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

// PMapInt64Uint64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint64(f func(int64) uint64, list []int64, optional ...Optional) []uint64 {
	if f == nil {
		return []uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint64NoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint64PreserveOrder(f, list, worker)
}

func pMapInt64Uint64PreserveOrder(f func(int64) uint64, list []int64, worker int) []uint64 {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]int64) {
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

func pMapInt64Uint64NoOrder(f func(int64) uint64, list []int64, worker int) []uint64 {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan uint64, chJobs chan int64) {
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

// PMapInt64Uint32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint32(f func(int64) uint32, list []int64, optional ...Optional) []uint32 {
	if f == nil {
		return []uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint32NoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint32PreserveOrder(f, list, worker)
}

func pMapInt64Uint32PreserveOrder(f func(int64) uint32, list []int64, worker int) []uint32 {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]int64) {
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

func pMapInt64Uint32NoOrder(f func(int64) uint32, list []int64, worker int) []uint32 {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan uint32, chJobs chan int64) {
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

// PMapInt64Uint16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint16(f func(int64) uint16, list []int64, optional ...Optional) []uint16 {
	if f == nil {
		return []uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint16NoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint16PreserveOrder(f, list, worker)
}

func pMapInt64Uint16PreserveOrder(f func(int64) uint16, list []int64, worker int) []uint16 {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]int64) {
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

func pMapInt64Uint16NoOrder(f func(int64) uint16, list []int64, worker int) []uint16 {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan uint16, chJobs chan int64) {
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

// PMapInt64Uint8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Uint8(f func(int64) uint8, list []int64, optional ...Optional) []uint8 {
	if f == nil {
		return []uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Uint8NoOrder(f, list, worker)
		}
	}

	return pMapInt64Uint8PreserveOrder(f, list, worker)
}

func pMapInt64Uint8PreserveOrder(f func(int64) uint8, list []int64, worker int) []uint8 {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]int64) {
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

func pMapInt64Uint8NoOrder(f func(int64) uint8, list []int64, worker int) []uint8 {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan uint8, chJobs chan int64) {
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

// PMapInt64Str applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Str(f func(int64) string, list []int64, optional ...Optional) []string {
	if f == nil {
		return []string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64StrNoOrder(f, list, worker)
		}
	}

	return pMapInt64StrPreserveOrder(f, list, worker)
}

func pMapInt64StrPreserveOrder(f func(int64) string, list []int64, worker int) []string {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]int64) {
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

func pMapInt64StrNoOrder(f func(int64) string, list []int64, worker int) []string {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan string, chJobs chan int64) {
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

// PMapInt64Bool applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Bool(f func(int64) bool, list []int64, optional ...Optional) []bool {
	if f == nil {
		return []bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64BoolNoOrder(f, list, worker)
		}
	}

	return pMapInt64BoolPreserveOrder(f, list, worker)
}

func pMapInt64BoolPreserveOrder(f func(int64) bool, list []int64, worker int) []bool {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]int64) {
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

func pMapInt64BoolNoOrder(f func(int64) bool, list []int64, worker int) []bool {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan bool, chJobs chan int64) {
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

// PMapInt64Float32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Float32(f func(int64) float32, list []int64, optional ...Optional) []float32 {
	if f == nil {
		return []float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Float32NoOrder(f, list, worker)
		}
	}

	return pMapInt64Float32PreserveOrder(f, list, worker)
}

func pMapInt64Float32PreserveOrder(f func(int64) float32, list []int64, worker int) []float32 {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]int64) {
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

func pMapInt64Float32NoOrder(f func(int64) float32, list []int64, worker int) []float32 {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan float32, chJobs chan int64) {
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

// PMapInt64Float64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt64Float64(f func(int64) float64, list []int64, optional ...Optional) []float64 {
	if f == nil {
		return []float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt64Float64NoOrder(f, list, worker)
		}
	}

	return pMapInt64Float64PreserveOrder(f, list, worker)
}

func pMapInt64Float64PreserveOrder(f func(int64) float64, list []int64, worker int) []float64 {
	chJobs := make(chan map[int]int64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]int64) {
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

func pMapInt64Float64NoOrder(f func(int64) float64, list []int64, worker int) []float64 {
	chJobs := make(chan int64, len(list))
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

		go func(chResult chan float64, chJobs chan int64) {
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

// PMapInt32Int applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Int(f func(int32) int, list []int32, optional ...Optional) []int {
	if f == nil {
		return []int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32IntNoOrder(f, list, worker)
		}
	}

	return pMapInt32IntPreserveOrder(f, list, worker)
}

func pMapInt32IntPreserveOrder(f func(int32) int, list []int32, worker int) []int {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]int32) {
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

func pMapInt32IntNoOrder(f func(int32) int, list []int32, worker int) []int {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan int, chJobs chan int32) {
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

// PMapInt32Int64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Int64(f func(int32) int64, list []int32, optional ...Optional) []int64 {
	if f == nil {
		return []int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Int64NoOrder(f, list, worker)
		}
	}

	return pMapInt32Int64PreserveOrder(f, list, worker)
}

func pMapInt32Int64PreserveOrder(f func(int32) int64, list []int32, worker int) []int64 {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]int32) {
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

func pMapInt32Int64NoOrder(f func(int32) int64, list []int32, worker int) []int64 {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan int64, chJobs chan int32) {
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

// PMapInt32Int16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Int16(f func(int32) int16, list []int32, optional ...Optional) []int16 {
	if f == nil {
		return []int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Int16NoOrder(f, list, worker)
		}
	}

	return pMapInt32Int16PreserveOrder(f, list, worker)
}

func pMapInt32Int16PreserveOrder(f func(int32) int16, list []int32, worker int) []int16 {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]int32) {
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

func pMapInt32Int16NoOrder(f func(int32) int16, list []int32, worker int) []int16 {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan int16, chJobs chan int32) {
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

// PMapInt32Int8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Int8(f func(int32) int8, list []int32, optional ...Optional) []int8 {
	if f == nil {
		return []int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Int8NoOrder(f, list, worker)
		}
	}

	return pMapInt32Int8PreserveOrder(f, list, worker)
}

func pMapInt32Int8PreserveOrder(f func(int32) int8, list []int32, worker int) []int8 {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]int32) {
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

func pMapInt32Int8NoOrder(f func(int32) int8, list []int32, worker int) []int8 {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan int8, chJobs chan int32) {
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

// PMapInt32Uint applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint(f func(int32) uint, list []int32, optional ...Optional) []uint {
	if f == nil {
		return []uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32UintNoOrder(f, list, worker)
		}
	}

	return pMapInt32UintPreserveOrder(f, list, worker)
}

func pMapInt32UintPreserveOrder(f func(int32) uint, list []int32, worker int) []uint {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]int32) {
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

func pMapInt32UintNoOrder(f func(int32) uint, list []int32, worker int) []uint {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan uint, chJobs chan int32) {
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

// PMapInt32Uint64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint64(f func(int32) uint64, list []int32, optional ...Optional) []uint64 {
	if f == nil {
		return []uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint64NoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint64PreserveOrder(f, list, worker)
}

func pMapInt32Uint64PreserveOrder(f func(int32) uint64, list []int32, worker int) []uint64 {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]int32) {
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

func pMapInt32Uint64NoOrder(f func(int32) uint64, list []int32, worker int) []uint64 {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan uint64, chJobs chan int32) {
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

// PMapInt32Uint32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint32(f func(int32) uint32, list []int32, optional ...Optional) []uint32 {
	if f == nil {
		return []uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint32NoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint32PreserveOrder(f, list, worker)
}

func pMapInt32Uint32PreserveOrder(f func(int32) uint32, list []int32, worker int) []uint32 {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]int32) {
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

func pMapInt32Uint32NoOrder(f func(int32) uint32, list []int32, worker int) []uint32 {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan uint32, chJobs chan int32) {
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

// PMapInt32Uint16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint16(f func(int32) uint16, list []int32, optional ...Optional) []uint16 {
	if f == nil {
		return []uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint16NoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint16PreserveOrder(f, list, worker)
}

func pMapInt32Uint16PreserveOrder(f func(int32) uint16, list []int32, worker int) []uint16 {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]int32) {
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

func pMapInt32Uint16NoOrder(f func(int32) uint16, list []int32, worker int) []uint16 {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan uint16, chJobs chan int32) {
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

// PMapInt32Uint8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Uint8(f func(int32) uint8, list []int32, optional ...Optional) []uint8 {
	if f == nil {
		return []uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Uint8NoOrder(f, list, worker)
		}
	}

	return pMapInt32Uint8PreserveOrder(f, list, worker)
}

func pMapInt32Uint8PreserveOrder(f func(int32) uint8, list []int32, worker int) []uint8 {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]int32) {
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

func pMapInt32Uint8NoOrder(f func(int32) uint8, list []int32, worker int) []uint8 {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan uint8, chJobs chan int32) {
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

// PMapInt32Str applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Str(f func(int32) string, list []int32, optional ...Optional) []string {
	if f == nil {
		return []string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32StrNoOrder(f, list, worker)
		}
	}

	return pMapInt32StrPreserveOrder(f, list, worker)
}

func pMapInt32StrPreserveOrder(f func(int32) string, list []int32, worker int) []string {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]int32) {
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

func pMapInt32StrNoOrder(f func(int32) string, list []int32, worker int) []string {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan string, chJobs chan int32) {
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

// PMapInt32Bool applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Bool(f func(int32) bool, list []int32, optional ...Optional) []bool {
	if f == nil {
		return []bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32BoolNoOrder(f, list, worker)
		}
	}

	return pMapInt32BoolPreserveOrder(f, list, worker)
}

func pMapInt32BoolPreserveOrder(f func(int32) bool, list []int32, worker int) []bool {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]int32) {
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

func pMapInt32BoolNoOrder(f func(int32) bool, list []int32, worker int) []bool {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan bool, chJobs chan int32) {
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

// PMapInt32Float32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Float32(f func(int32) float32, list []int32, optional ...Optional) []float32 {
	if f == nil {
		return []float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Float32NoOrder(f, list, worker)
		}
	}

	return pMapInt32Float32PreserveOrder(f, list, worker)
}

func pMapInt32Float32PreserveOrder(f func(int32) float32, list []int32, worker int) []float32 {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]int32) {
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

func pMapInt32Float32NoOrder(f func(int32) float32, list []int32, worker int) []float32 {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan float32, chJobs chan int32) {
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

// PMapInt32Float64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt32Float64(f func(int32) float64, list []int32, optional ...Optional) []float64 {
	if f == nil {
		return []float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt32Float64NoOrder(f, list, worker)
		}
	}

	return pMapInt32Float64PreserveOrder(f, list, worker)
}

func pMapInt32Float64PreserveOrder(f func(int32) float64, list []int32, worker int) []float64 {
	chJobs := make(chan map[int]int32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]int32) {
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

func pMapInt32Float64NoOrder(f func(int32) float64, list []int32, worker int) []float64 {
	chJobs := make(chan int32, len(list))
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

		go func(chResult chan float64, chJobs chan int32) {
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

// PMapInt16Int applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Int(f func(int16) int, list []int16, optional ...Optional) []int {
	if f == nil {
		return []int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16IntNoOrder(f, list, worker)
		}
	}

	return pMapInt16IntPreserveOrder(f, list, worker)
}

func pMapInt16IntPreserveOrder(f func(int16) int, list []int16, worker int) []int {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]int16) {
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

func pMapInt16IntNoOrder(f func(int16) int, list []int16, worker int) []int {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan int, chJobs chan int16) {
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

// PMapInt16Int64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Int64(f func(int16) int64, list []int16, optional ...Optional) []int64 {
	if f == nil {
		return []int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Int64NoOrder(f, list, worker)
		}
	}

	return pMapInt16Int64PreserveOrder(f, list, worker)
}

func pMapInt16Int64PreserveOrder(f func(int16) int64, list []int16, worker int) []int64 {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]int16) {
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

func pMapInt16Int64NoOrder(f func(int16) int64, list []int16, worker int) []int64 {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan int64, chJobs chan int16) {
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

// PMapInt16Int32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Int32(f func(int16) int32, list []int16, optional ...Optional) []int32 {
	if f == nil {
		return []int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Int32NoOrder(f, list, worker)
		}
	}

	return pMapInt16Int32PreserveOrder(f, list, worker)
}

func pMapInt16Int32PreserveOrder(f func(int16) int32, list []int16, worker int) []int32 {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]int16) {
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

func pMapInt16Int32NoOrder(f func(int16) int32, list []int16, worker int) []int32 {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan int32, chJobs chan int16) {
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

// PMapInt16Int8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Int8(f func(int16) int8, list []int16, optional ...Optional) []int8 {
	if f == nil {
		return []int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Int8NoOrder(f, list, worker)
		}
	}

	return pMapInt16Int8PreserveOrder(f, list, worker)
}

func pMapInt16Int8PreserveOrder(f func(int16) int8, list []int16, worker int) []int8 {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]int16) {
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

func pMapInt16Int8NoOrder(f func(int16) int8, list []int16, worker int) []int8 {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan int8, chJobs chan int16) {
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

// PMapInt16Uint applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint(f func(int16) uint, list []int16, optional ...Optional) []uint {
	if f == nil {
		return []uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16UintNoOrder(f, list, worker)
		}
	}

	return pMapInt16UintPreserveOrder(f, list, worker)
}

func pMapInt16UintPreserveOrder(f func(int16) uint, list []int16, worker int) []uint {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]int16) {
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

func pMapInt16UintNoOrder(f func(int16) uint, list []int16, worker int) []uint {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan uint, chJobs chan int16) {
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

// PMapInt16Uint64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint64(f func(int16) uint64, list []int16, optional ...Optional) []uint64 {
	if f == nil {
		return []uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint64NoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint64PreserveOrder(f, list, worker)
}

func pMapInt16Uint64PreserveOrder(f func(int16) uint64, list []int16, worker int) []uint64 {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]int16) {
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

func pMapInt16Uint64NoOrder(f func(int16) uint64, list []int16, worker int) []uint64 {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan uint64, chJobs chan int16) {
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

// PMapInt16Uint32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint32(f func(int16) uint32, list []int16, optional ...Optional) []uint32 {
	if f == nil {
		return []uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint32NoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint32PreserveOrder(f, list, worker)
}

func pMapInt16Uint32PreserveOrder(f func(int16) uint32, list []int16, worker int) []uint32 {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]int16) {
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

func pMapInt16Uint32NoOrder(f func(int16) uint32, list []int16, worker int) []uint32 {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan uint32, chJobs chan int16) {
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

// PMapInt16Uint16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint16(f func(int16) uint16, list []int16, optional ...Optional) []uint16 {
	if f == nil {
		return []uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint16NoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint16PreserveOrder(f, list, worker)
}

func pMapInt16Uint16PreserveOrder(f func(int16) uint16, list []int16, worker int) []uint16 {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]int16) {
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

func pMapInt16Uint16NoOrder(f func(int16) uint16, list []int16, worker int) []uint16 {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan uint16, chJobs chan int16) {
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

// PMapInt16Uint8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Uint8(f func(int16) uint8, list []int16, optional ...Optional) []uint8 {
	if f == nil {
		return []uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Uint8NoOrder(f, list, worker)
		}
	}

	return pMapInt16Uint8PreserveOrder(f, list, worker)
}

func pMapInt16Uint8PreserveOrder(f func(int16) uint8, list []int16, worker int) []uint8 {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]int16) {
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

func pMapInt16Uint8NoOrder(f func(int16) uint8, list []int16, worker int) []uint8 {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan uint8, chJobs chan int16) {
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

// PMapInt16Str applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Str(f func(int16) string, list []int16, optional ...Optional) []string {
	if f == nil {
		return []string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16StrNoOrder(f, list, worker)
		}
	}

	return pMapInt16StrPreserveOrder(f, list, worker)
}

func pMapInt16StrPreserveOrder(f func(int16) string, list []int16, worker int) []string {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]int16) {
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

func pMapInt16StrNoOrder(f func(int16) string, list []int16, worker int) []string {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan string, chJobs chan int16) {
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

// PMapInt16Bool applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Bool(f func(int16) bool, list []int16, optional ...Optional) []bool {
	if f == nil {
		return []bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16BoolNoOrder(f, list, worker)
		}
	}

	return pMapInt16BoolPreserveOrder(f, list, worker)
}

func pMapInt16BoolPreserveOrder(f func(int16) bool, list []int16, worker int) []bool {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]int16) {
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

func pMapInt16BoolNoOrder(f func(int16) bool, list []int16, worker int) []bool {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan bool, chJobs chan int16) {
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

// PMapInt16Float32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Float32(f func(int16) float32, list []int16, optional ...Optional) []float32 {
	if f == nil {
		return []float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Float32NoOrder(f, list, worker)
		}
	}

	return pMapInt16Float32PreserveOrder(f, list, worker)
}

func pMapInt16Float32PreserveOrder(f func(int16) float32, list []int16, worker int) []float32 {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]int16) {
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

func pMapInt16Float32NoOrder(f func(int16) float32, list []int16, worker int) []float32 {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan float32, chJobs chan int16) {
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

// PMapInt16Float64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt16Float64(f func(int16) float64, list []int16, optional ...Optional) []float64 {
	if f == nil {
		return []float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt16Float64NoOrder(f, list, worker)
		}
	}

	return pMapInt16Float64PreserveOrder(f, list, worker)
}

func pMapInt16Float64PreserveOrder(f func(int16) float64, list []int16, worker int) []float64 {
	chJobs := make(chan map[int]int16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]int16) {
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

func pMapInt16Float64NoOrder(f func(int16) float64, list []int16, worker int) []float64 {
	chJobs := make(chan int16, len(list))
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

		go func(chResult chan float64, chJobs chan int16) {
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

// PMapInt8Int applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Int(f func(int8) int, list []int8, optional ...Optional) []int {
	if f == nil {
		return []int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8IntNoOrder(f, list, worker)
		}
	}

	return pMapInt8IntPreserveOrder(f, list, worker)
}

func pMapInt8IntPreserveOrder(f func(int8) int, list []int8, worker int) []int {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]int8) {
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

func pMapInt8IntNoOrder(f func(int8) int, list []int8, worker int) []int {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan int, chJobs chan int8) {
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

// PMapInt8Int64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Int64(f func(int8) int64, list []int8, optional ...Optional) []int64 {
	if f == nil {
		return []int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Int64NoOrder(f, list, worker)
		}
	}

	return pMapInt8Int64PreserveOrder(f, list, worker)
}

func pMapInt8Int64PreserveOrder(f func(int8) int64, list []int8, worker int) []int64 {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]int8) {
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

func pMapInt8Int64NoOrder(f func(int8) int64, list []int8, worker int) []int64 {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan int64, chJobs chan int8) {
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

// PMapInt8Int32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Int32(f func(int8) int32, list []int8, optional ...Optional) []int32 {
	if f == nil {
		return []int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Int32NoOrder(f, list, worker)
		}
	}

	return pMapInt8Int32PreserveOrder(f, list, worker)
}

func pMapInt8Int32PreserveOrder(f func(int8) int32, list []int8, worker int) []int32 {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]int8) {
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

func pMapInt8Int32NoOrder(f func(int8) int32, list []int8, worker int) []int32 {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan int32, chJobs chan int8) {
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

// PMapInt8Int16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Int16(f func(int8) int16, list []int8, optional ...Optional) []int16 {
	if f == nil {
		return []int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Int16NoOrder(f, list, worker)
		}
	}

	return pMapInt8Int16PreserveOrder(f, list, worker)
}

func pMapInt8Int16PreserveOrder(f func(int8) int16, list []int8, worker int) []int16 {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]int8) {
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

func pMapInt8Int16NoOrder(f func(int8) int16, list []int8, worker int) []int16 {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan int16, chJobs chan int8) {
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

// PMapInt8Uint applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint(f func(int8) uint, list []int8, optional ...Optional) []uint {
	if f == nil {
		return []uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8UintNoOrder(f, list, worker)
		}
	}

	return pMapInt8UintPreserveOrder(f, list, worker)
}

func pMapInt8UintPreserveOrder(f func(int8) uint, list []int8, worker int) []uint {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]int8) {
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

func pMapInt8UintNoOrder(f func(int8) uint, list []int8, worker int) []uint {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan uint, chJobs chan int8) {
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

// PMapInt8Uint64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint64(f func(int8) uint64, list []int8, optional ...Optional) []uint64 {
	if f == nil {
		return []uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint64NoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint64PreserveOrder(f, list, worker)
}

func pMapInt8Uint64PreserveOrder(f func(int8) uint64, list []int8, worker int) []uint64 {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]int8) {
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

func pMapInt8Uint64NoOrder(f func(int8) uint64, list []int8, worker int) []uint64 {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan uint64, chJobs chan int8) {
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

// PMapInt8Uint32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint32(f func(int8) uint32, list []int8, optional ...Optional) []uint32 {
	if f == nil {
		return []uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint32NoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint32PreserveOrder(f, list, worker)
}

func pMapInt8Uint32PreserveOrder(f func(int8) uint32, list []int8, worker int) []uint32 {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]int8) {
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

func pMapInt8Uint32NoOrder(f func(int8) uint32, list []int8, worker int) []uint32 {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan uint32, chJobs chan int8) {
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

// PMapInt8Uint16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint16(f func(int8) uint16, list []int8, optional ...Optional) []uint16 {
	if f == nil {
		return []uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint16NoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint16PreserveOrder(f, list, worker)
}

func pMapInt8Uint16PreserveOrder(f func(int8) uint16, list []int8, worker int) []uint16 {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]int8) {
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

func pMapInt8Uint16NoOrder(f func(int8) uint16, list []int8, worker int) []uint16 {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan uint16, chJobs chan int8) {
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

// PMapInt8Uint8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Uint8(f func(int8) uint8, list []int8, optional ...Optional) []uint8 {
	if f == nil {
		return []uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Uint8NoOrder(f, list, worker)
		}
	}

	return pMapInt8Uint8PreserveOrder(f, list, worker)
}

func pMapInt8Uint8PreserveOrder(f func(int8) uint8, list []int8, worker int) []uint8 {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]int8) {
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

func pMapInt8Uint8NoOrder(f func(int8) uint8, list []int8, worker int) []uint8 {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan uint8, chJobs chan int8) {
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

// PMapInt8Str applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Str(f func(int8) string, list []int8, optional ...Optional) []string {
	if f == nil {
		return []string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8StrNoOrder(f, list, worker)
		}
	}

	return pMapInt8StrPreserveOrder(f, list, worker)
}

func pMapInt8StrPreserveOrder(f func(int8) string, list []int8, worker int) []string {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]int8) {
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

func pMapInt8StrNoOrder(f func(int8) string, list []int8, worker int) []string {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan string, chJobs chan int8) {
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

// PMapInt8Bool applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Bool(f func(int8) bool, list []int8, optional ...Optional) []bool {
	if f == nil {
		return []bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8BoolNoOrder(f, list, worker)
		}
	}

	return pMapInt8BoolPreserveOrder(f, list, worker)
}

func pMapInt8BoolPreserveOrder(f func(int8) bool, list []int8, worker int) []bool {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]int8) {
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

func pMapInt8BoolNoOrder(f func(int8) bool, list []int8, worker int) []bool {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan bool, chJobs chan int8) {
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

// PMapInt8Float32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Float32(f func(int8) float32, list []int8, optional ...Optional) []float32 {
	if f == nil {
		return []float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Float32NoOrder(f, list, worker)
		}
	}

	return pMapInt8Float32PreserveOrder(f, list, worker)
}

func pMapInt8Float32PreserveOrder(f func(int8) float32, list []int8, worker int) []float32 {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]int8) {
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

func pMapInt8Float32NoOrder(f func(int8) float32, list []int8, worker int) []float32 {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan float32, chJobs chan int8) {
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

// PMapInt8Float64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapInt8Float64(f func(int8) float64, list []int8, optional ...Optional) []float64 {
	if f == nil {
		return []float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapInt8Float64NoOrder(f, list, worker)
		}
	}

	return pMapInt8Float64PreserveOrder(f, list, worker)
}

func pMapInt8Float64PreserveOrder(f func(int8) float64, list []int8, worker int) []float64 {
	chJobs := make(chan map[int]int8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]int8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]int8) {
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

func pMapInt8Float64NoOrder(f func(int8) float64, list []int8, worker int) []float64 {
	chJobs := make(chan int8, len(list))
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

		go func(chResult chan float64, chJobs chan int8) {
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

// PMapUintInt applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt(f func(uint) int, list []uint, optional ...Optional) []int {
	if f == nil {
		return []int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintIntNoOrder(f, list, worker)
		}
	}

	return pMapUintIntPreserveOrder(f, list, worker)
}

func pMapUintIntPreserveOrder(f func(uint) int, list []uint, worker int) []int {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]uint) {
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

func pMapUintIntNoOrder(f func(uint) int, list []uint, worker int) []int {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan int, chJobs chan uint) {
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

// PMapUintInt64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt64(f func(uint) int64, list []uint, optional ...Optional) []int64 {
	if f == nil {
		return []int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt64NoOrder(f, list, worker)
		}
	}

	return pMapUintInt64PreserveOrder(f, list, worker)
}

func pMapUintInt64PreserveOrder(f func(uint) int64, list []uint, worker int) []int64 {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]uint) {
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

func pMapUintInt64NoOrder(f func(uint) int64, list []uint, worker int) []int64 {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan int64, chJobs chan uint) {
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

// PMapUintInt32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt32(f func(uint) int32, list []uint, optional ...Optional) []int32 {
	if f == nil {
		return []int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt32NoOrder(f, list, worker)
		}
	}

	return pMapUintInt32PreserveOrder(f, list, worker)
}

func pMapUintInt32PreserveOrder(f func(uint) int32, list []uint, worker int) []int32 {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]uint) {
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

func pMapUintInt32NoOrder(f func(uint) int32, list []uint, worker int) []int32 {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan int32, chJobs chan uint) {
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

// PMapUintInt16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt16(f func(uint) int16, list []uint, optional ...Optional) []int16 {
	if f == nil {
		return []int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt16NoOrder(f, list, worker)
		}
	}

	return pMapUintInt16PreserveOrder(f, list, worker)
}

func pMapUintInt16PreserveOrder(f func(uint) int16, list []uint, worker int) []int16 {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]uint) {
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

func pMapUintInt16NoOrder(f func(uint) int16, list []uint, worker int) []int16 {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan int16, chJobs chan uint) {
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

// PMapUintInt8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintInt8(f func(uint) int8, list []uint, optional ...Optional) []int8 {
	if f == nil {
		return []int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintInt8NoOrder(f, list, worker)
		}
	}

	return pMapUintInt8PreserveOrder(f, list, worker)
}

func pMapUintInt8PreserveOrder(f func(uint) int8, list []uint, worker int) []int8 {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]uint) {
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

func pMapUintInt8NoOrder(f func(uint) int8, list []uint, worker int) []int8 {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan int8, chJobs chan uint) {
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

// PMapUintUint64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint64(f func(uint) uint64, list []uint, optional ...Optional) []uint64 {
	if f == nil {
		return []uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint64NoOrder(f, list, worker)
		}
	}

	return pMapUintUint64PreserveOrder(f, list, worker)
}

func pMapUintUint64PreserveOrder(f func(uint) uint64, list []uint, worker int) []uint64 {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]uint) {
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

func pMapUintUint64NoOrder(f func(uint) uint64, list []uint, worker int) []uint64 {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan uint64, chJobs chan uint) {
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

// PMapUintUint32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint32(f func(uint) uint32, list []uint, optional ...Optional) []uint32 {
	if f == nil {
		return []uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint32NoOrder(f, list, worker)
		}
	}

	return pMapUintUint32PreserveOrder(f, list, worker)
}

func pMapUintUint32PreserveOrder(f func(uint) uint32, list []uint, worker int) []uint32 {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]uint) {
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

func pMapUintUint32NoOrder(f func(uint) uint32, list []uint, worker int) []uint32 {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan uint32, chJobs chan uint) {
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

// PMapUintUint16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint16(f func(uint) uint16, list []uint, optional ...Optional) []uint16 {
	if f == nil {
		return []uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint16NoOrder(f, list, worker)
		}
	}

	return pMapUintUint16PreserveOrder(f, list, worker)
}

func pMapUintUint16PreserveOrder(f func(uint) uint16, list []uint, worker int) []uint16 {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]uint) {
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

func pMapUintUint16NoOrder(f func(uint) uint16, list []uint, worker int) []uint16 {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan uint16, chJobs chan uint) {
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

// PMapUintUint8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintUint8(f func(uint) uint8, list []uint, optional ...Optional) []uint8 {
	if f == nil {
		return []uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintUint8NoOrder(f, list, worker)
		}
	}

	return pMapUintUint8PreserveOrder(f, list, worker)
}

func pMapUintUint8PreserveOrder(f func(uint) uint8, list []uint, worker int) []uint8 {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]uint) {
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

func pMapUintUint8NoOrder(f func(uint) uint8, list []uint, worker int) []uint8 {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan uint8, chJobs chan uint) {
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

// PMapUintStr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintStr(f func(uint) string, list []uint, optional ...Optional) []string {
	if f == nil {
		return []string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintStrNoOrder(f, list, worker)
		}
	}

	return pMapUintStrPreserveOrder(f, list, worker)
}

func pMapUintStrPreserveOrder(f func(uint) string, list []uint, worker int) []string {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]uint) {
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

func pMapUintStrNoOrder(f func(uint) string, list []uint, worker int) []string {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan string, chJobs chan uint) {
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

// PMapUintBool applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintBool(f func(uint) bool, list []uint, optional ...Optional) []bool {
	if f == nil {
		return []bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintBoolNoOrder(f, list, worker)
		}
	}

	return pMapUintBoolPreserveOrder(f, list, worker)
}

func pMapUintBoolPreserveOrder(f func(uint) bool, list []uint, worker int) []bool {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]uint) {
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

func pMapUintBoolNoOrder(f func(uint) bool, list []uint, worker int) []bool {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan bool, chJobs chan uint) {
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

// PMapUintFloat32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintFloat32(f func(uint) float32, list []uint, optional ...Optional) []float32 {
	if f == nil {
		return []float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintFloat32NoOrder(f, list, worker)
		}
	}

	return pMapUintFloat32PreserveOrder(f, list, worker)
}

func pMapUintFloat32PreserveOrder(f func(uint) float32, list []uint, worker int) []float32 {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]uint) {
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

func pMapUintFloat32NoOrder(f func(uint) float32, list []uint, worker int) []float32 {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan float32, chJobs chan uint) {
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

// PMapUintFloat64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUintFloat64(f func(uint) float64, list []uint, optional ...Optional) []float64 {
	if f == nil {
		return []float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUintFloat64NoOrder(f, list, worker)
		}
	}

	return pMapUintFloat64PreserveOrder(f, list, worker)
}

func pMapUintFloat64PreserveOrder(f func(uint) float64, list []uint, worker int) []float64 {
	chJobs := make(chan map[int]uint, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]uint) {
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

func pMapUintFloat64NoOrder(f func(uint) float64, list []uint, worker int) []float64 {
	chJobs := make(chan uint, len(list))
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

		go func(chResult chan float64, chJobs chan uint) {
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

// PMapUint64Int applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int(f func(uint64) int, list []uint64, optional ...Optional) []int {
	if f == nil {
		return []int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64IntNoOrder(f, list, worker)
		}
	}

	return pMapUint64IntPreserveOrder(f, list, worker)
}

func pMapUint64IntPreserveOrder(f func(uint64) int, list []uint64, worker int) []int {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]uint64) {
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

func pMapUint64IntNoOrder(f func(uint64) int, list []uint64, worker int) []int {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan int, chJobs chan uint64) {
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

// PMapUint64Int64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int64(f func(uint64) int64, list []uint64, optional ...Optional) []int64 {
	if f == nil {
		return []int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int64NoOrder(f, list, worker)
		}
	}

	return pMapUint64Int64PreserveOrder(f, list, worker)
}

func pMapUint64Int64PreserveOrder(f func(uint64) int64, list []uint64, worker int) []int64 {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]uint64) {
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

func pMapUint64Int64NoOrder(f func(uint64) int64, list []uint64, worker int) []int64 {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan int64, chJobs chan uint64) {
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

// PMapUint64Int32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int32(f func(uint64) int32, list []uint64, optional ...Optional) []int32 {
	if f == nil {
		return []int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int32NoOrder(f, list, worker)
		}
	}

	return pMapUint64Int32PreserveOrder(f, list, worker)
}

func pMapUint64Int32PreserveOrder(f func(uint64) int32, list []uint64, worker int) []int32 {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]uint64) {
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

func pMapUint64Int32NoOrder(f func(uint64) int32, list []uint64, worker int) []int32 {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan int32, chJobs chan uint64) {
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

// PMapUint64Int16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int16(f func(uint64) int16, list []uint64, optional ...Optional) []int16 {
	if f == nil {
		return []int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int16NoOrder(f, list, worker)
		}
	}

	return pMapUint64Int16PreserveOrder(f, list, worker)
}

func pMapUint64Int16PreserveOrder(f func(uint64) int16, list []uint64, worker int) []int16 {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]uint64) {
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

func pMapUint64Int16NoOrder(f func(uint64) int16, list []uint64, worker int) []int16 {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan int16, chJobs chan uint64) {
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

// PMapUint64Int8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Int8(f func(uint64) int8, list []uint64, optional ...Optional) []int8 {
	if f == nil {
		return []int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Int8NoOrder(f, list, worker)
		}
	}

	return pMapUint64Int8PreserveOrder(f, list, worker)
}

func pMapUint64Int8PreserveOrder(f func(uint64) int8, list []uint64, worker int) []int8 {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]uint64) {
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

func pMapUint64Int8NoOrder(f func(uint64) int8, list []uint64, worker int) []int8 {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan int8, chJobs chan uint64) {
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

// PMapUint64Uint applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Uint(f func(uint64) uint, list []uint64, optional ...Optional) []uint {
	if f == nil {
		return []uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64UintNoOrder(f, list, worker)
		}
	}

	return pMapUint64UintPreserveOrder(f, list, worker)
}

func pMapUint64UintPreserveOrder(f func(uint64) uint, list []uint64, worker int) []uint {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]uint64) {
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

func pMapUint64UintNoOrder(f func(uint64) uint, list []uint64, worker int) []uint {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan uint, chJobs chan uint64) {
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

// PMapUint64Uint32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Uint32(f func(uint64) uint32, list []uint64, optional ...Optional) []uint32 {
	if f == nil {
		return []uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Uint32NoOrder(f, list, worker)
		}
	}

	return pMapUint64Uint32PreserveOrder(f, list, worker)
}

func pMapUint64Uint32PreserveOrder(f func(uint64) uint32, list []uint64, worker int) []uint32 {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]uint64) {
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

func pMapUint64Uint32NoOrder(f func(uint64) uint32, list []uint64, worker int) []uint32 {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan uint32, chJobs chan uint64) {
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

// PMapUint64Uint16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Uint16(f func(uint64) uint16, list []uint64, optional ...Optional) []uint16 {
	if f == nil {
		return []uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Uint16NoOrder(f, list, worker)
		}
	}

	return pMapUint64Uint16PreserveOrder(f, list, worker)
}

func pMapUint64Uint16PreserveOrder(f func(uint64) uint16, list []uint64, worker int) []uint16 {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]uint64) {
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

func pMapUint64Uint16NoOrder(f func(uint64) uint16, list []uint64, worker int) []uint16 {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan uint16, chJobs chan uint64) {
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

// PMapUint64Uint8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Uint8(f func(uint64) uint8, list []uint64, optional ...Optional) []uint8 {
	if f == nil {
		return []uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Uint8NoOrder(f, list, worker)
		}
	}

	return pMapUint64Uint8PreserveOrder(f, list, worker)
}

func pMapUint64Uint8PreserveOrder(f func(uint64) uint8, list []uint64, worker int) []uint8 {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]uint64) {
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

func pMapUint64Uint8NoOrder(f func(uint64) uint8, list []uint64, worker int) []uint8 {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan uint8, chJobs chan uint64) {
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

// PMapUint64Str applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Str(f func(uint64) string, list []uint64, optional ...Optional) []string {
	if f == nil {
		return []string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64StrNoOrder(f, list, worker)
		}
	}

	return pMapUint64StrPreserveOrder(f, list, worker)
}

func pMapUint64StrPreserveOrder(f func(uint64) string, list []uint64, worker int) []string {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]uint64) {
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

func pMapUint64StrNoOrder(f func(uint64) string, list []uint64, worker int) []string {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan string, chJobs chan uint64) {
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

// PMapUint64Bool applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Bool(f func(uint64) bool, list []uint64, optional ...Optional) []bool {
	if f == nil {
		return []bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64BoolNoOrder(f, list, worker)
		}
	}

	return pMapUint64BoolPreserveOrder(f, list, worker)
}

func pMapUint64BoolPreserveOrder(f func(uint64) bool, list []uint64, worker int) []bool {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]uint64) {
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

func pMapUint64BoolNoOrder(f func(uint64) bool, list []uint64, worker int) []bool {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan bool, chJobs chan uint64) {
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

// PMapUint64Float32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Float32(f func(uint64) float32, list []uint64, optional ...Optional) []float32 {
	if f == nil {
		return []float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Float32NoOrder(f, list, worker)
		}
	}

	return pMapUint64Float32PreserveOrder(f, list, worker)
}

func pMapUint64Float32PreserveOrder(f func(uint64) float32, list []uint64, worker int) []float32 {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]uint64) {
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

func pMapUint64Float32NoOrder(f func(uint64) float32, list []uint64, worker int) []float32 {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan float32, chJobs chan uint64) {
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

// PMapUint64Float64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint64Float64(f func(uint64) float64, list []uint64, optional ...Optional) []float64 {
	if f == nil {
		return []float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint64Float64NoOrder(f, list, worker)
		}
	}

	return pMapUint64Float64PreserveOrder(f, list, worker)
}

func pMapUint64Float64PreserveOrder(f func(uint64) float64, list []uint64, worker int) []float64 {
	chJobs := make(chan map[int]uint64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]uint64) {
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

func pMapUint64Float64NoOrder(f func(uint64) float64, list []uint64, worker int) []float64 {
	chJobs := make(chan uint64, len(list))
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

		go func(chResult chan float64, chJobs chan uint64) {
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

// PMapUint32Int applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int(f func(uint32) int, list []uint32, optional ...Optional) []int {
	if f == nil {
		return []int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32IntNoOrder(f, list, worker)
		}
	}

	return pMapUint32IntPreserveOrder(f, list, worker)
}

func pMapUint32IntPreserveOrder(f func(uint32) int, list []uint32, worker int) []int {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]uint32) {
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

func pMapUint32IntNoOrder(f func(uint32) int, list []uint32, worker int) []int {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan int, chJobs chan uint32) {
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

// PMapUint32Int64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int64(f func(uint32) int64, list []uint32, optional ...Optional) []int64 {
	if f == nil {
		return []int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int64NoOrder(f, list, worker)
		}
	}

	return pMapUint32Int64PreserveOrder(f, list, worker)
}

func pMapUint32Int64PreserveOrder(f func(uint32) int64, list []uint32, worker int) []int64 {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]uint32) {
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

func pMapUint32Int64NoOrder(f func(uint32) int64, list []uint32, worker int) []int64 {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan int64, chJobs chan uint32) {
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

// PMapUint32Int32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int32(f func(uint32) int32, list []uint32, optional ...Optional) []int32 {
	if f == nil {
		return []int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int32NoOrder(f, list, worker)
		}
	}

	return pMapUint32Int32PreserveOrder(f, list, worker)
}

func pMapUint32Int32PreserveOrder(f func(uint32) int32, list []uint32, worker int) []int32 {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]uint32) {
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

func pMapUint32Int32NoOrder(f func(uint32) int32, list []uint32, worker int) []int32 {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan int32, chJobs chan uint32) {
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

// PMapUint32Int16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int16(f func(uint32) int16, list []uint32, optional ...Optional) []int16 {
	if f == nil {
		return []int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int16NoOrder(f, list, worker)
		}
	}

	return pMapUint32Int16PreserveOrder(f, list, worker)
}

func pMapUint32Int16PreserveOrder(f func(uint32) int16, list []uint32, worker int) []int16 {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]uint32) {
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

func pMapUint32Int16NoOrder(f func(uint32) int16, list []uint32, worker int) []int16 {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan int16, chJobs chan uint32) {
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

// PMapUint32Int8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Int8(f func(uint32) int8, list []uint32, optional ...Optional) []int8 {
	if f == nil {
		return []int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Int8NoOrder(f, list, worker)
		}
	}

	return pMapUint32Int8PreserveOrder(f, list, worker)
}

func pMapUint32Int8PreserveOrder(f func(uint32) int8, list []uint32, worker int) []int8 {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]uint32) {
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

func pMapUint32Int8NoOrder(f func(uint32) int8, list []uint32, worker int) []int8 {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan int8, chJobs chan uint32) {
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

// PMapUint32Uint applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Uint(f func(uint32) uint, list []uint32, optional ...Optional) []uint {
	if f == nil {
		return []uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32UintNoOrder(f, list, worker)
		}
	}

	return pMapUint32UintPreserveOrder(f, list, worker)
}

func pMapUint32UintPreserveOrder(f func(uint32) uint, list []uint32, worker int) []uint {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]uint32) {
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

func pMapUint32UintNoOrder(f func(uint32) uint, list []uint32, worker int) []uint {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan uint, chJobs chan uint32) {
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

// PMapUint32Uint64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Uint64(f func(uint32) uint64, list []uint32, optional ...Optional) []uint64 {
	if f == nil {
		return []uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Uint64NoOrder(f, list, worker)
		}
	}

	return pMapUint32Uint64PreserveOrder(f, list, worker)
}

func pMapUint32Uint64PreserveOrder(f func(uint32) uint64, list []uint32, worker int) []uint64 {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]uint32) {
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

func pMapUint32Uint64NoOrder(f func(uint32) uint64, list []uint32, worker int) []uint64 {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan uint64, chJobs chan uint32) {
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

// PMapUint32Uint16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Uint16(f func(uint32) uint16, list []uint32, optional ...Optional) []uint16 {
	if f == nil {
		return []uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Uint16NoOrder(f, list, worker)
		}
	}

	return pMapUint32Uint16PreserveOrder(f, list, worker)
}

func pMapUint32Uint16PreserveOrder(f func(uint32) uint16, list []uint32, worker int) []uint16 {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]uint32) {
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

func pMapUint32Uint16NoOrder(f func(uint32) uint16, list []uint32, worker int) []uint16 {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan uint16, chJobs chan uint32) {
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

// PMapUint32Uint8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Uint8(f func(uint32) uint8, list []uint32, optional ...Optional) []uint8 {
	if f == nil {
		return []uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Uint8NoOrder(f, list, worker)
		}
	}

	return pMapUint32Uint8PreserveOrder(f, list, worker)
}

func pMapUint32Uint8PreserveOrder(f func(uint32) uint8, list []uint32, worker int) []uint8 {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]uint32) {
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

func pMapUint32Uint8NoOrder(f func(uint32) uint8, list []uint32, worker int) []uint8 {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan uint8, chJobs chan uint32) {
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

// PMapUint32Str applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Str(f func(uint32) string, list []uint32, optional ...Optional) []string {
	if f == nil {
		return []string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32StrNoOrder(f, list, worker)
		}
	}

	return pMapUint32StrPreserveOrder(f, list, worker)
}

func pMapUint32StrPreserveOrder(f func(uint32) string, list []uint32, worker int) []string {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]uint32) {
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

func pMapUint32StrNoOrder(f func(uint32) string, list []uint32, worker int) []string {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan string, chJobs chan uint32) {
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

// PMapUint32Bool applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Bool(f func(uint32) bool, list []uint32, optional ...Optional) []bool {
	if f == nil {
		return []bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32BoolNoOrder(f, list, worker)
		}
	}

	return pMapUint32BoolPreserveOrder(f, list, worker)
}

func pMapUint32BoolPreserveOrder(f func(uint32) bool, list []uint32, worker int) []bool {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]uint32) {
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

func pMapUint32BoolNoOrder(f func(uint32) bool, list []uint32, worker int) []bool {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan bool, chJobs chan uint32) {
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

// PMapUint32Float32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Float32(f func(uint32) float32, list []uint32, optional ...Optional) []float32 {
	if f == nil {
		return []float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Float32NoOrder(f, list, worker)
		}
	}

	return pMapUint32Float32PreserveOrder(f, list, worker)
}

func pMapUint32Float32PreserveOrder(f func(uint32) float32, list []uint32, worker int) []float32 {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]uint32) {
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

func pMapUint32Float32NoOrder(f func(uint32) float32, list []uint32, worker int) []float32 {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan float32, chJobs chan uint32) {
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

// PMapUint32Float64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint32Float64(f func(uint32) float64, list []uint32, optional ...Optional) []float64 {
	if f == nil {
		return []float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint32Float64NoOrder(f, list, worker)
		}
	}

	return pMapUint32Float64PreserveOrder(f, list, worker)
}

func pMapUint32Float64PreserveOrder(f func(uint32) float64, list []uint32, worker int) []float64 {
	chJobs := make(chan map[int]uint32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]uint32) {
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

func pMapUint32Float64NoOrder(f func(uint32) float64, list []uint32, worker int) []float64 {
	chJobs := make(chan uint32, len(list))
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

		go func(chResult chan float64, chJobs chan uint32) {
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

// PMapUint16Int applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int(f func(uint16) int, list []uint16, optional ...Optional) []int {
	if f == nil {
		return []int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16IntNoOrder(f, list, worker)
		}
	}

	return pMapUint16IntPreserveOrder(f, list, worker)
}

func pMapUint16IntPreserveOrder(f func(uint16) int, list []uint16, worker int) []int {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]uint16) {
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

func pMapUint16IntNoOrder(f func(uint16) int, list []uint16, worker int) []int {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan int, chJobs chan uint16) {
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

// PMapUint16Int64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int64(f func(uint16) int64, list []uint16, optional ...Optional) []int64 {
	if f == nil {
		return []int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int64NoOrder(f, list, worker)
		}
	}

	return pMapUint16Int64PreserveOrder(f, list, worker)
}

func pMapUint16Int64PreserveOrder(f func(uint16) int64, list []uint16, worker int) []int64 {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]uint16) {
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

func pMapUint16Int64NoOrder(f func(uint16) int64, list []uint16, worker int) []int64 {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan int64, chJobs chan uint16) {
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

// PMapUint16Int32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int32(f func(uint16) int32, list []uint16, optional ...Optional) []int32 {
	if f == nil {
		return []int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int32NoOrder(f, list, worker)
		}
	}

	return pMapUint16Int32PreserveOrder(f, list, worker)
}

func pMapUint16Int32PreserveOrder(f func(uint16) int32, list []uint16, worker int) []int32 {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]uint16) {
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

func pMapUint16Int32NoOrder(f func(uint16) int32, list []uint16, worker int) []int32 {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan int32, chJobs chan uint16) {
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

// PMapUint16Int16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int16(f func(uint16) int16, list []uint16, optional ...Optional) []int16 {
	if f == nil {
		return []int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int16NoOrder(f, list, worker)
		}
	}

	return pMapUint16Int16PreserveOrder(f, list, worker)
}

func pMapUint16Int16PreserveOrder(f func(uint16) int16, list []uint16, worker int) []int16 {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]uint16) {
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

func pMapUint16Int16NoOrder(f func(uint16) int16, list []uint16, worker int) []int16 {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan int16, chJobs chan uint16) {
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

// PMapUint16Int8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Int8(f func(uint16) int8, list []uint16, optional ...Optional) []int8 {
	if f == nil {
		return []int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Int8NoOrder(f, list, worker)
		}
	}

	return pMapUint16Int8PreserveOrder(f, list, worker)
}

func pMapUint16Int8PreserveOrder(f func(uint16) int8, list []uint16, worker int) []int8 {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]uint16) {
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

func pMapUint16Int8NoOrder(f func(uint16) int8, list []uint16, worker int) []int8 {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan int8, chJobs chan uint16) {
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

// PMapUint16Uint applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Uint(f func(uint16) uint, list []uint16, optional ...Optional) []uint {
	if f == nil {
		return []uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16UintNoOrder(f, list, worker)
		}
	}

	return pMapUint16UintPreserveOrder(f, list, worker)
}

func pMapUint16UintPreserveOrder(f func(uint16) uint, list []uint16, worker int) []uint {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]uint16) {
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

func pMapUint16UintNoOrder(f func(uint16) uint, list []uint16, worker int) []uint {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan uint, chJobs chan uint16) {
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

// PMapUint16Uint64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Uint64(f func(uint16) uint64, list []uint16, optional ...Optional) []uint64 {
	if f == nil {
		return []uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Uint64NoOrder(f, list, worker)
		}
	}

	return pMapUint16Uint64PreserveOrder(f, list, worker)
}

func pMapUint16Uint64PreserveOrder(f func(uint16) uint64, list []uint16, worker int) []uint64 {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]uint16) {
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

func pMapUint16Uint64NoOrder(f func(uint16) uint64, list []uint16, worker int) []uint64 {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan uint64, chJobs chan uint16) {
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

// PMapUint16Uint32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Uint32(f func(uint16) uint32, list []uint16, optional ...Optional) []uint32 {
	if f == nil {
		return []uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Uint32NoOrder(f, list, worker)
		}
	}

	return pMapUint16Uint32PreserveOrder(f, list, worker)
}

func pMapUint16Uint32PreserveOrder(f func(uint16) uint32, list []uint16, worker int) []uint32 {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]uint16) {
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

func pMapUint16Uint32NoOrder(f func(uint16) uint32, list []uint16, worker int) []uint32 {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan uint32, chJobs chan uint16) {
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

// PMapUint16Uint8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Uint8(f func(uint16) uint8, list []uint16, optional ...Optional) []uint8 {
	if f == nil {
		return []uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Uint8NoOrder(f, list, worker)
		}
	}

	return pMapUint16Uint8PreserveOrder(f, list, worker)
}

func pMapUint16Uint8PreserveOrder(f func(uint16) uint8, list []uint16, worker int) []uint8 {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]uint16) {
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

func pMapUint16Uint8NoOrder(f func(uint16) uint8, list []uint16, worker int) []uint8 {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan uint8, chJobs chan uint16) {
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

// PMapUint16Str applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Str(f func(uint16) string, list []uint16, optional ...Optional) []string {
	if f == nil {
		return []string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16StrNoOrder(f, list, worker)
		}
	}

	return pMapUint16StrPreserveOrder(f, list, worker)
}

func pMapUint16StrPreserveOrder(f func(uint16) string, list []uint16, worker int) []string {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]uint16) {
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

func pMapUint16StrNoOrder(f func(uint16) string, list []uint16, worker int) []string {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan string, chJobs chan uint16) {
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

// PMapUint16Bool applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Bool(f func(uint16) bool, list []uint16, optional ...Optional) []bool {
	if f == nil {
		return []bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16BoolNoOrder(f, list, worker)
		}
	}

	return pMapUint16BoolPreserveOrder(f, list, worker)
}

func pMapUint16BoolPreserveOrder(f func(uint16) bool, list []uint16, worker int) []bool {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]uint16) {
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

func pMapUint16BoolNoOrder(f func(uint16) bool, list []uint16, worker int) []bool {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan bool, chJobs chan uint16) {
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

// PMapUint16Float32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Float32(f func(uint16) float32, list []uint16, optional ...Optional) []float32 {
	if f == nil {
		return []float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Float32NoOrder(f, list, worker)
		}
	}

	return pMapUint16Float32PreserveOrder(f, list, worker)
}

func pMapUint16Float32PreserveOrder(f func(uint16) float32, list []uint16, worker int) []float32 {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]uint16) {
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

func pMapUint16Float32NoOrder(f func(uint16) float32, list []uint16, worker int) []float32 {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan float32, chJobs chan uint16) {
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

// PMapUint16Float64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint16Float64(f func(uint16) float64, list []uint16, optional ...Optional) []float64 {
	if f == nil {
		return []float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint16Float64NoOrder(f, list, worker)
		}
	}

	return pMapUint16Float64PreserveOrder(f, list, worker)
}

func pMapUint16Float64PreserveOrder(f func(uint16) float64, list []uint16, worker int) []float64 {
	chJobs := make(chan map[int]uint16, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint16{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]uint16) {
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

func pMapUint16Float64NoOrder(f func(uint16) float64, list []uint16, worker int) []float64 {
	chJobs := make(chan uint16, len(list))
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

		go func(chResult chan float64, chJobs chan uint16) {
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

// PMapUint8Int applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int(f func(uint8) int, list []uint8, optional ...Optional) []int {
	if f == nil {
		return []int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8IntNoOrder(f, list, worker)
		}
	}

	return pMapUint8IntPreserveOrder(f, list, worker)
}

func pMapUint8IntPreserveOrder(f func(uint8) int, list []uint8, worker int) []int {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]uint8) {
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

func pMapUint8IntNoOrder(f func(uint8) int, list []uint8, worker int) []int {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan int, chJobs chan uint8) {
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

// PMapUint8Int64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int64(f func(uint8) int64, list []uint8, optional ...Optional) []int64 {
	if f == nil {
		return []int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int64NoOrder(f, list, worker)
		}
	}

	return pMapUint8Int64PreserveOrder(f, list, worker)
}

func pMapUint8Int64PreserveOrder(f func(uint8) int64, list []uint8, worker int) []int64 {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]uint8) {
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

func pMapUint8Int64NoOrder(f func(uint8) int64, list []uint8, worker int) []int64 {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan int64, chJobs chan uint8) {
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

// PMapUint8Int32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int32(f func(uint8) int32, list []uint8, optional ...Optional) []int32 {
	if f == nil {
		return []int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int32NoOrder(f, list, worker)
		}
	}

	return pMapUint8Int32PreserveOrder(f, list, worker)
}

func pMapUint8Int32PreserveOrder(f func(uint8) int32, list []uint8, worker int) []int32 {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]uint8) {
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

func pMapUint8Int32NoOrder(f func(uint8) int32, list []uint8, worker int) []int32 {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan int32, chJobs chan uint8) {
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

// PMapUint8Int16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int16(f func(uint8) int16, list []uint8, optional ...Optional) []int16 {
	if f == nil {
		return []int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int16NoOrder(f, list, worker)
		}
	}

	return pMapUint8Int16PreserveOrder(f, list, worker)
}

func pMapUint8Int16PreserveOrder(f func(uint8) int16, list []uint8, worker int) []int16 {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]uint8) {
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

func pMapUint8Int16NoOrder(f func(uint8) int16, list []uint8, worker int) []int16 {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan int16, chJobs chan uint8) {
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

// PMapUint8Int8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Int8(f func(uint8) int8, list []uint8, optional ...Optional) []int8 {
	if f == nil {
		return []int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Int8NoOrder(f, list, worker)
		}
	}

	return pMapUint8Int8PreserveOrder(f, list, worker)
}

func pMapUint8Int8PreserveOrder(f func(uint8) int8, list []uint8, worker int) []int8 {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]uint8) {
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

func pMapUint8Int8NoOrder(f func(uint8) int8, list []uint8, worker int) []int8 {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan int8, chJobs chan uint8) {
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

// PMapUint8Uint applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Uint(f func(uint8) uint, list []uint8, optional ...Optional) []uint {
	if f == nil {
		return []uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8UintNoOrder(f, list, worker)
		}
	}

	return pMapUint8UintPreserveOrder(f, list, worker)
}

func pMapUint8UintPreserveOrder(f func(uint8) uint, list []uint8, worker int) []uint {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]uint8) {
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

func pMapUint8UintNoOrder(f func(uint8) uint, list []uint8, worker int) []uint {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan uint, chJobs chan uint8) {
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

// PMapUint8Uint64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Uint64(f func(uint8) uint64, list []uint8, optional ...Optional) []uint64 {
	if f == nil {
		return []uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Uint64NoOrder(f, list, worker)
		}
	}

	return pMapUint8Uint64PreserveOrder(f, list, worker)
}

func pMapUint8Uint64PreserveOrder(f func(uint8) uint64, list []uint8, worker int) []uint64 {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]uint8) {
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

func pMapUint8Uint64NoOrder(f func(uint8) uint64, list []uint8, worker int) []uint64 {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan uint64, chJobs chan uint8) {
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

// PMapUint8Uint32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Uint32(f func(uint8) uint32, list []uint8, optional ...Optional) []uint32 {
	if f == nil {
		return []uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Uint32NoOrder(f, list, worker)
		}
	}

	return pMapUint8Uint32PreserveOrder(f, list, worker)
}

func pMapUint8Uint32PreserveOrder(f func(uint8) uint32, list []uint8, worker int) []uint32 {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]uint8) {
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

func pMapUint8Uint32NoOrder(f func(uint8) uint32, list []uint8, worker int) []uint32 {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan uint32, chJobs chan uint8) {
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

// PMapUint8Uint16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Uint16(f func(uint8) uint16, list []uint8, optional ...Optional) []uint16 {
	if f == nil {
		return []uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Uint16NoOrder(f, list, worker)
		}
	}

	return pMapUint8Uint16PreserveOrder(f, list, worker)
}

func pMapUint8Uint16PreserveOrder(f func(uint8) uint16, list []uint8, worker int) []uint16 {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]uint8) {
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

func pMapUint8Uint16NoOrder(f func(uint8) uint16, list []uint8, worker int) []uint16 {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan uint16, chJobs chan uint8) {
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

// PMapUint8Str applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Str(f func(uint8) string, list []uint8, optional ...Optional) []string {
	if f == nil {
		return []string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8StrNoOrder(f, list, worker)
		}
	}

	return pMapUint8StrPreserveOrder(f, list, worker)
}

func pMapUint8StrPreserveOrder(f func(uint8) string, list []uint8, worker int) []string {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]uint8) {
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

func pMapUint8StrNoOrder(f func(uint8) string, list []uint8, worker int) []string {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan string, chJobs chan uint8) {
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

// PMapUint8Bool applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Bool(f func(uint8) bool, list []uint8, optional ...Optional) []bool {
	if f == nil {
		return []bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8BoolNoOrder(f, list, worker)
		}
	}

	return pMapUint8BoolPreserveOrder(f, list, worker)
}

func pMapUint8BoolPreserveOrder(f func(uint8) bool, list []uint8, worker int) []bool {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]uint8) {
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

func pMapUint8BoolNoOrder(f func(uint8) bool, list []uint8, worker int) []bool {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan bool, chJobs chan uint8) {
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

// PMapUint8Float32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Float32(f func(uint8) float32, list []uint8, optional ...Optional) []float32 {
	if f == nil {
		return []float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Float32NoOrder(f, list, worker)
		}
	}

	return pMapUint8Float32PreserveOrder(f, list, worker)
}

func pMapUint8Float32PreserveOrder(f func(uint8) float32, list []uint8, worker int) []float32 {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]uint8) {
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

func pMapUint8Float32NoOrder(f func(uint8) float32, list []uint8, worker int) []float32 {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan float32, chJobs chan uint8) {
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

// PMapUint8Float64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapUint8Float64(f func(uint8) float64, list []uint8, optional ...Optional) []float64 {
	if f == nil {
		return []float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapUint8Float64NoOrder(f, list, worker)
		}
	}

	return pMapUint8Float64PreserveOrder(f, list, worker)
}

func pMapUint8Float64PreserveOrder(f func(uint8) float64, list []uint8, worker int) []float64 {
	chJobs := make(chan map[int]uint8, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]uint8{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]uint8) {
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

func pMapUint8Float64NoOrder(f func(uint8) float64, list []uint8, worker int) []float64 {
	chJobs := make(chan uint8, len(list))
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

		go func(chResult chan float64, chJobs chan uint8) {
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

// PMapStrInt applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt(f func(string) int, list []string, optional ...Optional) []int {
	if f == nil {
		return []int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrIntNoOrder(f, list, worker)
		}
	}

	return pMapStrIntPreserveOrder(f, list, worker)
}

func pMapStrIntPreserveOrder(f func(string) int, list []string, worker int) []int {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]string) {
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

func pMapStrIntNoOrder(f func(string) int, list []string, worker int) []int {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan int, chJobs chan string) {
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

// PMapStrInt64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt64(f func(string) int64, list []string, optional ...Optional) []int64 {
	if f == nil {
		return []int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt64NoOrder(f, list, worker)
		}
	}

	return pMapStrInt64PreserveOrder(f, list, worker)
}

func pMapStrInt64PreserveOrder(f func(string) int64, list []string, worker int) []int64 {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]string) {
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

func pMapStrInt64NoOrder(f func(string) int64, list []string, worker int) []int64 {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan int64, chJobs chan string) {
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

// PMapStrInt32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt32(f func(string) int32, list []string, optional ...Optional) []int32 {
	if f == nil {
		return []int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt32NoOrder(f, list, worker)
		}
	}

	return pMapStrInt32PreserveOrder(f, list, worker)
}

func pMapStrInt32PreserveOrder(f func(string) int32, list []string, worker int) []int32 {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]string) {
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

func pMapStrInt32NoOrder(f func(string) int32, list []string, worker int) []int32 {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan int32, chJobs chan string) {
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

// PMapStrInt16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt16(f func(string) int16, list []string, optional ...Optional) []int16 {
	if f == nil {
		return []int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt16NoOrder(f, list, worker)
		}
	}

	return pMapStrInt16PreserveOrder(f, list, worker)
}

func pMapStrInt16PreserveOrder(f func(string) int16, list []string, worker int) []int16 {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]string) {
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

func pMapStrInt16NoOrder(f func(string) int16, list []string, worker int) []int16 {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan int16, chJobs chan string) {
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

// PMapStrInt8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrInt8(f func(string) int8, list []string, optional ...Optional) []int8 {
	if f == nil {
		return []int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrInt8NoOrder(f, list, worker)
		}
	}

	return pMapStrInt8PreserveOrder(f, list, worker)
}

func pMapStrInt8PreserveOrder(f func(string) int8, list []string, worker int) []int8 {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]string) {
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

func pMapStrInt8NoOrder(f func(string) int8, list []string, worker int) []int8 {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan int8, chJobs chan string) {
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

// PMapStrUint applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint(f func(string) uint, list []string, optional ...Optional) []uint {
	if f == nil {
		return []uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUintNoOrder(f, list, worker)
		}
	}

	return pMapStrUintPreserveOrder(f, list, worker)
}

func pMapStrUintPreserveOrder(f func(string) uint, list []string, worker int) []uint {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]string) {
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

func pMapStrUintNoOrder(f func(string) uint, list []string, worker int) []uint {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan uint, chJobs chan string) {
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

// PMapStrUint64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint64(f func(string) uint64, list []string, optional ...Optional) []uint64 {
	if f == nil {
		return []uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint64NoOrder(f, list, worker)
		}
	}

	return pMapStrUint64PreserveOrder(f, list, worker)
}

func pMapStrUint64PreserveOrder(f func(string) uint64, list []string, worker int) []uint64 {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]string) {
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

func pMapStrUint64NoOrder(f func(string) uint64, list []string, worker int) []uint64 {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan uint64, chJobs chan string) {
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

// PMapStrUint32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint32(f func(string) uint32, list []string, optional ...Optional) []uint32 {
	if f == nil {
		return []uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint32NoOrder(f, list, worker)
		}
	}

	return pMapStrUint32PreserveOrder(f, list, worker)
}

func pMapStrUint32PreserveOrder(f func(string) uint32, list []string, worker int) []uint32 {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]string) {
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

func pMapStrUint32NoOrder(f func(string) uint32, list []string, worker int) []uint32 {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan uint32, chJobs chan string) {
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

// PMapStrUint16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint16(f func(string) uint16, list []string, optional ...Optional) []uint16 {
	if f == nil {
		return []uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint16NoOrder(f, list, worker)
		}
	}

	return pMapStrUint16PreserveOrder(f, list, worker)
}

func pMapStrUint16PreserveOrder(f func(string) uint16, list []string, worker int) []uint16 {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]string) {
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

func pMapStrUint16NoOrder(f func(string) uint16, list []string, worker int) []uint16 {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan uint16, chJobs chan string) {
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

// PMapStrUint8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrUint8(f func(string) uint8, list []string, optional ...Optional) []uint8 {
	if f == nil {
		return []uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrUint8NoOrder(f, list, worker)
		}
	}

	return pMapStrUint8PreserveOrder(f, list, worker)
}

func pMapStrUint8PreserveOrder(f func(string) uint8, list []string, worker int) []uint8 {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]string) {
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

func pMapStrUint8NoOrder(f func(string) uint8, list []string, worker int) []uint8 {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan uint8, chJobs chan string) {
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

// PMapStrBool applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrBool(f func(string) bool, list []string, optional ...Optional) []bool {
	if f == nil {
		return []bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrBoolNoOrder(f, list, worker)
		}
	}

	return pMapStrBoolPreserveOrder(f, list, worker)
}

func pMapStrBoolPreserveOrder(f func(string) bool, list []string, worker int) []bool {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]string) {
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

func pMapStrBoolNoOrder(f func(string) bool, list []string, worker int) []bool {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan bool, chJobs chan string) {
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

// PMapStrFloat32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrFloat32(f func(string) float32, list []string, optional ...Optional) []float32 {
	if f == nil {
		return []float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrFloat32NoOrder(f, list, worker)
		}
	}

	return pMapStrFloat32PreserveOrder(f, list, worker)
}

func pMapStrFloat32PreserveOrder(f func(string) float32, list []string, worker int) []float32 {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]string) {
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

func pMapStrFloat32NoOrder(f func(string) float32, list []string, worker int) []float32 {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan float32, chJobs chan string) {
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

// PMapStrFloat64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapStrFloat64(f func(string) float64, list []string, optional ...Optional) []float64 {
	if f == nil {
		return []float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapStrFloat64NoOrder(f, list, worker)
		}
	}

	return pMapStrFloat64PreserveOrder(f, list, worker)
}

func pMapStrFloat64PreserveOrder(f func(string) float64, list []string, worker int) []float64 {
	chJobs := make(chan map[int]string, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]string{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]string) {
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

func pMapStrFloat64NoOrder(f func(string) float64, list []string, worker int) []float64 {
	chJobs := make(chan string, len(list))
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

		go func(chResult chan float64, chJobs chan string) {
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

// PMapBoolInt applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt(f func(bool) int, list []bool, optional ...Optional) []int {
	if f == nil {
		return []int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolIntNoOrder(f, list, worker)
		}
	}

	return pMapBoolIntPreserveOrder(f, list, worker)
}

func pMapBoolIntPreserveOrder(f func(bool) int, list []bool, worker int) []int {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]bool) {
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

func pMapBoolIntNoOrder(f func(bool) int, list []bool, worker int) []int {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan int, chJobs chan bool) {
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

// PMapBoolInt64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt64(f func(bool) int64, list []bool, optional ...Optional) []int64 {
	if f == nil {
		return []int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt64NoOrder(f, list, worker)
		}
	}

	return pMapBoolInt64PreserveOrder(f, list, worker)
}

func pMapBoolInt64PreserveOrder(f func(bool) int64, list []bool, worker int) []int64 {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]bool) {
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

func pMapBoolInt64NoOrder(f func(bool) int64, list []bool, worker int) []int64 {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan int64, chJobs chan bool) {
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

// PMapBoolInt32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt32(f func(bool) int32, list []bool, optional ...Optional) []int32 {
	if f == nil {
		return []int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt32NoOrder(f, list, worker)
		}
	}

	return pMapBoolInt32PreserveOrder(f, list, worker)
}

func pMapBoolInt32PreserveOrder(f func(bool) int32, list []bool, worker int) []int32 {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]bool) {
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

func pMapBoolInt32NoOrder(f func(bool) int32, list []bool, worker int) []int32 {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan int32, chJobs chan bool) {
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

// PMapBoolInt16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt16(f func(bool) int16, list []bool, optional ...Optional) []int16 {
	if f == nil {
		return []int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt16NoOrder(f, list, worker)
		}
	}

	return pMapBoolInt16PreserveOrder(f, list, worker)
}

func pMapBoolInt16PreserveOrder(f func(bool) int16, list []bool, worker int) []int16 {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]bool) {
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

func pMapBoolInt16NoOrder(f func(bool) int16, list []bool, worker int) []int16 {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan int16, chJobs chan bool) {
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

// PMapBoolInt8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolInt8(f func(bool) int8, list []bool, optional ...Optional) []int8 {
	if f == nil {
		return []int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolInt8NoOrder(f, list, worker)
		}
	}

	return pMapBoolInt8PreserveOrder(f, list, worker)
}

func pMapBoolInt8PreserveOrder(f func(bool) int8, list []bool, worker int) []int8 {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]bool) {
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

func pMapBoolInt8NoOrder(f func(bool) int8, list []bool, worker int) []int8 {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan int8, chJobs chan bool) {
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

// PMapBoolUint applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint(f func(bool) uint, list []bool, optional ...Optional) []uint {
	if f == nil {
		return []uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUintNoOrder(f, list, worker)
		}
	}

	return pMapBoolUintPreserveOrder(f, list, worker)
}

func pMapBoolUintPreserveOrder(f func(bool) uint, list []bool, worker int) []uint {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]bool) {
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

func pMapBoolUintNoOrder(f func(bool) uint, list []bool, worker int) []uint {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan uint, chJobs chan bool) {
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

// PMapBoolUint64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint64(f func(bool) uint64, list []bool, optional ...Optional) []uint64 {
	if f == nil {
		return []uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint64NoOrder(f, list, worker)
		}
	}

	return pMapBoolUint64PreserveOrder(f, list, worker)
}

func pMapBoolUint64PreserveOrder(f func(bool) uint64, list []bool, worker int) []uint64 {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]bool) {
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

func pMapBoolUint64NoOrder(f func(bool) uint64, list []bool, worker int) []uint64 {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan uint64, chJobs chan bool) {
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

// PMapBoolUint32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint32(f func(bool) uint32, list []bool, optional ...Optional) []uint32 {
	if f == nil {
		return []uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint32NoOrder(f, list, worker)
		}
	}

	return pMapBoolUint32PreserveOrder(f, list, worker)
}

func pMapBoolUint32PreserveOrder(f func(bool) uint32, list []bool, worker int) []uint32 {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]bool) {
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

func pMapBoolUint32NoOrder(f func(bool) uint32, list []bool, worker int) []uint32 {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan uint32, chJobs chan bool) {
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

// PMapBoolUint16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint16(f func(bool) uint16, list []bool, optional ...Optional) []uint16 {
	if f == nil {
		return []uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint16NoOrder(f, list, worker)
		}
	}

	return pMapBoolUint16PreserveOrder(f, list, worker)
}

func pMapBoolUint16PreserveOrder(f func(bool) uint16, list []bool, worker int) []uint16 {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]bool) {
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

func pMapBoolUint16NoOrder(f func(bool) uint16, list []bool, worker int) []uint16 {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan uint16, chJobs chan bool) {
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

// PMapBoolUint8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolUint8(f func(bool) uint8, list []bool, optional ...Optional) []uint8 {
	if f == nil {
		return []uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolUint8NoOrder(f, list, worker)
		}
	}

	return pMapBoolUint8PreserveOrder(f, list, worker)
}

func pMapBoolUint8PreserveOrder(f func(bool) uint8, list []bool, worker int) []uint8 {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]bool) {
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

func pMapBoolUint8NoOrder(f func(bool) uint8, list []bool, worker int) []uint8 {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan uint8, chJobs chan bool) {
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

// PMapBoolStr applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolStr(f func(bool) string, list []bool, optional ...Optional) []string {
	if f == nil {
		return []string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolStrNoOrder(f, list, worker)
		}
	}

	return pMapBoolStrPreserveOrder(f, list, worker)
}

func pMapBoolStrPreserveOrder(f func(bool) string, list []bool, worker int) []string {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]bool) {
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

func pMapBoolStrNoOrder(f func(bool) string, list []bool, worker int) []string {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan string, chJobs chan bool) {
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

// PMapBoolFloat32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolFloat32(f func(bool) float32, list []bool, optional ...Optional) []float32 {
	if f == nil {
		return []float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolFloat32NoOrder(f, list, worker)
		}
	}

	return pMapBoolFloat32PreserveOrder(f, list, worker)
}

func pMapBoolFloat32PreserveOrder(f func(bool) float32, list []bool, worker int) []float32 {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]bool) {
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

func pMapBoolFloat32NoOrder(f func(bool) float32, list []bool, worker int) []float32 {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan float32, chJobs chan bool) {
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

// PMapBoolFloat64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapBoolFloat64(f func(bool) float64, list []bool, optional ...Optional) []float64 {
	if f == nil {
		return []float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapBoolFloat64NoOrder(f, list, worker)
		}
	}

	return pMapBoolFloat64PreserveOrder(f, list, worker)
}

func pMapBoolFloat64PreserveOrder(f func(bool) float64, list []bool, worker int) []float64 {
	chJobs := make(chan map[int]bool, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]bool{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]bool) {
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

func pMapBoolFloat64NoOrder(f func(bool) float64, list []bool, worker int) []float64 {
	chJobs := make(chan bool, len(list))
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

		go func(chResult chan float64, chJobs chan bool) {
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

// PMapFloat32Int applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int(f func(float32) int, list []float32, optional ...Optional) []int {
	if f == nil {
		return []int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32IntNoOrder(f, list, worker)
		}
	}

	return pMapFloat32IntPreserveOrder(f, list, worker)
}

func pMapFloat32IntPreserveOrder(f func(float32) int, list []float32, worker int) []int {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]float32) {
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

func pMapFloat32IntNoOrder(f func(float32) int, list []float32, worker int) []int {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan int, chJobs chan float32) {
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

// PMapFloat32Int64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int64(f func(float32) int64, list []float32, optional ...Optional) []int64 {
	if f == nil {
		return []int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int64NoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int64PreserveOrder(f, list, worker)
}

func pMapFloat32Int64PreserveOrder(f func(float32) int64, list []float32, worker int) []int64 {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]float32) {
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

func pMapFloat32Int64NoOrder(f func(float32) int64, list []float32, worker int) []int64 {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan int64, chJobs chan float32) {
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

// PMapFloat32Int32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int32(f func(float32) int32, list []float32, optional ...Optional) []int32 {
	if f == nil {
		return []int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int32NoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int32PreserveOrder(f, list, worker)
}

func pMapFloat32Int32PreserveOrder(f func(float32) int32, list []float32, worker int) []int32 {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]float32) {
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

func pMapFloat32Int32NoOrder(f func(float32) int32, list []float32, worker int) []int32 {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan int32, chJobs chan float32) {
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

// PMapFloat32Int16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int16(f func(float32) int16, list []float32, optional ...Optional) []int16 {
	if f == nil {
		return []int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int16NoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int16PreserveOrder(f, list, worker)
}

func pMapFloat32Int16PreserveOrder(f func(float32) int16, list []float32, worker int) []int16 {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]float32) {
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

func pMapFloat32Int16NoOrder(f func(float32) int16, list []float32, worker int) []int16 {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan int16, chJobs chan float32) {
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

// PMapFloat32Int8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Int8(f func(float32) int8, list []float32, optional ...Optional) []int8 {
	if f == nil {
		return []int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Int8NoOrder(f, list, worker)
		}
	}

	return pMapFloat32Int8PreserveOrder(f, list, worker)
}

func pMapFloat32Int8PreserveOrder(f func(float32) int8, list []float32, worker int) []int8 {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]float32) {
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

func pMapFloat32Int8NoOrder(f func(float32) int8, list []float32, worker int) []int8 {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan int8, chJobs chan float32) {
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

// PMapFloat32Uint applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint(f func(float32) uint, list []float32, optional ...Optional) []uint {
	if f == nil {
		return []uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32UintNoOrder(f, list, worker)
		}
	}

	return pMapFloat32UintPreserveOrder(f, list, worker)
}

func pMapFloat32UintPreserveOrder(f func(float32) uint, list []float32, worker int) []uint {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]float32) {
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

func pMapFloat32UintNoOrder(f func(float32) uint, list []float32, worker int) []uint {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan uint, chJobs chan float32) {
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

// PMapFloat32Uint64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint64(f func(float32) uint64, list []float32, optional ...Optional) []uint64 {
	if f == nil {
		return []uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint64NoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint64PreserveOrder(f, list, worker)
}

func pMapFloat32Uint64PreserveOrder(f func(float32) uint64, list []float32, worker int) []uint64 {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]float32) {
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

func pMapFloat32Uint64NoOrder(f func(float32) uint64, list []float32, worker int) []uint64 {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan uint64, chJobs chan float32) {
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

// PMapFloat32Uint32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint32(f func(float32) uint32, list []float32, optional ...Optional) []uint32 {
	if f == nil {
		return []uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint32NoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint32PreserveOrder(f, list, worker)
}

func pMapFloat32Uint32PreserveOrder(f func(float32) uint32, list []float32, worker int) []uint32 {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]float32) {
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

func pMapFloat32Uint32NoOrder(f func(float32) uint32, list []float32, worker int) []uint32 {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan uint32, chJobs chan float32) {
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

// PMapFloat32Uint16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint16(f func(float32) uint16, list []float32, optional ...Optional) []uint16 {
	if f == nil {
		return []uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint16NoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint16PreserveOrder(f, list, worker)
}

func pMapFloat32Uint16PreserveOrder(f func(float32) uint16, list []float32, worker int) []uint16 {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]float32) {
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

func pMapFloat32Uint16NoOrder(f func(float32) uint16, list []float32, worker int) []uint16 {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan uint16, chJobs chan float32) {
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

// PMapFloat32Uint8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Uint8(f func(float32) uint8, list []float32, optional ...Optional) []uint8 {
	if f == nil {
		return []uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Uint8NoOrder(f, list, worker)
		}
	}

	return pMapFloat32Uint8PreserveOrder(f, list, worker)
}

func pMapFloat32Uint8PreserveOrder(f func(float32) uint8, list []float32, worker int) []uint8 {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]float32) {
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

func pMapFloat32Uint8NoOrder(f func(float32) uint8, list []float32, worker int) []uint8 {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan uint8, chJobs chan float32) {
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

// PMapFloat32Str applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Str(f func(float32) string, list []float32, optional ...Optional) []string {
	if f == nil {
		return []string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32StrNoOrder(f, list, worker)
		}
	}

	return pMapFloat32StrPreserveOrder(f, list, worker)
}

func pMapFloat32StrPreserveOrder(f func(float32) string, list []float32, worker int) []string {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]float32) {
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

func pMapFloat32StrNoOrder(f func(float32) string, list []float32, worker int) []string {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan string, chJobs chan float32) {
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

// PMapFloat32Bool applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Bool(f func(float32) bool, list []float32, optional ...Optional) []bool {
	if f == nil {
		return []bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32BoolNoOrder(f, list, worker)
		}
	}

	return pMapFloat32BoolPreserveOrder(f, list, worker)
}

func pMapFloat32BoolPreserveOrder(f func(float32) bool, list []float32, worker int) []bool {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]float32) {
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

func pMapFloat32BoolNoOrder(f func(float32) bool, list []float32, worker int) []bool {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan bool, chJobs chan float32) {
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

// PMapFloat32Float64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat32Float64(f func(float32) float64, list []float32, optional ...Optional) []float64 {
	if f == nil {
		return []float64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat32Float64NoOrder(f, list, worker)
		}
	}

	return pMapFloat32Float64PreserveOrder(f, list, worker)
}

func pMapFloat32Float64PreserveOrder(f func(float32) float64, list []float32, worker int) []float64 {
	chJobs := make(chan map[int]float32, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float32{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float64, chJobs chan map[int]float32) {
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

func pMapFloat32Float64NoOrder(f func(float32) float64, list []float32, worker int) []float64 {
	chJobs := make(chan float32, len(list))
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

		go func(chResult chan float64, chJobs chan float32) {
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

// PMapFloat64Int applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int(f func(float64) int, list []float64, optional ...Optional) []int {
	if f == nil {
		return []int{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64IntNoOrder(f, list, worker)
		}
	}

	return pMapFloat64IntPreserveOrder(f, list, worker)
}

func pMapFloat64IntPreserveOrder(f func(float64) int, list []float64, worker int) []int {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int, chJobs chan map[int]float64) {
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

func pMapFloat64IntNoOrder(f func(float64) int, list []float64, worker int) []int {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan int, chJobs chan float64) {
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

// PMapFloat64Int64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int64(f func(float64) int64, list []float64, optional ...Optional) []int64 {
	if f == nil {
		return []int64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int64NoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int64PreserveOrder(f, list, worker)
}

func pMapFloat64Int64PreserveOrder(f func(float64) int64, list []float64, worker int) []int64 {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int64, chJobs chan map[int]float64) {
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

func pMapFloat64Int64NoOrder(f func(float64) int64, list []float64, worker int) []int64 {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan int64, chJobs chan float64) {
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

// PMapFloat64Int32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int32(f func(float64) int32, list []float64, optional ...Optional) []int32 {
	if f == nil {
		return []int32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int32NoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int32PreserveOrder(f, list, worker)
}

func pMapFloat64Int32PreserveOrder(f func(float64) int32, list []float64, worker int) []int32 {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int32, chJobs chan map[int]float64) {
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

func pMapFloat64Int32NoOrder(f func(float64) int32, list []float64, worker int) []int32 {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan int32, chJobs chan float64) {
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

// PMapFloat64Int16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int16(f func(float64) int16, list []float64, optional ...Optional) []int16 {
	if f == nil {
		return []int16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int16NoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int16PreserveOrder(f, list, worker)
}

func pMapFloat64Int16PreserveOrder(f func(float64) int16, list []float64, worker int) []int16 {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int16, chJobs chan map[int]float64) {
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

func pMapFloat64Int16NoOrder(f func(float64) int16, list []float64, worker int) []int16 {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan int16, chJobs chan float64) {
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

// PMapFloat64Int8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Int8(f func(float64) int8, list []float64, optional ...Optional) []int8 {
	if f == nil {
		return []int8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Int8NoOrder(f, list, worker)
		}
	}

	return pMapFloat64Int8PreserveOrder(f, list, worker)
}

func pMapFloat64Int8PreserveOrder(f func(float64) int8, list []float64, worker int) []int8 {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]int8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]int8, chJobs chan map[int]float64) {
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

func pMapFloat64Int8NoOrder(f func(float64) int8, list []float64, worker int) []int8 {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan int8, chJobs chan float64) {
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

// PMapFloat64Uint applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint(f func(float64) uint, list []float64, optional ...Optional) []uint {
	if f == nil {
		return []uint{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64UintNoOrder(f, list, worker)
		}
	}

	return pMapFloat64UintPreserveOrder(f, list, worker)
}

func pMapFloat64UintPreserveOrder(f func(float64) uint, list []float64, worker int) []uint {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint, chJobs chan map[int]float64) {
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

func pMapFloat64UintNoOrder(f func(float64) uint, list []float64, worker int) []uint {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan uint, chJobs chan float64) {
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

// PMapFloat64Uint64 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint64(f func(float64) uint64, list []float64, optional ...Optional) []uint64 {
	if f == nil {
		return []uint64{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint64NoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint64PreserveOrder(f, list, worker)
}

func pMapFloat64Uint64PreserveOrder(f func(float64) uint64, list []float64, worker int) []uint64 {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint64, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint64, chJobs chan map[int]float64) {
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

func pMapFloat64Uint64NoOrder(f func(float64) uint64, list []float64, worker int) []uint64 {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan uint64, chJobs chan float64) {
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

// PMapFloat64Uint32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint32(f func(float64) uint32, list []float64, optional ...Optional) []uint32 {
	if f == nil {
		return []uint32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint32NoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint32PreserveOrder(f, list, worker)
}

func pMapFloat64Uint32PreserveOrder(f func(float64) uint32, list []float64, worker int) []uint32 {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint32, chJobs chan map[int]float64) {
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

func pMapFloat64Uint32NoOrder(f func(float64) uint32, list []float64, worker int) []uint32 {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan uint32, chJobs chan float64) {
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

// PMapFloat64Uint16 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint16(f func(float64) uint16, list []float64, optional ...Optional) []uint16 {
	if f == nil {
		return []uint16{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint16NoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint16PreserveOrder(f, list, worker)
}

func pMapFloat64Uint16PreserveOrder(f func(float64) uint16, list []float64, worker int) []uint16 {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint16, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint16, chJobs chan map[int]float64) {
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

func pMapFloat64Uint16NoOrder(f func(float64) uint16, list []float64, worker int) []uint16 {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan uint16, chJobs chan float64) {
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

// PMapFloat64Uint8 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Uint8(f func(float64) uint8, list []float64, optional ...Optional) []uint8 {
	if f == nil {
		return []uint8{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Uint8NoOrder(f, list, worker)
		}
	}

	return pMapFloat64Uint8PreserveOrder(f, list, worker)
}

func pMapFloat64Uint8PreserveOrder(f func(float64) uint8, list []float64, worker int) []uint8 {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]uint8, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]uint8, chJobs chan map[int]float64) {
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

func pMapFloat64Uint8NoOrder(f func(float64) uint8, list []float64, worker int) []uint8 {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan uint8, chJobs chan float64) {
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

// PMapFloat64Str applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Str(f func(float64) string, list []float64, optional ...Optional) []string {
	if f == nil {
		return []string{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64StrNoOrder(f, list, worker)
		}
	}

	return pMapFloat64StrPreserveOrder(f, list, worker)
}

func pMapFloat64StrPreserveOrder(f func(float64) string, list []float64, worker int) []string {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]string, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]string, chJobs chan map[int]float64) {
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

func pMapFloat64StrNoOrder(f func(float64) string, list []float64, worker int) []string {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan string, chJobs chan float64) {
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

// PMapFloat64Bool applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Bool(f func(float64) bool, list []float64, optional ...Optional) []bool {
	if f == nil {
		return []bool{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64BoolNoOrder(f, list, worker)
		}
	}

	return pMapFloat64BoolPreserveOrder(f, list, worker)
}

func pMapFloat64BoolPreserveOrder(f func(float64) bool, list []float64, worker int) []bool {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]bool, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]bool, chJobs chan map[int]float64) {
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

func pMapFloat64BoolNoOrder(f func(float64) bool, list []float64, worker int) []bool {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan bool, chJobs chan float64) {
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

// PMapFloat64Float32 applies the function(1st argument) on each item in the list and returns a new list.
//  Order of new list is guaranteed. This feature can be disabled by passing: Optional{RandomOrder: true} to gain performance
//  Run in parallel. no_of_goroutines = no_of_items_in_list or 3rd argument can be passed to fix the number of goroutines.
//
// Takes 3 inputs. 3rd argument is optional
//  1. Function - takes 1 input
//  2. List
//  3. optional argument - fp.Optional{FixedPool: <some_number>}
func PMapFloat64Float32(f func(float64) float32, list []float64, optional ...Optional) []float32 {
	if f == nil {
		return []float32{}
	}

	var worker = len(list)
	if optional != nil && len(optional) > 0 {
		if optional[0].FixedPool > 0 && optional[0].FixedPool < worker {
			worker = optional[0].FixedPool
		}

		if optional[0].RandomOrder == true {
			return pMapFloat64Float32NoOrder(f, list, worker)
		}
	}

	return pMapFloat64Float32PreserveOrder(f, list, worker)
}

func pMapFloat64Float32PreserveOrder(f func(float64) float32, list []float64, worker int) []float32 {
	chJobs := make(chan map[int]float64, len(list))
	go func() {
		for i, v := range list {
			chJobs <- map[int]float64{i: v}
		}
		close(chJobs)
	}()

	chResult := make(chan map[int]float32, worker/3)

	var wg sync.WaitGroup

	for i := 0; i < worker; i++ {
		wg.Add(1)

		go func(chResult chan map[int]float32, chJobs chan map[int]float64) {
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

func pMapFloat64Float32NoOrder(f func(float64) float32, list []float64, worker int) []float32 {
	chJobs := make(chan float64, len(list))
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

		go func(chResult chan float32, chJobs chan float64) {
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
