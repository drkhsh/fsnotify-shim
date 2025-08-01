package fsnotify

import (
	"errors"
)

var (
	ErrEventOverflow = errors.New("fsnotify queue overflow")
)

type Event struct {
	Name string
	Op   Op
}

func (e Event) String() string {
	return ""
}

type Op uint32

const (
	Create Op = 1 << iota
	Write
	Remove
	Rename
	Chmod
)

func (o Op) Has(h Op) bool { return o&h != 0 }
func (e Event) Has(op Op) bool { return e.Op.Has(op) }

type Watcher struct {
	Events chan Event
	Errors chan error
}

func NewWatcher() (*Watcher, error) {
	return &Watcher{
		Events: make(chan Event),
		Errors: make(chan error),
	}, nil
}

func (w *Watcher) Add(name string) error {
	return nil
}

func (w *Watcher) Close() error {
	return nil
}

func (w *Watcher) Remove(name string) error {
	return nil
}
