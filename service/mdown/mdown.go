package mdown

import (
	"io/ioutil"
	"strings"

	"github.com/a8m/mark"
	"github.com/gernest/legend/core"
	"golang.org/x/net/context"
)

type Markdown struct{}

func (m *Markdown) Init(ctx context.Context, id string) error {
	return nil
}

func (m *Markdown) Serve(talk core.Talk) {
	data, err := ioutil.ReadAll(talk.Request())
	if err != nil {
		//TODO Do error handling?
	}
	out := mark.Render(string(data))
	talk.Respond() <- strings.NewReader(out)
}

func (s *Markdown) Info() core.ServiceInfo {
	return nil
}
