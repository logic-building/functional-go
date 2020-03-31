// Code generated by 'gofp'. DO NOT EDIT.
package employer
import "sync" 
import "github.com/logic-building/functional-go/internal/employee" 

func Map(f func(Employer) Employer, list []Employer) []Employer {
	if f == nil {
		return []Employer{}
	}
	newList := make([]Employer, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapPtr(f func(*Employer) *Employer, list []*Employer) []*Employer {
	if f == nil {
		return []*Employer{}
	}
	newList := make([]*Employer, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func Filter(f func(Employer) bool, list []Employer) []Employer {
	if f == nil {
		return []Employer{}
	}
	var newList []Employer
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterPtr(f func(*Employer) bool, list []*Employer) []*Employer {
	if f == nil {
		return []*Employer{}
	}
	var newList []*Employer
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func Remove(f func(Employer) bool, list []Employer) []Employer {
	if f == nil {
		return []Employer{}
	}
	var newList []Employer
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func Some(f func(Employer) bool, list []Employer) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomePtr(f func(*Employer) bool, list []*Employer) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func Every(f func(Employer) bool, list []Employer) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

func DropWhile(f func(Employer) bool, list []Employer) []Employer {
	if f == nil {
		return []Employer{}
	}
	var newList []Employer
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]Employer, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func TakeWhile(f func(Employer) bool, list []Employer) []Employer {
	if f == nil {
		return []Employer{}
	}
	var newList []Employer
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

func PMap(f func(Employer) Employer, list []Employer) []Employer {
	if f == nil {
		return []Employer{}
	}

	ch := make(chan map[int]Employer)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]Employer, i int, v Employer) {
			defer wg.Done()
			ch <- map[int]Employer{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]Employer, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

func FilterMap(fFilter func(Employer) bool, fMap func(Employer) Employer, list []Employer) []Employer {
	if fFilter == nil || fMap == nil {
		return []Employer{}
	}
	var newList []Employer
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func FilterMapPtr(fFilter func(*Employer) bool, fMap func(*Employer) *Employer, list []*Employer) []*Employer {
	if fFilter == nil || fMap == nil {
		return []*Employer{}
	}
	var newList []*Employer
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func Rest(l []Employer) []Employer {
	if l == nil {
		return []Employer{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []Employer{}
	}

	newList := make([]Employer, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

func Reduce(f func(Employer, Employer) Employer, list []Employer, initializer ...Employer) Employer {
	var init Employer 
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}
	
	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return Reduce(f, list[1:], r)
}

// DropLast drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLast(list []Employer) []Employer {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []Employer{}
	}

	newList := make([]Employer, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}

func MapEmployee(f func(employee.Employee) employee.Employee, list []employee.Employee) []employee.Employee {
	if f == nil {
		return []employee.Employee{}
	}
	newList := make([]employee.Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func MapEmployeePtr(f func(*employee.Employee) *employee.Employee, list []*employee.Employee) []*employee.Employee {
	if f == nil {
		return []*employee.Employee{}
	}
	newList := make([]*employee.Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

func FilterEmployee(f func(employee.Employee) bool, list []employee.Employee) []employee.Employee {
	if f == nil {
		return []employee.Employee{}
	}
	var newList []employee.Employee
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func FilterEmployeePtr(f func(*employee.Employee) bool, list []*employee.Employee) []*employee.Employee {
	if f == nil {
		return []*employee.Employee{}
	}
	var newList []*employee.Employee
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func RemoveEmployee(f func(employee.Employee) bool, list []employee.Employee) []employee.Employee {
	if f == nil {
		return []employee.Employee{}
	}
	var newList []employee.Employee
	for _, v := range list {
		if !f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

func SomeEmployee(f func(employee.Employee) bool, list []employee.Employee) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func SomeEmployeePtr(f func(*employee.Employee) bool, list []*employee.Employee) bool {
	if f == nil {
		return false
	}
	for _, v := range list {
		if f(v) {
			return true
		}
	}
	return false
}

func EveryEmployee(f func(employee.Employee) bool, list []employee.Employee) bool {
	if f == nil || len(list) == 0 {
		return false
	}
	for _, v := range list {
		if !f(v) {
			return false
		}
	}
	return true
}

func DropWhileEmployee(f func(employee.Employee) bool, list []employee.Employee) []employee.Employee {
	if f == nil {
		return []employee.Employee{}
	}
	var newList []employee.Employee
	for i, v := range list {
		if !f(v) {
			listLen := len(list)
			newList = make([]employee.Employee, listLen-i)
			j := 0
			for i < listLen {
				newList[j] = list[i]
				i++
				j++
			}
			return newList
		}
	}
	return newList
}

func TakeWhileEmployee(f func(employee.Employee) bool, list []employee.Employee) []employee.Employee {
	if f == nil {
		return []employee.Employee{}
	}
	var newList []employee.Employee
	for _, v := range list {
		if !f(v) {
			return newList
		}
		newList = append(newList, v)
	}
	return newList
}

func PMapEmployee(f func(employee.Employee) employee.Employee, list []employee.Employee) []employee.Employee {
	if f == nil {
		return []employee.Employee{}
	}

	ch := make(chan map[int]employee.Employee)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]employee.Employee, i int, v employee.Employee) {
			defer wg.Done()
			ch <- map[int]employee.Employee{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]employee.Employee, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

func FilterMapEmployee(fFilter func(employee.Employee) bool, fMap func(employee.Employee) employee.Employee, list []employee.Employee) []employee.Employee {
	if fFilter == nil || fMap == nil {
		return []employee.Employee{}
	}
	var newList []employee.Employee
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func FilterMapEmployeePtr(fFilter func(*employee.Employee) bool, fMap func(*employee.Employee) *employee.Employee, list []*employee.Employee) []*employee.Employee {
	if fFilter == nil || fMap == nil {
		return []*employee.Employee{}
	}
	var newList []*employee.Employee
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

func RestEmployee(l []employee.Employee) []employee.Employee {
	if l == nil {
		return []employee.Employee{}
	}

	len := len(l)
	if len == 0 || len == 1 {
		return []employee.Employee{}
	}

	newList := make([]employee.Employee, len-1)

	for i, v := range l[1:] {
		newList[i] = v
	}

	return newList
}

func ReduceEmployee(f func(employee.Employee, employee.Employee) employee.Employee, list []employee.Employee, initializer ...employee.Employee) employee.Employee {
	var init employee.Employee 
	lenList := len(list)

	if initializer != nil {
		init = initializer[0]
	} else if lenList > 0 {
		init = list[0]
		if lenList == 1 {
			return list[0]
		}
		if lenList >= 2 {
			list = list[1:]
		}
	}
	
	if lenList == 0 {
		return init
	}
	r := f(init, list[0])
	return ReduceEmployee(f, list[1:], r)
}

// DropLastEmployee drops last item from the list and returns new list.
// Returns empty list if there is only one item in the list or list empty
func DropLastEmployee(list []employee.Employee) []employee.Employee {
	listLen := len(list)

	if list == nil || listLen == 0 || listLen == 1 {
		return []employee.Employee{}
	}

	newList := make([]employee.Employee, listLen-1)

	for i := 0; i < listLen-1; i++ {
		newList[i] = list[i]
	}
	return newList
}


// MapEmployerEmployee takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapEmployerEmployee(f func(Employer) employee.Employee, list []Employer) []employee.Employee {
	if f == nil {
		return []employee.Employee{}
	}
	newList := make([]employee.Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapEmployerEmployee takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapEmployerEmployeePtr(f func(*Employer) *employee.Employee, list []*Employer) []*employee.Employee {
	if f == nil {
		return []*employee.Employee{}
	}
	newList := make([]*employee.Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapEmployerEmployee applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: Employer output type: employee.Employee
//	2. List
//
// Returns
//	New List of type employee.Employee
//	Empty list if all arguments are nil or either one is nil
func PMapEmployerEmployee(f func(Employer) employee.Employee, list []Employer) []employee.Employee {
	if f == nil {
		return []employee.Employee{}
	}

	ch := make(chan map[int]employee.Employee)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]employee.Employee, i int, v Employer) {
			defer wg.Done()
			ch <- map[int]employee.Employee{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]employee.Employee, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapEmployerEmployee filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - Employer and returns true/false.
//	2. Function: takes Employer as argument and returns employee.Employee
// 	3. List of type Employer
//
// Returns:
//	New List of type employee.Employee
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapEmployerEmployee(fFilter func(Employer) bool, fMap func(Employer) employee.Employee, list []Employer) []employee.Employee {
	if fFilter == nil || fMap == nil {
		return []employee.Employee{}
	}
	var newList []employee.Employee
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// MapEmployerInt takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapEmployerInt(f func(Employer) int, list []Employer) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapEmployerInt takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapEmployerIntPtr(f func(*Employer) *int, list []*Employer) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapEmployerInt applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: Employer output type: int
//	2. List
//
// Returns
//	New List of type int
//	Empty list if all arguments are nil or either one is nil
func PMapEmployerInt(f func(Employer) int, list []Employer) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v Employer) {
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

// FilterMapEmployerInt filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - Employer and returns true/false.
//	2. Function: takes Employer as argument and returns int
// 	3. List of type Employer
//
// Returns:
//	New List of type int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapEmployerInt(fFilter func(Employer) bool, fMap func(Employer) int, list []Employer) []int {
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

// MapEmployeeEmployer takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapEmployeeEmployer(f func(employee.Employee) Employer, list []employee.Employee) []Employer {
	if f == nil {
		return []Employer{}
	}
	newList := make([]Employer, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapEmployeeEmployer takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapEmployeeEmployerPtr(f func(*employee.Employee) *Employer, list []*employee.Employee) []*Employer {
	if f == nil {
		return []*Employer{}
	}
	newList := make([]*Employer, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapEmployeeEmployer applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: employee.Employee output type: Employer
//	2. List
//
// Returns
//	New List of type Employer
//	Empty list if all arguments are nil or either one is nil
func PMapEmployeeEmployer(f func(employee.Employee) Employer, list []employee.Employee) []Employer {
	if f == nil {
		return []Employer{}
	}

	ch := make(chan map[int]Employer)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]Employer, i int, v employee.Employee) {
			defer wg.Done()
			ch <- map[int]Employer{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]Employer, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapEmployeeEmployer filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - employee.Employee and returns true/false.
//	2. Function: takes employee.Employee as argument and returns Employer
// 	3. List of type employee.Employee
//
// Returns:
//	New List of type Employer
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapEmployeeEmployer(fFilter func(employee.Employee) bool, fMap func(employee.Employee) Employer, list []employee.Employee) []Employer {
	if fFilter == nil || fMap == nil {
		return []Employer{}
	}
	var newList []Employer
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// MapEmployeeInt takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapEmployeeInt(f func(employee.Employee) int, list []employee.Employee) []int {
	if f == nil {
		return []int{}
	}
	newList := make([]int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapEmployeeInt takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapEmployeeIntPtr(f func(*employee.Employee) *int, list []*employee.Employee) []*int {
	if f == nil {
		return []*int{}
	}
	newList := make([]*int, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapEmployeeInt applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: employee.Employee output type: int
//	2. List
//
// Returns
//	New List of type int
//	Empty list if all arguments are nil or either one is nil
func PMapEmployeeInt(f func(employee.Employee) int, list []employee.Employee) []int {
	if f == nil {
		return []int{}
	}

	ch := make(chan map[int]int)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]int, i int, v employee.Employee) {
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

// FilterMapEmployeeInt filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - employee.Employee and returns true/false.
//	2. Function: takes employee.Employee as argument and returns int
// 	3. List of type employee.Employee
//
// Returns:
//	New List of type int
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapEmployeeInt(fFilter func(employee.Employee) bool, fMap func(employee.Employee) int, list []employee.Employee) []int {
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

// MapIntEmployer takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntEmployer(f func(int) Employer, list []int) []Employer {
	if f == nil {
		return []Employer{}
	}
	newList := make([]Employer, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntEmployer takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntEmployerPtr(f func(*int) *Employer, list []*int) []*Employer {
	if f == nil {
		return []*Employer{}
	}
	newList := make([]*Employer, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapIntEmployer applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: Employer
//	2. List
//
// Returns
//	New List of type Employer
//	Empty list if all arguments are nil or either one is nil
func PMapIntEmployer(f func(int) Employer, list []int) []Employer {
	if f == nil {
		return []Employer{}
	}

	ch := make(chan map[int]Employer)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]Employer, i int, v int) {
			defer wg.Done()
			ch <- map[int]Employer{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]Employer, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapIntEmployer filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int and returns true/false.
//	2. Function: takes int as argument and returns Employer
// 	3. List of type int
//
// Returns:
//	New List of type Employer
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntEmployer(fFilter func(int) bool, fMap func(int) Employer, list []int) []Employer {
	if fFilter == nil || fMap == nil {
		return []Employer{}
	}
	var newList []Employer
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

// MapIntEmployee takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntEmployee(f func(int) employee.Employee, list []int) []employee.Employee {
	if f == nil {
		return []employee.Employee{}
	}
	newList := make([]employee.Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// MapIntEmployee takes two inputs -
// 1. Function 2. List. Then It returns a new list after applying the function on each item of the list
func MapIntEmployeePtr(f func(*int) *employee.Employee, list []*int) []*employee.Employee {
	if f == nil {
		return []*employee.Employee{}
	}
	newList := make([]*employee.Employee, len(list))
	for i, v := range list {
		newList[i] = f(v)
	}
	return newList
}

// PMapIntEmployee applies the function(1st argument) on each item of the list and returns new list.
// Run in parallel. no_of_goroutines = no_of_items_in_list
//
// Takes 2 inputs
//	1. Function - takes 1 input type: int output type: employee.Employee
//	2. List
//
// Returns
//	New List of type employee.Employee
//	Empty list if all arguments are nil or either one is nil
func PMapIntEmployee(f func(int) employee.Employee, list []int) []employee.Employee {
	if f == nil {
		return []employee.Employee{}
	}

	ch := make(chan map[int]employee.Employee)
	var wg sync.WaitGroup

	for i, v := range list {
		wg.Add(1)

		go func(wg *sync.WaitGroup, ch chan map[int]employee.Employee, i int, v int) {
			defer wg.Done()
			ch <- map[int]employee.Employee{i: f(v)}
		}(&wg, ch, i, v)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	newList := make([]employee.Employee, len(list))
	for m := range ch {
		for k, v := range m {
			newList[k] = v
		}
	}
	return newList
}

// FilterMapIntEmployee filters given list, then apply function(2nd argument) on each item in the list and returns a new list
// Takes 3 inputs
//	1. Function: takes one input type - int and returns true/false.
//	2. Function: takes int as argument and returns employee.Employee
// 	3. List of type int
//
// Returns:
//	New List of type employee.Employee
//  Empty list if all there parameters are nil or either of parameter is nil
func FilterMapIntEmployee(fFilter func(int) bool, fMap func(int) employee.Employee, list []int) []employee.Employee {
	if fFilter == nil || fMap == nil {
		return []employee.Employee{}
	}
	var newList []employee.Employee
	for _, v := range list {
		if fFilter(v) {
			newList = append(newList, fMap(v))
		}
	}
	return newList
}

