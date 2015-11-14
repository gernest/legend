package orchestra

import (
	"testing"

	"github.com/gernest/legend/core"
	"golang.org/x/net/context"
)

type SimpleServ struct {
	id string
}

func (s *SimpleServ) Init(ctx context.Context, id string) {
	s.id = id
}
func (s *SimpleServ) Serve(tal core.Talk) {
	tal.Respond() <- tal.Request()
}

func TestStart(t *testing.T) {
	err := Start("bogus")
	if err == nil {
		t.Error("expected an error got nil instead")
	}
}
