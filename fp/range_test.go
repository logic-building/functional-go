package fp

import (
	"reflect"
	"testing"
)

func TestRangeInt(t *testing.T) {
	expectedList := []int{}
	actualList := RangeInt(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int{-2, -1, 0, 1}
	actualList = RangeInt(-2, 2)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int{0, 1}
	actualList = RangeInt(0, 2)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int{}
	actualList = RangeInt(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int{}
	actualList = RangeInt(5, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int{-5, -4, -3, -2}
	actualList = RangeInt(-5, -1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int{-5, -4, -3, -2, -1}
	actualList = RangeInt(-5, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int{-5, -4, -3, -2, -1, 0}
	actualList = RangeInt(-5, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int{}
	actualList = RangeInt(-5, 1, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int{}
	actualList = RangeInt(-5, 1, -1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int{}
	actualList = RangeInt(-5, 1, -1, -2, -3)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int{1, 2}
	actualList = RangeInt(1, 3, 1, -2, -3)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int{}
	actualList = RangeInt(1, 3, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt failed. expected=%v, actual=%v", expectedList, actualList)
	}

	actualList = RangeInt(1, 100, 10)
	expectedList = []int{1, 11, 21, 31, 41, 51, 61, 71, 81, 91}
	expectedLen := 10
	actualLen := len(actualList)
	if actualLen != expectedLen || !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("RangeInt failed. expected-list=%v, length=%d, actual-list=%v, actual-length=%d", expectedList, expectedLen, actualList, actualLen)
	}
}

func TestRangeInt64(t *testing.T) {
	expectedList := []int64{}
	actualList := RangeInt64(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int64{-2, -1, 0, 1}
	actualList = RangeInt64(-2, 2)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int64{0, 1}
	actualList = RangeInt64(0, 2)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int64{}
	actualList = RangeInt64(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int64{}
	actualList = RangeInt64(5, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int64{-5, -4, -3, -2}
	actualList = RangeInt64(-5, -1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int64{-5, -4, -3, -2, -1}
	actualList = RangeInt64(-5, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int64{-5, -4, -3, -2, -1, 0}
	actualList = RangeInt64(-5, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt64 failed. expected=%v, actual=%v", expectedList, actualList)
	}
	expectedList = []int64{}
	actualList = RangeInt64(-5, 1, -1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int64{}
	actualList = RangeInt64(-5, 1, -1, -2, -3)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int64{1, 2}
	actualList = RangeInt64(1, 3, 1, -2, -3)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int64{}
	actualList = RangeInt64(1, 3, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	actualList = RangeInt64(1, 100, 10)
	expectedList = []int64{1, 11, 21, 31, 41, 51, 61, 71, 81, 91}
	expectedLen := 10
	actualLen := len(actualList)
	if actualLen != expectedLen || !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt64 failed. expected-list=%v, length=%d, actual-list=%v, actual-length=%d", expectedList, expectedLen, actualList, actualLen)
	}
}

func TestRangeInt32(t *testing.T) {
	expectedList := []int32{}
	actualList := RangeInt32(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int32{-2, -1, 0, 1}
	actualList = RangeInt32(-2, 2)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int32{0, 1}
	actualList = RangeInt32(0, 2)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int32{}
	actualList = RangeInt32(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int32{}
	actualList = RangeInt32(5, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int32{-5, -4, -3, -2}
	actualList = RangeInt32(-5, -1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int32{-5, -4, -3, -2, -1}
	actualList = RangeInt32(-5, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int32{-5, -4, -3, -2, -1, 0}
	actualList = RangeInt32(-5, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int32{}
	actualList = RangeInt32(-5, 1, -1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int32{}
	actualList = RangeInt32(-5, 1, -1, -2, -3)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int32{1, 2}
	actualList = RangeInt32(1, 3, 1, -2, -3)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int32{}
	actualList = RangeInt32(1, 3, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	actualList = RangeInt32(1, 100, 10)
	expectedList = []int32{1, 11, 21, 31, 41, 51, 61, 71, 81, 91}
	expectedLen := 10
	actualLen := len(actualList)
	if actualLen != expectedLen || !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt32 failed. expected-list=%v, length=%d, actual-list=%v, actual-length=%d", expectedList, expectedLen, actualList, actualLen)
	}

}

func TestRangeInt16(t *testing.T) {
	expectedList := []int16{}
	actualList := RangeInt16(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int16{-2, -1, 0, 1}
	actualList = RangeInt16(-2, 2)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int16{0, 1}
	actualList = RangeInt16(0, 2)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int16{}
	actualList = RangeInt16(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int16{}
	actualList = RangeInt16(5, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int16{-5, -4, -3, -2}
	actualList = RangeInt16(-5, -1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int16{-5, -4, -3, -2, -1}
	actualList = RangeInt16(-5, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int16{-5, -4, -3, -2, -1, 0}
	actualList = RangeInt16(-5, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int16{}
	actualList = RangeInt16(-5, 1, -1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int16{}
	actualList = RangeInt16(-5, 1, -1, -2, -3)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int16{1, 2}
	actualList = RangeInt16(1, 3, 1, -2, -3)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int16{}
	actualList = RangeInt16(1, 3, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	actualList = RangeInt16(1, 100, 10)
	expectedList = []int16{1, 11, 21, 31, 41, 51, 61, 71, 81, 91}
	expectedLen := 10
	actualLen := len(actualList)
	if actualLen != expectedLen || !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt16 failed. expected-list=%v, length=%d, actual-list=%v, actual-length=%d", expectedList, expectedLen, actualList, actualLen)
	}
}

func TestRangeInt8(t *testing.T) {
	expectedList := []int8{}
	actualList := RangeInt8(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int8{-2, -1, 0, 1}
	actualList = RangeInt8(-2, 2)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int8{0, 1}
	actualList = RangeInt8(0, 2)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int8{}
	actualList = RangeInt8(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int8{}
	actualList = RangeInt8(5, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int8{-5, -4, -3, -2}
	actualList = RangeInt8(-5, -1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int8{-5, -4, -3, -2, -1}
	actualList = RangeInt8(-5, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int8{-5, -4, -3, -2, -1, 0}
	actualList = RangeInt8(-5, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int8{}
	actualList = RangeInt8(-5, 1, -1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int8{}
	actualList = RangeInt8(-5, 1, -1, -2, -3)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int8{1, 2}
	actualList = RangeInt8(1, 3, 1, -2, -3)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []int8{}
	actualList = RangeInt8(1, 3, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	actualList = RangeInt8(1, 100, 10)
	expectedList = []int8{1, 11, 21, 31, 41, 51, 61, 71, 81, 91}
	expectedLen := 10
	actualLen := len(actualList)
	if actualLen != expectedLen || !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeInt8 failed. expected-list=%v, length=%d, actual-list=%v, actual-length=%d", expectedList, expectedLen, actualList, actualLen)
	}
}

func TestRangeUint(t *testing.T) {
	expectedList := []uint{}
	actualList := RangeUint(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint{0, 1}
	actualList = RangeUint(0, 2)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint{}
	actualList = RangeUint(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint{}
	actualList = RangeUint(5, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint{1, 2}
	actualList = RangeUint(1, 3, 1, 2, 3)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint{}
	actualList = RangeUint(1, 3, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint failed. expected=%v, actual=%v", expectedList, actualList)
	}

	actualList = RangeUint(1, 100, 10)
	expectedList = []uint{1, 11, 21, 31, 41, 51, 61, 71, 81, 91}
	expectedLen := 10
	actualLen := len(actualList)
	if actualLen != expectedLen || !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint failed. expected-list=%v, length=%d, actual-list=%v, actual-length=%d", expectedList, expectedLen, actualList, actualLen)
	}
}

func TestRangeUint64(t *testing.T) {
	expectedList := []uint64{}
	actualList := RangeUint64(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint64{0, 1}
	actualList = RangeUint64(0, 2)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint64{}
	actualList = RangeUint64(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint64{}
	actualList = RangeUint64(5, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint64{1, 2}
	actualList = RangeUint64(1, 3, 1, 2, 3)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint64{}
	actualList = RangeUint64(1, 3, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint64 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	actualList = RangeUint64(1, 100, 10)
	expectedList = []uint64{1, 11, 21, 31, 41, 51, 61, 71, 81, 91}
	expectedLen := 10
	actualLen := len(actualList)
	if actualLen != expectedLen || !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint64 failed. expected-list=%v, length=%d, actual-list=%v, actual-length=%d", expectedList, expectedLen, actualList, actualLen)
	}
}

func TestRangeUint32(t *testing.T) {
	expectedList := []uint32{}
	actualList := RangeUint32(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint32{0, 1}
	actualList = RangeUint32(0, 2)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint32{}
	actualList = RangeUint32(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint32{}
	actualList = RangeUint32(5, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint32{1, 2}
	actualList = RangeUint32(1, 3, 1, 2, 3)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint32{}
	actualList = RangeUint32(1, 3, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint32 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	actualList = RangeUint32(1, 100, 10)
	expectedList = []uint32{1, 11, 21, 31, 41, 51, 61, 71, 81, 91}
	expectedLen := 10
	actualLen := len(actualList)
	if actualLen != expectedLen || !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint32 failed. expected-list=%v, length=%d, actual-list=%v, actual-length=%d", expectedList, expectedLen, actualList, actualLen)
	}
}

func TestRangeUint16(t *testing.T) {
	expectedList := []uint16{}
	actualList := RangeUint16(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint16{0, 1}
	actualList = RangeUint16(0, 2)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint16{}
	actualList = RangeUint16(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint16{}
	actualList = RangeUint16(5, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint16{1, 2}
	actualList = RangeUint16(1, 3, 1, 2, 3)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint16{}
	actualList = RangeUint16(1, 3, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint16 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	actualList = RangeUint16(1, 100, 10)
	expectedList = []uint16{1, 11, 21, 31, 41, 51, 61, 71, 81, 91}
	expectedLen := 10
	actualLen := len(actualList)
	if actualLen != expectedLen || !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint16 failed. expected-list=%v, length=%d, actual-list=%v, actual-length=%d", expectedList, expectedLen, actualList, actualLen)
	}
}

func TestRangeUint8(t *testing.T) {
	expectedList := []uint8{}
	actualList := RangeUint8(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint8{0, 1}
	actualList = RangeUint8(0, 2)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint8{}
	actualList = RangeUint8(1, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint8{}
	actualList = RangeUint8(5, 1)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint8{1, 2}
	actualList = RangeUint8(1, 3, 1, 2, 3)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	expectedList = []uint8{}
	actualList = RangeUint8(1, 3, 0)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint8 failed. expected=%v, actual=%v", expectedList, actualList)
	}

	actualList = RangeUint8(1, 100, 10)
	expectedList = []uint8{1, 11, 21, 31, 41, 51, 61, 71, 81, 91}
	expectedLen := 10
	actualLen := len(actualList)
	if actualLen != expectedLen || !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestRangeUint8 failed. expected-list=%v, length=%d, actual-list=%v, actual-length=%d", expectedList, expectedLen, actualList, actualLen)
	}
}
