# functional-go
## Simple but functional
### Install
```
go get github.com/logic-building/functional-go/list-op/
```

### dep Gopkg.toml entry
```
[[constraint]]
name = "github.com/logic-building/functional-go"
version = "2.0"
```

### Contains functions
```
Takes list as argument and returns Distinct list
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

Takes two function as argument and apply it on each item in the list and return filtered list
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

Takes function as argument and apply it on each item in the list and return true/false
EveryInt
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

Takes number and list as argument and returns true/false if it exists
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

Takes number or string and list as arguments and remove the given item and returns the list
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
RemoeFloats32
RemoveStr
RemoveStrIgnoreCase
RemoveStrs
RemoveStrsIgnoreCase

Set operations: Add, Remove, Clear, GetList, NewSetInt, Join, Intersection, Minus, Subset, Superset
SetInt
SetIntSync
SetInt64
SetInt64Sync
```

### Example1 - Map : return the list of the square of each items in the list
```
squareList := MapInt(squareInt, []int{1, 2, 3}) // see the map_test.go for detail

func squareInt(num int) int {
	return num * num
}


output
-------
[1, 4, 9]

```

### Example2 - Filter: filter all the even numbers in the list
```
filteredList := FilterInt(isEven, []int{1, 2, 3, 4})

func isEven(num int) bool {
	return num%2 == 0
}

output:
[2, 4]

```

### Example3 - FilterMap: Multiply all positive numbers in the list by 2
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
list1 := []int{8, 2, 10, 4}
EveryInt(isEven, list1) // Returns true
```

### Example5 - Some: Test if number presents in the list
```
list1 := []int{8, 2, 10, 4}
SomeInt(8, list1) // returns true
```

### Example6 - Max: Get max number in the list
```
list := []int{8, 2, 10, 4}
max := MaxInt(list) // returns 10
```

### Example7 - Min: Get min number in the list
```
list := []int{8, 2, 10, 4}
min := MinInt(list) // returns 2
```

### Example8 - Remove: Remove num or string from list
```
newList := RemoveInt(1, []int{1, 2, 3, 1}) // returns [2, 3]
```

### Example9 - Distinct: returns distinct list
```
list := []int{8, 2, 8, 0, 2, 0}
distinct := DistinctInt(list) // returns [8, 2, 0]
```

### Example10 - Set : Create set objects and apply union operation
```
	mySet1 := NewSetInt([]int{10, 20, 30, 20})
	mySet2 := NewSetInt([]int{30, 40, 50})
	mySet3 := mySet1.Union(mySet2) // Returns  [10, 20, 30, 40, 50]
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
pkg: functional-go/list-op
BenchmarkFilterInt64-8                   	 1000000	      5408 ns/op
BenchmarkFilterMapInt64-8                	 1000000	      6271 ns/op
BenchmarkMapInt64_PassedMethod_1_Arg-8   	 1000000	      3767 ns/op
BenchmarkMapInt64_PassedMethod_2_Arg-8   	 1000000	      3819 ns/op
BenchmarkMapStr-8                        	 1000000	     52498 ns/op
PASS
ok  	functional-go/list-op	71.776s
```
