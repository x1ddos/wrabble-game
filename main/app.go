package main

import (
	"github.com/crhym3/go-endpoints/endpoints"
	"github.com/crhym3/wrabble-game/wrabble"
)

func init() {
	if _, err := wrabble.RegisterAPIService(); err != nil {
		panic(err)
	}
	endpoints.HandleHTTP()
}
