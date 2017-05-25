package sm2

import (
	"crypto/elliptic"
	"math/big"
	"testing"
)

const (
	P   = "8542D69E4C044F18E8B92435BF6FF7DE457283915C45517D722EDB8B08F1DFC3"
	ea  = "6a2c9dff321ac97084e822f903bdd6a6d2d9c45ac672bd2cdede29cfa95032b4"
	eb  = "1cdd78abccd705a6054358b49306fa6a47229c68200e89686f46705adbc19a45"
	na  = "1b16389f19e985a863d1013cbbb221377298bf3695d294509350b1bb5fa1ad0f"
	sum = "1c7400cb2ed7ffda1725777d754d932d489dd318a3bf517dbf5be9f7c1fed36"
	dif = "4d4f25536543c3ca7fa4ca4470b6dc3c8bb727f2a66433c46f97b974cd8e986f"
	pro = "66623b4f7619549cabf902e5c88722635079f15b19a8c0be7be29fbac0941984"
	sq  = "5fecc72fc253796dca3f69ac7be238adc832bd50f65a9bdc1bd632ceccd041cf"
)

var curve *elliptic.CurveParams = nil

func getCurve() *elliptic.CurveParams {
	if curve == nil {
		curve = new(elliptic.CurveParams)
		curve.P, _ = new(big.Int).SetString(P, 16)
	}
	return curve
}

func NewElement(v string) *ECFieldElement {
	e := NewECFieldElement()
	e.curveParam = getCurve()
	if len(v) > 0 {
		e.value, _ = new(big.Int).SetString(v, 16)
	}
	return e
}

func TestAdd(t *testing.T) {
	e0 := NewElement(ea)
	e1 := NewElement(eb)
	t.Logf("a = 0x%x", e0.value)
	t.Logf("b = 0x%x", e1.value)
	s := NewElement("")
	s.Add(e0, e1)
	t.Logf("a + b = 0x%x", s.value)
	result, _ := new(big.Int).SetString(sum, 16)
	if s.value.Cmp(result) != 0 {
		t.Logf("expect result is 0x%x", result)
		t.Fail()
	}
}

func TestAddZero(t *testing.T) {
	e0 := NewElement(ea)
	e1 := NewElement("")
	s := NewElement("")
	s.Add(e0, e1)
	if s.value.Cmp(e0.value) != 0 {
		t.Error("ERROR: a + 0 != a")
	}
}

func TestSub(t *testing.T) {
	e0 := NewElement(ea)
	e1 := NewElement(eb)

	t.Logf("a = 0x%x", e0.value)
	t.Logf("b = 0x%x", e1.value)

	d := NewElement("")
	d.Sub(e0, e1)
	t.Logf("a - b = 0x%x", d.value)
	result, _ := new(big.Int).SetString(dif, 16)
	if d.value.Cmp(result) != 0 {
		t.Logf("expected result is 0x%x", result)
		t.Fail()
	}
}

func TestNeg(t *testing.T) {
	e := NewElement(ea)
	t.Logf("a = 0x%x", e.value)
	eN := NewElement("")
	eN.Neg(e)
	t.Logf("-a = 0x%x", eN.value)
	result, _ := new(big.Int).SetString(na, 16)
	if eN.value.Cmp(result) != 0 {
		t.Logf("expected result is 0x%x", result)
		t.Fail()
	}
}

func TestMul(t *testing.T) {
	e0 := NewElement(ea)
	e1 := NewElement(eb)

	t.Logf("a = 0x%x", e0.value)
	t.Logf("b = 0x%x", e1.value)

	eM := NewElement("")
	eM.Mul(e0, e1)
	t.Logf("a * b = 0x%x", eM.value)

	result, _ := new(big.Int).SetString(pro, 16)
	if eM.value.Cmp(result) != 0 {
		t.Logf("expected result is 0x%x", result)
		t.Fail()
	}
}

func TestDiv(t *testing.T) {
	e0 := NewElement(pro)
	e1 := NewElement(ea)
	t.Logf("a = 0x%x", e0.value)
	t.Logf("b = 0x%x", e1.value)

	eD := NewElement("")
	eD.Div(e0, e1)
	t.Logf("a / b = 0x%x", eD.value)

	result, _ := new(big.Int).SetString(eb, 16)
	if eD.value.Cmp(result) != 0 {
		t.Logf("expected result is 0x%x", result)
		t.Fail()
	}
}

func TestSqure(t *testing.T) {
	e := NewElement(ea)
	t.Logf("a = 0x%x", e.value)
	eS := e.Square()
	t.Logf("a^2 = 0x%x", eS.value)
	result, _ := new(big.Int).SetString(sq, 16)
	if eS.value.Cmp(result) != 0 {
		t.Logf("expected result is 0x%x", result)
		t.Fail()
	}
}

func TestSqrt(t *testing.T) {
	e := NewElement(sq)
	e.value.SetString(sq, 16)
	t.Logf("a = 0x%x", e.value)
	eSqrt := Sqrt(e.value, e.curveParam)
	result, _ := new(big.Int).SetString(ea, 16)
	t.Logf("Sqrt(a) = 0x%x", eSqrt)
	if eSqrt.Cmp(result) != 0 {
		t.Logf("expected result is 0x%x", result)
		t.Fail()
	}
}
