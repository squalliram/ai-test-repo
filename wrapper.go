package main

import (
	"github.com/splitio/go-client/v6/splitio/client"
)

type Wrapper struct {
	fmeClient *client.SplitClient
}

func (w Wrapper) Evaluate(key string, feature string) string {
	return w.fmeClient.Treatment(key, feature, nil)
}
