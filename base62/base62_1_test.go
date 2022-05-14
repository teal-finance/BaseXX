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

package base62

import (
	"encoding/hex"
	"math/rand"
	"testing"
	"time"
)

type testValues struct {
	dec []byte
	enc string
}

var tstEncoding = NewEncoding(btcDigits[:Base])

const n = 8192 // power of two to speed up the % modulo
var testPairs = make([]testValues, 0, n)

func init() {
	// If we do not seed the prng - it will default to a seed of (1)
	// https://golang.org/pkg/math/rand/#Seed
	rand.Seed(time.Now().UTC().UnixNano())
}

func initTestPairs() {
	if len(testPairs) > 0 {
		return
	}
	// pre-make the test pairs, so it doesn't take up benchmark time...
	for i := 0; i < n; i++ {
		data := make([]byte, 32)
		rand.Read(data)
		testPairs = append(testPairs, testValues{
			dec: data,
			enc: StdEncoding.EncodeToString(data),
		})
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
	t.Helper()
	for j := 1; j < 256; j++ {
		b := make([]byte, j)
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

func BenchmarkEncode(b *testing.B) {
	initTestPairs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		StdEncoding.Encode(testPairs[i%n].dec)
	}
}

func BenchmarkDecode(b *testing.B) {
	initTestPairs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := StdEncoding.DecodeString(testPairs[i%n].enc)
		if err != nil {
			b.Error(err)
		}
	}
}
