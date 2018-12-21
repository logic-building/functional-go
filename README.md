# functional-go
## Simple but functional
### Install
```
go get github.com/logic-building/functional-go/list-op/
```

### Contains functions
```
MapInt
MapInt64
MapInt32
MapInt16
MapInt8
MapFloat64
MapFloat32
MapStr
```

### Example 1: return the list of the square of each items in the list
```
squareList := MapInt(squareInt, []int{1, 2, 3}) // see the map_test.go for detail

func squareInt(num int) int {
	return num * num
}

output
-------
[1, 4, 9]
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
| Function                | Iterations | Time Taken Per Operation |
|-------------------------|------------|--------------------------|
| BenchmarkMapInt64       |  1000000   | 3791 ns/op               |
| BenchMarkMapStr         |  1000000   |59735 ns/op               |

