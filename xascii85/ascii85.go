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
	"fmt"
)

// Encode encodes a slice of bytes into an Ascii85 string
// allocating the destination buffer at the right size.
func Encode(bin []byte) string {
	max := ascii85.MaxEncodedLen(len(bin))
	str := make([]byte, max)
	n := ascii85.Encode(str, bin)
	return string(str[:n])
}

// Decode decodes an Ascii85 string into a slice of bytes
// allocating the destination buffer at the right size.
func Decode(str string) ([]byte, error) {
	max := MaxDecodedLen(len(str))
	bin := make([]byte, max)

	n, _, err := ascii85.Decode(bin, []byte(str), true)
	if err != nil {
		return nil, fmt.Errorf("ascii85.Decode %w", err)
	}

	return bin[:n], nil
}

// MaxDecodedSize returns the maximum length required to decode n binary bytes.
// Ascii85 encodes 4 bytes 0x0000 by only one byte "z".
func MaxDecodedLen(n int) int { return 4 * n }
