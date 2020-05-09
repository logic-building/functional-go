package basic

// Map<FTYPE>Ptr applies the function(1st argument) on each item of the list and returns new list
func PMapPtr() string {
	return `
// PMap<FTYPE>Ptr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMap<FTYPE>Ptr(f func(*<TYPE>) *<TYPE>, list []*<TYPE>) []*<TYPE> {
	if f == nil {
		return []*<TYPE>{}
	}

	ch := make(chan map[int]*<TYPE>)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*<TYPE>, i int, v *<TYPE>) {
			defer wg.Done()
			ch <- map[int]*<TYPE>{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]*<TYPE>, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
`
}

// Map<FTYPE>PtrErr applies the function(1st argument) on each item of the list and returns new list
func PMapPtrErr() string {
	return `
// PMap<FTYPE>PtrErr applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMap<FTYPE>PtrErr(f func(*<TYPE>) (*<TYPE>, error), list []*<TYPE>) ([]*<TYPE>, error) {
	if f == nil {
		return []*<TYPE>{}, nil
	}

	ch := make(chan map[int]*<TYPE>, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]*<TYPE>, i int, v *<TYPE>) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]*<TYPE>{i: r}
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

	newList := make([]*<TYPE>, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}
`
}

// Map<FTYPE>PtrErr applies the function(1st argument) on each item of the list and returns new list
func PMapErr() string {
	return `
// PMap<FTYPE>Err applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
func PMap<FTYPE>Err(f func(<TYPE>) (<TYPE>, error), list []<TYPE>) ([]<TYPE>, error) {
	if f == nil {
		return []<TYPE>{}, nil
	}

	ch := make(chan map[int]<TYPE>, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]<TYPE>, i int, v <TYPE>) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]<TYPE>{i: r}
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

	newList := make([]<TYPE>, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}
`
}
