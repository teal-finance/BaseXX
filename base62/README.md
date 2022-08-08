# `BaseXX/base62`

Pretty good Base62 encoder with customizable encoding alphabet.

## Purpose

Base62 is designed to use only alphanumeric characters
in opposition of Base64 that requires
to use at least two punctuation characters.
See <https://wikiless.org/wiki/Base62>.

## Encoding Alphabet

The default encoding alphabet `StdEncoding`
uses the 62 digits and letters in lower and upper cases:

    0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz

You can also provide your own ASCII alphabet
(containing or not punctuation characters).

## Comparison with other BaseN

Characters often used by common BaseXX encodings:

    Alphanumeric  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
    Base62        0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
    Base64        0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=
    Base58         123456789ABCDEFGH JKLMN PQRSTUVWXYZabcdefghijk mnopqrstuvwxyz
    Hexadecimal   0123456789ABCDEF

## Can be much faster

Performance can be much much improved.
Using the tips of <https://github.com/mtraver/base91>
(original work from Joachim Henke),
this `BaseXX/base92` may become 200 times faster on the encoding,
and 30 times faster on the decoding.

## Contributions welcome

This Base92 needs your help to become faster.
Please propose your enhancements,
or even a further refactoring.
Any contribution is welcome. ;-)
