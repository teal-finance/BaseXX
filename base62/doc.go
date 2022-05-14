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

Purpose

Base62 is designed to use only alphanumeric characters
in opposition of Base64 that requires
to use some punctuation characters.
See <https://wikiless.org/wiki/Base62>.

Encoding Alphabet

The default encoding alphabet `StdEncoding`
uses the 62 digits and letters in lower and upper cases:

	0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz

You can also provide your own ASCII alphabet
containing or not punctuation characters.

Comparison

Characters often used by common BaseXX encodings:

	Alphanumeric  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	Base62        0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	Base64        0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=
	Base58         123456789ABCDEFGH JKLMN PQRSTUVWXYZabcdefghijk mnopqrstuvwxyz
	Hexadecimal   0123456789ABCDEF

Can be much faster

Performance can be much much improved.
Using the tips of https://github.com/mtraver/base91
(original work from Joachim Henke),
this BaseXX/base92 may become 200 times faster on the encoding,
and 30 times faster on the decoding.

Contributions welcome

This Base92 needs your help to become faster.
Please propose your enhancements,
or even a further refactoring.
Any contribution is welcome. ;-)
*/
package base62
