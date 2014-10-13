package wrabble

type Word struct {
	Word  string `json:"word"`
	Dict  string `json:"dict"`
	Score int32  `json:"score"`
}
