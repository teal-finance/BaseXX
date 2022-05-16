# `BaseXX/base92`

Pretty good Base92 encoder with customizable encoding alphabet.

## Copyright

Copyright (c) 2022 Teal.Finance contributors

This file is part of Teal.Finance/BaseXX licensed under the MIT License.
See the LICENSE file or <https://opensource.org/licenses/MIT>.
SPDX-License-Identifier: MIT

## Encoding alphabet

The default encoding alphabet `StdEncoding`
complies with the constraints of the cookie tokens:
`StdEncoding` uses 92 characters
from 0x20 (space) to 0x7E `~` except `"`, `;` and `\`.

You can also provide your own customized alphabet
with `NewEncoding()`.

## Other Base alphabets

In the following example, the Base92 encoding alphabet
is not the default one: the space has not been represented
=> the double-quote `"` is used to fill the 92nd ASCII character.

    Base95  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~ (and space)
    Base92  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!"#$%& ()*+,-./:;<=>?@[ ]^_`{|}~
    Base91  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!"#$%& ()*+, ./:;<=>?@[ ]^_`{|}~
    Ascii85 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstu    z!"#$%&'()*+,-./:;<=>?@[\]^_`
    Base64  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz          +   /
    Base62  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
    Base58   123456789ABCDEFGH JKLMN PQRSTUVWXYZabcdefghijk mnopqrstuvwxyz
    Hexa    0123456789ABCDEF

ASCII order:

    (space)!"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\]^_`abcdefghijklmnopqrstuvwxyz{|}~(del)

Similar Base92 project

Currently (May 2022), there are only two Base92 implementations in Go:
https://pkg.go.dev/search?q=base92

This `BaseXX/base92` has similar performance than
[`smartgo/base92`](https://pkg.go.dev/github.com/unix-world/smartgo/base92)
for large input data, but faster for short samples.

## Can be much faster

Performance can be much much improved.
Using the tips of https://github.com/mtraver/base91
(original work from Joachim Henke),
this BaseXX/base92 may become 200 times faster on the encoding,
and 30 times faster on the decoding.

## Contributions welcome

This Base92 needs your help to become faster.
Please propose your enhancements,
or even a further refactoring.
Any contribution is welcome. ;-)
