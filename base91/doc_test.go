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
package base91_test

import (
	"fmt"

	"github.com/teal-finance/BaseXX/base91"
)

// Encode any binary data to a Base91 string
func ExampleEncode() {
	bin := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255}

	str := base91.Encode(bin)

	fmt.Println("Binary input: ", bin)
	fmt.Println("Base91 string:", str)
	// Output:
	// Binary input:  [0 1 2 3 4 5 6 7 8 9 255]
	// Base91 string: !#B6*yOw]cPi5
}

// Decode back the encoded Base91 string
func ExampleDecode() {
	bin, err := base91.Decode("!#B6*yOw]cPi5")

	fmt.Println("Binary:", bin)
	fmt.Println("Error: ", err)
	// Output:
	// Binary: [0 1 2 3 4 5 6 7 8 9 255]
	// Error:  <nil>
}

// With custom alphabet

func ExampleEncodeAlphabet() {
	var noSingleDoubleQuotes = base91.NewAlphabet("" +
		"abcdefghijklmnopqrstuvwxyz[]^_`!@#$%&()*+-<=>" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789/{|}~.,:;?")

	bin := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255}

	str := base91.EncodeAlphabet(bin, noSingleDoubleQuotes)
	bin, err := base91.DecodeAlphabet(str, noSingleDoubleQuotes)

	fmt.Println("Binary:", bin)
	fmt.Println("Base91:", str)
	fmt.Println("Error: ", err)
	// Output:
	// Binary: [0 1 2 3 4 5 6 7 8 9 255]
	// Base91: ab!ui~>|MSAYt
	// Error:  <nil>
}
