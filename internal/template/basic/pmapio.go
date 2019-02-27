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
}`
}
