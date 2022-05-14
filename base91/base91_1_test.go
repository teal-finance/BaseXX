// #region <editor-fold desc="Preamble">
// Copyright (c) 2017-2020 Denis Subbotin, Philip Schlump,
//                         Nika Jones, Steven Allen, MoonFruit
// Copyright (c) 2022      Teal.Finance contributors
//
// This file is a modified copy from https://github.com/mr-tron/base58
// The source code has been adapted to support other bases.
// This file is now part of BaseXX under the terms of the MIT License.
// SPDX-License-Identifier: MIT
//
// BaseXX is distributed WITHOUT ANY WARRANTY.
// See the LICENSE file alongside the source files
// or online at <https://opensource.org/licenses/MIT>.
// #endregion </editor-fold>

package base91

import (
	"bytes"
	"encoding/hex"
	"math/rand"
	"testing"

	equimBase91 "github.com/Equim-chan/base91-go"
	bproctorBase91 "github.com/bproctor/base91"

	// breezechenBase91 "github.com/breezechen/base91"

	majestrateBase91 "github.com/majestrate/base91"
	mtraverBase91 "github.com/mtraver/base91"
	// acBase91 "github.com/teal-finance/BaseXX/ac/base91"
)

const nn = 11
const nnn = 1024 // power of two to speed up the % modulo
var bin [][]byte

const benchChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~'"
//nst benchChar2 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~\""

var benchEncoding = NewEncoding(benchChars)


var tstEncoding = NewEncoding(btcDigits[:Base])

// func init() {
// 	// If we do not seed the prng - it will default to a seed of (1)
// 	// https://golang.org/pkg/math/rand/#Seed
// 	rand.Seed(time.Now().UTC().UnixNano())
// }

func setup[T string | []byte](data *[]T, encode func([]byte) T) {
	if len(*data) > 0 {
		return
	}
	setupBin()
	// pre-make the test pairs, so it doesn't take up benchmark time...
	*data = make([]T, 0, nnn)
	for i := 0; i < nnn; i++ {
		str := encode(bin[i])
		*data = append(*data, str)
	}
}

func setupBin() {
	if len(bin) > 0 {
		return
	}
	bin = make([][]byte, 0, nnn)
	for i := 0; i < nnn; i++ {
		b := make([]byte, rand.Intn(nn))
		rand.Read(b)
		bin = append(bin, b)
	}
}

func randEncoding() *Encoding {
	// Permutes [0, 127] and returns the first XX elements according to the BaseXX.
	var randomness [128]byte
	rand.Read(randomness[:])

	var bts [128]byte
	for i, r := range randomness {
		j := int(r) % (i + 1)
		bts[i] = bts[j]
		bts[j] = byte(i)
	}
	return NewEncoding(string(bts[:Base]))
}

var btcDigits = "" +
	"123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz" +
	" !0OIl()*+[\\]^_`{|}~;:#$<=>%&',-./?@"

func TestInvalidEncodingTooShort(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on alphabet being too short did not occur")
		}
	}()

	_ = NewEncoding(btcDigits[:Base-1]) // too short
}

func TestInvalidEncodingTooLong(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on alphabet being too long did not occur")
		}
	}()

	_ = NewEncoding(btcDigits) // too long
}

func TestInvalidEncodingNon127(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on alphabet containing non-ascii chars did not occur")
		}
	}()

	_ = NewEncoding("\xFF" + btcDigits[:Base-1]) // good length
}

func TestInvalidEncodingDup(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on alphabet containing duplicate chars did not occur")
		}
	}()

	_ = NewEncoding(btcDigits[:1] + btcDigits[:Base-1]) // good length, but 1st char duplicated
}

func TestFastEqTrivialEncodingAndDecoding(t *testing.T) {
	for k := 0; k < 10; k++ {
		testEncDecLoop(t, randEncoding())
	}
	testEncDecLoop(t, StdEncoding)
	testEncDecLoop(t, tstEncoding)
}

func testEncDecLoop(t *testing.T, enc *Encoding) {
	for j := 1; j < 256; j++ {
		var b = make([]byte, j)
		for i := 0; i < 100; i++ {
			rand.Read(b)
			fe := enc.EncodeToString(b)

			fd, err := enc.DecodeString(fe)
			if err != nil {
				t.Errorf("fast error: %v", err)
			}

			if hex.EncodeToString(b) != hex.EncodeToString(fd) {
				t.Errorf("decoding err: %s != %s", hex.EncodeToString(b), hex.EncodeToString(fd))
			}
		}
	}
}

var sAscii []string

func TestBDecode(t *testing.T) {
	setup(&sAscii, benchEncoding.EncodeToString)

	for i := 0; i < nnn; i++ {
		b, err := benchEncoding.DecodeString(sAscii[i])
		if err != nil {
			t.Fatalf("#%d err=%v ascii=%v", i, err, sAscii[i])
		}
		if bytes.Compare(b, bin[i]) != 0 {
			t.Errorf("#%d ascii: %v", i, sAscii[i])
			t.Errorf("#%d want: %x", i, bin[i])
			t.Errorf("#%d got : %x", i, b)
			t.FailNow()
		}
	}
}

func BenchmarkEncode(b *testing.B) {
	setupBin()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		benchEncoding.Encode(bin[i%nnn])
	}
}

func BenchmarkDecode(b *testing.B) {
	setup(&sAscii, benchEncoding.EncodeToString)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = benchEncoding.DecodeString(sAscii[i%nnn])
	}
}

