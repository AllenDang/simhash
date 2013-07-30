package simhash

const (
  HashSize = 32
)

func calculateSimHash(input string) int {
  tokeniser := NewOverlappingStringTokeniser(4, 3)
  hashedTokens := getHashTokens(tokeniser.Tokenise(input))
  vector := make([]int, HashSize)
  for i, _ := range vector {
    vector[i] = 0
  }

  for _, v := range hashedTokens {
    for i, _ := range vector {
      if isBitSet(uint32(v), uint32(i)) {
        vector[i] += 1
      } else {
        vector[i] -= 1
      }
    }
  }

  fingerprint := 0
  for i, v := range vector {
    if v > 0 {
      fingerprint += 1 << uint32(i)
    }
  }

  return fingerprint
}

func GetLikenessValue(needle, haystack string) float64 {
  needleSimHash := calculateSimHash(needle)
  hayStackSimHash := calculateSimHash(haystack)
  return float64(HashSize-getHammingDistance(needleSimHash, hayStackSimHash)) / float64(HashSize)
}
