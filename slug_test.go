package simpleslug

import "testing"

func TestSlug(t *testing.T) {

	englishSlugs := map[string]string{
		"Hello World!":             "hello-world",
		"This is a 6 words phrase": "this-is-a-6-words-phrase",
		"¡Texto en Español!":       "texto-en-espa-ol",
	}

	for key, value := range englishSlugs {
		if got, want := Slug(key), value; got != want {
			t.Errorf("Error with Slug. Want: '%v', Got: '%v'", want, got)
		}
	}

	customSlugs := map[string]string{
		"This is a 6 words phrase": "this-is-a-words-phrase",
		"¡Tengo texto en Español!": "tengo-texto-en-espa-ol",
	}

	customRegExp := "[^a-z]+"

	for key, value := range customSlugs {
		if got, want := SlugCustom(key, customRegExp), value; got != want {
			t.Errorf("Error with Slug. Want: '%v', Got: '%v'", want, got)
		}
	}

}
