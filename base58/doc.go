// #region <editor-fold desc="Preamble">
// Copyright (c) 2017 Alexey Kamenskiy @akamensky
// Copyright (c) 2022 Teal.Finance contributors
//
// This file has been inspired from the README
// of <https://github.com/akamensky/base58>
// under the terms of the Unlicense License.
// SPDX-License-Identifier: Unlicense
// #endregion </editor-fold>

/*
Package base58 is a pretty good Base58 encoder
with customizable encoding alphabet.

Purpose

There are many methods to encode/decode
binary data into printable format.
The most common onces are HEX and Base64.

While those are good approaches in some situations,
each of them has own limitations:
HEX doubles the memory footprint and
Base64 is hard to understand/read.
They still have a place to be used when
storage and readability are not of concern.

Base58 encoding serves double purpose:

1. Long data presented in short format (compression of sorts)

2. Human readable by removing ambiguous characters

To meet this requirements, Base58 avoiding punctuation ("+" and "/") and
ambiguous digits and letters such as "0", "O" (capital) and "o" (lower-case),

	Alphanumeric  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	Base64        0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=
	Base58         123456789ABCDEFGH JKLMN PQRSTUVWXYZabcdefghijk mnopqrstuvwxyz
	Hexadecimal   0123456789ABCDEF

Provided alphabets

This BaseXX/base58 package provides two encoding alphabets:

1. the one used to represent Bitcoin addresses (the default one)

2. the one used by Flickr (different lower/upper case order)

	BTC    123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz
    Flickr 123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ

Comparison

Characters often used by common BaseXX encodings:

	Hexa    0123456789ABCDEF
	Base58   123456789ABCDEFGH JKLMN PQRSTUVWXYZabcdefghijk mnopqrstuvwxyz
	Base62  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
	Base64  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=
	Base91  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=*-_~.,?!@#$%&()[]{|}<>^:`'
	Base92  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=*-_~.,?!@#$%&()[]{|}<>^:`'"

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
package base58
