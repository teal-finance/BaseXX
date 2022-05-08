# BaseXX

Go modules: &emsp; &emsp; &emsp; &emsp; &emsp; [![Go Reference](https://pkg.go.dev/badge/github.com/teal-finance/BaseXX.svg "Go documentation for BaseXX")](https://pkg.go.dev/github.com/teal-finance/BaseXX)

[`import "github.com/teal-finance/BaseXX/base58"`](./base58/)  
[`import "github.com/teal-finance/BaseXX/base62"`](./base62/)  
[`import "github.com/teal-finance/BaseXX/base91"`](./base91/)  
[`import "github.com/teal-finance/BaseXX/base92"`](./base92/)  
[`import "github.com/teal-finance/BaseXX/xascii85"`](./xascii85/)  
[`import "github.com/teal-finance/BaseXX/ac/base91"`](./ac/base91/)

## Fork of mr-tron/base58

Originally, BaseXX is a fork of <https://github.com/mr-tron/base58>
adapted to support other bases. The following folders contain
code copied from <https://github.com/mr-tron/base58>:

- [base58](./base58/)
- [base62](./base62/)
- [base91](./base91/)
- [base92](./base92/)
- [helper](./helper/) (common code)

All these previous BaseXX encoders
support customized encoding alphabet
without any performance tradeoff.

## Faster than unix-world/smartgo

BaseXX is very similar to the project
<https://github.com/unix-world/smartgo>
from Radu Ovidiu Ilies who has adapted
<https://github.com/akamensky/base58>
to support any base.

BaseXX shares with [SmartGo](https://github.com/unix-world/smartgo)
the ability to quickly create new base encoders,
by copying source files and changing the alphabet.

The main interest of BaseXX compared to
[SmartGo](https://github.com/unix-world/smartgo)
is the performance: BaseXX is five times faster.
See the [benchmark results](#benchmark).

## Ascii85 convenient abstraction

This repo also ships the [Ascii85](./xascii85/)
that is just a layer on top of the "encoding/ascii85"
standard library providing the same function signature:

```go
func Encode(bin []byte) string
func Decode(str string) ([]byte, error)
```

## Even faster Base91 by Antonino Catinello

Moreover, this repo contains a
[copy](<https://github.com/teal-finance/BaseXX/ac/base91>)
of <https://codeberg.org/ac/base91>.
The latter cannot be used because the module name
`"catinello.eu/base91"` is not reachable (tested in May 2022).

The implementation of [Antonino Catinello](https://codeberg.org/ac)
is much faster than BaseXX/base91:
The Base91 by Antonino encodes six times faster,
and decodes twice faster.
See the [benchmark results](#benchmark)

However his Base91 implementation does not allow to change
the alphabet (but this is possible with some adaptations).

## Compliance with cookie token standards

The default alphabet of [BaseXX/base58](./base58/),
[BaseXX/base62](./base62/), [BaseXX/base91](./base91/)
and [BaseXX/base92](./base92/) conforms with the
cookie token constraints:

- characters from 0x20 (space) to 0x7E (~) included
- except three characters: " ; and \

The two other encoders, [BaseXX/xascii85](./xascii85/)
and [BaseXX/ac/base91](./ac/base91/) do not support
cookie token encoding.

## Usage

In this [example](./example/base92.go)
replace `base92` by `base58`, `base62`,
`base91` or `xascii85`.

```go
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
```

## Benchmark

The benchmark shows this BaseXX project is almost faster than the
[original project](https://github.com/mr-tron/base58)
on Base58.

```go
$ go test -run=NO -bench=. -benchmem github.com/teal-finance/BaseXX/...
goos: linux
goarch: amd64
pkg: github.com/teal-finance/BaseXX/ac/base91
cpu: AMD Ryzen 9 3900X 12-Core Processor
BenchmarkEncode-24           228           5165122 ns/op         6642451 B/op         34 allocs/op
BenchmarkDecode-24           225           5753119 ns/op         5241609 B/op         33 allocs/op
PASS
ok      github.com/teal-finance/BaseXX/ac/base91        3.584s
goos: linux
goarch: amd64
pkg: github.com/teal-finance/BaseXX/base58
cpu: AMD Ryzen 9 3900X 12-Core Processor
BenchmarkEncode-24                       1000000              1586 ns/op              96 B/op          2 allocs/op
BenchmarkEncodeMrTronBase58-24            596448              1743 ns/op              96 B/op          2 allocs/op
BenchmarkDecode-24                       1588452               736.3 ns/op           127 B/op          2 allocs/op
BenchmarkDecodeMrTronBase58-24           1505677               809.6 ns/op           127 B/op          2 allocs/op
PASS
ok      github.com/teal-finance/BaseXX/base58   6.859s
goos: linux
goarch: amd64
pkg: github.com/teal-finance/BaseXX/base62
cpu: AMD Ryzen 9 3900X 12-Core Processor
BenchmarkEncode-24        901647              1915 ns/op              96 B/op          2 allocs/op
BenchmarkDecode-24       1573705               789.6 ns/op           127 B/op          2 allocs/op
PASS
ok      github.com/teal-finance/BaseXX/base62   3.783s
goos: linux
goarch: amd64
pkg: github.com/teal-finance/BaseXX/base91
cpu: AMD Ryzen 9 3900X 12-Core Processor
BenchmarkEncode-24               1000000              2074 ns/op              96 B/op          2 allocs/op
BenchmarkEncodeACBase91-24       3479275               331.1 ns/op           120 B/op          4 allocs/op
BenchmarkDecode-24               2181652               647.6 ns/op           124 B/op          2 allocs/op
BenchmarkDecodeACBase91-24       3946309               304.2 ns/op            56 B/op          3 allocs/op
PASS
ok      github.com/teal-finance/BaseXX/base91   7.094s
goos: linux
goarch: amd64
pkg: github.com/teal-finance/BaseXX/base92
cpu: AMD Ryzen 9 3900X 12-Core Processor
BenchmarkEncode-24                        692828              2008 ns/op              96 B/op          2 allocs/op
BenchmarkEncodeSmartgoBase92-24           136118              7472 ns/op            1377 B/op         78 allocs/op
BenchmarkDecode-24                       3127748               381.6 ns/op           122 B/op          2 allocs/op
BenchmarkDecodeSmartgoBase92-24           619510              1990 ns/op             232 B/op          8 allocs/op
PASS
ok      github.com/teal-finance/BaseXX/base92   6.437s
PASS
ok      github.com/teal-finance/BaseXX/xascii85 0.002s
```

## Feedback

If you have some suggestions, or need a new feature,
please contact us, using the
[issue tracking](https://github.com/teal-finance/BaseXX/issues),
or at Teal.Finance[à]pm.me or
[@TealFinance](https://twitter.com/TealFinance).

Feel free to propose a
[Pull Request](https://github.com/teal-finance/BaseXX/pulls),
your contributions are welcome. :wink:

## See also

Other similar projects:

- <https://github.com/mr-tron/base58>
- <https://github.com/unix-world/smartgo>
- <https://github.com/bproctor/base91>
- <https://github.com/breezechen/base91>
- <https://github.com/Equim-chan/base91-go>
- <https://github.com/majestrate/base91>
- <https://github.com/mtraver/base91>
