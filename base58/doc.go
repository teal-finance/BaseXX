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
Package base58 is a pretty good Base58 encoder
with customizable encoding alphabet.

Base58 aims to provide the largest possible alphabet, but
avoiding punctuation ("+" and "/") and ambiguous digits
and letters such as "0", "O" (capital) and "o" (lower-case).

The package provides two encoding alphabets:

- the alphabet used to represent Bitcoin addresses (the default one)
- the alphabet used by Flickr

Base58 Usage

To encode any binary data to a Base58 string:

	import "github.com/teal-finance/BaseXX/base58"

	func foo() {
		bin := []byte{12, 23, 24, 45, 56, 67, 78, 89}
		str := base58.Encode(bin)
	}

To decode back the encoded Base58 string:

		bin, err := base58.Decode(str)

With custom alphabet

	var myAlphabet = base58.NewAlphabet(
		"ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz123456789")

	func bar(bin []byte) ([]byte, error) {
		str := base58.EncodeAlphabet(bin, myAlphabet)
		bin, err := base58.DecodeAlphabet(str, myAlphabet)
		return bin, err
	}
*/
package base58
