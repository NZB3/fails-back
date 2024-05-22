package controller

import (
	"fmt"
	"log"
	"net/http"
)

type counter interface {
	Inc()
	Res()
	Value() int64
}

type controller struct {
	Counter counter
}

func NewController(counter counter) *controller {
	return &controller{
		Counter: counter,
	}
}

func (c *controller) GetDaysCount(w http.ResponseWriter, r *http.Request) {
	log.Println("GetDaysCount called")

	value := c.Counter.Value()

	w.Header().Set("Content-Type", "application/json")

	_, err := w.Write([]byte(fmt.Sprintf("{\"count\":%d}", value)))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func (c *controller) Reset(w http.ResponseWriter, r *http.Request) {
	log.Println("Reset called")
	c.Counter.Res()
	w.WriteHeader(http.StatusOK)
}
