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
Package base91 is a pretty good Base91 encoder
with customizable encoding alphabet.

This package base91 provides a standard alphabet
complying with the constraints of the cookie tokens:
contains the characters from 0x21 (!) to 0x7E (~)
except " ; and \. This is similar to the base92
standard alphabet, but without the space character (Ox20).
You can also provide your own customized alphabet.

For a faster Base91 encoder, you may prefer:
<https://github.com/teal-finance/BaseXX/ac/base91>
a fork of <https://codeberg.org/ac/base91>.
The latter cannot be used because the module name
"catinello.eu/base91" is not reachable (tested in May 2022).
However this Base91 encoder does not allow to change
its specific different alphabet.

Base91 Usage

To encode any binary data to a Base91 string:

	import "github.com/teal-finance/BaseXX/base91"

	func foo() {
		bin := []byte{12, 23, 24, 45, 56, 67, 78, 89}
		str := base91.Encode(bin)
	}

To decode back the encoded Base91 string:

		bin, err := base91.Decode(str)

With custom alphabet

	var myAlphabet = base91.NewAlphabet(
		"PQRSTUVWXYZ[]^_`abcdefghijklmnopqrstuvwxyz{|}~" +
		"!#$%&'()*+,-./0123456789:<=>?@ABCDEFGHIJKLMNO")

	func bar(bin []byte) ([]byte, error) {
		str := base91.EncodeAlphabet(bin, myAlphabet)
		bin, err := base91.DecodeAlphabet(str, myAlphabet)
		return bin, err
	}
*/
package base91
