package simhash

import (
  "fmt"
  "testing"
)

func TestFixedSizeStringTokeniser(t *testing.T) {
  tokeniser := NewFixedSizeStringTokeniser(5)

  tokens := tokeniser.Tokenise("hello world this is my way to hell")
  for _, tk := range tokens {
    fmt.Println(tk)
  }
}

func TestOverlappingStringTokeniser(t *testing.T) {
  tokeniser := NewOverlappingStringTokeniser(4, 3)

  tokens := tokeniser.Tokenise("hello world this is my way to hell")
  for _, tk := range tokens {
    fmt.Println(tk)
  }
}
