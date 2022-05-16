// Copyright (c) 2022 Teal.Finance contributors
// This file is part of Teal.Finance/BaseXX licensed under the MIT License.
// SPDX-License-Identifier: MIT

// Package xascii85 implements the standard Encoding interface
// on top of "encoding/ascii85".
package xascii85

import (
	"encoding/ascii85"
)

// Encoding is just an empty type
// to provide the same inferface than "encoding/base64".
type Encoding struct{}

// StdEncoding is an empty value just
// to provide the same inferface than "encoding/base64".
var StdEncoding = Encoding{}

// NewEncoding creates a fake Encoding just
// to provide the same inferface than "encoding/base64".
func NewEncoding(_ string) *Encoding {
	return &Encoding{}
}

// Encode encodes binary bytes into Ascii85 bytes.
func (Encoding) Encode(dst, src []byte) (n int) {
	return ascii85.Encode(dst, src)
}

// Decode decodes Ascii85-encoded bytes into a slice of bytes.
func (Encoding) Decode(dst, src []byte) (n int, err error) {
	n, _, err = ascii85.Decode(dst, src, true)
	return n, err
}

// EncodeToString encodes binary bytes into an Ascii85 string
// allocating the destination buffer at the right size.
func (Encoding) EncodeToString(src []byte) string {
	max := ascii85.MaxEncodedLen(len(src))
	dst := make([]byte, max)
	n := ascii85.Encode(dst, src)
	return string(dst[:n])
}

// DecodeString decodes an Ascii85 string into a slice of bytes
// allocating the destination buffer at the right size.
func (enc Encoding) DecodeString(s string) ([]byte, error) {
	src := []byte(s)
	max := enc.DecodedLen(len(src))
	dst := make([]byte, max)
	n, _, err := ascii85.Decode(dst, src, true)
	return dst[:n], err
}

// EncodedLen returns the maximum length in bytes required to encode n bytes.
func (Encoding) EncodedLen(n int) int { return ascii85.MaxEncodedLen(n) }

// DecodedLen returns the maximum length in bytes
// required to decode n Ascii85-encoded bytes.
// Ascii85 decodes 4 bytes 0x0000 from only one byte "z".
func (Encoding) DecodedLen(n int) int { return 4 * n }
