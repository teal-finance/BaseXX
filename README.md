# BaseXX

Go modules: &emsp; &emsp; &emsp; &emsp; &emsp; [![Go Reference](https://pkg.go.dev/badge/github.com/teal-finance/BaseXX.svg "Go documentation for BaseXX")](https://pkg.go.dev/github.com/teal-finance/BaseXX)

[`import "github.com/teal-finance/BaseXX/base58"`](./base58/)  
[`import "github.com/teal-finance/BaseXX/base62"`](./base62/)  
[`import "github.com/teal-finance/BaseXX/base91"`](./base91/)  
[`import "github.com/teal-finance/BaseXX/base92"`](./base92/)  
[`import "github.com/teal-finance/BaseXX/xascii85"`](./xascii85/)  
[`import "github.com/teal-finance/BaseXX/ac/base91"`](./ac/base91/)

Characters often used by common BaseXX encodings:

```
Hexa    0123456789ABCDEF
Base58   123456789ABCDEFGH JKLMN PQRSTUVWXYZabcdefghijk mnopqrstuvwxyz
Base62  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
Base64  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=
Base91  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=*-_~.,?!@#$%&()[]{|}<>^:`'
Base92  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=*-_~.,?!@#$%&()[]{|}<>^:`'"
```

## Fork of mr-tron/base58

Originally, BaseXX is a fork of <https://github.com/mr-tron/base58>
adapted to support other bases:

- [`base58`](./base58/)
- [`base62`](./base62/)
- [`base91`](./base91/)
- [`base92`](./base92/)
- [`xascii85`](./xascii85/)

All these packages, except `xascii85`,
support customized encoding alphabet
without any performance tradeoff.

The `xascii85` package is just a layer on top of `"encoding/ascii85"`
to provide the same API as the other packages.

## Common interface

All these packages aim to provide the following API:

```go
func NewEncoding(alphabet string) *Encoding

interface Encoding {
    Encode(bin []byte) []byte
    Decode(ascii []byte) ([]byte, error)

    EncodeToString(bin []byte) string
    DecodeString(ascii string) ([]byte, error)

    EncodedLen(n int) int
    DecodedLen(n int) int
}
```

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

## Slower than Base91 by Antonino Catinello

This repo contains a
[copy](<https://github.com/teal-finance/BaseXX/ac/base91>)
(almost unmodified) of <https://codeberg.org/ac/base91>.
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

In the following example replace `base92` by
`base58`, `base62`, `base91` (or `xascii85`).

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

    var noSpace = base92.NewEncoding(
        "!#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNO" +
        "PQRSTUVWXYZ[]^_`abcdefghijklmnopqrstuvwxyz{|}~")

    txt := noSpace.EncodeToString(bin)
    bin, err = noSpace.DecodeString(txt)
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
BenchmarkEncode-24           234           5098065 ns/op         6642443 B/op         34 allocs/op
BenchmarkDecode-24           217           5285605 ns/op         5241612 B/op         33 allocs/op
PASS
ok      github.com/teal-finance/BaseXX/ac/base91        3.456s
?       github.com/teal-finance/BaseXX/ac/base91/cmd/base91     [no test files]
goos: linux
goarch: amd64
pkg: github.com/teal-finance/BaseXX/base58
cpu: AMD Ryzen 9 3900X 12-Core Processor            
BenchmarkEncode-24                       1000000              1706 ns/op              96 B/op          2 allocs/op
BenchmarkEncodeMrTronBase58-24           1000000              1802 ns/op              96 B/op          2 allocs/op
BenchmarkDecode-24                       1611724               733.3 ns/op           127 B/op          2 allocs/op
BenchmarkDecodeMrTronBase58-24           1632877               741.9 ns/op           127 B/op          2 allocs/op
PASS
ok      github.com/teal-finance/BaseXX/base58   7.445s
goos: linux
goarch: amd64
pkg: github.com/teal-finance/BaseXX/base62
cpu: AMD Ryzen 9 3900X 12-Core Processor            
BenchmarkEncode-24        908791              1360 ns/op              48 B/op          1 allocs/op
BenchmarkDecode-24       1524541               759.1 ns/op           127 B/op          2 allocs/op
PASS
ok      github.com/teal-finance/BaseXX/base62   4.193s
goos: linux
goarch: amd64
pkg: github.com/teal-finance/BaseXX/base91
cpu: AMD Ryzen 9 3900X 12-Core Processor            
BenchmarkEncode-24               1000000              1383 ns/op              48 B/op          1 allocs/op
BenchmarkEncodeACBase91-24       2823261               422.7 ns/op           168 B/op          5 allocs/op
BenchmarkDecode-24               1709204               691.3 ns/op           124 B/op          2 allocs/op
BenchmarkDecodeACBase91-24       2849089               416.9 ns/op           104 B/op          4 allocs/op
PASS
ok      github.com/teal-finance/BaseXX/base91   6.562s
goos: linux
goarch: amd64
pkg: github.com/teal-finance/BaseXX/base92
cpu: AMD Ryzen 9 3900X 12-Core Processor            
BenchmarkEncode-24                        891435              1349 ns/op              48 B/op          1 allocs/op
BenchmarkEncodeSmartgoBase92-24           140588              8100 ns/op            1377 B/op         78 allocs/op
BenchmarkDecode-24                       1795351               665.8 ns/op           122 B/op          2 allocs/op
BenchmarkDecodeSmartgoBase92-24           335845              3459 ns/op             232 B/op          8 allocs/op
PASS
ok      github.com/teal-finance/BaseXX/base92   6.344s
?       github.com/teal-finance/BaseXX/encoding [no test files]
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
