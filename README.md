# CleverGo Authentication Middleware
[![Build Status](https://img.shields.io/travis/clevergo/authmidware?style=for-the-badge)](https://travis-ci.org/clevergo/authmidware)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/authmidware?style=for-the-badge)](https://coveralls.io/github/clevergo/authmidware?branch=master)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/clevergo.tech/authmidware?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/authmidware?style=for-the-badge)](https://goreportcard.com/report/github.com/clevergo/authmidware)
[![Release](https://img.shields.io/github/release/clevergo/authmidware.svg?style=for-the-badge)](https://github.com/clevergo/authmidware/releases)

## Usage

```go
import (
    "clevergo.tech/auth"
    "clevergo.tech/auth/authenticators"
    "clevergo.tech/authmidware"
    "clevergo.tech/clevergo"
)
```

```go
var store auth.IdentityStore
authenticator := authenticators.NewBasicAuth(store)
app := clevergo.New()
app.Use(authmidware.New(authenticator))
```

Checkout [example](https://github.com/clevergo/examples/tree/master/auth) for details.
