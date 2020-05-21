package basic

// PMapIO is template to generate itself for different combination of data type.
func PMapIO() string {
	return `
// PMap<FINPUT_TYPE><FOUTPUT_TYPE> applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: <INPUT_TYPE> output type: <OUTPUT_TYPE>
//	2. List
//
// Returns
//	New List of type <OUTPUT_TYPE>
//	Empty list if all arguments are nil or either one is nil
func PMap<FINPUT_TYPE><FOUTPUT_TYPE>(f func(<INPUT_TYPE>) <OUTPUT_TYPE>, list []<INPUT_TYPE>) []<OUTPUT_TYPE> {
	if f == nil {
		return []<OUTPUT_TYPE>{}
	}

	ch := make(chan map[int]<OUTPUT_TYPE>)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]<OUTPUT_TYPE>, i int, v <INPUT_TYPE>) {
			defer wg.Done()
			ch <- map[int]<OUTPUT_TYPE>{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]<OUTPUT_TYPE>, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}
`
}

// PMapIOErr is template to generate itself for different combination of data type.
func PMapIOErr() string {
	return `
// PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err applies the function(1st argument) on each item of the list and returns new list and error.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: <INPUT_TYPE> output type: (<OUTPUT_TYPE>, error)
//	2. List
//
// Returns
//	New List of type (<OUTPUT_TYPE>, error)
//	Empty list if all arguments are nil or either one is nil
func PMap<FINPUT_TYPE><FOUTPUT_TYPE>Err(f func(<INPUT_TYPE>) (<OUTPUT_TYPE>, error), list []<INPUT_TYPE>) ([]<OUTPUT_TYPE>, error) {
	if f == nil {
		return []<OUTPUT_TYPE>{}, nil
	}

	ch := make(chan map[int]<OUTPUT_TYPE>, len(list))
	errCh := make(chan error, len(list))
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]<OUTPUT_TYPE>, i int, v <INPUT_TYPE>) {
			defer wg.Done()
			if len(errCh) >= 1 {
				return
			}
			r, err := f(v)
			if err != nil {
				errCh <- err
				return
			}
			ch <- map[int]<OUTPUT_TYPE>{i: r}
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

	newList := make([]<OUTPUT_TYPE>, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList, nil
}
`
}
