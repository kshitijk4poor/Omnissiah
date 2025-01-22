package tokenizer

type Pair struct {
	First  int
	Second int
}

type Vocab struct {
	// Maps token_id to token_str (e.g., {11246: "some"})
	Vocab map[int]string
	// Maps token_str to token_id (e.g., {"some": 11246})
	VocabInv map[string]int
}

type BPE struct {
	str2tok      map[string]int
	tok2str      map[int]string
	mergedTokens map[Pair]int
}

// NewBPE creates a new BPE instance with initialized maps
func NewBPE() *BPE {
	return &BPE{
		str2tok:      make(map[string]int),
		tok2str:      make(map[int]string),
		mergedTokens: make(map[Pair]int),
	}
}
