package tokenizer

type Pair struct {
	First  int
	Second int
}

type UniqueChars struct {
	Chars map[rune]struct{}
}

type BPE struct {
	str2tok      map[string]uint
	tok2str      map[uint]string
	mergedTokens map[Pair]uint
}

func NewBPE() *BPE {
	return &BPE{
		str2tok:      make(map[string]uint),
		tok2str:      make(map[uint]string),
		mergedTokens: make(map[Pair]uint),
	}
}

func (b *BPE) preprocess(text string) string {
	var processed_text string

	for i, char := range text {
		if char == ' ' && i != 0 {
			processed_text += "Ġ"
		}
		if char != ' ' {
			processed_text += string(char)
		}
	}
	return processed_text
}

func (b *BPE) BuildVocabFromText(text string) *UniqueChars {
	chars := &UniqueChars{
		Chars: make(map[rune]struct{}),
	}

	// Initialize with first 256 ASCII characters
	for i := 0; i < 256; i++ {
		chars.Chars[rune(i)] = struct{}{}
	}

	// Add characters from text that aren't already included
	for _, char := range text {
		chars.Chars[char] = struct{}{}
	}

	// Ensure 'Ġ' is included
	chars.Chars['Ġ'] = struct{}{}

	return chars
}
