// Copyright (c) 2019-2021 Antonino Catinello
// SPDX-License-Identifier: BSD-3-Clause

// Command line tool to de-/encode base91.
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	// "catinello.eu/base91"
	"github.com/teal-finance/BaseXX/ac/base91"
)

var version string

const wrap int = 127 // 127 + add newline char

func main() {
	alen := len(os.Args)

	if alen >= 2 {
		// decode
		if os.Args[1] == "-d" || os.Args[1] == "--decode" {
			if alen == 3 {
				if os.Args[2] != "-" {
					content, err := ioutil.ReadFile(os.Args[2])
					if err != nil {
						// error
						fmt.Fprintln(os.Stderr, err.Error())
					} else {
						fmt.Print(base91.Decode(content))
					}
				} else if os.Args[2] == "-" {
					var input []byte
					scanner := bufio.NewScanner(os.Stdin)
					for scanner.Scan() {
						input = append(input, scanner.Bytes()...)
					}

					if err := scanner.Err(); err != nil {
						fmt.Println(err)
					}

					fmt.Print(base91.Decode(input))
				}
			} else {
				// help
				help()
			}
		} else if os.Args[1] == "-h" {
			// help
			help()
		} else {
			// encode
			content, err := ioutil.ReadFile(os.Args[1])
			if err != nil {
				// error
				fmt.Fprintln(os.Stderr, err.Error())
			} else {
				fmt.Print(newLine(base91.EncodeToString(content)))
			}
		}
	} else if alen == 1 {
		var input []byte
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input = append(input, scanner.Bytes()...)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println(err)
		}

		fmt.Println(newLine(base91.EncodeToString(input)))
	} else {
		// help
		help()
	}
}

func help() {
	fmt.Println("base91 - Binary to ASCII text encoding.")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  base91 [OPTIONS] <FILE>")
	fmt.Println()
	fmt.Println("Parameter:")
	fmt.Println("  <FILE>               		| Path to file or - for stdin.")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -d | --decode <FILE> 		| Decode mode.")
	fmt.Println()
	fmt.Println("  --help               		| Show this help.")
	fmt.Println("  --license            		| Print license.")
	fmt.Println("  --version            		| Print version.")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  echo \"bla\"  | base91		| Encode")
	fmt.Println("  echo \"<izI\" | base91 -d - 	| Decode")
	fmt.Println()
	fmt.Println("Website:")
	fmt.Println("  https://codeberg.org/ac/base91")
	fmt.Println()
	fmt.Println("License:")
	fmt.Println("  BSD License ©  Antonino Catinello")
	fmt.Println()
	fmt.Println("Version:")
	fmt.Println("  " + version)
}

func newLine(content string) string {
	var tmp string

	for i := 0; i < len(content); i++ {
		if i%wrap == 0 && i != 0 {
			tmp = tmp + "\n"
		}
		tmp = tmp + string(content[i])
	}

	return tmp
}
