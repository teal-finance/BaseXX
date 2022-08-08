// Copyright (c) 2017-2020 Denis Subbotin, Philip Schlump,
//                         Nika Jones, Steven Allen, MoonFruit
// Copyright (c) 2022      Teal.Finance contributors
//
// This file is a modified copy from https://github.com/mr-tron/base58
// The source code has been adapted to support other bases.
// This file is now part of BaseXX under the terms of the MIT License.
// SPDX-License-Identifier: MIT
// See the LICENSE file or https://opensource.org/licenses/MIT

// Package encoding provides common functions for all BaseXX implementations.
package encoding

import (
	"log"
	"math"
)

// Encoding alphabet is an optimized form of the encoding characters.
type Encoding struct {
	EncChars []byte
	DecMap   [128]int8
}

// NewEncoding creates a new alphabet mapping.
//
// It panics if the passed string does not meet all requirements:
// its length (in bytes) must be the same as the base,
// all runes must be valid ASCII characters,
// and all characters must be different.
// Encoder string with non-printable characters are accepted.
func NewEncoding(encoder string, base int) *Encoding {
	if len(encoder) != base {
		log.Panicf("Base%d: alphabet must be %d long, but got %d characters", base, base, len(encoder))
	}

	ret := new(Encoding)
	ret.EncChars = []byte(encoder)
	if len(ret.EncChars) != base {
		log.Panicf("Base%d: alphabet must be %d bytes long, but got %d bytes", base, base, len(ret.EncChars))
	}

	for i := range ret.DecMap {
		ret.DecMap[i] = -1
	}

	distinct := 0
	for i, b := range ret.EncChars {
		if ret.DecMap[b] == -1 {
			distinct++
		}
		ret.DecMap[b] = int8(i)
	}

	if distinct != base {
		log.Panicf("Base%d: want %d distinct ASCII characters, "+
			"but provided alphabet has %d", base, base, distinct)
	}

	return ret
}

// PanicIfBadApproximation exits when a BaseXX is not well configured.
func PanicIfBadApproximation(base, a, b int) {
	want := math.Log(256) / math.Log(float64(base))
	got := float64(a) / float64(b)

	epsilon := math.Abs(want - got)
	if epsilon > 0.1 {
		log.Panicf("\n\nBase%d: want=%.4g but got=%.4g from fraction %d/%d. \n"+
			"Please find better numerator and denominator. \n\n"+
			"You may use: \n"+
			"go run github.com/teal-finance/diophantine/cmd@latest -list %.22g\n\n",
			base, want, got, a, b, want)
	}
}
