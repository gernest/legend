package orchestra

import (
	"fmt"
	"sync"

	"github.com/gernest/legend/core"
	"github.com/gernest/legend/registry"
	"github.com/gernest/legend/util"
	"golang.org/x/net/context"
)

var symphony = newRunner()

var mu sync.Mutex

type runner struct {
	running services
	add     chan core.Service
	stop    chan string
	done    chan struct{}
	stats   chan stats
}

const (
	statsServiceClosed = iota
	statsServiceErrored
)

type stats struct {
	id      int
	err     error
	service core.Service
}

func newRunner() *runner {
	r := &runner{
		running: make(services),
		add:     make(chan core.Service, 10),
		stop:    make(chan string, 10),
		done:    make(chan struct{}, 1),
		stats:   make(chan stats),
	}
	go r.run()
	return r
}

func (r *runner) run() {
	ctx, cancel := context.WithCancel(context.Background())
STOP:
	for {
		select {
		case <-r.done:
			cancel()
			break STOP
		case serv := <-r.add:
			go func(c context.Context, s core.Service) {
				serviceUUID := util.GenUUID()

				ctxWith, servCancel := context.WithCancel(ctx)

				err := s.Init(ctxWith, serviceUUID)
				if err != nil {
					//TODO handle?
				}
				info := s.Info()
				serviceName := info.Name() + "." + fmt.Sprint(info.Status()) + "." + info.ID()
				go r.monitorService(ctxWith, servCancel, s)
				r.running.add(s.Info())
				registry.Register(serviceName, s)
			}(ctx, serv)
		}
	}
}

func (r *runner) updateInfo(info core.ServiceInfo) {
	r.running.add(info)
}

func (r *runner) monitorService(ctx context.Context, cancel context.CancelFunc, service core.Service) {
STOP:
	for {
		select {
		case name := <-r.stop:
			sName := util.GetServiceName(service.Info())
			if name == sName {
				cancel()
				break STOP
			}
		case <-ctx.Done():
			cancel()
			break STOP
		}
	}
}

type services map[string]core.ServiceInfo

func (s services) add(info core.ServiceInfo) {
	mu.Lock()
	s[info.ID()] = info
	mu.Unlock()
}

func (s services) get(id string) core.ServiceInfo {
	mu.Lock()
	info := s.get(id)
	mu.Unlock()
	return info

}

func Start(name string) error {
	serv := registry.Get(name)
	if serv != nil {
		symphony.add <- serv
	}
	return core.ErrServiceNotFound
}

func Stop(name string) {
	symphony.stop <- name
}

func Done() {
	symphony.done <- struct{}{}
}

func Check(id string) core.ServiceInfo {
	return symphony.running.get(id)
}
