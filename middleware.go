// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package authmidware

import (
	"log"

	"github.com/clevergo/auth"
	"github.com/clevergo/clevergo"
)

// New returns a middleware with the given authenticator.
func New(authenticator auth.Authenticator) clevergo.MiddlewareFunc {
	return func(next clevergo.Handle) clevergo.Handle {
		return func(c *clevergo.Context) error {
			identity, err := authenticator.Authenticate(c.Request, c.Response)
			if err != nil {
				log.Println(err)
				authenticator.Challenge(c.Request, c.Response)
			} else {
				c.WithValue(auth.IdentityKey, identity)
			}
			return next(c)
		}
	}
}
