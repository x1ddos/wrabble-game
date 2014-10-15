package api

import (
	"github.com/crhym3/go-endpoints/endpoints"
)

func RegisterAllAPIServices() (*endpoints.RPCService, error) {
	wsrv, err := endpoints.RegisterServiceWithDefaults(&WrabbleService{})
	if err != nil {
		return nil, err
	}

	wsrv.Info().Name = "wrabble"
	wsrv.Info().Version = "v1beta"
	wsrv.Info().Default = true
	wsrv.Info().Description = "Wrabble game API"

	info := wsrv.MethodByName("GetWord").Info()
	info.Name, info.HTTPMethod, info.Path, info.Desc =
		"words.get", "GET", "words/{dict}/{word}", "Get a word from the dictionary."

	return wsrv, nil
}
