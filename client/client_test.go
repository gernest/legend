package client

import (
	"bytes"
	"testing"

	"golang.org/x/net/context"
)

func TestSend(t *testing.T) {
	bug := &bytes.Buffer{}
	rst := Send(context.Background(), "hello", bug)
	if rst != nil {
		t.Error("expected nil got %#v", rst)
	}
}
