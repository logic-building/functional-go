package list_op

import (
	"testing"
)

func TestMapInt(t *testing.T) {
	// Test : square the list
	expectedSquareList := []int{1, 4, 9}
	squareList := MapInt(squareInt, []int{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapInt failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int{6, 7, 8}
	partialAddInt64 := func(num int) int {
		return addInt(5, num)
	}
	sumList := MapInt(partialAddInt64, []int{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapInt failed")
	}
}

func squareInt(num int) int {
	return num * num
}

func addInt(num1, num2 int) int {
	return num1 + num2
}

func TestMapInt64(t *testing.T) {
	// Test : square the list
	expectedSquareList := []int64{1, 4, 9}
	squareList := MapInt64(squareInt64, []int64{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapInt64 failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int64{6, 7, 8}
	partialAddInt64 := func(num int64) int64 {
		return addInt64(5, num)
	}
	sumList := MapInt64(partialAddInt64, []int64{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapInt64 failed")
	}
}

func squareInt64(num int64) int64 {
	return num * num
}

func addInt64(num1, num2 int64) int64 {
	return num1 + num2
}

func TestMapInt32(t *testing.T) {
	// Test : square the list
	expectedSquareList := []int32{1, 4, 9}
	squareList := MapInt32(squareInt32, []int32{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapInt32 failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int32{6, 7, 8}
	partialAddInt32 := func(num int32) int32 {
		return addInt32(5, num)
	}
	sumList := MapInt32(partialAddInt32, []int32{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapInt32 failed")
	}
}

func squareInt32(num int32) int32 {
	return num * num
}

func addInt32(num1, num2 int32) int32 {
	return num1 + num2
}

func TestMapInt16(t *testing.T) {
	// Test : square the list
	expectedSquareList := []int16{1, 4, 9}
	squareList := MapInt16(squareInt16, []int16{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapInt16 failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int16{6, 7, 8}
	partialAddInt16 := func(num int16) int16 {
		return addInt16(5, num)
	}
	sumList := MapInt16(partialAddInt16, []int16{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapInt16 failed")
	}
}

func squareInt16(num int16) int16 {
	return num * num
}

func addInt16(num1, num2 int16) int16 {
	return num1 + num2
}

func TestMapInt8(t *testing.T) {
	// Test : square the list
	expectedSquareList := []int8{1, 4, 9}
	squareList := MapInt8(squareInt8, []int8{1, 2, 3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapInt8 failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []int8{6, 7, 8}
	partialAddInt8 := func(num int8) int8 {
		return addInt8(5, num)
	}
	sumList := MapInt8(partialAddInt8, []int8{1, 2, 3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapInt8 failed")
	}
}

func squareInt8(num int8) int8 {
	return num * num
}

func addInt8(num1, num2 int8) int8 {
	return num1 + num2
}

func TestMapFloat64(t *testing.T) {
	// Test : square the list
	expectedSquareList := []float64{1.2100000000000002, 4.840000000000001, 10.889999999999999}
	squareList := MapFloat64(squareFloat64, []float64{1.1, 2.2, 3.3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapFloat64 failed")
	}

	// Test: add 5 to each item in the list
	expectedSumList := []float64{6.1, 7.2, 8.3}
	partialAddFloat64 := func(num float64) float64 {
		return addFloat64(5, num)
	}
	sumList := MapFloat64(partialAddFloat64, []float64{1.1, 2.2, 3.3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapFloat64 failed")
	}
}

func squareFloat64(num float64) float64 {
	return num * num
}

func addFloat64(num1, num2 float64) float64 {
	return num1 + num2
}

func TestMapFloat32(t *testing.T) {
	// Test : square the list
	expectedSquareList := []float32{1.2100000000000002, 4.840000000000001, 10.889999}
	squareList := MapFloat32(squareFloat32, []float32{1.1, 2.2, 3.3})

	if squareList[0] != expectedSquareList[0] || squareList[1] != expectedSquareList[1] || squareList[2] != expectedSquareList[2] {
		t.Errorf("MapFloat32 failed. actual_list=%v, expected_list=%v", squareList, expectedSquareList)
	}

	// Test: add 5 to each item in the list
	expectedSumList := []float32{6.1, 7.2, 8.3}
	partialAddFloat32 := func(num float32) float32 {
		return addFloat32(5, num)
	}
	sumList := MapFloat32(partialAddFloat32, []float32{1.1, 2.2, 3.3})
	if sumList[0] != expectedSumList[0] || sumList[1] != expectedSumList[1] || sumList[2] != expectedSumList[2] {
		t.Errorf("MapFloat32 failed")
	}
}

func squareFloat32(num float32) float32 {
	return num * num
}

func addFloat32(num1, num2 float32) float32 {
	return num1 + num2
}
