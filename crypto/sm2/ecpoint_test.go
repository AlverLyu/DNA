package sm2

import (
	"crypto/elliptic"
	"math/big"
	"testing"
)

const (
	x1  = "0x6A2C9DFF321AC97084E822F903BDD6A6D2D9C45AC672BD2CDEDE29CFA95032B4"
	y1  = "0x67E8360AAF63A6BCBD9D855BE1612B5E4D6A9D9ED07C82F25CEB9929A50AD91F"
	y1n = "0x1D5AA0939CA0A85C2B1B9ED9DE0ECC7FF807E5F28BC8CE8B1543426163E706A4"

	x2 = "0x1CDD78ABCCD705A6054358B49306FA6A47229C68200E89686F46705ADBC19A45"
	y2 = "0x4F9CA9715179C725054E60A1BFA5E14CEEF635D6F7C4C093C41DF3BD80D99A89"

	x12 = "0x3F8F725C7D8B61A4545100F228F6E5838438F4E30C978AEE5605CD288274B422"
	y12 = "0x82C4841272109E0E3BC5CE22ACF8536BD221ACF28354E2121CCDD0351B8F248B"

	sumX = "0x63BED55B23EF616699B4D29D92ABE6817DAAC2B8596FCBF8E5E9AE06B734761B"
	sumY = "0x1086F91578676A007132503A98F9EF54F303BEF3A34F4D1CC93FD9D01A362B14"

	kx10b = "0xEC89E0146C4F98AD404169578954023161BC4FAF3C5B1B7197369FFD97C7675"
	ky10b = "0x22A4855C11C5C2EE5E50CF8C931D633478537DFF277D39DBCF4AE7EDFD234D13"

	kx32b = "0x5EE38C2401C8484B25201D4D3D89A44A9854BC324D7926BC999E0CF0D66C2310"
	ky32b = "0x531FC602A0903E540E5CACA28F6094B3E5F379C00605F2F3BBFEC6E15B7A059B"

	kx100b = "0xF2A1658D93B74B181D45B6F28A917D22B1DCD74A99AF18885FCBFC6CD270860"
	ky100b = "0x40EA9D1364816A1B8F36B159C68E334645165FF425B091346F440A881971CB37"

	kx200b = "0x2DDE01321F01ACA2F3BF5A7F29E49080C3A1C815478A6A33EF6B7860D7FE6EE0"
	ky200b = "0x50CEB7652CD39F62B2EEBD562BFFE5E7EFF1046B602F91A47766FA0DD0401C90"
)

func newTestPoint(x, y string) *ECPoint {
	if EcParams == nil {
		EcParams = new(elliptic.CurveParams)
		EcParams.P, _ = new(big.Int).SetString("8542D69E4C044F18E8B92435BF6FF7DE457283915C45517D722EDB8B08F1DFC3", 16)
		EcParams.N, _ = new(big.Int).SetString("8542D69E4C044F18E8B92435BF6FF7DD297720630485628D5AE74EE7C32E79B7", 16)
		Sm2ParamA, _ = new(big.Int).SetString("787968B4FA32C3FD2417842E73BBFEFF2F3C848B6831D7E0EC65228B3937E498", 16)
		EcParams.B, _ = new(big.Int).SetString("63E4C6D3B23B0C849CF84241484BFE48F61D59A5B16BA06E6E12D1DA27C5249A", 16)

		EcParams.Gx, _ = new(big.Int).SetString("421DEBD61B62EAB6746434EBC3CC315E32220B3BADD50BDC4C4E6C147FEDD43D", 16)
		EcParams.Gy, _ = new(big.Int).SetString("0680512BCBB42C07D47349D2153B70C4E5D7FDFCBFA36EA1A85841B9E46E09A2", 16)
		EcParams.BitSize = 256
		EcParams.Name = "sm2"

	}
	ecpoint := NewECPoint()
	if ecpoint == nil {
		return nil
	}
	if len(x) > 0 {
		ecpoint.X.value.SetString(x, 0)
	}
	if len(y) > 0 {
		ecpoint.Y.value.SetString(y, 0)
	}
	return ecpoint
}

func printXY(t *testing.T, pt *ECPoint) {
	t.Logf("\tX = 0x%X", pt.X.value)
	t.Logf("\tY = 0x%X", pt.Y.value)
}

func equal(pt0 *ECPoint, pt1 *ECPoint) bool {
	return pt0.X.value.Cmp(pt1.X.value) == 0 &&
		pt0.Y.value.Cmp(pt1.Y.value) == 0
}

func TestNegPoint(t *testing.T) {
	pt := newTestPoint(x1, y1)
	t.Logf("Point A:")
	t.Logf("\tX = 0x%X", pt.X.value)
	t.Logf("\tY = 0x%X", pt.Y.value)

	ptNeg := newTestPoint("", "")
	ptNeg.Neg(pt)
	t.Logf("Point -A:")
	t.Logf("\tX = 0x%X", ptNeg.X.value)
	t.Logf("\tY = 0x%X", ptNeg.Y.value)

	result := newTestPoint(x1, y1n)

	if !equal(result, ptNeg) {
		t.Logf("Expected result is:")
		printXY(t, result)
		t.Fail()
	}

	ptInfinite := newTestPoint("", "")
	ptNeg.Neg(ptInfinite)
	if !equal(ptNeg, ptInfinite) {
		t.Logf("ERROR: -O != O")
		t.Fail()
	}
}

