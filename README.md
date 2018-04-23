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
ok/00/urn:a:b______________________________________/-4            5000000              1824 ns/op             244 B/op          5 allocs/op
ok/01/URN:foo:a123,456_____________________________/-4            3000000              2315 ns/op             244 B/op          5 allocs/op
ok/02/urn:foo:a123%2C456___________________________/-4            3000000              2021 ns/op             244 B/op          5 allocs/op
ok/03/urn:ietf:params:scim:schemas:core:2.0:User___/-4            2000000              2860 ns/op             244 B/op          5 allocs/op
ok/04/urn:ietf:params:scim:schemas:extension:enterp/-4            2000000              3463 ns/op             244 B/op          5 allocs/op
ok/05/urn:ietf:params:scim:schemas:extension:enterp/-4            2000000              4181 ns/op             244 B/op          5 allocs/op
ok/06/urn:burnout:nss______________________________/-4            5000000              1867 ns/op             244 B/op          5 allocs/op
ok/07/urn:abcdefghilmnopqrstuvzabcdefghilm:x_______/-4            5000000              1839 ns/op             244 B/op          5 allocs/op
ok/08/urn:urnurnurn:urn____________________________/-4            5000000              1841 ns/op             244 B/op          5 allocs/op
ok/09/urn:ciao:@!=%2C(xyz)+a,b.*@g=$_'_____________/-4            3000000              2480 ns/op             244 B/op          5 allocs/op
ok/10/URN:x:abc%1Dz%2F%3az_________________________/-4            5000000              2027 ns/op             244 B/op          5 allocs/op
no/11/URN:-xxx:x___________________________________/-4           10000000               641 ns/op             196 B/op          4 allocs/op
no/12/urn::colon:nss_______________________________/-4           10000000               603 ns/op             196 B/op          4 allocs/op
no/13/urn:abcdefghilmnopqrstuvzabcdefghilmn:specifi/-4           10000000               661 ns/op             196 B/op          4 allocs/op
no/14/URN:a!?:x____________________________________/-4           10000000               627 ns/op             196 B/op          4 allocs/op
no/15/urn:urn:NSS__________________________________/-4           10000000               560 ns/op             196 B/op          4 allocs/op
no/16/urn:a:#______________________________________/-4           10000000               683 ns/op             196 B/op          4 allocs/op
no/17/urn:a:/______________________________________/-4           10000000               655 ns/op             196 B/op          4 allocs/op
no/18/urn:a:_______________________________________/-4           20000000               628 ns/op             196 B/op          4 allocs/op
no/19/urn:a:%______________________________________/-4           10000000               635 ns/op             196 B/op          4 allocs/op
```

---

[![Analytics](https://ga-beacon.appspot.com/UA-49657176-1/go-urn?flat)](https://github.com/igrigorik/ga-beacon)