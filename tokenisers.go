package simhash

type Tokeniser interface {
  Tokenise(input string) []string
}

type FixedSizeStringTokeniser struct {
  tokensize uint8
}

func NewFixedSizeStringTokeniser(tokensize uint8) *FixedSizeStringTokeniser {
  if tokensize < 2 || tokensize > 127 {
    panic("Token size must be between 2 and 127")
  }

  return &FixedSizeStringTokeniser{tokensize: tokensize}
}

func (t *FixedSizeStringTokeniser) Tokenise(input string) []string {
  var chunks []string
  offset := 0
  inputLen := len(input)

  for offset < inputLen {
    offset2 := offset + int(t.tokensize)
    if offset2 >= inputLen {
      chunks = append(chunks, input[offset:])
    } else {
      chunks = append(chunks, input[offset:offset2])
    }
    offset += int(t.tokensize)
  }

  return chunks
}

type OverlappingStringTokeniser struct {
  chunkSize, overlapSize uint8
}

func NewOverlappingStringTokeniser(chunkSize, overlapSize uint8) *OverlappingStringTokeniser {
  if chunkSize <= overlapSize {
    panic("Chunk size must be greater than overlap size.")
  }

  return &OverlappingStringTokeniser{chunkSize: chunkSize, overlapSize: overlapSize}
}

func (t *OverlappingStringTokeniser) Tokenise(input string) []string {
  var chunks []string
  inputLen := len(input)
  for position := 0; position < inputLen-int(t.chunkSize); position += int(t.chunkSize - t.overlapSize) {
    chunks = append(chunks, input[position:position+int(t.chunkSize)])
  }
  return chunks
}
