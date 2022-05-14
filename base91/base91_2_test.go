// #region <editor-fold desc="Preamble">
// Copyright (c) 2017-2020 Denis Subbotin, Philip Schlump,
//                         Nika Jones, Steven Allen, MoonFruit
// Copyright (c) 2022      Teal.Finance contributors
//
// This file is a modified copy from https://github.com/mr-tron/base58
// The source code has been adapted to support other bases.
// This file is now part of BaseXX under the terms of the MIT License.
// SPDX-License-Identifier: MIT
//
// BaseXX is distributed WITHOUT ANY WARRANTY.
// See the LICENSE file alongside the source files
// or online at <https://opensource.org/licenses/MIT>.
// #endregion </editor-fold>

package base91

import (
	"testing"
)

func TestBase58_test2(t *testing.T) {
	testAddr := []string{
		"1QCaxc8hutpdZ62iKZsn1TCG3nh7uPZojq",
		"1DhRmSGnhPjUaVPAj48zgPV9e2oRhAQFUb",
		"17LN2oPYRYsXS9TdYdXCCDvF2FegshLDU2",
		"14h2bDLZSuvRFhUL45VjPHJcW667mmRAAn",
	}

	for ii, vv := range testAddr {
		// num := Base58Decode([]byte(vv))
		// chk := Base58Encode(num)
		num, err := StdEncoding.DecodeString(vv)
		if err != nil {
			t.Errorf("Test %d, expected success, got error %s\n", ii, err)
		}
		chk := StdEncoding.EncodeToString(num)
		if vv != chk {
			t.Errorf("Test %d, expected=%s got=%s Address did base58 encode/decode correctly.", ii, vv, chk)
		}
	}
}
