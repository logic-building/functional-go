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
version = "1.1"
```

### Contains functions
```
Return list ---
MapInt
MapInt64
MapInt32
MapInt16
MapInt8
MapFloat64
MapFloat32
MapStr

Return list ---
FilterInt
FilterInt64
FilterInt32
FilterInt16
FilterInt8
FilterFloat64
FilterFloat32
FilterStr

Return list ---
FilterMapInt
FilterMapInt64
FilterMapInt32
FilterMapInt16
FilterMapInt8
FilterMapFloat64
FilterMapFloat32
FilterMapStr

Return true ---
EveryInt
EveryInt64
EveryInt32
EveryInt16
EveryInt8
EveryFloat64
EveryFloat32
EveryStr
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

### Exampl4 - Every: Test if every number in the list is even
```
list1 := []int{8, 2, 10, 4}
EveryInt(isEven, list1) // Returns true
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
BenchmarkFilterInt64-8                   	 1000000	      5284 ns/op
BenchmarkFilterMapInt64-8                	 1000000	      6007 ns/op
BenchmarkMapInt64_PassedMethod_1_Arg-8   	 1000000	      3735 ns/op
BenchmarkMapInt64_PassedMethod_2_Arg-8   	 1000000	      3748 ns/op
BenchmarkMapStr-8                        	 1000000	     51698 ns/op
PASS
ok  	functional-go/list-op	70.486s
```
