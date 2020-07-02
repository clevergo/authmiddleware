# CleverGo Authentication Middleware
[![Build Status](https://img.shields.io/travis/clevergo/authmiddleware?style=for-the-badge)](https://travis-ci.org/clevergo/authmiddleware)
[![Coverage Status](https://img.shields.io/coveralls/github/clevergo/authmiddleware?style=for-the-badge)](https://coveralls.io/github/clevergo/authmiddleware?branch=master)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white&style=for-the-badge)](https://pkg.go.dev/clevergo.tech/authmiddleware?tab=doc)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/authmiddleware?style=for-the-badge)](https://goreportcard.com/report/github.com/clevergo/authmiddleware)
[![Release](https://img.shields.io/github/release/clevergo/authmiddleware.svg?style=for-the-badge)](https://github.com/clevergo/authmiddleware/releases)

## Usage

```go
import (
    "clevergo.tech/auth"
    "clevergo.tech/auth/authenticators"
    "clevergo.tech/authmiddleware"
    "clevergo.tech/clevergo"
)
```

```go
var store auth.IdentityStore
authenticator := authenticators.NewBasicAuth(store)
app := clevergo.New()
app.Use(authmiddleware.New(authenticator))
```

Checkout [example](https://github.com/clevergo/examples/tree/master/auth) for details.
