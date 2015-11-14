package view

import (
	"bytes"
	"html/template"

	"github.com/gernest/legend/core"
	"github.com/gernest/legend/util"
	"golang.org/x/net/context"
)

type Renderer struct {
	tpl *template.Template
}

func (d *Renderer) Init(ctx context.Context, id string) error {
	return nil
}

func (r *Renderer) Serve(talk core.Talk) {
	var (
		name           = "name"
		templateContex = "data"
	)

	obj := util.NewJSONHelper()
	rst := util.NewJSONHelper()

	err := obj.Decode(talk.Request())
	if err != nil {
		//TODO?
	}
	if obj.HasKey(name) {
		tname := obj.GetString(name)
		data := obj.Get(templateContex)

		out := &bytes.Buffer{}
		rerr := r.tpl.ExecuteTemplate(out, tname, data)
		if rerr != nil {
			//TODO?
		}
		rst.Set("output", out.String())
	}

	rawOut, err := rst.Encode()
	if err != nil {
		// TODO handle error?
	}
	talk.Respond() <- bytes.NewReader(rawOut)
}
func (s *Renderer) Info() core.ServiceInfo {
	return nil
}
