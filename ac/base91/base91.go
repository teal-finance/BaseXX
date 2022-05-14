// Copyright (c) 2019-2021 Antonino Catinello
// SPDX-License-Identifier: BSD-3-Clause

// Based on http://base91.sourceforge.net/
package base91 // import "catinello.eu/base91"

// Encoding table holds all the characters for base91 encoding - slice is faster than an array.
var enctab = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~'")

// Decoding table maps all the characters back to their integer values - array is faster than a map
// This array represents all 91 characters with values below 91.
var dectab = [...]byte{
	91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
	91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
	91, 62, 90, 63, 64, 65, 66, 91, 67, 68, 69, 70, 71, 91, 72, 73,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 74, 75, 76, 77, 78, 79,
	80, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 81, 91, 82, 83, 84,
	85, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
	41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 86, 87, 88, 89, 91,
	91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
	91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
	91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
	91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
	91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
	91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
	91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
	91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91, 91,
}

// EncodeToString encodes the given byte array and returns a string.
func EncodeToString(d []byte) string {
	return string(Encode(d))
}

// Encode returns the base91 encoded string.
func Encode(d []byte) []byte {
	var n, b uint
	var o []byte

	for i := 0; i < len(d); i++ {
		b |= uint(d[i]) << n
		n += 8

		if n > 13 {
			v := b & 8191

			if v > 88 {
				b >>= 13
				n -= 13
			} else {
				v = b & 16383
				b >>= 14
				n -= 14
			}

			o = append(o, enctab[v%91], enctab[v/91])
		}
	}

	if n > 0 {
		o = append(o, enctab[b%91])

		if n > 7 || b > 90 {
			o = append(o, enctab[b/91])
		}
	}

	return o
}

// DecodeString decodes a given byte array are returns a string.
func DecodeString(d string) []byte {
	return Decode([]byte(d))
}

// Decode decodes a base91 encoded string and returns the result.
func Decode(d []byte) []byte {
	var b, n uint
	var o []byte
	v := -1

	for i := 0; i < len(d); i++ {
		c := dectab[d[i]]
		if c > 90 {
			continue
		}

		if v < 0 {
			v = int(c)
			continue
		}

		v += int(c) * 91
		b |= uint(v) << n

		if v&8191 > 88 {
			n += 13
		} else {
			n += 14
		}

		for {
			o = append(o, byte(b&255))
			b >>= 8
			n -= 8

			if n <= 7 {
				break
			}
		}

		v = -1
	}

	if v > -1 {
		o = append(o, byte((b|uint(v)<<n)&255))
	}

	return o
}
