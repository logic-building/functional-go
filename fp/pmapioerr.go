package fp

import "sync"

// PMapIntInt64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: (int64, error)
//	2. List
//
// Returns
//	New List of type (int64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapIntInt64Err(f func(int) (int64, error), list []int) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	ch := make(chan map[int]int64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v int) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapIntInt32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: (int32, error)
//	2. List
//
// Returns
//	New List of type (int32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapIntInt32Err(f func(int) (int32, error), list []int) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	ch := make(chan map[int]int32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v int) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapIntInt16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: (int16, error)
//	2. List
//
// Returns
//	New List of type (int16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapIntInt16Err(f func(int) (int16, error), list []int) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	ch := make(chan map[int]int16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v int) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapIntInt8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: (int8, error)
//	2. List
//
// Returns
//	New List of type (int8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapIntInt8Err(f func(int) (int8, error), list []int) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	ch := make(chan map[int]int8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v int) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapIntUintErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: (uint, error)
//	2. List
//
// Returns
//	New List of type (uint, error)
//	Empty list if all arguments are nil or either one is nil
func PMapIntUintErr(f func(int) (uint, error), list []int) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	ch := make(chan map[int]uint, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v int) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapIntUint64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: (uint64, error)
//	2. List
//
// Returns
//	New List of type (uint64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapIntUint64Err(f func(int) (uint64, error), list []int) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	ch := make(chan map[int]uint64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v int) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapIntUint32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: (uint32, error)
//	2. List
//
// Returns
//	New List of type (uint32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapIntUint32Err(f func(int) (uint32, error), list []int) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	ch := make(chan map[int]uint32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v int) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapIntUint16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: (uint16, error)
//	2. List
//
// Returns
//	New List of type (uint16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapIntUint16Err(f func(int) (uint16, error), list []int) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	ch := make(chan map[int]uint16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v int) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapIntUint8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: (uint8, error)
//	2. List
//
// Returns
//	New List of type (uint8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapIntUint8Err(f func(int) (uint8, error), list []int) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	ch := make(chan map[int]uint8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v int) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapIntStrErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: (string, error)
//	2. List
//
// Returns
//	New List of type (string, error)
//	Empty list if all arguments are nil or either one is nil
func PMapIntStrErr(f func(int) (string, error), list []int) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	ch := make(chan map[int]string, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v int) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]string{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapIntBoolErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: (bool, error)
//	2. List
//
// Returns
//	New List of type (bool, error)
//	Empty list if all arguments are nil or either one is nil
func PMapIntBoolErr(f func(int) (bool, error), list []int) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	ch := make(chan map[int]bool, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v int) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]bool{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapIntFloat32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: (float32, error)
//	2. List
//
// Returns
//	New List of type (float32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapIntFloat32Err(f func(int) (float32, error), list []int) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	ch := make(chan map[int]float32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float32, i int, v int) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapIntFloat64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: (float64, error)
//	2. List
//
// Returns
//	New List of type (float64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapIntFloat64Err(f func(int) (float64, error), list []int) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	ch := make(chan map[int]float64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float64, i int, v int) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt64IntErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: (int, error)
//	2. List
//
// Returns
//	New List of type (int, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt64IntErr(f func(int64) (int, error), list []int64) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	ch := make(chan map[int]int, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v int64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt64Int32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: (int32, error)
//	2. List
//
// Returns
//	New List of type (int32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Int32Err(f func(int64) (int32, error), list []int64) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	ch := make(chan map[int]int32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v int64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt64Int16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: (int16, error)
//	2. List
//
// Returns
//	New List of type (int16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Int16Err(f func(int64) (int16, error), list []int64) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	ch := make(chan map[int]int16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v int64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt64Int8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: (int8, error)
//	2. List
//
// Returns
//	New List of type (int8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Int8Err(f func(int64) (int8, error), list []int64) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	ch := make(chan map[int]int8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v int64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt64UintErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: (uint, error)
//	2. List
//
// Returns
//	New List of type (uint, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt64UintErr(f func(int64) (uint, error), list []int64) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	ch := make(chan map[int]uint, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v int64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt64Uint64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: (uint64, error)
//	2. List
//
// Returns
//	New List of type (uint64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Uint64Err(f func(int64) (uint64, error), list []int64) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	ch := make(chan map[int]uint64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v int64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt64Uint32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: (uint32, error)
//	2. List
//
// Returns
//	New List of type (uint32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Uint32Err(f func(int64) (uint32, error), list []int64) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	ch := make(chan map[int]uint32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v int64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt64Uint16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: (uint16, error)
//	2. List
//
// Returns
//	New List of type (uint16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Uint16Err(f func(int64) (uint16, error), list []int64) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	ch := make(chan map[int]uint16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v int64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt64Uint8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: (uint8, error)
//	2. List
//
// Returns
//	New List of type (uint8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Uint8Err(f func(int64) (uint8, error), list []int64) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	ch := make(chan map[int]uint8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v int64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt64StrErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: (string, error)
//	2. List
//
// Returns
//	New List of type (string, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt64StrErr(f func(int64) (string, error), list []int64) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	ch := make(chan map[int]string, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v int64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]string{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt64BoolErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: (bool, error)
//	2. List
//
// Returns
//	New List of type (bool, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt64BoolErr(f func(int64) (bool, error), list []int64) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	ch := make(chan map[int]bool, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v int64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]bool{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt64Float32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: (float32, error)
//	2. List
//
// Returns
//	New List of type (float32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Float32Err(f func(int64) (float32, error), list []int64) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	ch := make(chan map[int]float32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float32, i int, v int64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt64Float64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int64 output type: (float64, error)
//	2. List
//
// Returns
//	New List of type (float64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt64Float64Err(f func(int64) (float64, error), list []int64) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	ch := make(chan map[int]float64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float64, i int, v int64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt32IntErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: (int, error)
//	2. List
//
// Returns
//	New List of type (int, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt32IntErr(f func(int32) (int, error), list []int32) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	ch := make(chan map[int]int, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v int32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt32Int64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: (int64, error)
//	2. List
//
// Returns
//	New List of type (int64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Int64Err(f func(int32) (int64, error), list []int32) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	ch := make(chan map[int]int64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v int32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt32Int16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: (int16, error)
//	2. List
//
// Returns
//	New List of type (int16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Int16Err(f func(int32) (int16, error), list []int32) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	ch := make(chan map[int]int16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v int32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt32Int8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: (int8, error)
//	2. List
//
// Returns
//	New List of type (int8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Int8Err(f func(int32) (int8, error), list []int32) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	ch := make(chan map[int]int8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v int32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt32UintErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: (uint, error)
//	2. List
//
// Returns
//	New List of type (uint, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt32UintErr(f func(int32) (uint, error), list []int32) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	ch := make(chan map[int]uint, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v int32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt32Uint64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: (uint64, error)
//	2. List
//
// Returns
//	New List of type (uint64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Uint64Err(f func(int32) (uint64, error), list []int32) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	ch := make(chan map[int]uint64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v int32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt32Uint32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: (uint32, error)
//	2. List
//
// Returns
//	New List of type (uint32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Uint32Err(f func(int32) (uint32, error), list []int32) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	ch := make(chan map[int]uint32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v int32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt32Uint16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: (uint16, error)
//	2. List
//
// Returns
//	New List of type (uint16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Uint16Err(f func(int32) (uint16, error), list []int32) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	ch := make(chan map[int]uint16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v int32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt32Uint8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: (uint8, error)
//	2. List
//
// Returns
//	New List of type (uint8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Uint8Err(f func(int32) (uint8, error), list []int32) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	ch := make(chan map[int]uint8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v int32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt32StrErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: (string, error)
//	2. List
//
// Returns
//	New List of type (string, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt32StrErr(f func(int32) (string, error), list []int32) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	ch := make(chan map[int]string, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v int32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]string{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt32BoolErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: (bool, error)
//	2. List
//
// Returns
//	New List of type (bool, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt32BoolErr(f func(int32) (bool, error), list []int32) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	ch := make(chan map[int]bool, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v int32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]bool{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt32Float32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: (float32, error)
//	2. List
//
// Returns
//	New List of type (float32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Float32Err(f func(int32) (float32, error), list []int32) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	ch := make(chan map[int]float32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float32, i int, v int32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt32Float64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int32 output type: (float64, error)
//	2. List
//
// Returns
//	New List of type (float64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt32Float64Err(f func(int32) (float64, error), list []int32) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	ch := make(chan map[int]float64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float64, i int, v int32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt16IntErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: (int, error)
//	2. List
//
// Returns
//	New List of type (int, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt16IntErr(f func(int16) (int, error), list []int16) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	ch := make(chan map[int]int, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v int16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt16Int64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: (int64, error)
//	2. List
//
// Returns
//	New List of type (int64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Int64Err(f func(int16) (int64, error), list []int16) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	ch := make(chan map[int]int64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v int16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt16Int32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: (int32, error)
//	2. List
//
// Returns
//	New List of type (int32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Int32Err(f func(int16) (int32, error), list []int16) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	ch := make(chan map[int]int32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v int16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt16Int8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: (int8, error)
//	2. List
//
// Returns
//	New List of type (int8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Int8Err(f func(int16) (int8, error), list []int16) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	ch := make(chan map[int]int8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v int16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt16UintErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: (uint, error)
//	2. List
//
// Returns
//	New List of type (uint, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt16UintErr(f func(int16) (uint, error), list []int16) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	ch := make(chan map[int]uint, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v int16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt16Uint64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: (uint64, error)
//	2. List
//
// Returns
//	New List of type (uint64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Uint64Err(f func(int16) (uint64, error), list []int16) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	ch := make(chan map[int]uint64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v int16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt16Uint32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: (uint32, error)
//	2. List
//
// Returns
//	New List of type (uint32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Uint32Err(f func(int16) (uint32, error), list []int16) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	ch := make(chan map[int]uint32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v int16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt16Uint16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: (uint16, error)
//	2. List
//
// Returns
//	New List of type (uint16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Uint16Err(f func(int16) (uint16, error), list []int16) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	ch := make(chan map[int]uint16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v int16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt16Uint8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: (uint8, error)
//	2. List
//
// Returns
//	New List of type (uint8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Uint8Err(f func(int16) (uint8, error), list []int16) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	ch := make(chan map[int]uint8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v int16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt16StrErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: (string, error)
//	2. List
//
// Returns
//	New List of type (string, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt16StrErr(f func(int16) (string, error), list []int16) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	ch := make(chan map[int]string, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v int16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]string{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt16BoolErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: (bool, error)
//	2. List
//
// Returns
//	New List of type (bool, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt16BoolErr(f func(int16) (bool, error), list []int16) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	ch := make(chan map[int]bool, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v int16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]bool{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt16Float32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: (float32, error)
//	2. List
//
// Returns
//	New List of type (float32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Float32Err(f func(int16) (float32, error), list []int16) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	ch := make(chan map[int]float32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float32, i int, v int16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt16Float64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int16 output type: (float64, error)
//	2. List
//
// Returns
//	New List of type (float64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt16Float64Err(f func(int16) (float64, error), list []int16) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	ch := make(chan map[int]float64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float64, i int, v int16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt8IntErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: (int, error)
//	2. List
//
// Returns
//	New List of type (int, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt8IntErr(f func(int8) (int, error), list []int8) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	ch := make(chan map[int]int, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v int8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt8Int64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: (int64, error)
//	2. List
//
// Returns
//	New List of type (int64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Int64Err(f func(int8) (int64, error), list []int8) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	ch := make(chan map[int]int64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v int8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt8Int32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: (int32, error)
//	2. List
//
// Returns
//	New List of type (int32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Int32Err(f func(int8) (int32, error), list []int8) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	ch := make(chan map[int]int32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v int8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt8Int16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: (int16, error)
//	2. List
//
// Returns
//	New List of type (int16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Int16Err(f func(int8) (int16, error), list []int8) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	ch := make(chan map[int]int16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v int8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt8UintErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: (uint, error)
//	2. List
//
// Returns
//	New List of type (uint, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt8UintErr(f func(int8) (uint, error), list []int8) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	ch := make(chan map[int]uint, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v int8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt8Uint64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: (uint64, error)
//	2. List
//
// Returns
//	New List of type (uint64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Uint64Err(f func(int8) (uint64, error), list []int8) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	ch := make(chan map[int]uint64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v int8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt8Uint32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: (uint32, error)
//	2. List
//
// Returns
//	New List of type (uint32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Uint32Err(f func(int8) (uint32, error), list []int8) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	ch := make(chan map[int]uint32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v int8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt8Uint16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: (uint16, error)
//	2. List
//
// Returns
//	New List of type (uint16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Uint16Err(f func(int8) (uint16, error), list []int8) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	ch := make(chan map[int]uint16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v int8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt8Uint8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: (uint8, error)
//	2. List
//
// Returns
//	New List of type (uint8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Uint8Err(f func(int8) (uint8, error), list []int8) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	ch := make(chan map[int]uint8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v int8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt8StrErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: (string, error)
//	2. List
//
// Returns
//	New List of type (string, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt8StrErr(f func(int8) (string, error), list []int8) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	ch := make(chan map[int]string, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v int8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]string{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt8BoolErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: (bool, error)
//	2. List
//
// Returns
//	New List of type (bool, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt8BoolErr(f func(int8) (bool, error), list []int8) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	ch := make(chan map[int]bool, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v int8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]bool{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt8Float32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: (float32, error)
//	2. List
//
// Returns
//	New List of type (float32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Float32Err(f func(int8) (float32, error), list []int8) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	ch := make(chan map[int]float32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float32, i int, v int8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapInt8Float64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int8 output type: (float64, error)
//	2. List
//
// Returns
//	New List of type (float64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapInt8Float64Err(f func(int8) (float64, error), list []int8) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	ch := make(chan map[int]float64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float64, i int, v int8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUintIntErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: (int, error)
//	2. List
//
// Returns
//	New List of type (int, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUintIntErr(f func(uint) (int, error), list []uint) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	ch := make(chan map[int]int, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v uint) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUintInt64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: (int64, error)
//	2. List
//
// Returns
//	New List of type (int64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUintInt64Err(f func(uint) (int64, error), list []uint) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	ch := make(chan map[int]int64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v uint) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUintInt32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: (int32, error)
//	2. List
//
// Returns
//	New List of type (int32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUintInt32Err(f func(uint) (int32, error), list []uint) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	ch := make(chan map[int]int32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v uint) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUintInt16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: (int16, error)
//	2. List
//
// Returns
//	New List of type (int16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUintInt16Err(f func(uint) (int16, error), list []uint) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	ch := make(chan map[int]int16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v uint) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUintInt8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: (int8, error)
//	2. List
//
// Returns
//	New List of type (int8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUintInt8Err(f func(uint) (int8, error), list []uint) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	ch := make(chan map[int]int8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v uint) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUintUint64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: (uint64, error)
//	2. List
//
// Returns
//	New List of type (uint64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUintUint64Err(f func(uint) (uint64, error), list []uint) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	ch := make(chan map[int]uint64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v uint) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUintUint32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: (uint32, error)
//	2. List
//
// Returns
//	New List of type (uint32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUintUint32Err(f func(uint) (uint32, error), list []uint) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	ch := make(chan map[int]uint32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v uint) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUintUint16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: (uint16, error)
//	2. List
//
// Returns
//	New List of type (uint16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUintUint16Err(f func(uint) (uint16, error), list []uint) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	ch := make(chan map[int]uint16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v uint) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUintUint8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: (uint8, error)
//	2. List
//
// Returns
//	New List of type (uint8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUintUint8Err(f func(uint) (uint8, error), list []uint) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	ch := make(chan map[int]uint8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v uint) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUintStrErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: (string, error)
//	2. List
//
// Returns
//	New List of type (string, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUintStrErr(f func(uint) (string, error), list []uint) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	ch := make(chan map[int]string, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v uint) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]string{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUintBoolErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: (bool, error)
//	2. List
//
// Returns
//	New List of type (bool, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUintBoolErr(f func(uint) (bool, error), list []uint) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	ch := make(chan map[int]bool, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v uint) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]bool{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUintFloat32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: (float32, error)
//	2. List
//
// Returns
//	New List of type (float32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUintFloat32Err(f func(uint) (float32, error), list []uint) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	ch := make(chan map[int]float32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float32, i int, v uint) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUintFloat64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint output type: (float64, error)
//	2. List
//
// Returns
//	New List of type (float64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUintFloat64Err(f func(uint) (float64, error), list []uint) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	ch := make(chan map[int]float64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float64, i int, v uint) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint64IntErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: (int, error)
//	2. List
//
// Returns
//	New List of type (int, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint64IntErr(f func(uint64) (int, error), list []uint64) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	ch := make(chan map[int]int, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v uint64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint64Int64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: (int64, error)
//	2. List
//
// Returns
//	New List of type (int64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Int64Err(f func(uint64) (int64, error), list []uint64) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	ch := make(chan map[int]int64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v uint64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint64Int32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: (int32, error)
//	2. List
//
// Returns
//	New List of type (int32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Int32Err(f func(uint64) (int32, error), list []uint64) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	ch := make(chan map[int]int32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v uint64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint64Int16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: (int16, error)
//	2. List
//
// Returns
//	New List of type (int16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Int16Err(f func(uint64) (int16, error), list []uint64) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	ch := make(chan map[int]int16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v uint64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint64Int8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: (int8, error)
//	2. List
//
// Returns
//	New List of type (int8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Int8Err(f func(uint64) (int8, error), list []uint64) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	ch := make(chan map[int]int8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v uint64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint64UintErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: (uint, error)
//	2. List
//
// Returns
//	New List of type (uint, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint64UintErr(f func(uint64) (uint, error), list []uint64) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	ch := make(chan map[int]uint, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v uint64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint64Uint32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: (uint32, error)
//	2. List
//
// Returns
//	New List of type (uint32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Uint32Err(f func(uint64) (uint32, error), list []uint64) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	ch := make(chan map[int]uint32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v uint64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint64Uint16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: (uint16, error)
//	2. List
//
// Returns
//	New List of type (uint16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Uint16Err(f func(uint64) (uint16, error), list []uint64) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	ch := make(chan map[int]uint16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v uint64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint64Uint8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: (uint8, error)
//	2. List
//
// Returns
//	New List of type (uint8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Uint8Err(f func(uint64) (uint8, error), list []uint64) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	ch := make(chan map[int]uint8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v uint64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint64StrErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: (string, error)
//	2. List
//
// Returns
//	New List of type (string, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint64StrErr(f func(uint64) (string, error), list []uint64) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	ch := make(chan map[int]string, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v uint64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]string{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint64BoolErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: (bool, error)
//	2. List
//
// Returns
//	New List of type (bool, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint64BoolErr(f func(uint64) (bool, error), list []uint64) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	ch := make(chan map[int]bool, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v uint64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]bool{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint64Float32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: (float32, error)
//	2. List
//
// Returns
//	New List of type (float32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Float32Err(f func(uint64) (float32, error), list []uint64) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	ch := make(chan map[int]float32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float32, i int, v uint64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint64Float64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint64 output type: (float64, error)
//	2. List
//
// Returns
//	New List of type (float64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint64Float64Err(f func(uint64) (float64, error), list []uint64) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	ch := make(chan map[int]float64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float64, i int, v uint64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint32IntErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: (int, error)
//	2. List
//
// Returns
//	New List of type (int, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint32IntErr(f func(uint32) (int, error), list []uint32) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	ch := make(chan map[int]int, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v uint32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint32Int64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: (int64, error)
//	2. List
//
// Returns
//	New List of type (int64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Int64Err(f func(uint32) (int64, error), list []uint32) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	ch := make(chan map[int]int64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v uint32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint32Int32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: (int32, error)
//	2. List
//
// Returns
//	New List of type (int32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Int32Err(f func(uint32) (int32, error), list []uint32) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	ch := make(chan map[int]int32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v uint32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint32Int16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: (int16, error)
//	2. List
//
// Returns
//	New List of type (int16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Int16Err(f func(uint32) (int16, error), list []uint32) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	ch := make(chan map[int]int16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v uint32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint32Int8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: (int8, error)
//	2. List
//
// Returns
//	New List of type (int8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Int8Err(f func(uint32) (int8, error), list []uint32) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	ch := make(chan map[int]int8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v uint32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint32UintErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: (uint, error)
//	2. List
//
// Returns
//	New List of type (uint, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint32UintErr(f func(uint32) (uint, error), list []uint32) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	ch := make(chan map[int]uint, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v uint32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint32Uint64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: (uint64, error)
//	2. List
//
// Returns
//	New List of type (uint64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Uint64Err(f func(uint32) (uint64, error), list []uint32) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	ch := make(chan map[int]uint64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v uint32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint32Uint16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: (uint16, error)
//	2. List
//
// Returns
//	New List of type (uint16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Uint16Err(f func(uint32) (uint16, error), list []uint32) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	ch := make(chan map[int]uint16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v uint32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint32Uint8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: (uint8, error)
//	2. List
//
// Returns
//	New List of type (uint8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Uint8Err(f func(uint32) (uint8, error), list []uint32) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	ch := make(chan map[int]uint8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v uint32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint32StrErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: (string, error)
//	2. List
//
// Returns
//	New List of type (string, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint32StrErr(f func(uint32) (string, error), list []uint32) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	ch := make(chan map[int]string, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v uint32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]string{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint32BoolErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: (bool, error)
//	2. List
//
// Returns
//	New List of type (bool, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint32BoolErr(f func(uint32) (bool, error), list []uint32) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	ch := make(chan map[int]bool, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v uint32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]bool{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint32Float32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: (float32, error)
//	2. List
//
// Returns
//	New List of type (float32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Float32Err(f func(uint32) (float32, error), list []uint32) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	ch := make(chan map[int]float32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float32, i int, v uint32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint32Float64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint32 output type: (float64, error)
//	2. List
//
// Returns
//	New List of type (float64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint32Float64Err(f func(uint32) (float64, error), list []uint32) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	ch := make(chan map[int]float64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float64, i int, v uint32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint16IntErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: (int, error)
//	2. List
//
// Returns
//	New List of type (int, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint16IntErr(f func(uint16) (int, error), list []uint16) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	ch := make(chan map[int]int, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v uint16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint16Int64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: (int64, error)
//	2. List
//
// Returns
//	New List of type (int64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Int64Err(f func(uint16) (int64, error), list []uint16) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	ch := make(chan map[int]int64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v uint16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint16Int32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: (int32, error)
//	2. List
//
// Returns
//	New List of type (int32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Int32Err(f func(uint16) (int32, error), list []uint16) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	ch := make(chan map[int]int32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v uint16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint16Int16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: (int16, error)
//	2. List
//
// Returns
//	New List of type (int16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Int16Err(f func(uint16) (int16, error), list []uint16) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	ch := make(chan map[int]int16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v uint16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint16Int8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: (int8, error)
//	2. List
//
// Returns
//	New List of type (int8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Int8Err(f func(uint16) (int8, error), list []uint16) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	ch := make(chan map[int]int8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v uint16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint16UintErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: (uint, error)
//	2. List
//
// Returns
//	New List of type (uint, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint16UintErr(f func(uint16) (uint, error), list []uint16) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	ch := make(chan map[int]uint, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v uint16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint16Uint64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: (uint64, error)
//	2. List
//
// Returns
//	New List of type (uint64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Uint64Err(f func(uint16) (uint64, error), list []uint16) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	ch := make(chan map[int]uint64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v uint16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint16Uint32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: (uint32, error)
//	2. List
//
// Returns
//	New List of type (uint32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Uint32Err(f func(uint16) (uint32, error), list []uint16) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	ch := make(chan map[int]uint32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v uint16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint16Uint8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: (uint8, error)
//	2. List
//
// Returns
//	New List of type (uint8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Uint8Err(f func(uint16) (uint8, error), list []uint16) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	ch := make(chan map[int]uint8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v uint16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint16StrErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: (string, error)
//	2. List
//
// Returns
//	New List of type (string, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint16StrErr(f func(uint16) (string, error), list []uint16) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	ch := make(chan map[int]string, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v uint16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]string{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint16BoolErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: (bool, error)
//	2. List
//
// Returns
//	New List of type (bool, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint16BoolErr(f func(uint16) (bool, error), list []uint16) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	ch := make(chan map[int]bool, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v uint16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]bool{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint16Float32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: (float32, error)
//	2. List
//
// Returns
//	New List of type (float32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Float32Err(f func(uint16) (float32, error), list []uint16) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	ch := make(chan map[int]float32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float32, i int, v uint16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint16Float64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint16 output type: (float64, error)
//	2. List
//
// Returns
//	New List of type (float64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint16Float64Err(f func(uint16) (float64, error), list []uint16) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	ch := make(chan map[int]float64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float64, i int, v uint16) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint8IntErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: (int, error)
//	2. List
//
// Returns
//	New List of type (int, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint8IntErr(f func(uint8) (int, error), list []uint8) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	ch := make(chan map[int]int, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v uint8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint8Int64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: (int64, error)
//	2. List
//
// Returns
//	New List of type (int64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Int64Err(f func(uint8) (int64, error), list []uint8) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	ch := make(chan map[int]int64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v uint8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint8Int32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: (int32, error)
//	2. List
//
// Returns
//	New List of type (int32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Int32Err(f func(uint8) (int32, error), list []uint8) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	ch := make(chan map[int]int32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v uint8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint8Int16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: (int16, error)
//	2. List
//
// Returns
//	New List of type (int16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Int16Err(f func(uint8) (int16, error), list []uint8) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	ch := make(chan map[int]int16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v uint8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint8Int8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: (int8, error)
//	2. List
//
// Returns
//	New List of type (int8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Int8Err(f func(uint8) (int8, error), list []uint8) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	ch := make(chan map[int]int8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v uint8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint8UintErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: (uint, error)
//	2. List
//
// Returns
//	New List of type (uint, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint8UintErr(f func(uint8) (uint, error), list []uint8) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	ch := make(chan map[int]uint, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v uint8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint8Uint64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: (uint64, error)
//	2. List
//
// Returns
//	New List of type (uint64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Uint64Err(f func(uint8) (uint64, error), list []uint8) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	ch := make(chan map[int]uint64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v uint8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint8Uint32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: (uint32, error)
//	2. List
//
// Returns
//	New List of type (uint32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Uint32Err(f func(uint8) (uint32, error), list []uint8) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	ch := make(chan map[int]uint32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v uint8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint8Uint16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: (uint16, error)
//	2. List
//
// Returns
//	New List of type (uint16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Uint16Err(f func(uint8) (uint16, error), list []uint8) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	ch := make(chan map[int]uint16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v uint8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint8StrErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: (string, error)
//	2. List
//
// Returns
//	New List of type (string, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint8StrErr(f func(uint8) (string, error), list []uint8) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	ch := make(chan map[int]string, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v uint8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]string{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint8BoolErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: (bool, error)
//	2. List
//
// Returns
//	New List of type (bool, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint8BoolErr(f func(uint8) (bool, error), list []uint8) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	ch := make(chan map[int]bool, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v uint8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]bool{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint8Float32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: (float32, error)
//	2. List
//
// Returns
//	New List of type (float32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Float32Err(f func(uint8) (float32, error), list []uint8) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	ch := make(chan map[int]float32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float32, i int, v uint8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapUint8Float64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: uint8 output type: (float64, error)
//	2. List
//
// Returns
//	New List of type (float64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapUint8Float64Err(f func(uint8) (float64, error), list []uint8) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	ch := make(chan map[int]float64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float64, i int, v uint8) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapStrIntErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: (int, error)
//	2. List
//
// Returns
//	New List of type (int, error)
//	Empty list if all arguments are nil or either one is nil
func PMapStrIntErr(f func(string) (int, error), list []string) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	ch := make(chan map[int]int, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v string) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapStrInt64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: (int64, error)
//	2. List
//
// Returns
//	New List of type (int64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapStrInt64Err(f func(string) (int64, error), list []string) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	ch := make(chan map[int]int64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v string) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapStrInt32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: (int32, error)
//	2. List
//
// Returns
//	New List of type (int32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapStrInt32Err(f func(string) (int32, error), list []string) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	ch := make(chan map[int]int32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v string) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapStrInt16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: (int16, error)
//	2. List
//
// Returns
//	New List of type (int16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapStrInt16Err(f func(string) (int16, error), list []string) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	ch := make(chan map[int]int16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v string) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapStrInt8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: (int8, error)
//	2. List
//
// Returns
//	New List of type (int8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapStrInt8Err(f func(string) (int8, error), list []string) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	ch := make(chan map[int]int8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v string) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapStrUintErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: (uint, error)
//	2. List
//
// Returns
//	New List of type (uint, error)
//	Empty list if all arguments are nil or either one is nil
func PMapStrUintErr(f func(string) (uint, error), list []string) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	ch := make(chan map[int]uint, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v string) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapStrUint64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: (uint64, error)
//	2. List
//
// Returns
//	New List of type (uint64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapStrUint64Err(f func(string) (uint64, error), list []string) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	ch := make(chan map[int]uint64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v string) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapStrUint32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: (uint32, error)
//	2. List
//
// Returns
//	New List of type (uint32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapStrUint32Err(f func(string) (uint32, error), list []string) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	ch := make(chan map[int]uint32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v string) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapStrUint16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: (uint16, error)
//	2. List
//
// Returns
//	New List of type (uint16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapStrUint16Err(f func(string) (uint16, error), list []string) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	ch := make(chan map[int]uint16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v string) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapStrUint8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: (uint8, error)
//	2. List
//
// Returns
//	New List of type (uint8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapStrUint8Err(f func(string) (uint8, error), list []string) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	ch := make(chan map[int]uint8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v string) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapStrBoolErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: (bool, error)
//	2. List
//
// Returns
//	New List of type (bool, error)
//	Empty list if all arguments are nil or either one is nil
func PMapStrBoolErr(f func(string) (bool, error), list []string) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	ch := make(chan map[int]bool, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v string) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]bool{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapStrFloat32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: (float32, error)
//	2. List
//
// Returns
//	New List of type (float32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapStrFloat32Err(f func(string) (float32, error), list []string) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	ch := make(chan map[int]float32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float32, i int, v string) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapStrFloat64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: string output type: (float64, error)
//	2. List
//
// Returns
//	New List of type (float64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapStrFloat64Err(f func(string) (float64, error), list []string) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	ch := make(chan map[int]float64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float64, i int, v string) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapBoolIntErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: (int, error)
//	2. List
//
// Returns
//	New List of type (int, error)
//	Empty list if all arguments are nil or either one is nil
func PMapBoolIntErr(f func(bool) (int, error), list []bool) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	ch := make(chan map[int]int, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v bool) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapBoolInt64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: (int64, error)
//	2. List
//
// Returns
//	New List of type (int64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapBoolInt64Err(f func(bool) (int64, error), list []bool) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	ch := make(chan map[int]int64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v bool) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapBoolInt32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: (int32, error)
//	2. List
//
// Returns
//	New List of type (int32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapBoolInt32Err(f func(bool) (int32, error), list []bool) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	ch := make(chan map[int]int32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v bool) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapBoolInt16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: (int16, error)
//	2. List
//
// Returns
//	New List of type (int16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapBoolInt16Err(f func(bool) (int16, error), list []bool) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	ch := make(chan map[int]int16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v bool) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapBoolInt8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: (int8, error)
//	2. List
//
// Returns
//	New List of type (int8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapBoolInt8Err(f func(bool) (int8, error), list []bool) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	ch := make(chan map[int]int8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v bool) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapBoolUintErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: (uint, error)
//	2. List
//
// Returns
//	New List of type (uint, error)
//	Empty list if all arguments are nil or either one is nil
func PMapBoolUintErr(f func(bool) (uint, error), list []bool) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	ch := make(chan map[int]uint, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v bool) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapBoolUint64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: (uint64, error)
//	2. List
//
// Returns
//	New List of type (uint64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapBoolUint64Err(f func(bool) (uint64, error), list []bool) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	ch := make(chan map[int]uint64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v bool) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapBoolUint32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: (uint32, error)
//	2. List
//
// Returns
//	New List of type (uint32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapBoolUint32Err(f func(bool) (uint32, error), list []bool) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	ch := make(chan map[int]uint32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v bool) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapBoolUint16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: (uint16, error)
//	2. List
//
// Returns
//	New List of type (uint16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapBoolUint16Err(f func(bool) (uint16, error), list []bool) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	ch := make(chan map[int]uint16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v bool) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapBoolUint8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: (uint8, error)
//	2. List
//
// Returns
//	New List of type (uint8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapBoolUint8Err(f func(bool) (uint8, error), list []bool) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	ch := make(chan map[int]uint8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v bool) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapBoolStrErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: (string, error)
//	2. List
//
// Returns
//	New List of type (string, error)
//	Empty list if all arguments are nil or either one is nil
func PMapBoolStrErr(f func(bool) (string, error), list []bool) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	ch := make(chan map[int]string, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v bool) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]string{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapBoolFloat32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: (float32, error)
//	2. List
//
// Returns
//	New List of type (float32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapBoolFloat32Err(f func(bool) (float32, error), list []bool) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	ch := make(chan map[int]float32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float32, i int, v bool) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapBoolFloat64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: bool output type: (float64, error)
//	2. List
//
// Returns
//	New List of type (float64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapBoolFloat64Err(f func(bool) (float64, error), list []bool) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	ch := make(chan map[int]float64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float64, i int, v bool) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat32IntErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float32 output type: (int, error)
//	2. List
//
// Returns
//	New List of type (int, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32IntErr(f func(float32) (int, error), list []float32) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	ch := make(chan map[int]int, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v float32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat32Int64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float32 output type: (int64, error)
//	2. List
//
// Returns
//	New List of type (int64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Int64Err(f func(float32) (int64, error), list []float32) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	ch := make(chan map[int]int64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v float32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat32Int32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float32 output type: (int32, error)
//	2. List
//
// Returns
//	New List of type (int32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Int32Err(f func(float32) (int32, error), list []float32) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	ch := make(chan map[int]int32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v float32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat32Int16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float32 output type: (int16, error)
//	2. List
//
// Returns
//	New List of type (int16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Int16Err(f func(float32) (int16, error), list []float32) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	ch := make(chan map[int]int16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v float32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat32Int8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float32 output type: (int8, error)
//	2. List
//
// Returns
//	New List of type (int8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Int8Err(f func(float32) (int8, error), list []float32) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	ch := make(chan map[int]int8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v float32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat32UintErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float32 output type: (uint, error)
//	2. List
//
// Returns
//	New List of type (uint, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32UintErr(f func(float32) (uint, error), list []float32) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	ch := make(chan map[int]uint, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v float32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat32Uint64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float32 output type: (uint64, error)
//	2. List
//
// Returns
//	New List of type (uint64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Uint64Err(f func(float32) (uint64, error), list []float32) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	ch := make(chan map[int]uint64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v float32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat32Uint32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float32 output type: (uint32, error)
//	2. List
//
// Returns
//	New List of type (uint32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Uint32Err(f func(float32) (uint32, error), list []float32) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	ch := make(chan map[int]uint32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v float32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat32Uint16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float32 output type: (uint16, error)
//	2. List
//
// Returns
//	New List of type (uint16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Uint16Err(f func(float32) (uint16, error), list []float32) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	ch := make(chan map[int]uint16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v float32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat32Uint8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float32 output type: (uint8, error)
//	2. List
//
// Returns
//	New List of type (uint8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Uint8Err(f func(float32) (uint8, error), list []float32) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	ch := make(chan map[int]uint8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v float32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat32StrErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float32 output type: (string, error)
//	2. List
//
// Returns
//	New List of type (string, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32StrErr(f func(float32) (string, error), list []float32) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	ch := make(chan map[int]string, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v float32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]string{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat32BoolErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float32 output type: (bool, error)
//	2. List
//
// Returns
//	New List of type (bool, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32BoolErr(f func(float32) (bool, error), list []float32) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	ch := make(chan map[int]bool, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v float32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]bool{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat32Float64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float32 output type: (float64, error)
//	2. List
//
// Returns
//	New List of type (float64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat32Float64Err(f func(float32) (float64, error), list []float32) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	ch := make(chan map[int]float64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float64, i int, v float32) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat64IntErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float64 output type: (int, error)
//	2. List
//
// Returns
//	New List of type (int, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64IntErr(f func(float64) (int, error), list []float64) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	ch := make(chan map[int]int, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v float64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat64Int64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float64 output type: (int64, error)
//	2. List
//
// Returns
//	New List of type (int64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Int64Err(f func(float64) (int64, error), list []float64) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	ch := make(chan map[int]int64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v float64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat64Int32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float64 output type: (int32, error)
//	2. List
//
// Returns
//	New List of type (int32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Int32Err(f func(float64) (int32, error), list []float64) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	ch := make(chan map[int]int32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v float64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat64Int16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float64 output type: (int16, error)
//	2. List
//
// Returns
//	New List of type (int16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Int16Err(f func(float64) (int16, error), list []float64) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	ch := make(chan map[int]int16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v float64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat64Int8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float64 output type: (int8, error)
//	2. List
//
// Returns
//	New List of type (int8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Int8Err(f func(float64) (int8, error), list []float64) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	ch := make(chan map[int]int8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v float64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]int8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]int8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat64UintErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float64 output type: (uint, error)
//	2. List
//
// Returns
//	New List of type (uint, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64UintErr(f func(float64) (uint, error), list []float64) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	ch := make(chan map[int]uint, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v float64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat64Uint64Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float64 output type: (uint64, error)
//	2. List
//
// Returns
//	New List of type (uint64, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Uint64Err(f func(float64) (uint64, error), list []float64) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	ch := make(chan map[int]uint64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v float64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint64{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint64, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat64Uint32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float64 output type: (uint32, error)
//	2. List
//
// Returns
//	New List of type (uint32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Uint32Err(f func(float64) (uint32, error), list []float64) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	ch := make(chan map[int]uint32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v float64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat64Uint16Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float64 output type: (uint16, error)
//	2. List
//
// Returns
//	New List of type (uint16, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Uint16Err(f func(float64) (uint16, error), list []float64) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	ch := make(chan map[int]uint16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v float64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint16{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint16, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat64Uint8Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float64 output type: (uint8, error)
//	2. List
//
// Returns
//	New List of type (uint8, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Uint8Err(f func(float64) (uint8, error), list []float64) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	ch := make(chan map[int]uint8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v float64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]uint8{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]uint8, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat64StrErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float64 output type: (string, error)
//	2. List
//
// Returns
//	New List of type (string, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64StrErr(f func(float64) (string, error), list []float64) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	ch := make(chan map[int]string, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v float64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]string{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]string, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat64BoolErr applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float64 output type: (bool, error)
//	2. List
//
// Returns
//	New List of type (bool, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64BoolErr(f func(float64) (bool, error), list []float64) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	ch := make(chan map[int]bool, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v float64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]bool{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]bool, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}

// PMapFloat64Float32Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: float64 output type: (float32, error)
//	2. List
//
// Returns
//	New List of type (float32, error)
//	Empty list if all arguments are nil or either one is nil
func PMapFloat64Float32Err(f func(float64) (float32, error), list []float64) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	ch := make(chan map[int]float32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float32, i int, v float64) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]float32{i: r}
		}(&wg, ch, i, v)
	}

	wg.Wait()
	close(ch)
	close(errCh)
	
	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	newList := make([]float32, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}
