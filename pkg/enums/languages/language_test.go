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

package languages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseStringToLanguage(t *testing.T) {
	t.Run("should success parse string to language", func(t *testing.T) {
		assert.Equal(t, Go, ParseStringToLanguage("Go"))
		assert.Equal(t, CSharp, ParseStringToLanguage("C#"))
		assert.Equal(t, Dart, ParseStringToLanguage("Dart"))
		assert.Equal(t, Ruby, ParseStringToLanguage("Ruby"))
		assert.Equal(t, Python, ParseStringToLanguage("Python"))
		assert.Equal(t, Java, ParseStringToLanguage("Java"))
		assert.Equal(t, Kotlin, ParseStringToLanguage("Kotlin"))
		assert.Equal(t, Javascript, ParseStringToLanguage("Javascript"))
		assert.Equal(t, Typescript, ParseStringToLanguage("Typescript"))
		assert.Equal(t, Leaks, ParseStringToLanguage("Leaks"))
		assert.Equal(t, HCL, ParseStringToLanguage("HCL"))
		assert.Equal(t, C, ParseStringToLanguage("C"))
		assert.Equal(t, PHP, ParseStringToLanguage("PHP"))
		assert.Equal(t, HTML, ParseStringToLanguage("HTML"))
		assert.Equal(t, Generic, ParseStringToLanguage("Generic"))
		assert.Equal(t, Yaml, ParseStringToLanguage("YAML"))
		assert.Equal(t, Elixir, ParseStringToLanguage("Elixir"))
		assert.Equal(t, Shell, ParseStringToLanguage("Shell"))
		assert.Equal(t, Unknown, ParseStringToLanguage(""))
	})
}

func TestValues(t *testing.T) {
	t.Run("should return 19 valid values", func(t *testing.T) {
		assert.Len(t, Values(), 19)
	})
}

func TestMapEnableLanguages(t *testing.T) {
	t.Run("should map success get all languages from map", func(t *testing.T) {
		mapLanguages := mapEnableLanguages()

		assert.Equal(t, Go, mapLanguages["Go"])
		assert.Equal(t, CSharp, mapLanguages["C#"])
		assert.Equal(t, Dart, mapLanguages["Dart"])
		assert.Equal(t, Ruby, mapLanguages["Ruby"])
		assert.Equal(t, Python, mapLanguages["Python"])
		assert.Equal(t, Java, mapLanguages["Java"])
		assert.Equal(t, Kotlin, mapLanguages["Kotlin"])
		assert.Equal(t, Javascript, mapLanguages["Javascript"])
		assert.Equal(t, Typescript, mapLanguages["Typescript"])
		assert.Equal(t, Leaks, mapLanguages["Leaks"])
		assert.Equal(t, HCL, mapLanguages["HCL"])
		assert.Equal(t, C, mapLanguages["C"])
		assert.Equal(t, PHP, mapLanguages["PHP"])
		assert.Equal(t, HTML, mapLanguages["HTML"])
		assert.Equal(t, Generic, mapLanguages["Generic"])
		assert.Equal(t, Yaml, mapLanguages["YAML"])
		assert.Equal(t, Elixir, mapLanguages["Elixir"])
		assert.Equal(t, Shell, mapLanguages["Shell"])
		assert.Equal(t, Unknown, mapLanguages["Unknown"])
	})
}

func TestToString(t *testing.T) {
	t.Run("should success parse to string", func(t *testing.T) {
		assert.Equal(t, "Go", Go.ToString())
		assert.Equal(t, "C#", CSharp.ToString())
		assert.Equal(t, "Dart", Dart.ToString())
		assert.Equal(t, "Ruby", Ruby.ToString())
		assert.Equal(t, "Python", Python.ToString())
		assert.Equal(t, "Java", Java.ToString())
		assert.Equal(t, "Kotlin", Kotlin.ToString())
		assert.Equal(t, "Javascript", Javascript.ToString())
		assert.Equal(t, "Typescript", Typescript.ToString())
		assert.Equal(t, "Leaks", Leaks.ToString())
		assert.Equal(t, "HCL", HCL.ToString())
		assert.Equal(t, "C", C.ToString())
		assert.Equal(t, "PHP", PHP.ToString())
		assert.Equal(t, "HTML", HTML.ToString())
		assert.Equal(t, "Generic", Generic.ToString())
		assert.Equal(t, "YAML", Yaml.ToString())
		assert.Equal(t, "Elixir", Elixir.ToString())
		assert.Equal(t, "Shell", Shell.ToString())
		assert.Equal(t, "Unknown", Unknown.ToString())
	})
}

func TestGetCustomImagesKeyByLanguage(t *testing.T) {
	t.Run("should success get horusec config json key by language", func(t *testing.T) {
		assert.Equal(t, "go", Go.GetCustomImagesKeyByLanguage())
		assert.Equal(t, "csharp", CSharp.GetCustomImagesKeyByLanguage())
		assert.Equal(t, "dart", Dart.GetCustomImagesKeyByLanguage())
		assert.Equal(t, "ruby", Ruby.GetCustomImagesKeyByLanguage())
		assert.Equal(t, "python", Python.GetCustomImagesKeyByLanguage())
		assert.Equal(t, "java", Java.GetCustomImagesKeyByLanguage())
		assert.Equal(t, "kotlin", Kotlin.GetCustomImagesKeyByLanguage())
		assert.Equal(t, "javascript", Javascript.GetCustomImagesKeyByLanguage())
		assert.Equal(t, "leaks", Leaks.GetCustomImagesKeyByLanguage())
		assert.Equal(t, "hcl", HCL.GetCustomImagesKeyByLanguage())
		assert.Equal(t, "c", C.GetCustomImagesKeyByLanguage())
		assert.Equal(t, "php", PHP.GetCustomImagesKeyByLanguage())
		assert.Equal(t, "yaml", Yaml.GetCustomImagesKeyByLanguage())
		assert.Equal(t, "elixir", Elixir.GetCustomImagesKeyByLanguage())
		assert.Equal(t, "shell", Shell.GetCustomImagesKeyByLanguage())
	})
}