func TestAPlusO(t *testing.T) {
	pt1 := newTestPoint(x1, y1)
	t.Logf("Point A:")
	printXY(t, pt1)

	ptInfinite := newTestPoint("0", "0")

	sum := newTestPoint("", "")
	sum.Add(pt1, ptInfinite)
	t.Logf("A + O:")
	printXY(t, sum)
	if !equal(pt1, sum) {
		t.Error("ERROR: A + O != A")
	}
}

func TestOPlusA(t *testing.T) {
	pt1 := newTestPoint(x1, y1)
	t.Logf("Point A:")
	printXY(t, pt1)

	ptInfinite := newTestPoint("0", "0")

	sum := newTestPoint("", "")

	sum.Add(ptInfinite, pt1)
	t.Logf("O + A:")
	printXY(t, sum)
	if !equal(sum, pt1) {
		t.Error("ERROR: O + A != A")
	}
}

func TestAddNeg(t *testing.T) {
	pt1 := newTestPoint(x1, y1)
	t.Logf("Point A:")
	printXY(t, pt1)

	pt2 := newTestPoint("", "")
	pt2.Neg(pt1)
	t.Logf("-A:")
	printXY(t, pt2)

	pt2.Add(pt1, pt2)
	t.Logf("A + -A:")
	printXY(t, pt2)
	if !pt2.IsInfinity() {
		t.Error("ERROR: A + -A != O")
	}
}

func TestAddPoint(t *testing.T) {
	pt1 := newTestPoint(x1, y1)
	t.Logf("Point A:")
	t.Logf("\tX = 0x%X", pt1.X.value)
	t.Logf("\tY = 0x%X", pt1.Y.value)

	pt2 := newTestPoint(x2, y2)
	t.Logf("Point B:")
	t.Logf("\tX = 0x%X", pt2.X.value)
	t.Logf("\tY = 0x%X", pt2.Y.value)

	sum := newTestPoint("", "")
	sum.Add(pt1, pt2)
	t.Logf("A + B:")
	t.Logf("\tX = 0x%X", sum.X.value)
	t.Logf("\tY = 0x%X", sum.Y.value)

	res := newTestPoint(sumX, sumY)
	if sum.X.value.Cmp(res.X.value) != 0 || sum.Y.value.Cmp(res.Y.value) != 0 {
		t.Logf("Expected result is:")
		t.Logf("\tX = 0x%X", res.X.value)
		t.Logf("\tY = 0x%X", res.Y.value)
		t.Fail()
	}
}

func TestTwice(t *testing.T) {
	pt := newTestPoint(x1, y1)
	t.Logf("Point A:")
	t.Logf("\tX = 0x%X", pt.X.value)
	t.Logf("\tY = 0x%X", pt.Y.value)

	pt2 := pt.Twice()
	t.Logf("B = 2A:")
	t.Logf("\tX = 0x%X", pt2.X.value)
	t.Logf("\tY = 0x%X", pt2.Y.value)

	ptResX, _ := new(big.Int).SetString(x12, 0)
	ptResY, _ := new(big.Int).SetString(y12, 0)
	if pt2.X.value.Cmp(ptResX) != 0 || pt2.Y.value.Cmp(ptResY) != 0 {
		t.Logf("Expected result is:")
		t.Logf("\tX = 0x%X", ptResX)
		t.Logf("\tY = 0x%X", ptResY)
		t.Fail()
	}
}

func TestNPoint(t *testing.T) {
	pt := newTestPoint(x1, y1)
	t.Logf("Point A:")
	printXY(t, pt)

	k := big.NewInt(1024)
	npt := Multiply(pt, k)
	t.Logf("2^10 * A:")
	printXY(t, npt)

	res := newTestPoint(kx10b, ky10b)
	if !equal(res, npt) {
		t.Logf("Expected result is:")
		printXY(t, res)
		t.Fail()
	}

	k.SetString("4294967296", 10)
	npt = Multiply(pt, k)
	t.Logf("2^32 * A:")
	printXY(t, npt)

	res = newTestPoint(kx32b, ky32b)
	if !equal(res, npt) {
		t.Logf("Expected result is:")
		printXY(t, res)
		t.Fail()
	}

	k.SetString("1267650600228229401496703205376", 10)
	npt = Multiply(pt, k)
	t.Logf("2^100 * A:")
	printXY(t, npt)

	res = newTestPoint(kx100b, ky100b)
	if !equal(res, npt) {
		t.Logf("Expected result is:")
		printXY(t, res)
		t.Fail()
	}

	k.SetString("1606938044258990275541962092341162602522202993782792835301376", 10)
	npt = Multiply(pt, k)
	t.Logf("2^200 * A:")
	printXY(t, npt)

	res = newTestPoint(kx200b, ky200b)
	if !equal(res, npt) {
		t.Logf("Expected result is:")
		printXY(t, res)
		t.Fail()
	}
}
