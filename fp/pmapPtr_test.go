package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestPmapIntPtr(t *testing.T) {
	// Test : square the list
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v6 int = 6
	var v7 int = 7
	var v8 int = 8
	var v9 int = 9

	expectedSquareList := []*int{&v1, &v4, &v9}
	squareList := PMapIntPtr(squareIntPtr, []*int{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapIntPtr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*int{&v6, &v7, &v8}
	partialAddIntPtr := func(num *int) *int {
		r := 5 + *num
		return &r
	}
	sumList := PMapIntPtr(partialAddIntPtr, []*int{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapIntPtr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMapIntPtr(nil, nil)) > 0 {
		t.Errorf("PMapIntPtr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMapIntPtr(nil, []*int{})) > 0 {
		t.Errorf("PMapIntPtr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func squareIntPtr(num *int) *int {
	r := *num * *num
	return &r
}

func TestPmapIntPtr2(t *testing.T) {
	// Test : square the list
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v9 int = 9

	expectedSquareList := []*int{&v1, &v4, &v9}
	squareList := PMapIntPtr(squareIntPtr, []*int{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapIntPtr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapInt64Ptr(t *testing.T) {
	// Test : square the list
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v6 int64 = 6
	var v7 int64 = 7
	var v8 int64 = 8
	var v9 int64 = 9

	expectedSquareList := []*int64{&v1, &v4, &v9}
	squareList := PMapInt64Ptr(squareInt64Ptr, []*int64{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapInt64Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*int64{&v6, &v7, &v8}
	partialAddInt64Ptr := func(num *int64) *int64 {
		r := 5 + *num
		return &r
	}
	sumList := PMapInt64Ptr(partialAddInt64Ptr, []*int64{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapInt64Ptr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMapInt64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt64Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMapInt64Ptr(nil, []*int64{})) > 0 {
		t.Errorf("PMapInt64Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func squareInt64Ptr(num *int64) *int64 {
	r := *num * *num
	return &r
}

func TestPmapInt64Ptr2(t *testing.T) {
	// Test : square the list
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v9 int64 = 9

	expectedSquareList := []*int64{&v1, &v4, &v9}
	squareList := PMapInt64Ptr(squareInt64Ptr, []*int64{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapInt64Ptr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapInt32Ptr(t *testing.T) {
	// Test : square the list
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v6 int32 = 6
	var v7 int32 = 7
	var v8 int32 = 8
	var v9 int32 = 9

	expectedSquareList := []*int32{&v1, &v4, &v9}
	squareList := PMapInt32Ptr(squareInt32Ptr, []*int32{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapInt32Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*int32{&v6, &v7, &v8}
	partialAddInt32Ptr := func(num *int32) *int32 {
		r := 5 + *num
		return &r
	}
	sumList := PMapInt32Ptr(partialAddInt32Ptr, []*int32{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapInt32Ptr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMapInt32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt32Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMapInt32Ptr(nil, []*int32{})) > 0 {
		t.Errorf("PMapInt32Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func squareInt32Ptr(num *int32) *int32 {
	r := *num * *num
	return &r
}

func TestPmapInt32Ptr2(t *testing.T) {
	// Test : square the list
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v9 int32 = 9

	expectedSquareList := []*int32{&v1, &v4, &v9}
	squareList := PMapInt32Ptr(squareInt32Ptr, []*int32{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapInt32Ptr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapInt16Ptr(t *testing.T) {
	// Test : square the list
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v6 int16 = 6
	var v7 int16 = 7
	var v8 int16 = 8
	var v9 int16 = 9

	expectedSquareList := []*int16{&v1, &v4, &v9}
	squareList := PMapInt16Ptr(squareInt16Ptr, []*int16{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapInt16Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*int16{&v6, &v7, &v8}
	partialAddInt16Ptr := func(num *int16) *int16 {
		r := 5 + *num
		return &r
	}
	sumList := PMapInt16Ptr(partialAddInt16Ptr, []*int16{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapInt16Ptr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMapInt16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt16Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMapInt16Ptr(nil, []*int16{})) > 0 {
		t.Errorf("PMapInt16Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func squareInt16Ptr(num *int16) *int16 {
	r := *num * *num
	return &r
}

func TestPmapInt16Ptr2(t *testing.T) {
	// Test : square the list
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v9 int16 = 9

	expectedSquareList := []*int16{&v1, &v4, &v9}
	squareList := PMapInt16Ptr(squareInt16Ptr, []*int16{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapInt16Ptr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapInt8Ptr(t *testing.T) {
	// Test : square the list
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v6 int8 = 6
	var v7 int8 = 7
	var v8 int8 = 8
	var v9 int8 = 9

	expectedSquareList := []*int8{&v1, &v4, &v9}
	squareList := PMapInt8Ptr(squareInt8Ptr, []*int8{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapInt8Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*int8{&v6, &v7, &v8}
	partialAddInt8Ptr := func(num *int8) *int8 {
		r := 5 + *num
		return &r
	}
	sumList := PMapInt8Ptr(partialAddInt8Ptr, []*int8{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapInt8Ptr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMapInt8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapInt8Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMapInt8Ptr(nil, []*int8{})) > 0 {
		t.Errorf("PMapInt8Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func squareInt8Ptr(num *int8) *int8 {
	r := *num * *num
	return &r
}

func TestPmapInt8Ptr2(t *testing.T) {
	// Test : square the list
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v9 int8 = 9

	expectedSquareList := []*int8{&v1, &v4, &v9}
	squareList := PMapInt8Ptr(squareInt8Ptr, []*int8{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapInt8Ptr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapUintPtr(t *testing.T) {
	// Test : square the list
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v6 uint = 6
	var v7 uint = 7
	var v8 uint = 8
	var v9 uint = 9

	expectedSquareList := []*uint{&v1, &v4, &v9}
	squareList := PMapUintPtr(squareUintPtr, []*uint{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapUintPtr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*uint{&v6, &v7, &v8}
	partialAddUintPtr := func(num *uint) *uint {
		r := 5 + *num
		return &r
	}
	sumList := PMapUintPtr(partialAddUintPtr, []*uint{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapUintPtr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMapUintPtr(nil, nil)) > 0 {
		t.Errorf("PMapUintPtr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMapUintPtr(nil, []*uint{})) > 0 {
		t.Errorf("PMapUintPtr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func squareUintPtr(num *uint) *uint {
	r := *num * *num
	return &r
}

func TestPmapUintPtr2(t *testing.T) {
	// Test : square the list
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v9 uint = 9

	expectedSquareList := []*uint{&v1, &v4, &v9}
	squareList := PMapUintPtr(squareUintPtr, []*uint{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapUintPtr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapUint64Ptr(t *testing.T) {
	// Test : square the list
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v6 uint64 = 6
	var v7 uint64 = 7
	var v8 uint64 = 8
	var v9 uint64 = 9

	expectedSquareList := []*uint64{&v1, &v4, &v9}
	squareList := PMapUint64Ptr(squareUint64Ptr, []*uint64{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapUint64Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*uint64{&v6, &v7, &v8}
	partialAddUint64Ptr := func(num *uint64) *uint64 {
		r := 5 + *num
		return &r
	}
	sumList := PMapUint64Ptr(partialAddUint64Ptr, []*uint64{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapUint64Ptr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMapUint64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint64Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMapUint64Ptr(nil, []*uint64{})) > 0 {
		t.Errorf("PMapUint64Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func squareUint64Ptr(num *uint64) *uint64 {
	r := *num * *num
	return &r
}

func TestPmapUint64Ptr2(t *testing.T) {
	// Test : square the list
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v9 uint64 = 9

	expectedSquareList := []*uint64{&v1, &v4, &v9}
	squareList := PMapUint64Ptr(squareUint64Ptr, []*uint64{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapUint64Ptr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapUint32Ptr(t *testing.T) {
	// Test : square the list
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v6 uint32 = 6
	var v7 uint32 = 7
	var v8 uint32 = 8
	var v9 uint32 = 9

	expectedSquareList := []*uint32{&v1, &v4, &v9}
	squareList := PMapUint32Ptr(squareUint32Ptr, []*uint32{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapUint32Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*uint32{&v6, &v7, &v8}
	partialAddUint32Ptr := func(num *uint32) *uint32 {
		r := 5 + *num
		return &r
	}
	sumList := PMapUint32Ptr(partialAddUint32Ptr, []*uint32{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapUint32Ptr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMapUint32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint32Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMapUint32Ptr(nil, []*uint32{})) > 0 {
		t.Errorf("PMapUint32Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func squareUint32Ptr(num *uint32) *uint32 {
	r := *num * *num
	return &r
}

func TestPmapUint32Ptr2(t *testing.T) {
	// Test : square the list
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v9 uint32 = 9

	expectedSquareList := []*uint32{&v1, &v4, &v9}
	squareList := PMapUint32Ptr(squareUint32Ptr, []*uint32{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapUint32Ptr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapUint16Ptr(t *testing.T) {
	// Test : square the list
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v6 uint16 = 6
	var v7 uint16 = 7
	var v8 uint16 = 8
	var v9 uint16 = 9

	expectedSquareList := []*uint16{&v1, &v4, &v9}
	squareList := PMapUint16Ptr(squareUint16Ptr, []*uint16{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapUint16Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*uint16{&v6, &v7, &v8}
	partialAddUint16Ptr := func(num *uint16) *uint16 {
		r := 5 + *num
		return &r
	}
	sumList := PMapUint16Ptr(partialAddUint16Ptr, []*uint16{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapUint16Ptr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMapUint16Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint16Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMapUint16Ptr(nil, []*uint16{})) > 0 {
		t.Errorf("PMapUint16Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func squareUint16Ptr(num *uint16) *uint16 {
	r := *num * *num
	return &r
}

func TestPmapUint16Ptr2(t *testing.T) {
	// Test : square the list
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v9 uint16 = 9

	expectedSquareList := []*uint16{&v1, &v4, &v9}
	squareList := PMapUint16Ptr(squareUint16Ptr, []*uint16{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapUint16Ptr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapUint8Ptr(t *testing.T) {
	// Test : square the list
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v6 uint8 = 6
	var v7 uint8 = 7
	var v8 uint8 = 8
	var v9 uint8 = 9

	expectedSquareList := []*uint8{&v1, &v4, &v9}
	squareList := PMapUint8Ptr(squareUint8Ptr, []*uint8{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapUint8Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*uint8{&v6, &v7, &v8}
	partialAddUint8Ptr := func(num *uint8) *uint8 {
		r := 5 + *num
		return &r
	}
	sumList := PMapUint8Ptr(partialAddUint8Ptr, []*uint8{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapUint8Ptr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMapUint8Ptr(nil, nil)) > 0 {
		t.Errorf("PMapUint8Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMapUint8Ptr(nil, []*uint8{})) > 0 {
		t.Errorf("PMapUint8Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func squareUint8Ptr(num *uint8) *uint8 {
	r := *num * *num
	return &r
}

func TestPmapUint8Ptr2(t *testing.T) {
	// Test : square the list
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v9 uint8 = 9

	expectedSquareList := []*uint8{&v1, &v4, &v9}
	squareList := PMapUint8Ptr(squareUint8Ptr, []*uint8{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapUint8Ptr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapStrPtr(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v6 string = "51"
	var v7 string = "52"
	var v8 string = "53"

	var v1SquareExpected string = "11"
	var v2SquareExpected string = "22"
	var v3SquareExpected string = "33"

	expectedSquareList := []*string{&v1SquareExpected, &v2SquareExpected, &v3SquareExpected}
	squareList := PMapStrPtr(squareStrPtr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapStrPtr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*string{&v6, &v7, &v8}
	partialAddStrPtr := func(num *string) *string {
		r := "5" + *num
		return &r
	}
	sumList := PMapStrPtr(partialAddStrPtr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapStrPtr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMapStrPtr(nil, nil)) > 0 {
		t.Errorf("PMapStrPtr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMapStrPtr(nil, []*string{})) > 0 {
		t.Errorf("PMapStrPtr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func squareStrPtr(num *string) *string {
	r := *num + *num
	return &r
}

func TestPmapStrPtr2(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v11 string = "33"
	var v4 string = "22"
	var v9 string = "33"

	appendStr := func(v *string) *string {
		r := *v + *v
		return &r
	}

	expectedSquareList := []*string{&v11, &v4, &v9}
	squareList := PMapStrPtr(appendStr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapStrPtr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPMapBoolPtr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expectedSumList := []*bool{&vf, &vf}

	newList := PMapBoolPtr(inverseBoolPtr, []*bool{&vt, &vt}, Optional{FixedPool: 1})
	if *newList[0] != *expectedSumList[0] || *newList[1] != *expectedSumList[1] {
		t.Errorf("MapBoolPtr failed")
	}

	if len(PMapBoolPtr(nil, nil)) > 0 {
		t.Errorf("MapBoolPtr failed.")
	}
}

func TestPmapBool2Ptr(t *testing.T) {
	// Test : square the list
	var vt bool = true
	var vf bool = false

	inverse := func(v *bool) *bool {
		if *v == true {
			return &vf
		}
		return &vt
	}

	expectedList := []*bool{&vt, &vt, &vt}
	actualList := PMapBoolPtr(inverse, []*bool{&vf, &vf, &vf}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedList {
		for _, x := range actualList {
			if *v == *x {
				count++
				break
			}
		}
	}

	if count != len(expectedList) {
		t.Errorf("PMapBool failed.expected len=%v, actual len=%v", len(expectedList), count)
	}
}

func TestPmapFloat32Ptr(t *testing.T) {
	// Test : square the list
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v6 float32 = 6
	var v7 float32 = 7
	var v8 float32 = 8
	var v9 float32 = 9

	expectedSquareList := []*float32{&v1, &v4, &v9}
	squareList := PMapFloat32Ptr(squareFloat32Ptr, []*float32{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapFloat32Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*float32{&v6, &v7, &v8}
	partialAddFloat32Ptr := func(num *float32) *float32 {
		r := 5 + *num
		return &r
	}
	sumList := PMapFloat32Ptr(partialAddFloat32Ptr, []*float32{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapFloat32Ptr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMapFloat32Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat32Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMapFloat32Ptr(nil, []*float32{})) > 0 {
		t.Errorf("PMapFloat32Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func squareFloat32Ptr(num *float32) *float32 {
	r := *num * *num
	return &r
}

func TestPmapFloat32Ptr2(t *testing.T) {
	// Test : square the list
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v9 float32 = 9

	expectedSquareList := []*float32{&v1, &v4, &v9}
	squareList := PMapFloat32Ptr(squareFloat32Ptr, []*float32{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapFloat32Ptr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapFloat64Ptr(t *testing.T) {
	// Test : square the list
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v6 float64 = 6
	var v7 float64 = 7
	var v8 float64 = 8
	var v9 float64 = 9

	expectedSquareList := []*float64{&v1, &v4, &v9}
	squareList := PMapFloat64Ptr(squareFloat64Ptr, []*float64{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapFloat64Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*float64{&v6, &v7, &v8}
	partialAddFloat64Ptr := func(num *float64) *float64 {
		r := 5 + *num
		return &r
	}
	sumList := PMapFloat64Ptr(partialAddFloat64Ptr, []*float64{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapFloat64Ptr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}

	if len(PMapFloat64Ptr(nil, nil)) > 0 {
		t.Errorf("PMapFloat64Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	if len(PMapFloat64Ptr(nil, []*float64{})) > 0 {
		t.Errorf("PMapFloat64Ptr failed.expected=%v, actual=%v", expectedSquareList, sumList)
		t.Errorf(reflect.String.String())
	}
}

func squareFloat64Ptr(num *float64) *float64 {
	r := *num * *num
	return &r
}

func TestPmapFloat64Ptr2(t *testing.T) {
	// Test : square the list
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v9 float64 = 9

	expectedSquareList := []*float64{&v1, &v4, &v9}
	squareList := PMapFloat64Ptr(squareFloat64Ptr, []*float64{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapFloat64Ptr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}
