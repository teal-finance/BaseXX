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
package base58_test

import (
	"fmt"

	"github.com/teal-finance/BaseXX/base58"
)

// Encode any binary data to a Base58 string
func ExampleEncoding_Encode() {
	bin := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 254, 255}

	str := base58.StdEncoding.EncodeToString(bin)

	fmt.Println("Binary input: ", bin)
	fmt.Println("Base58 string:", str)
	// Output:
	// Binary input:  [0 1 2 3 4 5 6 7 8 9 254 255]
	// Base58 string: 1FVk6iLh9oT6ivJ
}

// Decode back the encoded Base58 string
func ExampleEncoding_DecodeString() {
	bin, err := base58.StdEncoding.DecodeString("1FVk6iLh9oT6ivJ")

	fmt.Println("Binary:", bin)
	fmt.Println("Error: ", err)
	// Output:
	// Binary: [0 1 2 3 4 5 6 7 8 9 254 255]
	// Error:  <nil>
}

// With custom alphabet
func ExampleEncoding_EncodeToString() {
	var myEncoding = base58.NewEncoding(
		"ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz123456789")

	bin := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 254, 255}

	str := myEncoding.EncodeToString(bin)
	bin, err := myEncoding.DecodeString(str)

	fmt.Println("Binary:", bin)
	fmt.Println("Base58:", str)
	fmt.Println("Error: ", err)
	// Output:
	// Binary: [0 1 2 3 4 5 6 7 8 9 254 255]
	// Base58: AQeuFsVrJxcFs5T
	// Error:  <nil>
}
