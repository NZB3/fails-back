package counter

import "sync"

type Counter struct {
	c  int64
	mu *sync.Mutex
}

func New(values ...int64) Counter {
	var value int64
	if len(values) == 0 {
		value = 0
	}
	value = values[0]
	return Counter{
		c:  value,
		mu: &sync.Mutex{},
	}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.c++
}

func (c *Counter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c
}

func (c *Counter) Res() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.c = 0
}
