package fp

import "sync"

// PMapIntErr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapIntErr(f func(int) (int, error), list []int) ([]int, error) {
	if f == nil {
		return []int{}, nil
	}

	ch := make(chan map[int]int, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v int) {
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

// PMapInt64Err applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapInt64Err(f func(int64) (int64, error), list []int64) ([]int64, error) {
	if f == nil {
		return []int64{}, nil
	}

	ch := make(chan map[int]int64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int64, i int, v int64) {
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

// PMapInt32Err applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapInt32Err(f func(int32) (int32, error), list []int32) ([]int32, error) {
	if f == nil {
		return []int32{}, nil
	}

	ch := make(chan map[int]int32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int32, i int, v int32) {
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

// PMapInt16Err applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapInt16Err(f func(int16) (int16, error), list []int16) ([]int16, error) {
	if f == nil {
		return []int16{}, nil
	}

	ch := make(chan map[int]int16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int16, i int, v int16) {
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

// PMapInt8Err applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapInt8Err(f func(int8) (int8, error), list []int8) ([]int8, error) {
	if f == nil {
		return []int8{}, nil
	}

	ch := make(chan map[int]int8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int8, i int, v int8) {
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

// PMapUintErr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapUintErr(f func(uint) (uint, error), list []uint) ([]uint, error) {
	if f == nil {
		return []uint{}, nil
	}

	ch := make(chan map[int]uint, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint, i int, v uint) {
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

// PMapUint64Err applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapUint64Err(f func(uint64) (uint64, error), list []uint64) ([]uint64, error) {
	if f == nil {
		return []uint64{}, nil
	}

	ch := make(chan map[int]uint64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint64, i int, v uint64) {
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

// PMapUint32Err applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapUint32Err(f func(uint32) (uint32, error), list []uint32) ([]uint32, error) {
	if f == nil {
		return []uint32{}, nil
	}

	ch := make(chan map[int]uint32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint32, i int, v uint32) {
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

// PMapUint16Err applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapUint16Err(f func(uint16) (uint16, error), list []uint16) ([]uint16, error) {
	if f == nil {
		return []uint16{}, nil
	}

	ch := make(chan map[int]uint16, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint16, i int, v uint16) {
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

// PMapUint8Err applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapUint8Err(f func(uint8) (uint8, error), list []uint8) ([]uint8, error) {
	if f == nil {
		return []uint8{}, nil
	}

	ch := make(chan map[int]uint8, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]uint8, i int, v uint8) {
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

// PMapStrErr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapStrErr(f func(string) (string, error), list []string) ([]string, error) {
	if f == nil {
		return []string{}, nil
	}

	ch := make(chan map[int]string, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]string, i int, v string) {
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

// PMapBoolErr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapBoolErr(f func(bool) (bool, error), list []bool) ([]bool, error) {
	if f == nil {
		return []bool{}, nil
	}

	ch := make(chan map[int]bool, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]bool, i int, v bool) {
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

// PMapFloat32Err applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapFloat32Err(f func(float32) (float32, error), list []float32) ([]float32, error) {
	if f == nil {
		return []float32{}, nil
	}

	ch := make(chan map[int]float32, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float32, i int, v float32) {
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

// PMapFloat64Err applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMapFloat64Err(f func(float64) (float64, error), list []float64) ([]float64, error) {
	if f == nil {
		return []float64{}, nil
	}

	ch := make(chan map[int]float64, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]float64, i int, v float64) {
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
