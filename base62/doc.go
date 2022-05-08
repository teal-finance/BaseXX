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
Package base62 is a pretty good Base62 encoder
with customizable encoding alphabet.

Base62 is designed to use only alphanumeric characters
in opposition of Base64 that requires
to use some punctuation characters.
See <https://wikiless.org/wiki/Base62>.

This package base62 provides the standard alphabet
using the 62 digits and letters in lower and upper case:

	0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz

You can also provide your own alphabet
containing or not the punctuation characters.

Base62 Usage

To encode any binary data to a Base62 string:

	import "github.com/teal-finance/BaseXX/base62"

	func foo() {
		bin := []byte{12, 23, 24, 45, 56, 67, 78, 89}
		str := base62.Encode(bin)
	}

To decode back the encoded Base62 string:

		bin, err := base62.Decode(str)

With custom alphabet

	var myAlphabet = base62.NewAlphabet(
		"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	func bar(bin []byte) ([]byte, error) {
		str := base62.EncodeAlphabet(bin, myAlphabet)
		bin, err := base62.DecodeAlphabet(str, myAlphabet)
		return bin, err
	}
*/
package base62
