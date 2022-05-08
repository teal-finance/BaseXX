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

// Package helper provides common functions for the packages
// at "github.com/teal-finance/garcon/basexx".
package helper

import (
	"log"
	"math"
)

// Alphabet is an optimized form of the encoding characters.
type Alphabet struct {
	Decode [128]int8
	Encode []byte
}

// NewAlphabet creates a new alphabet.
//
// It panics if the passed string is not 92 bytes long, isn't valid ASCII,
// or does not contain 92 distinct characters.
func NewAlphabet(s string, base int) *Alphabet {
	if len(s) != base {
		log.Panicf("Base%d: alphabet must be %d long, but got %d characters", base, base, len(s))
	}

	ret := new(Alphabet)
	ret.Encode = []byte(s)
	if len(ret.Encode) != base {
		log.Panicf("Base%d: alphabet must be %d bytes long, but got %d bytes", base, base, len(ret.Encode))
	}

	for i := range ret.Decode {
		ret.Decode[i] = -1
	}

	distinct := 0
	for i, b := range ret.Encode {
		if ret.Decode[b] == -1 {
			distinct++
		}
		ret.Decode[b] = int8(i)
	}

	if distinct != base {
		log.Panicf("Base%d: want %d distinct ASCII characters, "+
			"but provided alphabet has %d", base, base, distinct)
	}

	return ret
}

// PanicIfBadApproximation exits when a BaseXX is not well configured
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
