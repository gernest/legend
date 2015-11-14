package registry

import (
	"sync"

	"github.com/gernest/legend/core"
	"github.com/gernest/legend/util"
)

var reg = newRegistrar()

type registrar struct {
	mu       sync.Mutex
	services map[string]core.Service
}

func newRegistrar() *registrar {
	return &registrar{services: make(map[string]core.Service)}
}
func (r *registrar) register(id string, service core.Service) {
	r.mu.Lock()
	r.services[id] = service
	r.mu.Unlock()
}

func (r *registrar) remove(id string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.services, id)
}

func (r *registrar) get(id string) core.Service {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.services[id]
}

func (r *registrar) all() []string {
	r.mu.Lock()
	defer r.mu.Unlock()
	var rst []string
	for k := range r.services {
		rst = append(rst, k)
	}
	return rst
}

func Register(id string, service core.Service) {
	reg.register(id, service)
}

func Remove(id string) {
	reg.remove(id)
}

func Get(id string) core.Service {
	return reg.get(id)
}

func GetAll() []string {
	return reg.all()
}

func GetRunning(name string) core.Service {
	for _, k := range GetAll() {
		info, err := util.ParseServiceID(k)
		if err != nil {
			return nil
		}
		if info.Name() == name && info.ID() != "" {
			return Get(k)
		}
	}
	return nil
}
