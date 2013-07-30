package simhash

import (
  "hash/crc32"
  "io"
)

func getHashTokens(tokens []string) []uint32 {
  hashedTokens := make([]uint32, len(tokens))
  ieee := crc32.NewIEEE()
  for i, token := range tokens {
    ieee.Reset()
    io.WriteString(ieee, token)
    hashedTokens[i] = ieee.Sum32()
  }
  return hashedTokens
}

func isBitSet(b, pos uint32) bool {
  return (b & (1 << pos)) != 0
}

func getHammingDistance(firstValue, secondValue int) int {
  hammingBits := firstValue ^ secondValue
  hammingValue := 0
  for i := 0; i < 32; i++ {
    if isBitSet(uint32(hammingBits), uint32(i)) {
      hammingValue += 1
    }
  }
  return hammingValue
}
