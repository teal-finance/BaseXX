# `BaseXX/base91`

Pretty good Base91 encoder with customizable encoding alphabet.

## Copyright

Copyright (c) 2022 Teal.Finance contributors

This file is part of Teal.Finance/BaseXX licensed under the MIT License.
See the LICENSE file or <https://opensource.org/licenses/MIT>.
SPDX-License-Identifier: MIT

## Encoding alphabet

The default encoding alphabet `StdEncoding`
complies with the constraints of the cookie tokens:
`StdEncoding` uses 91 characters
from 0x21 `!` to 0x7E `~` except `"` `;` and `\`.

You can also provide your own customized alphabet
with `NewEncoding()`.

The following example uses the semicolon `;`
and drops the single `'` and double quotes `"`.

```go
noQuotes := base91.NewEncoding("" +
 "abcdefghijklmnopqrstuvwxyz" +
 "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
 "0123456789/{|}~.,:;?[]^_`!@#$%&()*+-<=>")
```

## Other base alphabets

Characters often used by common BaseXX encodings:

    Base92  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=*-_~.,?!@#$%&()[]{|}<>^:`'"
    Base91  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=*-_~.,?!@#$%&()[]{|}<>^:`'
    Base64  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz+/=
    Base62  0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz
    Base58   123456789ABCDEFGH JKLMN PQRSTUVWXYZabcdefghijk mnopqrstuvwxyz
    Hexa    0123456789ABCDEF

## Slower than the Michael Traver's implementation

The implementation <https://github.com/mtraver/base91>
from Joachim Henke and Michael Traver
is much cleaner and faster:

* Standard Encoding interface
* 190 times faster encoding
* 35 times faster decoding

For a cleaner and faster Base91 encoder,
you should use: <https://github.com/mtraver/base91>

## Can be much faster

Performance can be much improved.
Using the tips of <https://github.com/mtraver/base91>
(original work from Joachim Henke),
this `BaseXX/base91` may become 200 times faster on the encoding,
and 30 times faster on the decoding.

## Contributions welcome

This `BaseXX/base91` needs your help to become faster.
Please propose your enhancements,
or even a further refactoring.
Any contribution is welcome. ;-)
