# Goat [![Build Status](https://travis-ci.org/cycloidio/goat.svg?branch=develop)](https://travis-ci.org/cycloidio/goat) [![Coverage Status](https://coveralls.io/repos/github/cycloidio/goat/badge.svg)](https://coveralls.io/github/cycloidio/goat)

## What is Goat?

Goat is a golang SDK for auth0 API.

It tries to answer the lack of SDK in golang for auth0.
Currently most of the users method have been implemented but some methods are still missing optional parameters.

Any contributions are welcome!

## Getting started

To get started, first make sure you've properly set up your Golang environment and then run the
```bash
$ go get github.com/cycloidio/goat/auth0
```
to get the latest version of the [goat](https://github.com/cycloidio/goat/).

### Initialisation

In order to use the SDK:
```go
import 	"github.com/cycloidio/goat/auth0"

var auth0Domain = https://cycloid.eu.auth0.com // your auth0 domain
var auth0APIBasePath = /api/v2                 // the API base path
var auth0Token = "XXXXXX"                      // the token you want to use

var auth0 = goat.NewAuth0(auth0Domain, auth0APIBasePath, auth0Token)
```

The idea behind the token variable is to be able to use different tokens, based on your needs instead of having a token that can do everything. 

### Calls

You can then call any of the methods that have been implemented or call arbitrary a method via the `Call` method.
Example of re-implementing the `GetUser` method:

```go
// This method
json, err := auth0.GetUser("github|testing-user")
// is identical to
json, err := auth0.Call("/users/github|testing-user", http.MethodGet, nil)
```

## License

The MIT License (MIT)

Copyright (c) 2016 cycloid.io

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
