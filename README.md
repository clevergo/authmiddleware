# CleverGo Authentication Middleware
[![Build Status](https://travis-ci.org/clevergo/authmidware.svg?branch=master)](https://travis-ci.org/clevergo/authmidware)
[![Coverage Status](https://coveralls.io/repos/github/clevergo/authmidware/badge.svg?branch=master)](https://coveralls.io/github/clevergo/authmidware?branch=master)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue)](https://pkg.go.dev/github.com/clevergo/authmidware)
[![Go Report Card](https://goreportcard.com/badge/github.com/clevergo/authmidware)](https://goreportcard.com/report/github.com/clevergo/authmidware)
[![Release](https://img.shields.io/github/release/clevergo/authmidware.svg?style=flat-square)](https://github.com/clevergo/authmidware/releases)
[![Sourcegraph](https://sourcegraph.com/github.com/clevergo/authmidware/-/badge.svg)](https://sourcegraph.com/github.com/clevergo/authmidware?badge)

## Usage

```go
import (
    "github.com/clevergo/auth"
    "github.com/clevergo/auth/authenticators"
    "github.com/clevergo/authmidware"
)
```

```go
var store auth.IdentityStore
authenticator := authenticators.NewBasicAuth(store)
router.Use(authmidware.New(authenticator))
```
