package wrabble

type Word struct {
	Word  string `json:"word" datastore:"-"`
	Dict  string `json:"dict" datastore:"-"`
	Len   int    `json:"len" datastore:"l"`
	Score int    `json:"score" datastore:"s"`
}
