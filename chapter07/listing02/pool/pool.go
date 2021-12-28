package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m       sync.Mutex
	r       chan io.Closer
	factory func() (io.Closer, error)
	closed  bool
}

var PoolClosedERR = errors.New("pool  closed!!")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("wrong param")
	}

	return &Pool{
		r:       make(chan io.Closer, size),
		factory: fn,
	}, nil
}

func (p *Pool) Aquire() (io.Closer, error) { // 这里返回指针还是值？？？   接口值  就没必要指针了
	select {
	case res, ok := <-p.r:
		if !ok {
			log.Println("pool closed")
			return nil, PoolClosedERR
		}
		log.Println("share resource")
		return res, nil
	default:
		log.Println("new resource")
		return p.factory()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		r.Close()
		return
	}

	select {
	case p.r <- r:
		return
	default:
		r.Close()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return
	}

	p.closed = true

	close(p.r)

	for r := range p.r {
		r.Close()
	}
}
