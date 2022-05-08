// #region <editor-fold desc="Preamble">
// Copyright (c) 2022 Teal.Finance contributors
//
// This file is part of BaseXX under the terms of the MIT License.
// SPDX-License-Identifier: MIT
//
// BaseXX is distributed WITHOUT ANY WARRANTY.
// See the LICENSE file alongside the source files
// or online at <https://opensource.org/licenses/MIT>.
// #endregion </editor-fold>

/*
Package base92 is a pretty good Base92 encoder
with customizable encoding alphabet.

The package provides a standard alphabet
complying with the constraints of the cookie tokens:
the standard alphabet uses 92 characters
from 0x20 (space) to 0x7E (~) except " ; and \.

You can also provide your own customized alphabet.

Base92 Usage

To encode any binary data to a Base92 string:

	import "github.com/teal-finance/BaseXX/base92"

	func foo() {
		bin := []byte{12, 23, 24, 45, 56, 67, 78, 89}
		str := base92.Encode(bin)
	}

To decode back the encoded Base92 string:

		bin, err := base92.Decode(str)

With custom alphabet

	var noSpace = base92.NewAlphabet(
		"!#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNO" +
		"PQRSTUVWXYZ[]^_`abcdefghijklmnopqrstuvwxyz{|}~")

	func bar(bin []byte) ([]byte, error) {
		str := base92.EncodeAlphabet(bin, noSpace)
		bin, err := base92.DecodeAlphabet(str, noSpace)
		return bin, err
	}
*/
package base92
