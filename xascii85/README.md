# `xascii85`

Standard Encoding interface on top of `"encoding/ascii85"`.

## Copyright

Copyright (c) 2022 Teal.Finance contributors

This file is part of Teal.Finance/BaseXX licensed under the MIT License.
See the LICENSE file or <https://opensource.org/licenses/MIT>.
SPDX-License-Identifier: MIT

## Ascii85 advantages

The main idea is to encode by chunk of 4 bytes, instead of 3 bytes for Base64.

There are 95 printable ASCII characters including the space.
To represent 4 bytes, 5 printable ASCII characters are required:

     95⁵ = 7 737 809 375   <-- Minimum 5 printable ASCII characters
    256⁴ = 4 294 967 296
     95⁴ =    81 450 625

The minimum set is 85 characters:

     86⁵ = 4 704 270 176
     85⁵ = 4 437 053 125   <-- Minimum 85 different ASCII characters
    256⁴ = 4 294 967 296
     84⁵ = 4 182 119 424
     83⁵ = 3 939 040 643

Therefore, 85 is the minimal number of different characters,
to encode any sequence of 4 bytes as 5 printable ASCII characters.

## Interface

The idea is to provide the same interface as "encoding/base64".
See <https://pkg.go.dev/encoding/base64>

```go
func NewEncoding(encoder string) *Encoding

interface Encoding {
    Decode(dst, src []byte) (n int, err error)
    Encode(dst, src []byte) (n int)
    // Here Encode() returns the number of written bytes.
    // This is different with encoding/base64.
    // Ascii85 encoded length cannot be known from just
    // the number of bytes to encode, whereas it can with Base64.
    
    DecodeString(s string) ([]byte, error)
    EncodeToString(src []byte) string
    
    DecodedLen(n int) int // Returns the Max.
    EncodedLen(n int) int // Returns the Max.
    
    // Not implemented.
    // Strict() *Encoding
    // WithPadding(padding rune) *Encoding
}
```

## Definition in PostScript documentation

Asci85 encodes binary data in an ASCII base-85 representation.
This encoding uses nearly all of the printable ASCII character set.
The resulting expansion factor is 4:5, making this encoding
much more efficient than hexadecimal.

## Encoding alphabet

ASCII characters from 0x21 `!` through 0x75 `u`.

## Comparison to other encodings

    Base95  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~ (and space)
    Base94  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~
    Base92  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!"#$%& ()*+,-./:;<=>?@[ ]^_`{|}~
    Base91  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!"#$%& ()*+, ./:;<=>?@[ ]^_`{|}~
    Ascii85 0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstu    z!"#$%&'()*+,-./:;<=>?@[\]^_`
    Z85     0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz! #$%& ()*+ -./: <=>?@[ ]^  { }
    Base70  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz          + -./           _    ~
    Base64  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz          +   /
    Base62  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
    Base58   123456789ABCDEFGH JKLMN PQRSTUVWXYZabcdefghijk mnopqrstuvwxyz
    Hexa    0123456789ABCDEF

## Specification in PostScript documentation

The `ASCII85Encode` filter encodes binary data in the ASCII base-85 encoding.
Generally, for every 4 bytes of binary data, it produces 5 ASCII printing
characters in the range `!` through `u`. It inserts a newline in the encoded
output at least once every 80 characters, thereby limiting the lengths of lines.

When the `ASCII85Encode` filter is closed, it writes the 2-character sequence `~>`
as an EOD marker.

Binary data bytes are encoded in 4-tuples (groups of 4). Each 4-tuple is
used to produce a 5-tuple of ASCII characters. If the binary 4-tuple is
(b1 b2 b3 b4) and the encoded 5-tuple is (c1 c2 c3 c4 c5), then the relation
between them is

(b1 × 256³) + (b2 × 256²) + (b3 × 256¹) + b4 =
(c1 × 85⁴) + (c2 × 85³) + (c3 × 85²) + (c4 × 85¹) + c5

In other words, 4 bytes of binary data are interpreted as a base-256 number
and then converted into a base-85 number. The five “digits” of this number,
(c1 c2 c3 c4 c5), are then converted into ASCII characters by adding 33,
which is the ASCII code for `!`, to each. ASCII characters in the range `!` to `u`
are used, where `!` represents the value 0 and u represents the value 84.

As a special case, if all five digits are 0, they are represented by a
single character `z` instead of by `!!!!!`.
