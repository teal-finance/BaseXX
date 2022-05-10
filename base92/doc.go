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

The package provides a standard alphabet
complying with the constraints of the cookie tokens:
the standard alphabet uses 92 characters
from 0x20 (space) to 0x7E (~) except " ; and \.

You can also provide your own customized alphabet.

Comparison

Characters often used by common BaseXX encodings:

	Base92  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=*-_~.,?!@#$%&()[]{|}<>^:`'"
	Base91  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=*-_~.,?!@#$%&()[]{|}<>^:`'
	Base64  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=
	Base62  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	Base58   123456789ABCDEFGH JKLMN PQRSTUVWXYZabcdefghijk mnopqrstuvwxyz
	Hexa    0123456789ABCDEF

*/
package base92
