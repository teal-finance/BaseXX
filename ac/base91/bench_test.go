// SPDX-License-Identifier: BSD-3-Clause

package base91 // import "catinello.eu/base91"

import (
	"crypto/rand"
	"io"
	"testing"
)

func BenchmarkEncode(b *testing.B) {
	s := make([]byte, 1024*1024)
	if _, err := io.ReadFull(rand.Reader, s); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Encode(s)
	}
}

func BenchmarkDecode(b *testing.B) {
	s := make([]byte, 1024*1024)
	if _, err := rand.Read(s); err != nil {
		b.Fatal(err)
	}

	encoded := Encode(s)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Decode(encoded)
	}
}
