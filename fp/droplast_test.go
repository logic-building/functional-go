package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestDropLastInt(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	expectedList := []int{1, 2, 3, 4}
	actualList := DropLastInt(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []int{1, 2}
	expectedList = []int{1}
	actualList = DropLastInt(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []int{1}
	expectedList = []int{}
	actualList = DropLastInt(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []int{}
	expectedList = []int{}
	actualList = DropLastInt(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []int{}
	actualList = DropLastInt(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastInt64(t *testing.T) {
	list := []int64{1, 2, 3, 4, 5}
	expectedList := []int64{1, 2, 3, 4}
	actualList := DropLastInt64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []int64{1, 2}
	expectedList = []int64{1}
	actualList = DropLastInt64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []int64{1}
	expectedList = []int64{}
	actualList = DropLastInt64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []int64{}
	expectedList = []int64{}
	actualList = DropLastInt64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []int64{}
	actualList = DropLastInt64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastInt32(t *testing.T) {
	list := []int32{1, 2, 3, 4, 5}
	expectedList := []int32{1, 2, 3, 4}
	actualList := DropLastInt32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []int32{1, 2}
	expectedList = []int32{1}
	actualList = DropLastInt32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []int32{1}
	expectedList = []int32{}
	actualList = DropLastInt32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []int32{}
	expectedList = []int32{}
	actualList = DropLastInt32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []int32{}
	actualList = DropLastInt32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastInt16(t *testing.T) {
	list := []int16{1, 2, 3, 4, 5}
	expectedList := []int16{1, 2, 3, 4}
	actualList := DropLastInt16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []int16{1, 2}
	expectedList = []int16{1}
	actualList = DropLastInt16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []int16{1}
	expectedList = []int16{}
	actualList = DropLastInt16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []int16{}
	expectedList = []int16{}
	actualList = DropLastInt16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []int16{}
	actualList = DropLastInt16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastInt8(t *testing.T) {
	list := []int8{1, 2, 3, 4, 5}
	expectedList := []int8{1, 2, 3, 4}
	actualList := DropLastInt8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []int8{1, 2}
	expectedList = []int8{1}
	actualList = DropLastInt8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []int8{1}
	expectedList = []int8{}
	actualList = DropLastInt8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []int8{}
	expectedList = []int8{}
	actualList = DropLastInt8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []int8{}
	actualList = DropLastInt8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastInt8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastUint(t *testing.T) {
	list := []uint{1, 2, 3, 4, 5}
	expectedList := []uint{1, 2, 3, 4}
	actualList := DropLastUint(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []uint{1, 2}
	expectedList = []uint{1}
	actualList = DropLastUint(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []uint{1}
	expectedList = []uint{}
	actualList = DropLastUint(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []uint{}
	expectedList = []uint{}
	actualList = DropLastUint(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []uint{}
	actualList = DropLastUint(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastUint64(t *testing.T) {
	list := []uint64{1, 2, 3, 4, 5}
	expectedList := []uint64{1, 2, 3, 4}
	actualList := DropLastUint64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []uint64{1, 2}
	expectedList = []uint64{1}
	actualList = DropLastUint64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []uint64{1}
	expectedList = []uint64{}
	actualList = DropLastUint64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []uint64{}
	expectedList = []uint64{}
	actualList = DropLastUint64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []uint64{}
	actualList = DropLastUint64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastUint32(t *testing.T) {
	list := []uint32{1, 2, 3, 4, 5}
	expectedList := []uint32{1, 2, 3, 4}
	actualList := DropLastUint32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []uint32{1, 2}
	expectedList = []uint32{1}
	actualList = DropLastUint32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []uint32{1}
	expectedList = []uint32{}
	actualList = DropLastUint32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []uint32{}
	expectedList = []uint32{}
	actualList = DropLastUint32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []uint32{}
	actualList = DropLastUint32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastUint16(t *testing.T) {
	list := []uint16{1, 2, 3, 4, 5}
	expectedList := []uint16{1, 2, 3, 4}
	actualList := DropLastUint16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []uint16{1, 2}
	expectedList = []uint16{1}
	actualList = DropLastUint16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []uint16{1}
	expectedList = []uint16{}
	actualList = DropLastUint16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []uint16{}
	expectedList = []uint16{}
	actualList = DropLastUint16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []uint16{}
	actualList = DropLastUint16(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint16 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastUint8(t *testing.T) {
	list := []uint8{1, 2, 3, 4, 5}
	expectedList := []uint8{1, 2, 3, 4}
	actualList := DropLastUint8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []uint8{1, 2}
	expectedList = []uint8{1}
	actualList = DropLastUint8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []uint8{1}
	expectedList = []uint8{}
	actualList = DropLastUint8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []uint8{}
	expectedList = []uint8{}
	actualList = DropLastUint8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []uint8{}
	actualList = DropLastUint8(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastUint8 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastStr(t *testing.T) {
	list := []string{"1", "2", "3", "4", "5"}
	expectedList := []string{"1", "2", "3", "4"}
	actualList := DropLastStr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastStr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []string{"1", "2"}
	expectedList = []string{"1"}
	actualList = DropLastStr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastStr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []string{"1"}
	expectedList = []string{}
	actualList = DropLastStr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastStr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []string{}
	expectedList = []string{}
	actualList = DropLastStr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastStr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []string{}
	actualList = DropLastStr(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastStr failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastBool(t *testing.T) {
	list := []bool{true, true, true, true, false}
	expectedList := []bool{true, true, true, true}
	actualList := DropLastBool(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastBool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []bool{true, true}
	expectedList = []bool{true}
	actualList = DropLastBool(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastBool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []bool{true}
	expectedList = []bool{}
	actualList = DropLastBool(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastBool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []bool{}
	expectedList = []bool{}
	actualList = DropLastBool(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastBool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []bool{}
	actualList = DropLastBool(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastBool failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastFloat32(t *testing.T) {
	list := []float32{1, 2, 3, 4, 5}
	expectedList := []float32{1, 2, 3, 4}
	actualList := DropLastFloat32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastFloat32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []float32{1, 2}
	expectedList = []float32{1}
	actualList = DropLastFloat32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastFloat32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []float32{1}
	expectedList = []float32{}
	actualList = DropLastFloat32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastFloat32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []float32{}
	expectedList = []float32{}
	actualList = DropLastFloat32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastFloat32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []float32{}
	actualList = DropLastFloat32(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastFloat32 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}

func TestDropLastFloat64(t *testing.T) {
	list := []float64{1, 2, 3, 4, 5}
	expectedList := []float64{1, 2, 3, 4}
	actualList := DropLastFloat64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastFloat64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []float64{1, 2}
	expectedList = []float64{1}
	actualList = DropLastFloat64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastFloat64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []float64{1}
	expectedList = []float64{}
	actualList = DropLastFloat64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastFloat64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = []float64{}
	expectedList = []float64{}
	actualList = DropLastFloat64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastFloat64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}

	list = nil
	expectedList = []float64{}
	actualList = DropLastFloat64(list)
	if !reflect.DeepEqual(expectedList, actualList) {
		t.Errorf("TestDropLastFloat64 failed. acutal_list=%v, expected_list=%v", actualList, expectedList)
	}
}
