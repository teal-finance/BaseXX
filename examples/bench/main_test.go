// Copyright (c) 2022 Teal.Finance contributors
// This file is part of Teal.Finance/BaseXX licensed under the MIT License.
// SPDX-License-Identifier: MIT

package main

import (
	"strconv"
	"testing"
)

func Test_humanBytes(t *testing.T) {
	cases := []struct {
		want string
		size int
	}{
		{"954 MB", 1000 * 1000 * 1000},
		{"1023 MB", 1023 * 1024 * 1024},
		{"1 GB", 1024 * 1024 * 1024},
		{"1 GB", 1025 * 1024 * 1024},
		{"1 GB", 1030 * 1024 * 1024},
		{"1 GB", 1035 * 1024 * 1024},
		{"1 GB", 1040 * 1024 * 1024},
		{"1.0 GB", 1045 * 1024 * 1024},
		{"1.0 GB", 1050 * 1024 * 1024},
		{"1.1 GB", 1090 * 1024 * 1024},
		{"1.1 GB", 1100 * 1024 * 1024},
		{"9.3 GB", 10 * 1000 * 1000 * 1000},
		{"9.8 GB", 10 * 1000 * 1024 * 1024},
		{"10 GB", 10 * 1024 * 1024 * 1024},
		{"98 GB", 100 * 1000 * 1024 * 1024},
		{"977 GB", 1000 * 1000 * 1024 * 1024},
	}
	for i, c := range cases {
		name := c.want + " #" + strconv.Itoa(i)
		t.Run(name, func(t *testing.T) {
			if got := humanBytes(c.size); got != c.want {
				t.Errorf("humanBytes() = %v, want %v", got, c.want)
			}
		})
	}
}

func Test_humanDuration(t *testing.T) {
	cases := []struct {
		want string
		d    float64
	}{
		{"1.0 ns", 1},
		{"1.0 us", 1000},
		{"1.0 ms", 1000_000},
		{"1.0 s", 1000_000_000},
		{"50 s", 50_000_000_000},
		{"1.0 m", 60_000_000_000},
		{"60 m", 3599_000_000_000},
		{"1.0 h", 3600_000_000_000},
		{"1.0 d", 24 * 3600_000_000_000},
		{"1.0 y", 365 * 24 * 3600_000_000_000},
		{"1.0 c", 100 * 365 * 24 * 3600_000_000_000},
	}
	for _, c := range cases {
		t.Run(c.want, func(t *testing.T) {
			if got := humanDuration(c.d); got != c.want {
				t.Errorf("humanBytes() = %v, want %v", got, c.want)
			}
		})
	}
}
