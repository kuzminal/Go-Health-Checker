package main

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"
)

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Measurable
	Ping() string
	GetID() string
	Health() bool
}

type Checker struct {
	items []Checkable
}

func (c *Checker) Add(newItems ...Checkable) {
	c.items = append(c.items, newItems...)
}

func (c Checker) String() string {
	var result = make([]string, 5)
	for _, item := range c.items {
		result = append(result, item.GetID())
	}
	return fmt.Sprint(strings.Join(result, " "))
}

func (c *Checker) Check() {
	for _, item := range c.items {
		if !item.Health() {
			fmt.Println(item.GetID() + " не работает")
		}
	}
}

func (c *Checker) Run(ctx context.Context, wg *sync.WaitGroup, chAdd chan Checkable) {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			c.Check()
		case <-ctx.Done():
			fmt.Println("Проверки остановлены")
			wg.Done()
			return
		case item := <-chAdd:
			c.Add(item)
		}
	}
}

func (c *Checker) Stop(cancel context.CancelFunc) {
	cancel()
}
