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

// Encode any binary data to a Base62 string
func ExampleEncode() {
	bin := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255}

	str := base62.Encode(bin)

	fmt.Println("Binary input: ", bin)
	fmt.Println("Base62 string:", str)
	// Output:
	// Binary input:  [0 1 2 3 4 5 6 7 8 9 255]
	// Base62 string: 01TSm0PiyImxMV
}

// Decode back the encoded Base62 string
func ExampleDecode() {
	bin, err := base62.Decode("01TSm0PiyImxMV")

	fmt.Println("Binary:", bin)
	fmt.Println("Error: ", err)
	// Output:
	// Binary: [0 1 2 3 4 5 6 7 8 9 255]
	// Error:  <nil>
}

// With custom alphabet

func ExampleEncodeAlphabet() {
	var myAlphabet = base62.NewAlphabet(
		"0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	bin := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 254, 255}

	str := base62.EncodeAlphabet(bin, myAlphabet)
	bin, err := base62.DecodeAlphabet(str, myAlphabet)

	fmt.Println("Binary:", bin)
	fmt.Println("Base62:", str)
	fmt.Println("Error: ", err)
	// Output:
	// Binary: [0 1 2 3 4 5 6 7 8 9 254 255]
	// Base62: 065EOdIdGZA96TZ
	// Error:  <nil>
}
