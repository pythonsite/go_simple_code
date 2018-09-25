package tombv1

import (
	"fmt"
	"sync"
	"errors"
)


type Tomb struct {
	m 		sync.Mutex
	dying   chan struct{}
	dead    chan struct{}
	reason 	error
}


var (
	ErrStilAlive = errors.New("tomb: still alive")
	ErrDying = errors.New("tomb: dying")
)

func (t *Tomb) init() {
	t.m.Lock()
	if t.dead == nil {
		// 这里默认创建的是无缓冲区的chan
		t.dead = make(chan struct{})
		t.dying = make(chan struct{})
		t.reason = ErrStilAlive
	}
	t.m.Unlock()
}

func (t *Tomb) Dying () <- chan struct{} {
	t.init()
	return t.dying
}

// 将会阻塞直到这个goroutine是dead状态，并且返回这个goroutine的death的原因
func (t *Tomb) Wait() error {
	t.init()
	<- t.dead
	t.m.Lock()
	reason := t.reason
	t.m.Unlock()
	return reason
}

func (t *Tomb) Done() {
	t.Kill(nil)
	close(t.dead)
}

func (t *Tomb) Kill(reason error) {
	t.init()
	t.m.Lock()
	defer t.m.Unlock()
	if reason == ErrDying {
		if t.reason == ErrStilAlive {
			panic("tomb: Kill with ErrDying while still alive")
		}
		return
	}
	if t.reason == nil || t.reason == ErrStilAlive {
		t.reason = reason
	}
	select {
	case <- t.dying:
	default:
		close(t.dying)
	}
}


func (t *Tomb) Killf(f string, a...interface{}) error {
	err := fmt.Errorf(f,a...)
	t.Kill(err)
	return err
}


func (t *Tomb) Err()(reason error) {
	t.init()
	t.m.Lock()
	reason = t.reason
	t.m.Unlock()
	return
}
