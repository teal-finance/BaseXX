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
package base62_test

import (
	"fmt"

	"github.com/teal-finance/BaseXX/base62"
)

// Encode any binary data to a Base62 string.
func ExampleEncoding_Encode() {
	bin := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255}

	str := base62.StdEncoding.EncodeToString(bin)

	fmt.Println("Binary input: ", bin)
	fmt.Println("Base62 string:", str)
	// Output:
	// Binary input:  [0 1 2 3 4 5 6 7 8 9 255]
	// Base62 string: 01TSm0PiyImxMV
}

// Decode back the encoded Base62 string.
func ExampleEncoding_DecodeString() {
	bin, err := base62.StdEncoding.DecodeString("01TSm0PiyImxMV")

	fmt.Println("Binary:", bin)
	fmt.Println("Error: ", err)
	// Output:
	// Binary: [0 1 2 3 4 5 6 7 8 9 255]
	// Error:  <nil>
}

// With custom alphabet.
func ExampleEncoding_EncodeToString() {
	myEncoding := base62.NewEncoding(
		"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	bin := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 254, 255}

	str := myEncoding.EncodeToString(bin)
	bin, err := myEncoding.DecodeString(str)

	fmt.Println("Binary:", bin)
	fmt.Println("Base62:", str)
	fmt.Println("Error: ", err)
	// Output:
	// Binary: [0 1 2 3 4 5 6 7 8 9 254 255]
	// Base62: 065EOdIdGZA96TZ
	// Error:  <nil>
}
