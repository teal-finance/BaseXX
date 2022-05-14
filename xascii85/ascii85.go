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

// package xascii85 provides convenient Encode()
// and Decode() functions on top of "encoding/ascii85".
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

// EncodeToString encodes binary bytes into Ascii85 bytes.
func (enc *Encoding) EncodeToString(bin []byte) string {
	return string(enc.Encode(bin))
}

// EncodeToString encodes binary bytes into a Ascii85 string
// allocating the destination buffer at the right size.
func (enc *Encoding) Encode(bin []byte) []byte {
	max := ascii85.MaxEncodedLen(len(bin))
	str := make([]byte, max)
	n := ascii85.Encode(str, bin)
	return str[:n]
}

// DecodeString decodes an Ascii85 string into a slice of bytes
// allocating the destination buffer at the right size.
func (enc *Encoding) DecodeString(str string) ([]byte, error) {
	return enc.Decode([]byte(str))
}

// Decode decodes Ascii85-encoded bytes into a slice of bytes
// allocating the destination buffer at the right size.
func (enc *Encoding) Decode(a85 []byte) ([]byte, error) {
	max := enc.DecodedLen(len(a85))
	bin := make([]byte, max)
	n, _, err := ascii85.Decode(bin, a85, true)
	return bin[:n], err
}

// EncodedLen returns the maximum length in bytes required to encode n bytes.
func (enc *Encoding) EncodedLen(n int) int { return ascii85.MaxEncodedLen(n) }

// DecodedLen returns the maximum length in bytes required to decode n Ascii85-encoded bytes.
// Ascii85 encodes 4 bytes 0x0000 by only one byte "z".
func (enc *Encoding) DecodedLen(n int) int { return 4 * n }
