package counter

import "sync"

type observer interface {
	Update(value int)
}

type Counter struct {
	c         int
	mu        *sync.Mutex
	observers []observer
}

func New(value *int, observers ...observer) Counter {
	return Counter{
		c: func(value *int) int {
			if value == nil {
				return 0
			}
			return *value
		}(value),
		mu:        &sync.Mutex{},
		observers: observers,
	}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.c++
	c.notifyObservers(c.c)
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c
}

func (c *Counter) Res() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.c = 0
	c.notifyObservers(c.c)
}

func (c *Counter) notifyObservers(value int) {
	for _, observer := range c.observers {
		observer.Update(value)
	}
}
