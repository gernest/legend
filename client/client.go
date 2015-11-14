package client

import (
	"github.com/gernest/legend/core"
	"github.com/gernest/legend/registry"
	"golang.org/x/net/context"
)

func Send(ctx context.Context, name string, payload core.Payload) core.Payload {
	service := registry.Get(name)
	if service == nil {
		return nil
	}
	talk := core.NewTalk(payload)
	go service.Serve(talk)
	var rst core.Payload
STOP:
	for {
		select {
		case resp := <-talk.Respond():
			rst = resp
			break STOP
		case <-ctx.Done():
			break STOP
		}
	}
	return rst
}
