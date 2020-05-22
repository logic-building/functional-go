[![Docs](https://godoc.org/github.com/logic-building/functional-go?status.svg)](https://godoc.org/github.com/logic-building/functional-go)
[![CircleCI](https://circleci.com/gh/logic-building/functional-go.svg?style=svg)](https://circleci.com/gh/logic-building/functional-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/logic-building/functional-go)](https://goreportcard.com/report/github.com/logic-building/functional-go)

# functional-go : Functional programming in golang. This library is inspired by clojure.
## Simple but functional

### Install
```
go get github.com/logic-building/functional-go/fp/
go get github.com/logic-building/functional-go/set/
go install github.com/logic-building/functional-go/gofp

go get -u github.com/logic-building/functional-go/fp/
go get -u github.com/logic-building/functional-go/set/
go install github.com/logic-building/functional-go/gofp

```

### dep Gopkg.toml entry
```
[[constraint]]
name = "github.com/logic-building/functional-go"
version = "8.7.0"
```

### Quick Start
#### For Data types available in golang
```
import "github.com/logic-building/functional-go/fp"

fp.MapInt(square, []int{1, 2, 3, 4}) // Returns: [1 4 9 16]

func square(num int) int {
	return num * num
}

```
#### Four variants of function 1 is given above and 3 are given below
##### MapInt, MapIntPtr, MapIntErr, MapIntPtrErr
### MapIntPtr
```
package main

import (
	"fmt"
	"github.com/logic-building/functional-go/fp"
)

var v1 int = 1
var v2 int = 2
var v3 int = 3
var v4 int = 4

func main() {
	result := fp.MapIntPtr(square, []*int{&v1, &v2, &v3, &v4})
	fmt.Println(*result[0], *result[1], *result[2], *result[3])
}

func square(num *int) *int {
	r := *num * *num
	return &r
}

/*
output:
1 4 9 16
*/
```

### MapIntErr
```
package main

import (
	"errors"
	"fmt"
	"github.com/logic-building/functional-go/fp"
)

func main() {
	result, _ := fp.MapIntErr(square, []int{1, 2, 3, 4})
	fmt.Println(result[0], result[1], result[2], result[3])
}

func square(num int) (int, error) {
	if num == -1 {
		return 0, errors.New("-1 is not valid")
	}
	r := num * num
	return r, nil
}

/*
output:
1 4 9 16
*/
```

### MapIntPtrErr
```
package main

import (
	"errors"
	"fmt"

	"github.com/logic-building/functional-go/fp"
)

func main() {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4

	result, _ := fp.MapIntPtrErr(square, []*int{&v1, &v2, &v3, &v4})
	fmt.Println(*result[0], *result[1], *result[2], *result[3])
}

func square(num *int) (*int, error) {
	if *num == -1 {
		return nil, errors.New("-1 is not valid")
	}
	r := *num * *num
	return &r, nil
}
/*
output:
1 4 9 16
*/
```

####  Generate functional code locally in project for user defined data type
```
Design 1: Functional code distributed within different package
Design 2: All functional code in one place
```
### Design 1:  Functional code distributed within different package
```
1. Install "gofp" to generate code
   go get github.com/logic-building/functional-go/gofp
   go get -u github.com/logic-building/functional-go/gofp

   go install github.com/logic-building/functional-go/gofp

2. Add this line in a file where user defined data type exists
   //go:generate gofp -destination <file> -pkg <pkg> -type <Types separated by comma>

   // If the user defined type has to be imported, then use options "-imports" and value will be comma separated.
   // See the file internal/employer/employer.go for the example

example:
    package employee

   //go:generate gofp -destination fp.go -pkg employee -type "Employee, Teacher"
   type Employee struct {
   	id     int
   	name   string
   	salary float64
   }

   type Teacher struct {
   	id     int
   	name   string
   	salary float64
   }

Note:
   A. fp.go               :  generated code
   B. employee            :  package name
   C. "Employee, Teacher" :  User defined data types
   
Note: Another option : -mapfun "true"
   This allows to include function - merge,zip
   Caution: It will complain at runtime if struct contains slice or array 

3. Generate functional code
   go generate ./...

4. Now write your code

    emp1 := employee.Employee{1, "A", 1000}
   	emp2 := employee.Employee{2, "B", 1000}
   	emp3 := employee.Employee{3, "C", 1000}

   	empList := []employee.Employee{emp1, emp2, emp3}

   	newEmpList := employee.Map(incrementSalary, empList) //  Returns: [{1 A 1500} {2 B 1500} {3 C 1500}]

   func incrementSalary(emp employee.Employee) employee.Employee {
        emp.Salary = emp.Salary + 500
        return emp
   }

```

### Design 2: All functional code in one place
```
1. Install "gofp" to generate code
   go get -u github.com/logic-building/functional-go/gofp
   go install github.com/logic-building/functional-go/gofp

2. create package "gfp" : See the example - internal/gfp/gfp.go
   and two lines given below

   package gfp
   //go:generate gofp -destination fp.go -pkg gfp -type "employer.Employer, employee.Employee, int" -imports "github.com/logic-building/functional-go/internal/employee, github.com/logic-building/functional-go/internal/employer"


   -destination fp.go  : Generated file in package "gfp" will be "fp.go" which contains all the functional code
   -pkg gfp : Saying that package is "gfp"
   -type "employer.Employer, employee.Employee, int" : Generate functions for the combinations of these three data types
   -imports "github.com/logic-building/functional-go/internal/employee, github.com/logic-building/functional-go/internal/employer" : import statements for two user defined types
                "employee.Employee" and employer.Employer. The value for -imports should be comma separated


3. Generate functional code : Run the command given below from the root folder of your project
   go generate ./...

4. Now write your code

    gfp.MapEmployer
    or
    gfp.MapEmployee
    or
    gfp.MapEmployerEmployee

```


### Contains functions
```
Takes list as argument and returns Distinct list.  Returns empty list if 2nd parameter is empty or nil
DistinctInt
DistinctInt64
DistinctInt32
DistinctInt16
DistinctInt8
DistinctUint
DistinctUint64
DistinctUint32
DistinctUint16
DistinctUint8
DistinctFloat64
DistinctFloat32
DistinctStr

add Ptr at end for pointer versioon function
eg. DistinctIntPtr

Takes function as argument and apply it on each item in the list and return list
MapInt
MapInt64
MapInt32
MapInt16
MapInt8
MapUint
MapUint64
MapUint32
MapUint16
MapUint8
MapFloat64
MapFloat32
MapStr
   And also all basic combination such as
MapStrInt64
MapInt64Str
...

Pmap : For parallel processing
PmapInt
PmapInt64
PmapInt32
PmapInt16
PmapInt8
PmapUint
PmapUint64
PmapUint32
PmapUint16
PmapUint8
PmapFloat64
PmapFloat32
PmapStr
    And also all basic combination such as
PMapStrInt64
PMapInt64Str

add Ptr at end for pointer versioon function
eg. PMapInt64StrPtr

Takes function as argument and apply it on each item in the list and return filtered list
FilterInt
FilterInt64
FilterInt32
FilterInt16
FilterInt8
FilterUint
FilterUint64
FilterUint32
FilterUint16
FilterUint8
FilterFloat64
FilterFloat32
FilterStr

add Ptr at end for pointer versioon function
eg. FilterIntPtr

Takes two functions as argument and apply them on each item in the list and return the filtered list
FilterMapInt
FilterMapInt64
FilterMapInt32
FilterMapInt16
FilterMapInt8
FilterMapUint
FilterMapUint64
FilterMapUint32
FilterMapUint16
FilterMapUint8
FilterMapFloat64
FilterMapFloat32
FilterMapStr
    And also all basic combination such as
FilterMapStrInt64
FilterMapInt64Str

add Ptr at end for pointer versioon function
eg. FilterMapIntPtr

Takes list as parameter and returns new list which includes all the items except 1st one.
    Input: list
    output: New list includes all the items except 1st one.
            Returns new empty list if the input is either nil or empty or length 1 list
    Ex.
		list := []int{1, 2, 3, 4, 5}
		RestInt(list) // returns [2, 3, 4, 5]

RestInt
RestInt64
RestInt32
RestInt16
RestInt8
RestUint
RestUint64
RestUint32
RestUint16
RestUint8
RestFloat64
RestFloat32
RestStr

add Ptr at end for pointer versioon function
eg. RestIntPtr


Takes function as argument and apply it on each item in the list and return true/false. Returns false if 2nd argument is empty
EveryBool : Takes two arguments
             a. function without arguments and returns bool
             b. list of boolean values

        return:
           true if all the values in the list are true, otherwise false
           See example below


EveryInt :
       Takes two arguments
        a. function with 1 argument type int and returns bool
        b. list of int

        return: True if function-1st argument returns true for each item in the list, otherwise false
EveryInt64
EveryInt32
EveryInt16
EveryInt8
EveryUint
EveryUint64
EveryUint32
EveryUint16
EveryUint8
EveryFloat64
EveryFloat32
EveryStr

add Ptr at end for pointer versioon function
eg. EveryIntPtr

Takes a number and a list as arguments and returns true if number exists in the list else returns false
ExistsInt
ExistsInt64
ExistsInt32
ExistsInt16
ExistsInt8
ExistsUint
ExistsUint64
ExistsUint32
ExistsUint16
ExistsUint8
ExistsFloat64
ExistsFloat32
ExistsStr

add Ptr at end for pointer versioon function
eg. ExistsIntPtr

Takes list as argument and return max number
- Returns 0 if list is either empty or nil

MaxInt
MaxInt64
MaxInt32
MaxInt16
MaxInt8
MaxUint
MaxUint64
MaxUint32
MaxUint16
MaxUint8
MaxFloat64
MaxFloat32

add Ptr at end for pointer versioon function
eg. MaxIntPtr

Takes list as argument and return min number
- Returns 0 if list is either empty or nil
MinInt
MinInt64
MinInt32
MinInt8
MinUint
MinUint64
MinUint32
MinUint16
MinUint8
MinFloat64
MinFloat32

Takes List as argument and returns Min and Max
   - Returns 0, 0 if list is either empty or nil
   - Returns same result for min and max if there is only one item in the list
MinMaxInt
MinMaxInt64
MinMaxInt32
MinMaxInt8
MinMaxUint
MinMaxUint64
MinMaxUint32
MinMaxUint16
MinMaxUint8
MinMaxFloat64
MinMaxFloat32

Returns a new list after dropping single item or multiple items 
DropInt
DropInts
DropInt64
DropInts64
DropInt32
DropInts32
DropInt16
DropInts16
DropInt8
DropInts8
DropUint
DropUInts
DropUint64
DropUints64
DropUint32
DropUints32
DropUint16
DropUints16
DropUint8
DropUints8
DropFloat64
DropFloats64
DropFloat32
DropFloats32
DropStr
DropStrIgnoreCase
DropStrs
DropStrsIgnoreCase

DropWhile:
  Drop the items in the list based on condition
  Takes 2 inputs
    a. Function : takes 1 input - int or float or string and return true/false
    b. list
  returns:
      NewList

      see example

DropWhileInt
DropWhileInt64
DropWhileInt32
DropWhileInt16
DropWhileInt8
DropWhileUint
DropWhileUint64
DropWhileUint32
DropWhileUint16
DropWhileUint8
DropWhileFloat64
DropWhileFloat32
DropWhileStr

TakeWhile:
  take the items in the list based on condition

Takes 2 inputs
  a. Function : takes 1 input - int or float or string and return true/false
  b. list

  returns:
     NewList

        see example

TakeWhileInt
TakeWhileInt64
TakeWhileInt32
TakeWhileInt16
TakeWhileInt8
TakeWhileUint
TakeWhileUint64
TakeWhileUint32
TakeWhileUint16
TakeWhileUint8
TakeWhileFloat64
TakeWhileFloat32
TakeWhileStr

Takes function  and list as arguments and remove items based on function passed and return new list. See example
RemoveInt
RemoveInts
RemoveInt64
RemoveInts64
RemoveInt32
RemoveInts32
RemoveInt16
RemoveInts16
RemoveInt8
RemoveInts8
RemoveUint
RemoveUInts
RemoveUint64
RemoveUints64
RemoveUint32
RemoveUints32
RemoveUint16
RemoveUints16
RemoveUint8
RemoveUints8
RemoveFloat64
RemoveFloats64
RemoveFloat32
RemoveFloats32
RemoveStr
RemoveStrIgnoreCase
RemoveStrs
RemoveStrsIgnoreCase

Range function
Takes 3 inputs:
   a. lower
   b. upper
   c. hops(optional)

      Returns list of range between lower and upper value
      Returns empty list if 3rd arguments is either of the following
           a. 0
           b. Negative number

      Note: 3rd argument is considered as hops value if 3 or more than 3 arguments are passed.

RangeInt
RangeInt64
RangeInt32
RangeInt16
RangeInt8
RangeUint
RangeUint64
RangeUint32
RangeUint16
RangeUint8

SomeInt: takes two parameters
    A. Function - one parameter and returns boolean
    B. list
Returns:
   Bool

Ex.
    SomeInt(isEven, []int{1, 2, 3, 5}) : returns true
    SomeInt(nil, nil) : returns false
    SomeInt(isEven, []int{}) : returns false

SomeInt
SomeInt64
SomeInt32
SomeInt16
SomeInt8
SomeUint
SomeUint64
SomeUint32
SomeUint16
SomeUint8
SomeFloat64
SomeFloat32
SomeStr

Takes three inputs:
   A. function - takes two argument
   B. list
   c. initializer - optional
 Returns:
    single value

 Example
   list := []int{1, 2, 3, 4, 5}
   ReduceInt(plusInt, list) // returns: 15

   func plusInt(num1, num2 int) int {
	    return num1 + num2
   }

ReduceInt
ReduceInt64
ReduceInt32
ReduceInt16
ReduceInt8
ReduceUint
ReduceUint64
ReduceUint32
ReduceUint16
ReduceUint8
ReduceFloat64
ReduceFloat32
ReduceStr

Merge : Takes two inputs - map1 and map2 and returns new map after merging map1 and map2

MergeInt - takes input1: map<int,int>, input2: map<int, int>
MergeInt64
MergeInt32
MergeInt16
MergeInt8
MergeUint
MergeUInt64
MergeUint32
MergeUint16
MergeUint8
MergeFloat64
MergeFloat32
MergeBool
MergeStr

and there are also functions available for different combination of inputs in map

MergeIntStr - takes input1: map<int,string>, input2: map<int, string>
MergeStrFloat64
...

Example:
    map1 := map[int]int{1: 10, 2: 20, 3: 30}
    map2 := map[int]int{4: 40, 5: 50, 3: 300}
    MergeIntInt(map1, map2) // Returns: map[1:10 2:20 4:40 5:50 3:300]


Zip: Takes two inputs - list1 and list2 and returns new map after zipping two lists

ZipInt  : takes input1 - type "int", input2 - type "int"
ZipInt64 : takes input1 - type "int64", input2 - type "int64"
ZipInt32
ZipInt16
ZipInt8
ZipUint
ZipUint64
ZipUint32
ZipUint16
ZipUint8
ZipFloat64
ZipFloat32
ZipStr
ZipBool

and there are also functions available for different combination of inputs,

ZipIntStr : takes input1 - type "int", input2 - type "string"
ZipStrFloat64 : takes input1 - type "string", input2 - type "float64"
 ...

    Example:
        list1 := []int{1, 2, 3, 4}
        list2 := []int{10, 20, 30, 40}

        ZipInt(list1, list2) // returns map[1: 10, 2: 20, 3: 30, 4: 40]

Set operations: Add, Remove, Clear, GetList, NewSetInt, Join, Intersection, Minus, Subset, Superset
SetInt
SetIntSync
SetInt64
SetInt64Sync
SetInt32
SetInt32Sync
SetInt16
SetInt16Sync
SetInt8
SetInt8Sync
SetUint
SetUintSync
SetUint64
SetUint64Sync
SetUint32
SetUint32Sync
SetUint16
SetUint16Sync
SetUint8
SetUint8Sync
SetUintFloat64
SetUintFloat64Sync
SetUintFloat32
SetUintFloat32Sync
SetStr
SetStrSync
```

### Example1 - Map : return the list of the square of each items in the list
```
squareList := fp.MapInt(squareInt, []int{1, 2, 3}) // see the map_test.go for detail

func squareInt(num int) int {
	return num * num
}


output
-------
[1, 4, 9]

```

### Example2 - Filter: filter all the even numbers in the list
```
filteredList := fp.FilterInt(isEven, []int{1, 2, 3, 4})

func isEven(num int) bool {
	return num%2 == 0
}

output:
[2, 4]

```

### Example3 - fp.FilterMap: Multiply all positive numbers in the list by 2
```
filteredList := FilterMapInt(isPositive, multiplyBy2, []int{-1, 0, 2, 4})

func isPositive(num int) bool {
	return num > 0
}
func multiplyBy2(num int) int {
	return num * 2
}

output:
[4, 8]
```

### Example4 - Every: Test if every number in the list is even
```
list := []bool{true, true, true, true}
fp.EveryBool(fp.True, list)  // Returns true

list1 := []int{8, 2, 10, 4}
fp.EveryInt(isEven, list1) // Returns true
```

### Example5 - Exists: Test if number presents in the list
```
list1 := []int{8, 2, 10, 4}
fp.ExistsInt(8, list1) // returns true
```

### Example6 - Max: Get max number in the list
```
list := []int{8, 2, 10, 4}
max := fp.MaxInt(list) // returns 10
```

### Example7 - Min: Get min number in the list
```
list := []int{8, 2, 10, 4}
min := fp.MinInt(list) // returns 2
```

### Example8 - Returns a new list after dropping the given item
```
newList := fp.DropInt(1, []int{1, 2, 3, 1}) // returns [2, 3]

To drop multiple items:
newList := fp.DropInts([]int{1, 2}, []int{1, 2, 3, 1}) // returns [3]

```

### Example9 - Distinct: returns distinct list
```
list := []int{8, 2, 8, 0, 2, 0}
distinct := fp.DistinctInt(list) // returns [8, 2, 0]
```

### Example10 - Set : Create set objects and apply union operation
```
	mySet1 := set.NewInt([]int{10, 20, 30, 20})
	mySet2 := set.NewInt([]int{30, 40, 50})
	mySet3 := mySet1.Union(mySet2) // Returns  [10, 20, 30, 40, 50]
```

### Example11 - Range : accepts lower and end int and returns range list
```
	fp.Range(1, 5) // returns [1, 2, 3, 4]
	fp.Range(1, 5, 2) // returns [1, 3]
```

### Example12 - Remove :
```
    NewList := RemoveInt(isEven, []int{1, 2, 3, 4}) // returns [1, 3]

    func isEven(num int) bool {
    	return num%2 == 0
    }

```

### Example13 - DropWhileInt :
```
    NewList := DropWhile(isEven, []int{4, 2, 3, 4}) // returns [3, 4]

    func isEven(num int) bool {
    	return num%2 == 0
    }

```

### Example14 - TakeWhileInt :
```
    NewList := TakeWhile(isEven, []int{4, 2, 3, 4}) // returns [4, 2]

    func isEven(num int) bool {
    	return num%2 == 0
    }

```

### Test Coverage
```
Tests Passed : 1217
ok  	functional-go/fp	0.015s	coverage: 100.0% of statements
ok  	functional-go/set	0.031s	coverage: 100.0% of statements
```

### BenchMark test:
```
  Model Identifier:	MacBookPro11,5
  Processor Name:	Intel Core i7
  Processor Speed:	2.5 GHz
  Number of Processors:	1
  Total Number of Cores:	4
  Memory:	16 GB

  Note: Bench result can vary bassed on the function passed. Check the test file for the detail.
```

```
goos: darwin
goarch: amd64
pkg: functional-go/fp
BenchmarkFilterInt64-8                   	 1000000	      5030 ns/op
BenchmarkFilterMapInt64-8                	 1000000	      5865 ns/op
BenchmarkMapInt64_PassedMethod_1_Arg-8   	 1000000	      3597 ns/op
BenchmarkMapInt64_PassedMethod_2_Arg-8   	 1000000	      3622 ns/op
BenchmarkMapStr-8                        	 1000000	     49438 ns/op
PASS
ok  	functional-go/fp	67.567s
```
