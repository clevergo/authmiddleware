// Copyright 2020 CleverGo. All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// in the LICENSE file.

package authmiddleware

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"clevergo.tech/auth"
	"clevergo.tech/clevergo"
	"github.com/stretchr/testify/assert"
)

type nullIdentity struct{}

func (nullIdentity) GetID() string {
	return "null"
}

type nullAuthenticator struct {
	authenticated bool
}

func (na *nullAuthenticator) Authenticate(*http.Request, http.ResponseWriter) (auth.Identity, error) {
	if na.authenticated {
		return nullIdentity{}, nil
	}
	return nil, errors.New("unauthenticated")
}

func (na *nullAuthenticator) Challenge(*http.Request, http.ResponseWriter) {
}

func TestGetIdentity(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	identity := nullIdentity{}
	r = r.WithContext(context.WithValue(r.Context(), auth.IdentityKey, identity))
	assert.Equal(t, identity, GetIdentity(r.Context()))
}

func TestNew(t *testing.T) {
	cases := []struct {
		authenticated bool
	}{
		{true},
		{false},
	}
	for _, test := range cases {
		authenticator := &nullAuthenticator{authenticated: test.authenticated}
		m := New(authenticator)
		handled := false
		handle := func(c *clevergo.Context) error {
			handled = true
			identity := auth.GetIdentity(c.Request.Context())
			if test.authenticated {
				assert.NotNil(t, identity)
			} else {
				assert.Nil(t, identity)
			}
			return nil
		}
		c := &clevergo.Context{
			Request: httptest.NewRequest(http.MethodGet, "/", nil),
		}
		m(handle)(c)
		assert.True(t, handled)
	}
}
