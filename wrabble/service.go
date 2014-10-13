package wrabble

import (
	"net/http"

	"github.com/crhym3/go-endpoints/endpoints"
)

type GetWordReq struct {
	Word string `json:"word"`
	Dict string `json:"dict"`
}

type WrabbleService struct {
}

func (s *WrabbleService) GetWord(r *http.Request, req *GetWordReq, res *Word) error {
	res.Word = req.Word
	res.Dict = req.Dict
	res.Score = 0
	return nil
}

func RegisterAPIService() (*endpoints.RPCService, error) {
	wrabble := &WrabbleService{}
	srv, err := endpoints.RegisterServiceWithDefaults(wrabble)
	if err != nil {
		return nil, err
	}

	srv.Info().Name = "wrabble"
	srv.Info().Version = "v1beta"
	srv.Info().Default = true
	srv.Info().Description = "Wrabble game API"

	info := srv.MethodByName("GetWord").Info()
	info.Name, info.HTTPMethod, info.Path, info.Desc =
		"words.get", "GET", "words/{dict}/{word}", "Get a word from the dictionary."

	return srv, nil
}
