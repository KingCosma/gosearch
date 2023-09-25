# gosearch
An algorithm that searches for words in a slice.

example of this in action

```
package main

import (
  "fmt"

  "github.com/KingCosma/gosearch"
)

func main() {
  x := []string{"hello world", "hi mom", "wow golang is good", "golang is fast}
  y := "golang 4ever"
  fmt.Println(gosearch.SearchWords(x, y))
}


```
