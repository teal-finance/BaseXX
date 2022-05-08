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

package base58

import (
	"fmt"

	"github.com/teal-finance/BaseXX/helper"
)

const (
	Base = 58
	// approximation of ceil(log(256)/log(base)).
	numerator   = 11
	denominator = 8 // power of two -> speed up EncodeAlphabet()
)

// StdAlphabet is the default encoding alphabet, same as BTCAlphabet.
var StdAlphabet = BTCAlphabet

// BTCAlphabet is the Bitcoin Base58 alphabet.
var BTCAlphabet = NewAlphabet("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// FlickrAlphabet is the Flickr Base58 alphabet.
var FlickrAlphabet = NewAlphabet("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")

func init() {
	helper.PanicIfBadApproximation(Base, numerator, denominator)
}

func NewAlphabet(s string) *helper.Alphabet {
	return helper.NewAlphabet(s, Base)
}

// Encode encodes binary bytes into a Base58 string
// using the default alphabet.
func Encode(bin []byte) string {
	return EncodeAlphabet(bin, StdAlphabet)
}

// EncodeAlphabet encodes binary bytes into a Base58 string
// using the given alphabet.
func EncodeAlphabet(bin []byte, alphabet *helper.Alphabet) string {
	size := len(bin)

	zcount := 0
	for zcount < size && bin[zcount] == 0 {
		zcount++
	}

	// It is crucial to make this as short as possible, especially for
	// the usual case of bitcoin addrs
	size = zcount +
		// This is an integer simplification of
		// ceil(log(256)/log(base))
		(size-zcount)*numerator/denominator + 1

	out := make([]byte, size)

	var i, high int
	var carry uint32

	high = size - 1
	for _, b := range bin {
		i = size - 1
		for carry = uint32(b); i > high || carry != 0; i-- {
			carry += 256 * uint32(out[i])
			out[i] = byte(carry % uint32(Base))
			carry /= uint32(Base)
		}
		high = i
	}

	// Determine the additional "zero-gap" in the buffer (aside from zcount)
	for i = zcount; i < size && out[i] == 0; i++ {
	}

	// Now encode the values with actual alphabet in-place
	val := out[i-zcount:]
	size = len(val)
	for i = 0; i < size; i++ {
		out[i] = alphabet.Encode[val[i]]
	}

	return string(out[:size])
}

// Decode decodes a Base58 string into binary bytes
// using the default alphabet.
func Decode(str string) ([]byte, error) {
	return DecodeAlphabet(str, StdAlphabet)
}

// DecodeAlphabet decodes a Base58 string into binary bytes
// using the given alphabet.
func DecodeAlphabet(str string, alphabet *helper.Alphabet) ([]byte, error) {
	if len(str) == 0 {
		return nil, nil
	}

	zero := alphabet.Encode[0]
	strLen := len(str)

	var zcount int
	for i := 0; i < strLen && str[i] == zero; i++ {
		zcount++
	}

	var t, c uint64

	// the 32bit algo stretches the result up to 2 times
	binu := make([]byte, 2*((strLen*denominator/numerator)+1))
	outi := make([]uint32, (strLen+3)/4)

	for _, r := range str {
		if r > 127 {
			return nil, fmt.Errorf("Base%d: high-bit set on invalid digit", Base)
		}
		if alphabet.Decode[r] == -1 {
			return nil, fmt.Errorf("Base%d: invalid digit %q", Base, r)
		}

		c = uint64(alphabet.Decode[r])

		for j := len(outi) - 1; j >= 0; j-- {
			t = uint64(outi[j])*uint64(Base) + c
			c = t >> 32
			outi[j] = uint32(t & 0xffffffff)
		}
	}

	// initial mask depends on b92sz, on further loops it always starts at 24 bits
	mask := (uint(strLen%4) * 8)
	if mask == 0 {
		mask = 32
	}
	mask -= 8

	outLen := 0
	for j := 0; j < len(outi); j++ {
		for mask < 32 { // loop relies on uint overflow
			binu[outLen] = byte(outi[j] >> mask)
			mask -= 8
			outLen++
		}
		mask = 24
	}

	// find the most significant byte post-decode, if any
	for msb := zcount; msb < len(binu); msb++ {
		if binu[msb] > 0 {
			return binu[msb-zcount : outLen], nil
		}
	}

	// it's all zeroes
	return binu[:outLen], nil
}
