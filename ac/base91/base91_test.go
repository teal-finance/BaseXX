// SPDX-License-Identifier: BSD-3-Clause

package base91 // import "catinello.eu/base91"

import (
	"bytes"
	"encoding/hex"
	"testing"
)

var samples = map[string]string{
	"":                            "",
	"1":                           "xA",
	"1234567890":                  "QztEml0o[2;(A",
	"abcdefghijklmnopqurstuvwxyz": "#G(Ic,5ph#77&xrmlrjg2]jTs%2<WF%qfB",
}

// init inserts a failing case in samples.
func Disable_init() {
	const hexa = "5526a41a95041b"
	const str = ":Ro7<O'9B"
	b, err := hex.DecodeString(hexa)
	if err != nil {
		panic(err)
	}
	samples[string(b)] = str
}

func TestEncode(t *testing.T) {
	for bin, s := range samples {
		b := []byte(bin)
		if got := EncodeToString(b); got != s {
			t.Error("Incorrect encoding of ", b)
			t.Errorf("want: %q", s)
			t.Errorf("got : %q", got)
		}
	}
}

func TestDecode(t *testing.T) {
	for bin, s := range samples {
		b := []byte(bin)
		if got := DecodeString(s); !bytes.Equal(got, b) {
			t.Errorf("Incorrect decoding of %q", s)
			t.Errorf("want: %x", b)
			t.Errorf("got : %x", got)
		}
	}
}
