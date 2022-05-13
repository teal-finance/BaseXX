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
package base92_test

import (
	"fmt"

	"github.com/teal-finance/BaseXX/base92"
)

// Encode any binary data to a Base92 string
func ExampleEncoding_Encode() {
	bin := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255}

	str := base92.StdEncoding.EncodeToString(bin)

	fmt.Println("Binary input: ", bin)
	fmt.Printf("Base92 string: %q\n", str)
	// Output:
	// Binary input:  [0 1 2 3 4 5 6 7 8 9 255]
	// Base92 string: " !2V2aO7r^-Kf"
}

// Decode back the encoded Base92 string
func ExampleEncoding_DecodeString() {
	bin, err := base92.StdEncoding.DecodeString(" !2V2aO7r^-Kf")

	fmt.Println("Binary:", bin)
	fmt.Println("Error: ", err)
	// Output:
	// Binary: [0 1 2 3 4 5 6 7 8 9 255]
	// Error:  <nil>
}

// With custom alphabet
func ExampleEncoding_EncodeToString() {
	var noSingleDoubleQuotes = base92.NewEncoding("" +
		"abcdefghijklmnopqrstuvwxyz[]^_`!@#$%&()*+-<=> " +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789/{|}~.,:;?")

	bin := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255}

	str := noSingleDoubleQuotes.EncodeToString(bin)
	bin, err := noSingleDoubleQuotes.DecodeString(str)

	fmt.Println("Binary:", bin)
	fmt.Println("Base92:", str)
	fmt.Println("Error: ", err)
	// Output:
	// Binary: [0 1 2 3 4 5 6 7 8 9 255]
	// Base92: abrGrQ w7Nm-V
	// Error:  <nil>
}
