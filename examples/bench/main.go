// Copyright (c) 2022 Teal.Finance contributors
// This file is part of BaseXX (collection of Base58/Base62/Base91/Base92)
// licensed under the MIT License.
// SPDX-License-Identifier: MIT
// See the LICENSE file or <https://opensource.org/licenses/MIT>.

// Usage
//
//    go run github.com/teal-finance/BaseXX/examples/bench -h
//
// Bench all encoders
//
//    go run github.com/teal-finance/BaseXX/examples/bench -f path/file/to/encode
//
// Short sample
//
//    go run github.com/teal-finance/BaseXX/examples/bench -f <(echo 123456789-abcdefghijk)
//
// Bench in parallel
//
//    go run github.com/teal-finance/BaseXX/examples/bench -j -f <(echo 123456789-abcdefghijk)
//
// Detect race conditions
//
// Using -race flag slow down the test.
//
//    go run -race github.com/teal-finance/BaseXX/examples/bench -j -f <(echo 123456789-abcdefghijk)
//
// Increase number of loops
//
//    go run -race github.com/teal-finance/BaseXX/examples/bench -n 999000 -j -f <(echo 123456789-abcdefghijk)
//
package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	t91 "github.com/mtraver/base91"
	b91 "github.com/teal-finance/BaseXX/base91"
	b92 "github.com/teal-finance/BaseXX/base92"
	s92 "github.com/unix-world/smartgo/base92"

	"github.com/unix-world/smartgo"
	"golang.org/x/perf/benchmath"
)

const (
	T91 = iota
	B91
	B92
	S92
	N
)

func main() {
	// Command line arguments
	filename := flag.String("f", "README.md", "File to encode")
	loops := flag.Int("n", 0, "Number of encoding/decoding loops")
	parallel := flag.Bool("j", false, "Run the bench in separate go routines to detect race conditions")
	enableT91 := flag.Bool("traver", false, "Bench github.com/mtraver/base91")
	enableB91 := flag.Bool("91", false, "Bench BaseXX/base91")
	enableB92 := flag.Bool("92", false, "Bench BaseXX/base92")
	enableS92 := flag.Bool("smart", false, "Bench github.com/unix-world/smartgo/base92")
	flag.Parse()

	// Enable all by default
	if !*enableT91 && !*enableB91 && !*enableB92 && !*enableS92 {
		*enableT91 = true
		*enableB91 = true
		*enableB92 = true
		*enableS92 = true
	}

	txt, errMsg := smartgo.SafePathFileRead(*filename, true)
	if errMsg != "" {
		panic(errMsg)
	}

	bin := []byte(txt)
	log.Print("File: ", *filename, "\tSize: ", humanBytes(len(bin)))

	if *loops <= 0 {
		*loops = 1000_000 / len(bin)
		if *loops <= 0 {
			*loops = 1
		}
	}
	log.Print("Number of loops: ", *loops)

	names := [N]string{"mtraver/base91", "BaseXX/base91", "BaseXX/base92", "smartgo/base92"}
	enable := [N]bool{*enableT91, *enableB91, *enableB92, *enableS92}
	encodings := [N]encoding{t91.StdEncoding, b91.StdEncoding, b92.StdEncoding, nil}
	ns := [N][]float64{
		make([]float64, *loops),
		make([]float64, *loops),
		make([]float64, *loops),
		make([]float64, *loops),
	}

	var wg sync.WaitGroup

	for l := 0; l < *loops; l++ {
		for i := 0; i < N; i++ {
			if !enable[i] {
				continue
			}
			if *parallel {
				wg.Add(1)
				go func(ii, ll int) {
					ns[ii][ll] = bench(bin, encodings[ii])
					wg.Done()
				}(i, l)
			} else {
				ns[i][l] = bench(bin, encodings[i])
			}
		}
	}

	wg.Wait()

	assumption := benchmath.AssumeNormal
	for i := 0; i < N; i++ {
		if !enable[i] {
			continue
		}
		sample := benchmath.NewSample(ns[i], &benchmath.DefaultThresholds)
		summary := assumption.Summary(sample, 0.95)
		log.Printf("%15s%10s%5s",
			names[i],
			humanDuration(summary.Center),
			summary.PctRangeString())
	}
}

type encoding interface {
	DecodeString(s string) ([]byte, error)
	EncodeToString(src []byte) string
}

func bench(bin []byte, baseN encoding) float64 {
	start := time.Now()

	if baseN == nil {
		str := s92.Encode(bin)
		_, _ = s92.Decode(str)
	} else {
		str := baseN.EncodeToString(bin)
		_, _ = baseN.DecodeString(str)
	}

	return float64(time.Since(start).Nanoseconds())
}

// humanBytes converts a number of bytes into human readable units:
// KB (1024 bytes), MB (1024 KB), GB, TB, PB and EB.
func humanBytes(size int) string {
	const kilo = 1024
	if size < kilo {
		return fmt.Sprintf("%d B", size)
	}

	div, i := kilo, 0
	for n := size / kilo; n >= kilo; n /= kilo {
		div *= kilo
		i++
	}
	result := float64(size) / float64(div)

	// number of fractional digits
	n := 0
	if (result < 10) && (size%div > size/64) {
		n = 1
	}

	return fmt.Sprintf("%.*f %cB", n, result, "KMGTPE"[i])
}

// humanDuration converts a number of nanoseconds into human readable units:
// ns, us (microseconds), ms (milliseconds), s, m, h, d (day). y (year) and c (century).
func humanDuration(d float64) string {
	units := [...]float64{1000, 1000, 1000, 60, 60, 24, 365, 100, 999999999}

	if d < units[0] {
		return fmt.Sprintf("%.1f ns", d)
	}

	i := 0
	for d /= units[0]; d >= units[i+1]; d /= units[i] {
		i++
	}

	// number of fractional digits
	n := 0
	if d < 10 {
		n = 1
	}

	u := string("umsmhdyc"[i])
	if i < 2 {
		u += "s"
	}

	return fmt.Sprintf("%.*f %s", n, d, u)
}
