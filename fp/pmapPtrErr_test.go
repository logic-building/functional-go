package fp

import (
	"errors"
	"testing"
	"time"
)

func TestPmapIntPtrErr(t *testing.T) {
	// Test : square the list
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v6 int = 6
	var v7 int = 7
	var v8 int = 8
	var v9 int = 9
	var v0 int

	expectedSquareList := []*int{&v1, &v4, &v9}
	squareList, _ := PMapIntPtrErr(squareIntPtrErr, []*int{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapIntPtr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*int{&v6, &v7, &v8}
	partialAddIntPtrErr := func(num *int) (*int, error) {
		if *num == 0 {
			return nil, errors.New("0 is invalid for this test")
		}
		r := 5 + *num
		return &r, nil
	}
	sumList, _ := PMapIntPtrErr(partialAddIntPtrErr, []*int{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapIntPtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapIntPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapIntPtrErr(nil, []*int{})
	if len(r) > 0 {
		t.Errorf("PMapIntPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapIntPtrErr(squareIntPtrErr, []*int{&v1, &v0, &v0, &v0, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapIntPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapIntPtrErr(squareIntPtrErr, []*int{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapIntPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapIntPtrErr(squareIntPtrErr, []*int{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapIntPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareIntPtrErr(num *int) (*int, error) {
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapIntPtr2Err(t *testing.T) {
	// Test : square the list
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v9 int = 9
	var v0 int

	expectedSquareList := []*int{&v1, &v4, &v9}
	squareList, _ := PMapIntPtrErr(squareIntPtrErr, []*int{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapIntPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapIntPtrErr(squareIntPtrErr, []*int{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapIntPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapIntPtrErr(squareIntPtrErr, []*int{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMapIntPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapIntPtrErr(square2IntPtrErr, []*int{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapIntPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapIntPtrErr(square2IntPtrErr, []*int{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMapIntPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapIntPtrErr(square2IntPtrErr, []*int{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMapIntPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func square2IntPtrErr(num *int) (*int, error) {
	time.Sleep(50 * time.Millisecond)
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapInt64PtrErr(t *testing.T) {
	// Test : square the list
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v6 int64 = 6
	var v7 int64 = 7
	var v8 int64 = 8
	var v9 int64 = 9
	var v0 int64

	expectedSquareList := []*int64{&v1, &v4, &v9}
	squareList, _ := PMapInt64PtrErr(squareInt64PtrErr, []*int64{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapInt64Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*int64{&v6, &v7, &v8}
	partialAddInt64PtrErr := func(num *int64) (*int64, error) {
		if *num == 0 {
			return nil, errors.New("0 is invalid for this test")
		}
		r := 5 + *num
		return &r, nil
	}
	sumList, _ := PMapInt64PtrErr(partialAddInt64PtrErr, []*int64{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapInt64PtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapInt64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapInt64PtrErr(nil, []*int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapInt64PtrErr(squareInt64PtrErr, []*int64{&v1, &v0, &v0, &v0, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt64PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapInt64PtrErr(squareInt64PtrErr, []*int64{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapInt64PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapInt64PtrErr(squareInt64PtrErr, []*int64{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapInt64PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareInt64PtrErr(num *int64) (*int64, error) {
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapInt64Ptr2Err(t *testing.T) {
	// Test : square the list
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v9 int64 = 9
	var v0 int64

	expectedSquareList := []*int64{&v1, &v4, &v9}
	squareList, _ := PMapInt64PtrErr(squareInt64PtrErr, []*int64{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapInt64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapInt64PtrErr(squareInt64PtrErr, []*int64{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapInt64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt64PtrErr(squareInt64PtrErr, []*int64{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMapInt64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt64PtrErr(square2Int64PtrErr, []*int64{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapInt64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt64PtrErr(square2Int64PtrErr, []*int64{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMapInt64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt64PtrErr(square2Int64PtrErr, []*int64{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMapInt64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func square2Int64PtrErr(num *int64) (*int64, error) {
	time.Sleep(50 * time.Millisecond)
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapInt32PtrErr(t *testing.T) {
	// Test : square the list
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v6 int32 = 6
	var v7 int32 = 7
	var v8 int32 = 8
	var v9 int32 = 9
	var v0 int32

	expectedSquareList := []*int32{&v1, &v4, &v9}
	squareList, _ := PMapInt32PtrErr(squareInt32PtrErr, []*int32{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapInt32Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*int32{&v6, &v7, &v8}
	partialAddInt32PtrErr := func(num *int32) (*int32, error) {
		if *num == 0 {
			return nil, errors.New("0 is invalid for this test")
		}
		r := 5 + *num
		return &r, nil
	}
	sumList, _ := PMapInt32PtrErr(partialAddInt32PtrErr, []*int32{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapInt32PtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapInt32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapInt32PtrErr(nil, []*int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapInt32PtrErr(squareInt32PtrErr, []*int32{&v1, &v0, &v0, &v0, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt32PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapInt32PtrErr(squareInt32PtrErr, []*int32{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapInt32PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapInt32PtrErr(squareInt32PtrErr, []*int32{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapInt32PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareInt32PtrErr(num *int32) (*int32, error) {
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapInt32Ptr2Err(t *testing.T) {
	// Test : square the list
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v9 int32 = 9
	var v0 int32

	expectedSquareList := []*int32{&v1, &v4, &v9}
	squareList, _ := PMapInt32PtrErr(squareInt32PtrErr, []*int32{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapInt32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapInt32PtrErr(squareInt32PtrErr, []*int32{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapInt32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt32PtrErr(squareInt32PtrErr, []*int32{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMapInt32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt32PtrErr(square2Int32PtrErr, []*int32{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapInt32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt32PtrErr(square2Int32PtrErr, []*int32{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMapInt32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt32PtrErr(square2Int32PtrErr, []*int32{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMapInt32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func square2Int32PtrErr(num *int32) (*int32, error) {
	time.Sleep(50 * time.Millisecond)
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapInt16PtrErr(t *testing.T) {
	// Test : square the list
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v6 int16 = 6
	var v7 int16 = 7
	var v8 int16 = 8
	var v9 int16 = 9
	var v0 int16

	expectedSquareList := []*int16{&v1, &v4, &v9}
	squareList, _ := PMapInt16PtrErr(squareInt16PtrErr, []*int16{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapInt16Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*int16{&v6, &v7, &v8}
	partialAddInt16PtrErr := func(num *int16) (*int16, error) {
		if *num == 0 {
			return nil, errors.New("0 is invalid for this test")
		}
		r := 5 + *num
		return &r, nil
	}
	sumList, _ := PMapInt16PtrErr(partialAddInt16PtrErr, []*int16{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapInt16PtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapInt16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapInt16PtrErr(nil, []*int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapInt16PtrErr(squareInt16PtrErr, []*int16{&v1, &v0, &v0, &v0, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt16PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapInt16PtrErr(squareInt16PtrErr, []*int16{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapInt16PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapInt16PtrErr(squareInt16PtrErr, []*int16{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapInt16PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareInt16PtrErr(num *int16) (*int16, error) {
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapInt16Ptr2Err(t *testing.T) {
	// Test : square the list
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v9 int16 = 9
	var v0 int16

	expectedSquareList := []*int16{&v1, &v4, &v9}
	squareList, _ := PMapInt16PtrErr(squareInt16PtrErr, []*int16{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapInt16Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapInt16PtrErr(squareInt16PtrErr, []*int16{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapInt16Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt16PtrErr(squareInt16PtrErr, []*int16{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMapInt16Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt16PtrErr(square2Int16PtrErr, []*int16{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapInt16Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt16PtrErr(square2Int16PtrErr, []*int16{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMapInt16Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt16PtrErr(square2Int16PtrErr, []*int16{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMapInt16Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func square2Int16PtrErr(num *int16) (*int16, error) {
	time.Sleep(50 * time.Millisecond)
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapInt8PtrErr(t *testing.T) {
	// Test : square the list
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v6 int8 = 6
	var v7 int8 = 7
	var v8 int8 = 8
	var v9 int8 = 9
	var v0 int8

	expectedSquareList := []*int8{&v1, &v4, &v9}
	squareList, _ := PMapInt8PtrErr(squareInt8PtrErr, []*int8{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapInt8Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*int8{&v6, &v7, &v8}
	partialAddInt8PtrErr := func(num *int8) (*int8, error) {
		if *num == 0 {
			return nil, errors.New("0 is invalid for this test")
		}
		r := 5 + *num
		return &r, nil
	}
	sumList, _ := PMapInt8PtrErr(partialAddInt8PtrErr, []*int8{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapInt8PtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapInt8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapInt8PtrErr(nil, []*int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapInt8PtrErr(squareInt8PtrErr, []*int8{&v1, &v0, &v0, &v0, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapInt8PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapInt8PtrErr(squareInt8PtrErr, []*int8{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapInt8PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapInt8PtrErr(squareInt8PtrErr, []*int8{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapInt8PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareInt8PtrErr(num *int8) (*int8, error) {
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapInt8Ptr2Err(t *testing.T) {
	// Test : square the list
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v9 int8 = 9
	var v0 int8

	expectedSquareList := []*int8{&v1, &v4, &v9}
	squareList, _ := PMapInt8PtrErr(squareInt8PtrErr, []*int8{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapInt8Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapInt8PtrErr(squareInt8PtrErr, []*int8{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapInt8Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt8PtrErr(squareInt8PtrErr, []*int8{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMapInt8Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt8PtrErr(square2Int8PtrErr, []*int8{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapInt8Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt8PtrErr(square2Int8PtrErr, []*int8{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMapInt8Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapInt8PtrErr(square2Int8PtrErr, []*int8{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMapInt8Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func square2Int8PtrErr(num *int8) (*int8, error) {
	time.Sleep(50 * time.Millisecond)
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapUintPtrErr(t *testing.T) {
	// Test : square the list
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v6 uint = 6
	var v7 uint = 7
	var v8 uint = 8
	var v9 uint = 9
	var v0 uint

	expectedSquareList := []*uint{&v1, &v4, &v9}
	squareList, _ := PMapUintPtrErr(squareUintPtrErr, []*uint{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapUintPtr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*uint{&v6, &v7, &v8}
	partialAddUintPtrErr := func(num *uint) (*uint, error) {
		if *num == 0 {
			return nil, errors.New("0 is invalid for this test")
		}
		r := 5 + *num
		return &r, nil
	}
	sumList, _ := PMapUintPtrErr(partialAddUintPtrErr, []*uint{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapUintPtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapUintPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapUintPtrErr(nil, []*uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapUintPtrErr(squareUintPtrErr, []*uint{&v1, &v0, &v0, &v0, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUintPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapUintPtrErr(squareUintPtrErr, []*uint{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapUintPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapUintPtrErr(squareUintPtrErr, []*uint{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapUintPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareUintPtrErr(num *uint) (*uint, error) {
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapUintPtr2Err(t *testing.T) {
	// Test : square the list
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v9 uint = 9
	var v0 uint

	expectedSquareList := []*uint{&v1, &v4, &v9}
	squareList, _ := PMapUintPtrErr(squareUintPtrErr, []*uint{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapUintPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapUintPtrErr(squareUintPtrErr, []*uint{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapUintPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUintPtrErr(squareUintPtrErr, []*uint{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMapUintPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUintPtrErr(square2UintPtrErr, []*uint{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapUintPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUintPtrErr(square2UintPtrErr, []*uint{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMapUintPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUintPtrErr(square2UintPtrErr, []*uint{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMapUintPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func square2UintPtrErr(num *uint) (*uint, error) {
	time.Sleep(50 * time.Millisecond)
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapUint64PtrErr(t *testing.T) {
	// Test : square the list
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v6 uint64 = 6
	var v7 uint64 = 7
	var v8 uint64 = 8
	var v9 uint64 = 9
	var v0 uint64

	expectedSquareList := []*uint64{&v1, &v4, &v9}
	squareList, _ := PMapUint64PtrErr(squareUint64PtrErr, []*uint64{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapUint64Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*uint64{&v6, &v7, &v8}
	partialAddUint64PtrErr := func(num *uint64) (*uint64, error) {
		if *num == 0 {
			return nil, errors.New("0 is invalid for this test")
		}
		r := 5 + *num
		return &r, nil
	}
	sumList, _ := PMapUint64PtrErr(partialAddUint64PtrErr, []*uint64{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapUint64PtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapUint64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapUint64PtrErr(nil, []*uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapUint64PtrErr(squareUint64PtrErr, []*uint64{&v1, &v0, &v0, &v0, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint64PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapUint64PtrErr(squareUint64PtrErr, []*uint64{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapUint64PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapUint64PtrErr(squareUint64PtrErr, []*uint64{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapUint64PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareUint64PtrErr(num *uint64) (*uint64, error) {
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapUint64Ptr2Err(t *testing.T) {
	// Test : square the list
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v9 uint64 = 9
	var v0 uint64

	expectedSquareList := []*uint64{&v1, &v4, &v9}
	squareList, _ := PMapUint64PtrErr(squareUint64PtrErr, []*uint64{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapUint64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapUint64PtrErr(squareUint64PtrErr, []*uint64{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapUint64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint64PtrErr(squareUint64PtrErr, []*uint64{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMapUint64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint64PtrErr(square2Uint64PtrErr, []*uint64{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapUint64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint64PtrErr(square2Uint64PtrErr, []*uint64{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMapUint64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint64PtrErr(square2Uint64PtrErr, []*uint64{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMapUint64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func square2Uint64PtrErr(num *uint64) (*uint64, error) {
	time.Sleep(50 * time.Millisecond)
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapUint32PtrErr(t *testing.T) {
	// Test : square the list
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v6 uint32 = 6
	var v7 uint32 = 7
	var v8 uint32 = 8
	var v9 uint32 = 9
	var v0 uint32

	expectedSquareList := []*uint32{&v1, &v4, &v9}
	squareList, _ := PMapUint32PtrErr(squareUint32PtrErr, []*uint32{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapUint32Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*uint32{&v6, &v7, &v8}
	partialAddUint32PtrErr := func(num *uint32) (*uint32, error) {
		if *num == 0 {
			return nil, errors.New("0 is invalid for this test")
		}
		r := 5 + *num
		return &r, nil
	}
	sumList, _ := PMapUint32PtrErr(partialAddUint32PtrErr, []*uint32{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapUint32PtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapUint32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapUint32PtrErr(nil, []*uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapUint32PtrErr(squareUint32PtrErr, []*uint32{&v1, &v0, &v0, &v0, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint32PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapUint32PtrErr(squareUint32PtrErr, []*uint32{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapUint32PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapUint32PtrErr(squareUint32PtrErr, []*uint32{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapUint32PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareUint32PtrErr(num *uint32) (*uint32, error) {
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapUint32Ptr2Err(t *testing.T) {
	// Test : square the list
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v9 uint32 = 9
	var v0 uint32

	expectedSquareList := []*uint32{&v1, &v4, &v9}
	squareList, _ := PMapUint32PtrErr(squareUint32PtrErr, []*uint32{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapUint32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapUint32PtrErr(squareUint32PtrErr, []*uint32{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapUint32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint32PtrErr(squareUint32PtrErr, []*uint32{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMapUint32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint32PtrErr(square2Uint32PtrErr, []*uint32{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapUint32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint32PtrErr(square2Uint32PtrErr, []*uint32{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMapUint32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint32PtrErr(square2Uint32PtrErr, []*uint32{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMapUint32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func square2Uint32PtrErr(num *uint32) (*uint32, error) {
	time.Sleep(50 * time.Millisecond)
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapUint16PtrErr(t *testing.T) {
	// Test : square the list
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v6 uint16 = 6
	var v7 uint16 = 7
	var v8 uint16 = 8
	var v9 uint16 = 9
	var v0 uint16

	expectedSquareList := []*uint16{&v1, &v4, &v9}
	squareList, _ := PMapUint16PtrErr(squareUint16PtrErr, []*uint16{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapUint16Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*uint16{&v6, &v7, &v8}
	partialAddUint16PtrErr := func(num *uint16) (*uint16, error) {
		if *num == 0 {
			return nil, errors.New("0 is invalid for this test")
		}
		r := 5 + *num
		return &r, nil
	}
	sumList, _ := PMapUint16PtrErr(partialAddUint16PtrErr, []*uint16{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapUint16PtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapUint16PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapUint16PtrErr(nil, []*uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapUint16PtrErr(squareUint16PtrErr, []*uint16{&v1, &v0, &v0, &v0, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint16PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapUint16PtrErr(squareUint16PtrErr, []*uint16{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapUint16PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapUint16PtrErr(squareUint16PtrErr, []*uint16{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapUint16PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareUint16PtrErr(num *uint16) (*uint16, error) {
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapUint16Ptr2Err(t *testing.T) {
	// Test : square the list
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v9 uint16 = 9
	var v0 uint16

	expectedSquareList := []*uint16{&v1, &v4, &v9}
	squareList, _ := PMapUint16PtrErr(squareUint16PtrErr, []*uint16{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapUint16Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapUint16PtrErr(squareUint16PtrErr, []*uint16{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapUint16Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint16PtrErr(squareUint16PtrErr, []*uint16{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMapUint16Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint16PtrErr(square2Uint16PtrErr, []*uint16{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapUint16Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint16PtrErr(square2Uint16PtrErr, []*uint16{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMapUint16Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint16PtrErr(square2Uint16PtrErr, []*uint16{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMapUint16Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func square2Uint16PtrErr(num *uint16) (*uint16, error) {
	time.Sleep(50 * time.Millisecond)
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapUint8PtrErr(t *testing.T) {
	// Test : square the list
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v6 uint8 = 6
	var v7 uint8 = 7
	var v8 uint8 = 8
	var v9 uint8 = 9
	var v0 uint8

	expectedSquareList := []*uint8{&v1, &v4, &v9}
	squareList, _ := PMapUint8PtrErr(squareUint8PtrErr, []*uint8{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapUint8Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*uint8{&v6, &v7, &v8}
	partialAddUint8PtrErr := func(num *uint8) (*uint8, error) {
		if *num == 0 {
			return nil, errors.New("0 is invalid for this test")
		}
		r := 5 + *num
		return &r, nil
	}
	sumList, _ := PMapUint8PtrErr(partialAddUint8PtrErr, []*uint8{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapUint8PtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapUint8PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapUint8PtrErr(nil, []*uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapUint8PtrErr(squareUint8PtrErr, []*uint8{&v1, &v0, &v0, &v0, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapUint8PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapUint8PtrErr(squareUint8PtrErr, []*uint8{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapUint8PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapUint8PtrErr(squareUint8PtrErr, []*uint8{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapUint8PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareUint8PtrErr(num *uint8) (*uint8, error) {
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapUint8Ptr2Err(t *testing.T) {
	// Test : square the list
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v9 uint8 = 9
	var v0 uint8

	expectedSquareList := []*uint8{&v1, &v4, &v9}
	squareList, _ := PMapUint8PtrErr(squareUint8PtrErr, []*uint8{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapUint8Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapUint8PtrErr(squareUint8PtrErr, []*uint8{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapUint8Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint8PtrErr(squareUint8PtrErr, []*uint8{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMapUint8Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint8PtrErr(square2Uint8PtrErr, []*uint8{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapUint8Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint8PtrErr(square2Uint8PtrErr, []*uint8{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMapUint8Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapUint8PtrErr(square2Uint8PtrErr, []*uint8{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMapUint8Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func square2Uint8PtrErr(num *uint8) (*uint8, error) {
	time.Sleep(50 * time.Millisecond)
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapStrPtrErr(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v0 string = "0"

	var v51 string = "51"
	var v52 string = "52"
	var v53 string = "53"

	var v11 string = "11"
	var v22 string = "22"
	var v33 string = "33"

	expectedSquareList := []*string{&v11, &v22, &v33}
	squareList, _ := PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapStrPtr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*string{&v51, &v52, &v53}
	partialAddStrPtrErr := func(num *string) (*string, error) {
		if *num == "0" {
			return nil, errors.New("0 is invalid for this test")
		}
		r := "5" + *num
		return &r, nil
	}
	sumList, _ := PMapStrPtrErr(partialAddStrPtrErr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapStrPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapStrPtrErr(nil, []*string{})
	if len(r) > 0 {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v0, &v0, &v0, &v0, &v3}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapStrPtrErr(squareStrPtrErr, []*string{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapStrPtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareStrPtrErr(num *string) (*string, error) {
	if *num == "0" {
		return nil, errors.New("0  is invalid number")
	}
	r := *num + *num
	return &r, nil
}

func TestPmapStrPtr2Err(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v11 string = "33"
	var v4 string = "22"
	var v9 string = "33"
	var v0 string = "0"

	appendStr := func(v *string) (*string, error) {
		r := *v + *v
		if *v == "0" {
			return nil, errors.New("0 is not allowed in this test")
		}
		return &r, nil
	}

	appendStr2 := func(v *string) (*string, error) {
		time.Sleep(200 * time.Millisecond)
		r := *v + *v
		if *v == "0" {
			return nil, errors.New("0 is not allowed in this test")
		}
		return &r, nil
	}

	expectedSquareList := []*string{&v11, &v4, &v9}
	squareList, _ := PMapStrPtrErr(appendStr, []*string{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapStrPtrErr(squareStrPtrErr, []*string{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMap<FTYPE>Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrPtrErr(squareStrPtrErr, []*string{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrPtrErr(appendStr2, []*string{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrPtrErr(appendStr2, []*string{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapStrPtrErr(appendStr, []*string{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMapStrPtr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPMapBoolPtrErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expectedSumList := []*bool{&vf, &vf}

	newList, _ := PMapBoolPtrErr(inverseBoolPtrErr, []*bool{&vt, &vt}, Optional{FixedPool: 1})
	if *newList[0] != *expectedSumList[0] || *newList[1] != *expectedSumList[1] {
		t.Errorf("MapBoolPtrErr failed")
	}

	r, _ := PMapBoolPtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolPtrErr failed.")
	}

	_, err := PMapBoolPtrErr(inverseBoolPtrErr, []*bool{&vf, &vf, &vf, &vt}, Optional{FixedPool: 2})
	if err == nil {
		t.Errorf("MapBoolPtrErr failed")
	}
}

func TestPmapBool2PtrErr(t *testing.T) {
	// Test : square the list
	var vt bool = true
	var vf bool = false

	inverse := func(v *bool) (*bool, error) {
		if *v == true {
			return &vt, nil
		}

		if *v == false {
			return nil, errors.New("false is error")
		}
		return &vf, nil
	}

	inverse2 := func(v *bool) (*bool, error) {
		time.Sleep(50 * time.Millisecond)
		if *v == true {
			return &vt, nil
		}

		if *v == false {
			return nil, errors.New("false is error")
		}
		return &vf, nil
	}

	expectedList := []*bool{&vt, &vt, &vt}
	actualList, _ := PMapBoolPtrErr(inverse, []*bool{&vt, &vt, &vt}, Optional{FixedPool: 2, RandomOrder: true})

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
		t.Errorf("PMapBoolPtr2Err failed.expected len=%v, actual len=%v", len(expectedList), count)
	}

	_, err := PMapBoolPtrErr(inverse, []*bool{&vf, &vf, &vt}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolPtr2Err failed")
	}

	_, err = PMapBoolPtrErr(inverse, []*bool{&vf, &vf, &vf}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolPtr2Err failed")
	}

	_, err = PMapBoolPtrErr(inverse, []*bool{&vf, &vf, &vf, &vf, &vf, &vt}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolPtr2Err failed")
	}

	_, err = PMapBoolPtrErr(inverse2, []*bool{&vf, &vt}, Optional{FixedPool: 2, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolPtr2Err failed")
	}

	_, err = PMapBoolPtrErr(inverse, []*bool{&vf, &vf, &vf, &vf, &vf}, Optional{FixedPool: 1, RandomOrder: true})
	if err == nil {
		t.Errorf("PMapBoolPtr2Err failed")
	}
}

func TestPmapFloat32PtrErr(t *testing.T) {
	// Test : square the list
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v6 float32 = 6
	var v7 float32 = 7
	var v8 float32 = 8
	var v9 float32 = 9
	var v0 float32

	expectedSquareList := []*float32{&v1, &v4, &v9}
	squareList, _ := PMapFloat32PtrErr(squareFloat32PtrErr, []*float32{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapFloat32Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*float32{&v6, &v7, &v8}
	partialAddFloat32PtrErr := func(num *float32) (*float32, error) {
		if *num == 0 {
			return nil, errors.New("0 is invalid for this test")
		}
		r := 5 + *num
		return &r, nil
	}
	sumList, _ := PMapFloat32PtrErr(partialAddFloat32PtrErr, []*float32{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapFloat32PtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapFloat32PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapFloat32PtrErr(nil, []*float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapFloat32PtrErr(squareFloat32PtrErr, []*float32{&v1, &v0, &v0, &v0, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat32PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapFloat32PtrErr(squareFloat32PtrErr, []*float32{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapFloat32PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapFloat32PtrErr(squareFloat32PtrErr, []*float32{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapFloat32PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareFloat32PtrErr(num *float32) (*float32, error) {
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapFloat32Ptr2Err(t *testing.T) {
	// Test : square the list
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v9 float32 = 9
	var v0 float32

	expectedSquareList := []*float32{&v1, &v4, &v9}
	squareList, _ := PMapFloat32PtrErr(squareFloat32PtrErr, []*float32{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapFloat32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapFloat32PtrErr(squareFloat32PtrErr, []*float32{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapFloat32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapFloat32PtrErr(squareFloat32PtrErr, []*float32{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMapFloat32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapFloat32PtrErr(square2Float32PtrErr, []*float32{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapFloat32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapFloat32PtrErr(square2Float32PtrErr, []*float32{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMapFloat32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapFloat32PtrErr(square2Float32PtrErr, []*float32{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMapFloat32Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func square2Float32PtrErr(num *float32) (*float32, error) {
	time.Sleep(50 * time.Millisecond)
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapFloat64PtrErr(t *testing.T) {
	// Test : square the list
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v6 float64 = 6
	var v7 float64 = 7
	var v8 float64 = 8
	var v9 float64 = 9
	var v0 float64

	expectedSquareList := []*float64{&v1, &v4, &v9}
	squareList, _ := PMapFloat64PtrErr(squareFloat64PtrErr, []*float64{&v1, &v2, &v3}, Optional{FixedPool: 1})

	if *squareList[0] != *expectedSquareList[0] || *squareList[1] != *expectedSquareList[1] || *squareList[2] != *expectedSquareList[2] {
		t.Errorf("PMapFloat64Ptr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []*float64{&v6, &v7, &v8}
	partialAddFloat64PtrErr := func(num *float64) (*float64, error) {
		if *num == 0 {
			return nil, errors.New("0 is invalid for this test")
		}
		r := 5 + *num
		return &r, nil
	}
	sumList, _ := PMapFloat64PtrErr(partialAddFloat64PtrErr, []*float64{&v1, &v2, &v3}, Optional{FixedPool: 1})
	if *sumList[0] != *expectedSumList[0] || *sumList[1] != *expectedSumList[1] || *sumList[2] != *expectedSumList[2] {
		t.Errorf("PMapFloat64PtrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapFloat64PtrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapFloat64PtrErr(nil, []*float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapFloat64PtrErr(squareFloat64PtrErr, []*float64{&v1, &v0, &v0, &v0, &v3}, Optional{FixedPool: 1})
	if err == nil {
		t.Errorf("PMapFloat64PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapFloat64PtrErr(squareFloat64PtrErr, []*float64{&v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapFloat64PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err = PMapFloat64PtrErr(squareFloat64PtrErr, []*float64{&v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v1, &v2, &v3, &v0}, Optional{FixedPool: 3})
	if err == nil {
		t.Errorf("PMapFloat64PtrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareFloat64PtrErr(num *float64) (*float64, error) {
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}

func TestPmapFloat64Ptr2Err(t *testing.T) {
	// Test : square the list
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v9 float64 = 9
	var v0 float64

	expectedSquareList := []*float64{&v1, &v4, &v9}
	squareList, _ := PMapFloat64PtrErr(squareFloat64PtrErr, []*float64{&v1, &v2, &v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if *v == *x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapFloat64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err := PMapFloat64PtrErr(squareFloat64PtrErr, []*float64{&v0, &v0, &v3}, Optional{FixedPool: 2, RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapFloat64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapFloat64PtrErr(squareFloat64PtrErr, []*float64{&v0, &v0, &v0}, Optional{FixedPool: 2})
	
	if err == nil {
		t.Errorf("PMapFloat64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapFloat64PtrErr(square2Float64PtrErr, []*float64{&v0, &v3}, Optional{RandomOrder: true})
	
	if err == nil {
		t.Errorf("PMapFloat64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapFloat64PtrErr(square2Float64PtrErr, []*float64{&v0, &v3})
	
	if err == nil {
		t.Errorf("PMapFloat64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}

	_, err = PMapFloat64PtrErr(square2Float64PtrErr, []*float64{&v0, &v0, &v0, &v0, &v0, &v0}, Optional{FixedPool: 1})
	
	if err == nil {
		t.Errorf("PMapFloat64Ptr2Err failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func square2Float64PtrErr(num *float64) (*float64, error) {
	time.Sleep(50 * time.Millisecond)
	if *num == 0 {
		return nil, errors.New("0  is invalid number")
	}
	r := *num * *num
	return &r, nil
}
