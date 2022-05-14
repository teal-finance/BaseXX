# BaseXX

Go modules: &emsp; &emsp; &emsp; &emsp; &emsp; [![Go Reference](https://pkg.go.dev/badge/github.com/teal-finance/BaseXX.svg "Go documentation for BaseXX")](https://pkg.go.dev/github.com/teal-finance/BaseXX)

[`import "github.com/teal-finance/BaseXX/base58"`](./base58/)  
[`import "github.com/teal-finance/BaseXX/base62"`](./base62/)  
[`import "github.com/teal-finance/BaseXX/base91"`](./base91/)  
[`import "github.com/teal-finance/BaseXX/base92"`](./base92/)  
[`import "github.com/teal-finance/BaseXX/xascii85"`](./xascii85/)  

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

BaseXX is similar to <https://github.com/unix-world/smartgo>
from Radu Ovidiu Ilies who has adapted
<https://github.com/akamensky/base58> to support any base.

BaseXX shares with [SmartGo](https://github.com/unix-world/smartgo)
the ability to quickly create new base encoders,
by copying source files and changing few things.

The main interest of BaseXX compared to
[SmartGo](https://github.com/unix-world/smartgo)
is the performance: BaseXX is five times faster on Base92!
See the [benchmark results](#benchmark).

## Slower than Base91 by Joachim Henke and Michael Traver

The implementation [github.com/mtraver/base91](https://github.com/mtraver/base91)
is much cleaner and faster:

- Standard Encoding interface
- 190 times faster encoding
- 35 times faster decoding

See the [benchmark results](#benchmark).

is much faster than BaseXX/base91:
The Base91 by Antonino encodes six times faster,
and decodes twice faster.

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
[original project](https://github.com/mr-tron/base58) on Base58.

```go
goos: linux
goarch: amd64
pkg: github.com/teal-finance/BaseXX/ac/base91
cpu: AMD Ryzen 9 3900X 12-Core Processor            
BenchmarkEncode-24           236           5220578 ns/op
BenchmarkDecode-24           222           5695511 ns/op
PASS
ok      github.com/teal-finance/BaseXX/ac/base91        3.600s
goos: linux
goarch: amd64
pkg: github.com/teal-finance/BaseXX/base58
cpu: AMD Ryzen 9 3900X 12-Core Processor            
BenchmarkEncode-24                       1000000              1621 ns/op
BenchmarkEncodeMrTronBase58-24            663850              1734 ns/op
BenchmarkDecode-24                       1625545               745.3 ns/op
BenchmarkDecodeMrTronBase58-24           1602920               742.9 ns/op
PASS
ok      github.com/teal-finance/BaseXX/base58   6.733s
goos: linux
goarch: amd64
pkg: github.com/teal-finance/BaseXX/base62
cpu: AMD Ryzen 9 3900X 12-Core Processor            
BenchmarkEncode-24       1000000              1255 ns/op
BenchmarkDecode-24       1610145               784.7 ns/op
PASS
ok      github.com/teal-finance/BaseXX/base62   3.302s
goos: linux
goarch: amd64
pkg: github.com/teal-finance/BaseXX/base91
cpu: AMD Ryzen 9 3900X 12-Core Processor            
BenchmarkEncoding_Encode-24                         2821            402254 ns/op
BenchmarkEncoding_EncodeToString-24                 2972            411904 ns/op
BenchmarkEncoding_DecodeString-24                  15943             84473 ns/op
BenchmarkBproctorBase91_Decode-24                  39297             29221 ns/op
BenchmarkEquimBase91_Decode-24                    335587              3520 ns/op
BenchmarkMajestrateBase91_Decode-24               357110              3220 ns/op
BenchmarkMtraverBase91_EncodeToString-24          542314              2190 ns/op
BenchmarkMtraverBase91_Encode-24                 1308085               943.3 ns/op
BenchmarkMtraverBase91_DecodeString-24            462540              2426 ns/op
BenchmarkMtraverBase91_Decode-24                 1075700              1063 ns/op
PASS
ok      github.com/teal-finance/BaseXX/base91   14.924s
goos: linux
goarch: amd64
pkg: github.com/teal-finance/BaseXX/base92
cpu: AMD Ryzen 9 3900X 12-Core Processor            
BenchmarkEncode-24                        819314              1415 ns/op
BenchmarkEncodeSmartgoBase92-24           153697              8140 ns/op
BenchmarkDecode-24                       1799289               665.9 ns/op
BenchmarkDecodeSmartgoBase92-24           341623              3501 ns/op
PASS
ok      github.com/teal-finance/BaseXX/base92   6.616s
PASS
ok      github.com/teal-finance/BaseXX/xascii85 0.003s
```

## Can be much faster

Performance can be much much improved.
Using the tips of <https://github.com/mtraver/base91>
(original work from Joachim Henke),
this BaseXX repo may become 200 times faster on the encoding,
and 30 times faster on the decoding.

## Contributions welcome

This BaseXX repo needs your help to become faster.
Please propose your enhancements,
or even a further refactoring.
Any contribution is welcome. ;-)

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

See similar other projects.

### Many bases by Radu Ovidiu Ilies

The repo <https://github.com/mr-tron/base58> proposes many bases.
See section [Faster than unix-world/smartgo](#faster-than-unix-worldsmartgo).

### Base91 by Joachim Henke and Michael Traver

This is the cleanest and fastest Base91 Go-implementation I have found.

Michael Traver has ported to Go the excellent work from Joachim Henke.

See <https://github.com/mtraver/base91>.

### Base91 by Majestrate

Good Base91 implementation at
<https://github.com/majestrate/base91>.
I have not find any bug.

### Base91 by Equim

Good Base91 implementation at
<https://github.com/Equim-chan/base91-go>.
I have not find any bug.

### Base91 by Brad Proctor

The repo <https://github.com/bproctor/base91>
is a slow Base91 implementation.
Still faster than BaseXX/Base91.
I have not find any bug.

### Base91 by Antonino Catinello

The implementation by [Antonino Catinello](https://codeberg.org/ac)
is not the fastest, and contains bugs. This repo contains a
[fork](<https://github.com/teal-finance/BaseXX/ac/base91>)
(almost unmodified) of <https://codeberg.org/ac/base91>.
The latter cannot be used because the module name
`"catinello.eu/base91"` is not reachable (tested in May 2022).

The implementation of [Antonino Catinello](https://codeberg.org/ac)
is much faster than BaseXX/base91:
The Base91 by Antonino encodes six times faster,
and decodes twice faster.

This Base91 implementation does not customize the encoding alphabet.

### Base91 by Chris Snell and Breeze Chen

The repo <https://github.com/breezechen/base91>
is an unmaintained and buggy Base91 implementation.
