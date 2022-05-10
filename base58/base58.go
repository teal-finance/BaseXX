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

	"github.com/teal-finance/BaseXX/encoding"
)

const (
	Base = 58
	// approximation of ceil(log(256)/log(base)).
	numerator   = 11
	denominator = 8 // power of two -> speed up EncodeEncoding()
)

// StdEncoding is the default encoding alphabet, same as BTCEncoding.
var StdEncoding = BTCEncoding

// BTCEncoding is the Bitcoin Base58 enc.
var BTCEncoding = NewEncoding("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// FlickrEncoding is the Flickr Base58 enc.
var FlickrEncoding = NewEncoding("123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")

func init() {
	encoding.PanicIfBadApproximation(Base, numerator, denominator)
}

type Encoding encoding.Encoding

func NewEncoding(encoder string) *Encoding {
	e := encoding.NewEncoding(encoder, Base)
	return (*Encoding)(e)
}

// EncodeToString encodes binary bytes into Base58 bytes.
func (enc *Encoding) EncodeToString(bin []byte) string {
	return string(enc.Encode(bin))
}

// EncodeToString encodes binary bytes into a Base58 string.
func (enc *Encoding) Encode(bin []byte) []byte {
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
		out[i] = enc.EncChars[val[i]]
	}

	return out[:size]
}

// DecodeString decodes a Base58 string into binary bytes.
func (enc *Encoding) DecodeString(str string) ([]byte, error) {
	if len(str) == 0 {
		return nil, nil
	}

	zero := enc.EncChars[0]
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
		if enc.DecMap[r] == -1 {
			return nil, fmt.Errorf("Base%d: invalid digit %q", Base, r)
		}

		c = uint64(enc.DecMap[r])

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
