package simhash

import (
  "fmt"
  "testing"
)

func TestGetLikenessValue(t *testing.T) {
  needle := "Reading bytes into structs using reflection"
  hayStack := "Golang - mapping an variable length array to a struct"

  likeness := GetLikenessValue(needle, hayStack)
  fmt.Println("Likeness:", likeness)
}

func BenchmarkGetLikenessValue(b *testing.B) {
  needle := "Reading bytes into structs using reflection"
  hayStack := "Golang - mapping an variable length array to a struct"

  for i := 0; i < b.N; i++ {
    GetLikenessValue(needle, hayStack)
  }
}
