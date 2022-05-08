package main

import "github.com/teal-finance/BaseXX/base92"

func main() {
	// Encode any binary data
	bin := []byte{12, 23, 24, 45, 56, 67, 78, 89}
	str := base92.Encode(bin)

	// Decode back
	bin, err := base92.Decode(str)
	if err != nil {
		panic(err)
	}

	// Use custom alphabet, not applicable for xascii85

	var noSpace = base92.NewAlphabet(
		"!#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNO" +
		"PQRSTUVWXYZ[]^_`abcdefghijklmnopqrstuvwxyz{|}~")

	txt := base92.EncodeAlphabet(bin, noSpace)
	bin, err = base92.DecodeAlphabet(txt, noSpace)
}
