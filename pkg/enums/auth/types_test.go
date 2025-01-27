// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInvalid(t *testing.T) {
	t.Run("should true when invalid type", func(t *testing.T) {
		assert.True(t, AuthorizationType("test").IsInvalid())
	})

	t.Run("should false when valid type", func(t *testing.T) {
		assert.False(t, Horusec.IsInvalid())
		assert.False(t, Keycloak.IsInvalid())
		assert.False(t, Ldap.IsInvalid())
	})
}

func TestValues(t *testing.T) {
	t.Run("should 3 valid auth types", func(t *testing.T) {
		var testType AuthorizationType
		assert.Len(t, testType.Values(), 3)
	})
}

func TestToString(t *testing.T) {
	t.Run("should success parse to string", func(t *testing.T) {
		assert.Equal(t, "horusec", Horusec.ToString())
		assert.Equal(t, "ldap", Ldap.ToString())
		assert.Equal(t, "keycloak", Keycloak.ToString())
	})
}

func TestGetAuthTypeByString(t *testing.T) {
	t.Run("should return type correctly when exists", func(t *testing.T) {
		assert.Equal(t, Horusec, GetAuthTypeByString("horusec"))
		assert.Equal(t, Ldap, GetAuthTypeByString("ldap"))
		assert.Equal(t, Keycloak, GetAuthTypeByString("keycloak"))
	})

	t.Run("should return type empty when not exists", func(t *testing.T) {
		assert.Equal(t, AuthorizationType(""), GetAuthTypeByString("test"))
	})
}
