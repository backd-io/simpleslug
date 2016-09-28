package simpleslug

import (
	"regexp"
	"strings"
)

// Slug returns an English Slugged string
func Slug(origin string) string {
	return SlugCustom(origin, "[^a-z0-9]+")
}

// SlugCustom allows to set a regexp to return a custom slug
func SlugCustom(origin, regExp string) string {
	var re = regexp.MustCompile(regExp)
	return strings.Trim(re.ReplaceAllString(strings.ToLower(origin), "-"), "-")
}
