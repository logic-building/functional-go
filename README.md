# functional-go : Functional programming in golang. This library is inspired by clojure.
## Simple but functional

### Install
```
go get github.com/logic-building/functional-go/fp/
go get github.com/logic-building/functional-go/set/

```

### dep Gopkg.toml entry
```
[[constraint]]
name = "github.com/logic-building/functional-go"
version = "5.0"
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

Takes number and list as arguments and returns true/false if number exists in the list
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

Takes number or string and list as arguments and Drop the given item and returns the list
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

### Example5 - Some: Test if number presents in the list
```
list1 := []int{8, 2, 10, 4}
fp.SomeInt(8, list1) // returns true
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

### Example8 - Drop: Drop num or string from list
```
newList := fp.DropInt(1, []int{1, 2, 3, 1}) // returns [2, 3]
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

### Test Coverage
```
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
