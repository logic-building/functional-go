# functional-go
## Simple but functional
### Install
```
go get github.com/logic-building/functional-go/list-op/
```

### Gopkg.toml entry
```
[[constraint]]
name = "github.com/logic-building/functional-go"
version = "1.1"
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

FilterInt
FilterInt64
FilterInt32
FilterInt16
FilterInt8
FilterFloat64
FilterFloat32
FilterStr
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

### Example2: filter all the even numbers in the list
```
filteredList := FilterInt(isEven, []int{1, 2, 3, 4})

func isEven(num int) bool {
	return num%2 == 0
}

output:
[2, 4]

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
BenchmarkFilterInt64-8                   	 1000000	      5262 ns/op
BenchmarkMapInt64_PassedMethod_1_Arg-8   	 1000000	      3726 ns/op
BenchmarkMapInt64_PassedMethod_2_Arg-8   	 1000000	      3710 ns/op
BenchmarkMapStr-8                        	 1000000	     51892 ns/op
PASS
ok  	functional-go/list-op	64.604s
```
