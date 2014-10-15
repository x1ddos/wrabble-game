package main

import (
	"github.com/crhym3/go-endpoints/endpoints"
	"github.com/crhym3/wrabble-game/wrabble/api"
)

func init() {
	if _, err := api.RegisterAllAPIServices(); err != nil {
		panic(err)
	}
	endpoints.HandleHTTP()
}
