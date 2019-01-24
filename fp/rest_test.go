package fp

import (
	"reflect"
	"testing"
)

func TestRestInt(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	expectedList := []int{2, 3, 4, 5}
	newList := RestInt(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []int{1}
	expectedList = []int{}
	newList = RestInt(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []int{}
	expectedList = []int{}
	newList = RestInt(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []int{}
	newList = RestInt(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestRestInt64(t *testing.T) {
	list := []int64{1, 2, 3, 4, 5}
	expectedList := []int64{2, 3, 4, 5}
	newList := RestInt64(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []int64{1}
	expectedList = []int64{}
	newList = RestInt64(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []int64{}
	expectedList = []int64{}
	newList = RestInt64(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []int64{}
	newList = RestInt64(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt64 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestRestInt32(t *testing.T) {
	list := []int32{1, 2, 3, 4, 5}
	expectedList := []int32{2, 3, 4, 5}
	newList := RestInt32(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []int32{1}
	expectedList = []int32{}
	newList = RestInt32(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []int32{}
	expectedList = []int32{}
	newList = RestInt32(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []int32{}
	newList = RestInt32(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt32 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestRestInt16(t *testing.T) {
	list := []int16{1, 2, 3, 4, 5}
	expectedList := []int16{2, 3, 4, 5}
	newList := RestInt16(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []int16{1}
	expectedList = []int16{}
	newList = RestInt16(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []int16{}
	expectedList = []int16{}
	newList = RestInt16(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []int16{}
	newList = RestInt16(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt16 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestRestInt8(t *testing.T) {
	list := []int8{1, 2, 3, 4, 5}
	expectedList := []int8{2, 3, 4, 5}
	newList := RestInt8(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []int8{1}
	expectedList = []int8{}
	newList = RestInt8(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []int8{}
	expectedList = []int8{}
	newList = RestInt8(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []int8{}
	newList = RestInt8(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt8 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestRestUInt(t *testing.T) {
	list := []uint{1, 2, 3, 4, 5}
	expectedList := []uint{2, 3, 4, 5}
	newList := RestUint(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []uint{1}
	expectedList = []uint{}
	newList = RestUint(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []uint{}
	expectedList = []uint{}
	newList = RestUint(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []uint{}
	newList = RestUint(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestRestUInt64(t *testing.T) {
	list := []uint64{1, 2, 3, 4, 5}
	expectedList := []uint64{2, 3, 4, 5}
	newList := RestUint64(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []uint64{1}
	expectedList = []uint64{}
	newList = RestUint64(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []uint64{}
	expectedList = []uint64{}
	newList = RestUint64(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []uint64{}
	newList = RestUint64(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint64 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestRestUInt32(t *testing.T) {
	list := []uint32{1, 2, 3, 4, 5}
	expectedList := []uint32{2, 3, 4, 5}
	newList := RestUint32(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []uint32{1}
	expectedList = []uint32{}
	newList = RestUint32(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []uint32{}
	expectedList = []uint32{}
	newList = RestUint32(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []uint32{}
	newList = RestUint32(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint32 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestRestUInt16(t *testing.T) {
	list := []uint16{1, 2, 3, 4, 5}
	expectedList := []uint16{2, 3, 4, 5}
	newList := RestUint16(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []uint16{1}
	expectedList = []uint16{}
	newList = RestUint16(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []uint16{}
	expectedList = []uint16{}
	newList = RestUint16(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []uint16{}
	newList = RestUint16(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint16 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestRestUInt8(t *testing.T) {
	list := []uint8{1, 2, 3, 4, 5}
	expectedList := []uint8{2, 3, 4, 5}
	newList := RestUint8(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []uint8{1}
	expectedList = []uint8{}
	newList = RestUint8(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []uint8{}
	expectedList = []uint8{}
	newList = RestUint8(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []uint8{}
	newList = RestUint8(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint8 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestRestFloat64(t *testing.T) {
	list := []float64{1, 2, 3, 4, 5}
	expectedList := []float64{2, 3, 4, 5}
	newList := RestFloat64(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []float64{1}
	expectedList = []float64{}
	newList = RestFloat64(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []float64{}
	expectedList = []float64{}
	newList = RestFloat64(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []float64{}
	newList = RestFloat64(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat64 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestRestFloat32(t *testing.T) {
	list := []float32{1, 2, 3, 4, 5}
	expectedList := []float32{2, 3, 4, 5}
	newList := RestFloat32(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []float32{1}
	expectedList = []float32{}
	newList = RestFloat32(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []float32{}
	expectedList = []float32{}
	newList = RestFloat32(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []float32{}
	newList = RestFloat32(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat32 failed. expected=%v, actual=%v", expectedList, newList)
	}
}

func TestRestStr(t *testing.T) {
	list := []string{"Ravan", "Ram", "Shyam", "Sita", "Radha"}
	expectedList := []string{"Ram", "Shyam", "Sita", "Radha"}
	newList := RestStr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestStr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []string{"Ram"}
	expectedList = []string{}
	newList = RestStr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestStr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []string{}
	expectedList = []string{}
	newList = RestStr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestStr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []string{}
	newList = RestStr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestStr failed. expected=%v, actual=%v", expectedList, newList)
	}
}
