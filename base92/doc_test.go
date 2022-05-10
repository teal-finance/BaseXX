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
func ExampleEncode() {
	bin := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255}

	str := base92.Encode(bin)

	fmt.Println("Binary input: ", bin)
	fmt.Println("Base92 string:", str)
	// Output:
	// Binary input:  [0 1 2 3 4 5 6 7 8 9 255]
	// Base92 string: !2V2aO7r^-Kf
}

// Decode back the encoded Base92 string
func ExampleDecode() {
	bin, err := base92.Decode("!2V2aO7r^-Kf")

	fmt.Println("Binary:", bin)
	fmt.Println("Error: ", err)
	// Output:
	// Binary: [0 1 2 3 4 5 6 7 8 9 255]
	// Error:  <nil>
}

// With custom alphabet

func ExampleEncodeAlphabet() {
	var noSingleDoubleQuotes = base92.NewAlphabet("" +
		"abcdefghijklmnopqrstuvwxyz[]^_`!@#$%&()*+-<=> " +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789/{|}~.,:;?")

	bin := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 255}

	str := base92.EncodeAlphabet(bin, noSingleDoubleQuotes)
	bin, err := base92.DecodeAlphabet(str, noSingleDoubleQuotes)

	fmt.Println("Binary:", bin)
	fmt.Println("Base92:", str)
	fmt.Println("Error: ", err)
	// Output:
	// Binary: [0 1 2 3 4 5 6 7 8 9 255]
	// Base92: abrGrQ w7Nm-V
	// Error:  <nil>
}
