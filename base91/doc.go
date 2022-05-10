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

Encoding alphabet

This package base91 provides a standard alphabet
complying with the constraints of the cookie tokens:
contains the characters from 0x21 (!) to 0x7E (~)
except " ; and \. This is similar to the base92
standard alphabet, but without the space character (Ox20).
You can also provide your own customized alphabet.

Slower than codeberg.org/ac/base91

For a faster Base91 encoder, you may prefer:
<https://github.com/teal-finance/BaseXX/ac/base91>
a fork of <https://codeberg.org/ac/base91>.
The latter cannot be used because the module name
"catinello.eu/base91" is not reachable (tested in May 2022).
However this Base91 encoder does not allow to change
its specific different alphabet.

Comparison

Characters often used by common BaseXX encodings:

	Base91  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=!#$%&'()*,-.:<>?@[]^_`{|}~
	Base64  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=
	Base62  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	Base58   123456789ABCDEFGH JKLMN PQRSTUVWXYZabcdefghijk mnopqrstuvwxyz
	Hexa    0123456789ABCDEF

*/
package base91
