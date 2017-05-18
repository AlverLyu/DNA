package common

import (
	"bytes"
	"testing"
)

func TestSerialize(t *testing.T) {
	var f0 Fixed64 = 1000
	var f1 Fixed64
	buf := new(bytes.Buffer)
	err := f0.Serialize(buf)
	if err != nil {
		t.Logf("Serialize error: %s", err.Error())
		t.Fail()
	}

	f1.Deserialize(buf)
	if f1 != f0 {
		t.Logf("Deserialized value does not equals the original one")
		t.Fail()
	}
}

func TestString(t *testing.T) {
	var f Fixed64 = 1
	exp := "0.00000001"
	if f.String() != exp {
		t.Logf("The Fixed64 value is 1, the result is %s while expects %s", f.String(), exp)
		t.Fail()
	}

	f = -1
	exp = "-0.00000001"
	if f.String() != exp {
		t.Logf("The Fixed64 value is -1, the result is %s while expects %s", f.String(), exp)
		t.Fail()
	}

	f = 1234567890
	exp = "12.34567890"
	if f.String() != exp {
		t.Logf("The Fixed64 value is 1234567890, the result is %s while expects %s", f.String(), exp)
		t.Fail()
	}
}
