# simpleslug

This is a basic library to generate slugs from a string.


Simple Usage


```
package main

import (
  "fmt"
  "github.com/backd-io/simpleslug"
)

func main() {

  // Return a simple slug, by removing characters that not meet the `[^a-z0-9]+`
  // regexp
  fmt.Println(simpleslug.Slug("Hello World!"))

  // Also you can change the slug creation by changing the regexp


  customSlugs := []string{
    "This is a 6 words phrase",
    "¡Tengo texto en Español!",
  }

  customRegExp := "[^a-z]+"

  for _, value := range customSlugs {
    fmt.Println(simpleslug.SlugCustom(value, customRegExp))
  }

}
```
