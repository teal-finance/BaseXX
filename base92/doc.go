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

Encoding alphabet

The default encoding alphabet `StdEncoding`
complies with the constraints of the cookie tokens:
`StdEncoding` uses 92 characters
from 0x20 (space) to 0x7E (~) except " ; and \.

You can also provide your own customized alphabet
with `NewEncoding()`.

Other BaseXX alphabets

Characters often used by common BaseXX encodings:

	Base92  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=*-_~.,?!@#$%&()[]{|}<>^:`'"
	Base91  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=*-_~.,?!@#$%&()[]{|}<>^:`'
	Base64  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=
	Base62  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	Base58   123456789ABCDEFGH JKLMN PQRSTUVWXYZabcdefghijk mnopqrstuvwxyz
	Hexa    0123456789ABCDEF

In the previous encodings example, the Base92 encoding alphabet
is not the default one: the space has not been represented
=> the double-quote (") is used to fill the 92nd ASCII character.

The fastest Base92

Currently (May 2022), there are only two Base92 implementations in Go:
https://pkg.go.dev/search?q=base92

This BaseXX/base92 is the fastest, five times faster than
https://pkg.go.dev/github.com/unix-world/smartgo/base92

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
package base92
