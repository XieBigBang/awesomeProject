package pool

import (
	"log"
	"sync"
)

type resource struct {
	name string
}

func (r *resource) Close() {
	log.Printf("%v close", r.name)
}

type pool struct {
	m sync.Mutex
	r chan resource
}