/*
var acAscii [][]byte

func TestBDecodeACBase91(t *testing.T) {
	setup(&acAscii, acBase91.Encode)
	for i := 0; i < n; i++ {
		b := acBase91.Decode(acAscii[i])
		if bytes.Compare(b, bin[i]) != 0 {
			t.Errorf("#%d ascii: %v", i, string(acAscii[i]))
			t.Errorf("#%d want: %x", i, bin[i])
			t.Errorf("#%d got : %x", i, b)
			t.FailNow()
		}
	}
}

func BenchmarkEncodeACBase91(b *testing.B) {
	setupBin()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		acBase91.Encode(bin[i%n])
	}
}

func BenchmarkDecodeACBase91(b *testing.B) {
	setup(&dataValues, benchEncoding)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = acBase91.Decode(dataValues[i%n].bAscii)
	}
}
*/

var bproctorAscii [][]byte

func TestBDecodeBproctorBase91(t *testing.T) {
	setup(&bproctorAscii, bproctorBase91.Encode)
	for i := 0; i < nnn; i++ {
		b := bproctorBase91.Decode(bproctorAscii[i])
		if bytes.Compare(b, bin[i]) != 0 {
			t.Errorf("#%d ascii: %v", i, bproctorAscii[i])
			t.Errorf("#%d want: %x", i, bin[i])
			t.Errorf("#%d got : %x", i, b)
			t.FailNow()
		}
	}

}

func BenchmarkDecodeBproctorBase91(b *testing.B) {
	setup(&bproctorAscii, bproctorBase91.Encode)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = bproctorBase91.Decode(bproctorAscii[i%nnn])
	}
}

/*
var breezechenEncoding = breezechenBase91.StdEncoding // breezechenBase91.NewEncoding(benchChars)
var breezechenAscii [][]byte

func TestBDecodeBreezechenBase91(t *testing.T) {
	a := make([]byte, nn*2)
	setup(&breezechenAscii, func(b []byte) []byte {
		n := breezechenEncoding.Encode(a, b)
		var out []byte
		copy(out, a[:n])
		return out
	})

	b := make([]byte, nn*2)
	for i := 0; i < nnn; i++ {
		n, err := breezechenEncoding.Decode(b, breezechenAscii[i])
		if err != nil {
			t.Fatalf("#%d err=%v ascii=%v", i, err, string(breezechenAscii[i]))
		}
		if bytes.Compare(b[:n], bin[i]) != 0 {
			t.Errorf("#%d ascii: %v", i, breezechenAscii[i])
			t.Errorf("#%d want: %x", i, bin[i])
			t.Errorf("#%d got : %x", i, b[:n])
			t.FailNow()
		}
	}
}

func BenchmarkDecodeBreezechenBase91(b *testing.B) {
	setup(&dataValues, benchEncoding)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = breezechenEncoding.Decode(nil, dataValues[i%n].bAscii)
	}
}
*/

var equimAscii []string

func TestBDecodeEquimBase91(t *testing.T) {
	setup(&equimAscii, equimBase91.EncodeToString)
	for i := 0; i < nnn; i++ {
		b := equimBase91.DecodeString(equimAscii[i])
		if bytes.Compare(b, bin[i]) != 0 {
			t.Errorf("#%d ascii: %v", i, equimAscii[i])
			t.Errorf("#%d want: %x", i, bin[i])
			t.Errorf("#%d got : %x", i, b)
			t.FailNow()
		}
	}

}

func BenchmarkDecodeEquimBase91(b *testing.B) {
	setup(&equimAscii, equimBase91.EncodeToString)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		equimBase91.DecodeString(equimAscii[i%nnn])
	}
}

var majestrateAscii [][]byte

func TestBDecodeMajestrateBase91(t *testing.T) {
	setup(&majestrateAscii, majestrateBase91.Encode)
	for i := 0; i < nnn; i++ {
		b, err := majestrateBase91.Decode(majestrateAscii[i])
		if err != nil {
			t.Fatalf("#%d err=%v ascii=%v", i, err, majestrateAscii[i])
		}
		if bytes.Compare(b, bin[i]) != 0 {
			t.Errorf("#%d ascii: %v", i, majestrateAscii[i])
			t.Errorf("#%d want: %x", i, bin[i])
			t.Errorf("#%d got : %x", i, b)
			t.FailNow()
		}
	}
}

func BenchmarkDecodeMajestrateBase91(b *testing.B) {
	setup(&majestrateAscii, majestrateBase91.Encode)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = majestrateBase91.Decode(majestrateAscii[i%nnn])
	}
}

var mtraverEncoding = mtraverBase91.NewEncoding(benchChars)
var mtraverAscii []string

func TestBDecodeMtraverBase91Std(t *testing.T) {
	setup(&mtraverAscii, mtraverEncoding.EncodeToString)
	if len(mtraverAscii) == 0 {
		t.Fatal("len=0")
	}

	for i := 0; i < nnn; i++ {
		b, err := mtraverEncoding.DecodeString(mtraverAscii[i])
		if err != nil {
			t.Fatalf("#%d err=%v ascii=%v", i, err, mtraverAscii[i])
		}
		if bytes.Compare(b, bin[i]) != 0 {
			t.Errorf("#%d len=%d ascii: %v", i, len(mtraverAscii[i]), mtraverAscii[i])
			t.Errorf("#%d len=%d want: %x", i, len(bin[i]), bin[i])
			t.Errorf("#%d len=%d got : %x", i, len(b), b)
			t.FailNow()
		}
	}
}

func BenchmarkDecodeMtraverBase91(b *testing.B) {
	setup(&mtraverAscii, mtraverEncoding.EncodeToString)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = mtraverEncoding.DecodeString(mtraverAscii[i%nnn])
	}
}
