package ds

import (
	"github.com/crhym3/wrabble-game/wrabble"

	"appengine"
	"appengine/datastore"
)

const (
	EntityKindWord = "Word"
	EntityKindDict = "Dict"
)

var boggleScorer = &wrabble.BoggleScorer{}

func NewWord(c appengine.Context, dict, word string) (*datastore.Key, *wrabble.Word) {
	w := &wrabble.Word{
		Word:  word,
		Dict:  dict,
		Len:   len(word),
		Score: boggleScorer.Compute(word),
	}
	dkey := datastore.NewKey(c, EntityKindDict, dict, 0, nil)
	wkey := datastore.NewKey(c, EntityKindWord, word, 0, dkey)
	return wkey, w
}

func GetWord(c appengine.Context, dict, word string) (*wrabble.Word, error) {
	dkey := datastore.NewKey(c, EntityKindDict, dict, 0, nil)
	wkey := datastore.NewKey(c, EntityKindWord, word, 0, dkey)
	w := &wrabble.Word{Dict: dict, Word: word}
	var err error
	if err = datastore.Get(c, wkey, w); err == datastore.ErrNoSuchEntity {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return w, nil
}
