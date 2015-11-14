package store

import (
	"bytes"
	"encoding/json"
	"sync"

	"github.com/gernest/legend/core"
	"github.com/gernest/legend/util"
	"github.com/pborman/uuid"
	"golang.org/x/net/context"
)

var (
	data map[string][]byte
	mu   sync.Mutex
)

func init() {
	data = make(map[string][]byte)
}

type DB struct{}

func (d *DB) Init(ctx context.Context, id string) error {
	return nil
}

func (d *DB) Serve(talk core.Talk) {
	var (
		action  = "action"
		payload = "payload"
		keyID   = "id"
	)
	obj := util.NewJSONHelper()
	rst := util.NewJSONHelper()
	err := obj.Decode(talk.Request())
	if err != nil {
		// TODO handle error
	}
	id := uuid.NewRandom().String()

	if obj.HasKey(action) {
		switch obj.GetString(action) {
		case "save":
			if obj.HasKey(payload) {
				raw, err := json.Marshal(obj.Get(payload))
				if err != nil {
					// TODO handle error?
				}
				rst.Set("id", id)
				mu.Lock()
				data[id] = raw
				mu.Unlock()
			}
		case "get":
			if obj.HasKey(keyID) {
				id := obj.GetString(keyID)
				mu.Lock()
				out := data[id]
				mu.Unlock()
				rst.Set(payload, string(out))
				rst.Set(keyID, id)
			}
		}
	}

	rawOut, err := rst.Encode()
	if err != nil {
		// TODO handle error?
	}
	talk.Respond() <- bytes.NewReader(rawOut)
}
func (s *DB) Info() core.ServiceInfo {
	return nil
}
