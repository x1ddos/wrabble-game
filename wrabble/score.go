package wrabble

type Scorer interface {
	Compute(word string) int
}

type BoggleScorer struct{}

func (s *BoggleScorer) Compute(word string) int {
	l := len(word)
	switch {
	case l > 2 && l < 5:
		return 1
	case l == 5:
		return 2
	case l == 6:
		return 3
	case l == 7:
		return 5
	case l > 7:
		return 11
	}
	return 0
}
