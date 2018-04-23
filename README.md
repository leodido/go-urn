[![Build](https://img.shields.io/travis/leodido/go-urn/master.svg?style=for-the-badge)](https://travis-ci.org/leodido/go-urn) [![Coverage](https://img.shields.io/codecov/c/github/leodido/go-urn.svg?style=for-the-badge)](https://codecov.io/gh/leodido/go-urn) [![Documentation](https://img.shields.io/badge/godoc-reference-blue.svg?style=for-the-badge)](https://godoc.org/github.com/leodido/go-urn)

**A parser for URNs**.

> As seen on [RFC 2141](https://tools.ietf.org/html/rfc2141#ref-1).

[API documentation](https://godoc.org/github.com/leodido/go-urn).

## Installation

```
go get github.com/leodido/go-urn
```

Or

```
go get gopkg.in/leodido/go-urn.v1
```

## Perfs

```
ok/00/urn:a:b______________________________________/-4           50000000               197 ns/op             166 B/op          6 allocs/op
ok/01/URN:foo:a123,456_____________________________/-4           10000000               954 ns/op             272 B/op         34 allocs/op
ok/02/urn:foo:a123%2C456___________________________/-4           10000000              1057 ns/op             320 B/op         36 allocs/op
ok/03/urn:ietf:params:scim:schemas:core:2.0:User___/-4            2000000              4261 ns/op            1696 B/op        134 allocs/op
ok/04/urn:ietf:params:scim:schemas:extension:enterp/-4            1000000              6345 ns/op            3296 B/op        198 allocs/op
ok/05/urn:ietf:params:scim:schemas:extension:enterp/-4            1000000              9186 ns/op            5728 B/op        270 allocs/op
ok/06/urn:burnout:nss______________________________/-4           20000000               463 ns/op             192 B/op         14 allocs/op
ok/07/urn:abcdefghilmnopqrstuvzabcdefghilm:x_______/-4           30000000               311 ns/op             197 B/op          6 allocs/op
ok/08/urn:urnurnurn:urn____________________________/-4           20000000               461 ns/op             192 B/op         14 allocs/op
ok/09/urn:ciao:@!=%2C(xyz)+a,b.*@g=$_'_____________/-4            3000000              2771 ns/op             976 B/op         88 allocs/op
ok/10/URN:x:abc%1Dz%2F%3az_________________________/-4           10000000              1169 ns/op             336 B/op         38 allocs/op
no/11/URN:-xxx:x___________________________________/-4           20000000               386 ns/op             304 B/op          6 allocs/op
no/12/urn::colon:nss_______________________________/-4           20000000               416 ns/op             304 B/op          6 allocs/op
no/13/urn:abcdefghilmnopqrstuvzabcdefghilmn:specifi/-4           20000000               512 ns/op             304 B/op          6 allocs/op
no/14/URN:a!?:x____________________________________/-4           20000000               393 ns/op             304 B/op          6 allocs/op
no/15/urn:urn:NSS__________________________________/-4           20000000               366 ns/op             272 B/op          6 allocs/op
no/16/urn:white_space:NSS__________________________/-4           20000000               416 ns/op             304 B/op          6 allocs/op
no/17/urn:concat:no_spaces_________________________/-4           20000000               519 ns/op             312 B/op          9 allocs/op
no/18/urn:a:/______________________________________/-4           20000000               411 ns/op             304 B/op          7 allocs/op
no/19/urn:UrN:NSS__________________________________/-4           20000000               391 ns/op             272 B/op          6 allocs/op
```

Notice ragel implementation also provides at the meantime:

1. fine-grained and informative errors
2. normalization during parsing

---

[![Analytics](https://ga-beacon.appspot.com/UA-49657176-1/go-urn?flat)](https://github.com/igrigorik/ga-beacon)