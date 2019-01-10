package fp

import "testing"

func TestTrue(t *testing.T) {
	if True() == false {
		t.Errorf("TestTrue failed. expected=true, actual=false")
	}
}

func TestFalse(t *testing.T) {
	if False() == true {
		t.Errorf("TestFalse failed. expected=false, actual=true")
	}
}
