package core

import (
	"errors"
	"io"

	"golang.org/x/net/context"
)

var (
	ErrServiceNotFound = errors.New("service not found")
)

type Payload interface {
	io.Reader
}

type Talk interface {
	Request() Payload
	Respond() chan Payload
}

type Service interface {
	Init(ctx context.Context, id string) error
	Serve(tlk Talk)
	Info() ServiceInfo
}

const (
	StatusOffline = iota
	StatusOnline
	StatusStoped
)

type ServiceInfo interface {
	Name() string
	Status() int
	ID() string
}
