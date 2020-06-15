package fp

import (
	"errors"
	"testing"
)

func TestPmapIntErr(t *testing.T) {
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

	expectedSquareList := []int{v1, v4, v9}
	squareList, _ := PMapIntErr(squareIntErr, []int{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMapIntErr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int{v6, v7, v8}
	partialAddIntErr := func(num int) (int, error) {
		if num == 0 {
			return 0, errors.New("0 is invalid for this test")
		}
		r := 5 + num
		return r, nil
	}
	sumList, _ := PMapIntErr(partialAddIntErr, []int{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMapIntErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapIntErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapIntErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapIntErr(nil, []int{})
	if len(r) > 0 {
		t.Errorf("PMapIntErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapIntErr(squareIntErr, []int{v1, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMapIntErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareIntErr(num int) (int, error) {
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}


func TestPmapInt64Err(t *testing.T) {
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

	expectedSquareList := []int64{v1, v4, v9}
	squareList, _ := PMapInt64Err(squareInt64Err, []int64{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMapInt64Err failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int64{v6, v7, v8}
	partialAddInt64Err := func(num int64) (int64, error) {
		if num == 0 {
			return 0, errors.New("0 is invalid for this test")
		}
		r := 5 + num
		return r, nil
	}
	sumList, _ := PMapInt64Err(partialAddInt64Err, []int64{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMapInt64Err failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapInt64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt64Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapInt64Err(nil, []int64{})
	if len(r) > 0 {
		t.Errorf("PMapInt64Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapInt64Err(squareInt64Err, []int64{v1, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMapInt64Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareInt64Err(num int64) (int64, error) {
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}


func TestPmapInt32Err(t *testing.T) {
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

	expectedSquareList := []int32{v1, v4, v9}
	squareList, _ := PMapInt32Err(squareInt32Err, []int32{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMapInt32Err failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int32{v6, v7, v8}
	partialAddInt32Err := func(num int32) (int32, error) {
		if num == 0 {
			return 0, errors.New("0 is invalid for this test")
		}
		r := 5 + num
		return r, nil
	}
	sumList, _ := PMapInt32Err(partialAddInt32Err, []int32{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMapInt32Err failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapInt32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt32Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapInt32Err(nil, []int32{})
	if len(r) > 0 {
		t.Errorf("PMapInt32Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapInt32Err(squareInt32Err, []int32{v1, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMapInt32Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareInt32Err(num int32) (int32, error) {
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}


func TestPmapInt16Err(t *testing.T) {
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

	expectedSquareList := []int16{v1, v4, v9}
	squareList, _ := PMapInt16Err(squareInt16Err, []int16{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMapInt16Err failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int16{v6, v7, v8}
	partialAddInt16Err := func(num int16) (int16, error) {
		if num == 0 {
			return 0, errors.New("0 is invalid for this test")
		}
		r := 5 + num
		return r, nil
	}
	sumList, _ := PMapInt16Err(partialAddInt16Err, []int16{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMapInt16Err failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapInt16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt16Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapInt16Err(nil, []int16{})
	if len(r) > 0 {
		t.Errorf("PMapInt16Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapInt16Err(squareInt16Err, []int16{v1, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMapInt16Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareInt16Err(num int16) (int16, error) {
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}


func TestPmapInt8Err(t *testing.T) {
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

	expectedSquareList := []int8{v1, v4, v9}
	squareList, _ := PMapInt8Err(squareInt8Err, []int8{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMapInt8Err failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int8{v6, v7, v8}
	partialAddInt8Err := func(num int8) (int8, error) {
		if num == 0 {
			return 0, errors.New("0 is invalid for this test")
		}
		r := 5 + num
		return r, nil
	}
	sumList, _ := PMapInt8Err(partialAddInt8Err, []int8{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMapInt8Err failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapInt8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapInt8Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapInt8Err(nil, []int8{})
	if len(r) > 0 {
		t.Errorf("PMapInt8Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapInt8Err(squareInt8Err, []int8{v1, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMapInt8Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareInt8Err(num int8) (int8, error) {
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}


func TestPmapUintErr(t *testing.T) {
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

	expectedSquareList := []uint{v1, v4, v9}
	squareList, _ := PMapUintErr(squareUintErr, []uint{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMapUintErr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []uint{v6, v7, v8}
	partialAddUintErr := func(num uint) (uint, error) {
		if num == 0 {
			return 0, errors.New("0 is invalid for this test")
		}
		r := 5 + num
		return r, nil
	}
	sumList, _ := PMapUintErr(partialAddUintErr, []uint{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMapUintErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapUintErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUintErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapUintErr(nil, []uint{})
	if len(r) > 0 {
		t.Errorf("PMapUintErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapUintErr(squareUintErr, []uint{v1, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMapUintErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareUintErr(num uint) (uint, error) {
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}


func TestPmapUint64Err(t *testing.T) {
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

	expectedSquareList := []uint64{v1, v4, v9}
	squareList, _ := PMapUint64Err(squareUint64Err, []uint64{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMapUint64Err failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []uint64{v6, v7, v8}
	partialAddUint64Err := func(num uint64) (uint64, error) {
		if num == 0 {
			return 0, errors.New("0 is invalid for this test")
		}
		r := 5 + num
		return r, nil
	}
	sumList, _ := PMapUint64Err(partialAddUint64Err, []uint64{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMapUint64Err failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapUint64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint64Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapUint64Err(nil, []uint64{})
	if len(r) > 0 {
		t.Errorf("PMapUint64Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapUint64Err(squareUint64Err, []uint64{v1, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMapUint64Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareUint64Err(num uint64) (uint64, error) {
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}


func TestPmapUint32Err(t *testing.T) {
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

	expectedSquareList := []uint32{v1, v4, v9}
	squareList, _ := PMapUint32Err(squareUint32Err, []uint32{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMapUint32Err failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []uint32{v6, v7, v8}
	partialAddUint32Err := func(num uint32) (uint32, error) {
		if num == 0 {
			return 0, errors.New("0 is invalid for this test")
		}
		r := 5 + num
		return r, nil
	}
	sumList, _ := PMapUint32Err(partialAddUint32Err, []uint32{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMapUint32Err failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapUint32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint32Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapUint32Err(nil, []uint32{})
	if len(r) > 0 {
		t.Errorf("PMapUint32Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapUint32Err(squareUint32Err, []uint32{v1, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMapUint32Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareUint32Err(num uint32) (uint32, error) {
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}


func TestPmapUint16Err(t *testing.T) {
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

	expectedSquareList := []uint16{v1, v4, v9}
	squareList, _ := PMapUint16Err(squareUint16Err, []uint16{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMapUint16Err failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []uint16{v6, v7, v8}
	partialAddUint16Err := func(num uint16) (uint16, error) {
		if num == 0 {
			return 0, errors.New("0 is invalid for this test")
		}
		r := 5 + num
		return r, nil
	}
	sumList, _ := PMapUint16Err(partialAddUint16Err, []uint16{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMapUint16Err failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapUint16Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint16Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapUint16Err(nil, []uint16{})
	if len(r) > 0 {
		t.Errorf("PMapUint16Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapUint16Err(squareUint16Err, []uint16{v1, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMapUint16Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareUint16Err(num uint16) (uint16, error) {
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}


func TestPmapUint8Err(t *testing.T) {
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

	expectedSquareList := []uint8{v1, v4, v9}
	squareList, _ := PMapUint8Err(squareUint8Err, []uint8{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMapUint8Err failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []uint8{v6, v7, v8}
	partialAddUint8Err := func(num uint8) (uint8, error) {
		if num == 0 {
			return 0, errors.New("0 is invalid for this test")
		}
		r := 5 + num
		return r, nil
	}
	sumList, _ := PMapUint8Err(partialAddUint8Err, []uint8{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMapUint8Err failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapUint8Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapUint8Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapUint8Err(nil, []uint8{})
	if len(r) > 0 {
		t.Errorf("PMapUint8Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapUint8Err(squareUint8Err, []uint8{v1, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMapUint8Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareUint8Err(num uint8) (uint8, error) {
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}


func TestPmapStrErr(t *testing.T) {
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

	expectedSquareList := []string{v11, v22, v33}
	squareList, _ := PMapStrErr(squareStrErr, []string{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMapStrErr failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []string{v51, v52, v53}
	partialAddStrErr := func(num string) (string, error) {
		if num == "0" {
			return "", errors.New("0 is invalid for this test")
		}
		r := "5" + num
		return r, nil
	}
	sumList, _ := PMapStrErr(partialAddStrErr, []string{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMapStrErr failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapStrErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapStrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapStrErr(nil, []string{})
	if len(r) > 0 {
		t.Errorf("PMapStrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapStrErr(squareStrErr, []string{v1, v0, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMapStrErr failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareStrErr(num string) (string, error) {
	if num == "0" {
		return "0", errors.New("0  is invalid number")
	}
	r := num + num
	return r, nil
}


func TestPMapBoolErr(t *testing.T) {
	var vt bool = true
	var vf bool = false

	expectedSumList := []bool{vf}
	
	newList, _ := PMapBoolErr(inverseBoolErr, []bool{vt})
	if newList[0] != expectedSumList[0]  {
		t.Errorf("MapBoolErr failed")
	}

	r, _ := PMapBoolErr(nil, nil)
	if len(r) > 0 {
		t.Errorf("MapBoolErr failed.")
	}

	_, err := PMapBoolErr(inverseBoolErr, []bool{vf, vf, vf, vt})
	if err == nil {
		t.Errorf("MapBoolErr failed")
	}
}


func TestPmapFloat32Err(t *testing.T) {
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

	expectedSquareList := []float32{v1, v4, v9}
	squareList, _ := PMapFloat32Err(squareFloat32Err, []float32{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMapFloat32Err failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []float32{v6, v7, v8}
	partialAddFloat32Err := func(num float32) (float32, error) {
		if num == 0 {
			return 0, errors.New("0 is invalid for this test")
		}
		r := 5 + num
		return r, nil
	}
	sumList, _ := PMapFloat32Err(partialAddFloat32Err, []float32{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMapFloat32Err failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapFloat32Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat32Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapFloat32Err(nil, []float32{})
	if len(r) > 0 {
		t.Errorf("PMapFloat32Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapFloat32Err(squareFloat32Err, []float32{v1, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMapFloat32Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareFloat32Err(num float32) (float32, error) {
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}


func TestPmapFloat64Err(t *testing.T) {
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

	expectedSquareList := []float64{v1, v4, v9}
	squareList, _ := PMapFloat64Err(squareFloat64Err, []float64{v1, v2, v3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("PMapFloat64Err failed. expected=%v, actual=%v", expectedSquareList, squareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []float64{v6, v7, v8}
	partialAddFloat64Err := func(num float64) (float64, error) {
		if num == 0 {
			return 0, errors.New("0 is invalid for this test")
		}
		r := 5 + num
		return r, nil
	}
	sumList, _ := PMapFloat64Err(partialAddFloat64Err, []float64{v1, v2, v3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("PMapFloat64Err failed.expected=%v, actual=%v", expectedSumList, sumList)
	}
	r, _ := PMapFloat64Err(nil, nil)
	if len(r) > 0 {
		t.Errorf("PMapFloat64Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	r, _ = PMapFloat64Err(nil, []float64{})
	if len(r) > 0 {
		t.Errorf("PMapFloat64Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}

	_, err := PMapFloat64Err(squareFloat64Err, []float64{v1, v0, v0, v0, v3})
	if err == nil {
		t.Errorf("PMapFloat64Err failed.expected=%v, actual=%v", expectedSquareList, sumList)
	}
}

func squareFloat64Err(num float64) (float64, error) {
	if num == 0 {
		return 0, errors.New("0  is invalid number")
	}
	r := num * num
	return r, nil
}

