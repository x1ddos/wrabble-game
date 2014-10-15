package api

import (
	"net/http"
	"strings"

	"github.com/crhym3/go-endpoints/endpoints"
	"github.com/crhym3/wrabble-game/wrabble"
	"github.com/crhym3/wrabble-game/wrabble/ds"
)

type GetWordReq struct {
	Word string `json:"word"`
	Dict string `json:"dict"`
}

type WrabbleService struct {
}

func (s *WrabbleService) GetWord(r *http.Request, req *GetWordReq, res *wrabble.Word) error {
	c := endpoints.NewContext(r)
	word, err := ds.GetWord(c, req.Dict, strings.ToLower(req.Word))
	if err != nil {
		return err
	}
	if word == nil {
		return endpoints.NotFoundError
	}
	*res = *word
	return nil
}
