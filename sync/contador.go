package sync

import "sync"

type Contador struct {
	mu    sync.Mutex
	valor int
}

func (c *Contador) Incrementa() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.valor++
}

func (c *Contador) Valor() int {
	return c.valor
}
