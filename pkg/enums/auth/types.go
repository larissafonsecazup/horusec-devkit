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

type AuthorizationType string

const (
	Keycloak AuthorizationType = "keycloak"
	Ldap     AuthorizationType = "ldap"
	Horusec  AuthorizationType = "horusec"
)

func (a AuthorizationType) IsInvalid() bool {
	for _, v := range a.Values() {
		if v == a {
			return false
		}
	}

	return true
}

func (a AuthorizationType) Values() []AuthorizationType {
	return []AuthorizationType{
		Keycloak,
		Ldap,
		Horusec,
	}
}

func (a AuthorizationType) ToString() string {
	return string(a)
}

func GetAuthTypeByString(authType string) (a AuthorizationType) {
	for _, v := range a.Values() {
		if v.ToString() == authType {
			return v
		}
	}

	return ""
}
