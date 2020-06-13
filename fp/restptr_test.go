package fp

import (
	_ "errors"
	"reflect"
	"testing"
)

func TestRestIntPtr(t *testing.T) {
	var v1 int = 1
	var v2 int = 2
	var v3 int = 3
	var v4 int = 4
	var v5 int = 5
	

	list := []*int{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*int{&v2, &v3, &v4, &v5}
	newList := RestIntPtr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestIntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*int{&v1}
	expectedList = []*int{}
	newList = RestIntPtr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestIntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*int{}
	expectedList = []*int{}
	newList = RestIntPtr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestIntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*int{}
	newList = RestIntPtr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestIntPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}


func TestRestInt64Ptr(t *testing.T) {
	var v1 int64 = 1
	var v2 int64 = 2
	var v3 int64 = 3
	var v4 int64 = 4
	var v5 int64 = 5
	

	list := []*int64{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*int64{&v2, &v3, &v4, &v5}
	newList := RestInt64Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*int64{&v1}
	expectedList = []*int64{}
	newList = RestInt64Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*int64{}
	expectedList = []*int64{}
	newList = RestInt64Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*int64{}
	newList = RestInt64Ptr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}


func TestRestInt32Ptr(t *testing.T) {
	var v1 int32 = 1
	var v2 int32 = 2
	var v3 int32 = 3
	var v4 int32 = 4
	var v5 int32 = 5
	

	list := []*int32{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*int32{&v2, &v3, &v4, &v5}
	newList := RestInt32Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*int32{&v1}
	expectedList = []*int32{}
	newList = RestInt32Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*int32{}
	expectedList = []*int32{}
	newList = RestInt32Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*int32{}
	newList = RestInt32Ptr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}


func TestRestInt16Ptr(t *testing.T) {
	var v1 int16 = 1
	var v2 int16 = 2
	var v3 int16 = 3
	var v4 int16 = 4
	var v5 int16 = 5
	

	list := []*int16{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*int16{&v2, &v3, &v4, &v5}
	newList := RestInt16Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*int16{&v1}
	expectedList = []*int16{}
	newList = RestInt16Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*int16{}
	expectedList = []*int16{}
	newList = RestInt16Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*int16{}
	newList = RestInt16Ptr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}


func TestRestInt8Ptr(t *testing.T) {
	var v1 int8 = 1
	var v2 int8 = 2
	var v3 int8 = 3
	var v4 int8 = 4
	var v5 int8 = 5
	

	list := []*int8{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*int8{&v2, &v3, &v4, &v5}
	newList := RestInt8Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*int8{&v1}
	expectedList = []*int8{}
	newList = RestInt8Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*int8{}
	expectedList = []*int8{}
	newList = RestInt8Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*int8{}
	newList = RestInt8Ptr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestInt8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}


func TestRestUintPtr(t *testing.T) {
	var v1 uint = 1
	var v2 uint = 2
	var v3 uint = 3
	var v4 uint = 4
	var v5 uint = 5
	

	list := []*uint{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*uint{&v2, &v3, &v4, &v5}
	newList := RestUintPtr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*uint{&v1}
	expectedList = []*uint{}
	newList = RestUintPtr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*uint{}
	expectedList = []*uint{}
	newList = RestUintPtr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*uint{}
	newList = RestUintPtr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUintPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}


func TestRestUint64Ptr(t *testing.T) {
	var v1 uint64 = 1
	var v2 uint64 = 2
	var v3 uint64 = 3
	var v4 uint64 = 4
	var v5 uint64 = 5
	

	list := []*uint64{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*uint64{&v2, &v3, &v4, &v5}
	newList := RestUint64Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*uint64{&v1}
	expectedList = []*uint64{}
	newList = RestUint64Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*uint64{}
	expectedList = []*uint64{}
	newList = RestUint64Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*uint64{}
	newList = RestUint64Ptr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}


func TestRestUint32Ptr(t *testing.T) {
	var v1 uint32 = 1
	var v2 uint32 = 2
	var v3 uint32 = 3
	var v4 uint32 = 4
	var v5 uint32 = 5
	

	list := []*uint32{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*uint32{&v2, &v3, &v4, &v5}
	newList := RestUint32Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*uint32{&v1}
	expectedList = []*uint32{}
	newList = RestUint32Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*uint32{}
	expectedList = []*uint32{}
	newList = RestUint32Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*uint32{}
	newList = RestUint32Ptr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}


func TestRestUint16Ptr(t *testing.T) {
	var v1 uint16 = 1
	var v2 uint16 = 2
	var v3 uint16 = 3
	var v4 uint16 = 4
	var v5 uint16 = 5
	

	list := []*uint16{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*uint16{&v2, &v3, &v4, &v5}
	newList := RestUint16Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*uint16{&v1}
	expectedList = []*uint16{}
	newList = RestUint16Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*uint16{}
	expectedList = []*uint16{}
	newList = RestUint16Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*uint16{}
	newList = RestUint16Ptr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint16Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}


func TestRestUint8Ptr(t *testing.T) {
	var v1 uint8 = 1
	var v2 uint8 = 2
	var v3 uint8 = 3
	var v4 uint8 = 4
	var v5 uint8 = 5
	

	list := []*uint8{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*uint8{&v2, &v3, &v4, &v5}
	newList := RestUint8Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*uint8{&v1}
	expectedList = []*uint8{}
	newList = RestUint8Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*uint8{}
	expectedList = []*uint8{}
	newList = RestUint8Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*uint8{}
	newList = RestUint8Ptr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestUint8Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}


func TestRestStrPtr(t *testing.T) {
	var v1 string = "1"
	var v2 string = "2"
	var v3 string = "3"
	var v4 string = "4"
	var v5 string = "5"
	

	list := []*string{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*string{&v2, &v3, &v4, &v5}
	newList := RestStrPtr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*string{&v1}
	expectedList = []*string{}
	newList = RestStrPtr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*string{}
	expectedList = []*string{}
	newList = RestStrPtr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*string{}
	newList = RestStrPtr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestStrPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}


func TestRestBoolPtr(t *testing.T) {
	var vt bool = true
	var vf bool = false
	

	list := []*bool{&vt, &vf, &vf, &vf, &vt}
	expectedList := []*bool{&vf, &vf, &vf, &vt}
	newList := RestBoolPtr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestBoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*bool{&vt}
	expectedList = []*bool{}
	newList = RestBoolPtr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestBoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*bool{}
	expectedList = []*bool{}
	newList = RestBoolPtr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestBoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*bool{}
	newList = RestBoolPtr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestBoolPtr failed. expected=%v, actual=%v", expectedList, newList)
	}
}


func TestRestFloat32Ptr(t *testing.T) {
	var v1 float32 = 1
	var v2 float32 = 2
	var v3 float32 = 3
	var v4 float32 = 4
	var v5 float32 = 5
	

	list := []*float32{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*float32{&v2, &v3, &v4, &v5}
	newList := RestFloat32Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*float32{&v1}
	expectedList = []*float32{}
	newList = RestFloat32Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*float32{}
	expectedList = []*float32{}
	newList = RestFloat32Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*float32{}
	newList = RestFloat32Ptr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat32Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}


func TestRestFloat64Ptr(t *testing.T) {
	var v1 float64 = 1
	var v2 float64 = 2
	var v3 float64 = 3
	var v4 float64 = 4
	var v5 float64 = 5
	

	list := []*float64{&v1, &v2, &v3, &v4, &v5}
	expectedList := []*float64{&v2, &v3, &v4, &v5}
	newList := RestFloat64Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*float64{&v1}
	expectedList = []*float64{}
	newList = RestFloat64Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	list = []*float64{}
	expectedList = []*float64{}
	newList = RestFloat64Ptr(list)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}

	expectedList = []*float64{}
	newList = RestFloat64Ptr(nil)
	if !reflect.DeepEqual(expectedList, newList) {
		t.Errorf("RestFloat64Ptr failed. expected=%v, actual=%v", expectedList, newList)
	}
}

