package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestPmapInt2(t *testing.T) {
	// Test : square the list
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v9 int = 9

	expectedSquareList := []int{v1, v4, v9}
	squareList := PMapInt(squareInt, []int{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapInt failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapInt642(t *testing.T) {
	// Test : square the list
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v9 int64 = 9

	expectedSquareList := []int64{v1, v4, v9}
	squareList := PMapInt64(squareInt64, []int64{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapInt64 failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapInt322(t *testing.T) {
	// Test : square the list
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v9 int32 = 9

	expectedSquareList := []int32{v1, v4, v9}
	squareList := PMapInt32(squareInt32, []int32{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapInt32 failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapInt162(t *testing.T) {
	// Test : square the list
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v9 int16 = 9

	expectedSquareList := []int16{v1, v4, v9}
	squareList := PMapInt16(squareInt16, []int16{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapInt16 failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapInt82(t *testing.T) {
	// Test : square the list
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v9 int8 = 9

	expectedSquareList := []int8{v1, v4, v9}
	squareList := PMapInt8(squareInt8, []int8{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapInt8 failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapUint2(t *testing.T) {
	// Test : square the list
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v9 uint = 9

	expectedSquareList := []uint{v1, v4, v9}
	squareList := PMapUint(squareUint, []uint{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapUint failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapUint642(t *testing.T) {
	// Test : square the list
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v9 uint64 = 9

	expectedSquareList := []uint64{v1, v4, v9}
	squareList := PMapUint64(squareUint64, []uint64{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapUint64 failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapUint322(t *testing.T) {
	// Test : square the list
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v9 uint32 = 9

	expectedSquareList := []uint32{v1, v4, v9}
	squareList := PMapUint32(squareUint32, []uint32{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapUint32 failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapUint162(t *testing.T) {
	// Test : square the list
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v9 uint16 = 9

	expectedSquareList := []uint16{v1, v4, v9}
	squareList := PMapUint16(squareUint16, []uint16{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapUint16 failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapUint82(t *testing.T) {
	// Test : square the list
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v9 uint8 = 9

	expectedSquareList := []uint8{v1, v4, v9}
	squareList := PMapUint8(squareUint8, []uint8{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapUint8 failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapStr2(t *testing.T) {
	// Test : square the list
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v11 string = "33"
	var v4 string = "22"
	var v9 string = "33"

	appendStr := func(v string) string {
		r := v + v
		return r
	}

	expectedSquareList := []string{v11, v4, v9}
	squareList := PMapStr(appendStr, []string{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapStr failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapBool2(t *testing.T) {
	// Test : square the list
	var vt bool = true
	var vf bool = false

	inverse := func(v bool) bool {
		if v == true {
			return false
		}
		return true
	}

	expectedList := []bool{vt, vt, vt}
	actualList := PMapBool(inverse, []bool{vf, vf, vf}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedList {
		for _, x := range actualList {
			if v == x {
				count++
				break
			}
		}
	}

	if count != len(expectedList) {
		t.Errorf("PMapBool failed.expected len=%v, actual len=%v", len(expectedList), count)
		t.Errorf(reflect.String.String())
	}
}

func TestPmapFloat322(t *testing.T) {
	// Test : square the list
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v9 float32 = 9

	expectedSquareList := []float32{v1, v4, v9}
	squareList := PMapFloat32(squareFloat32, []float32{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapFloat32 failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}

func TestPmapFloat642(t *testing.T) {
	// Test : square the list
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v9 float64 = 9

	expectedSquareList := []float64{v1, v4, v9}
	squareList := PMapFloat64(squareFloat64, []float64{v1, v2, v3}, Optional{FixedPool: 2, RandomOrder: true})

	count := 0
	for _, v := range expectedSquareList {
		for _, x := range squareList {
			if v == x {
				count++
			}
		}
	}

	if count != len(expectedSquareList) {
		t.Errorf("PMapFloat64 failed.expected len=%v, actual len=%v", len(expectedSquareList), count)
	}
}
