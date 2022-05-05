package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	checker := Checker{}
	ch := make(chan Checkable)
	checker.Add(&GoMetrClient{URL: "yandex.ru", TimeOut: 1})
	checker.Add(&GoMetrClient{URL: "mail.ru", TimeOut: 1})
	checker.Add(&GoMetrClient{URL: "google.com", TimeOut: 5})
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println("Восстановление после сбоя.....")
				fmt.Println(err)
			}
		}()
		checker.Run(ctx, wg, ch)
	}()
	time.Sleep(6 * time.Second)
	ch <- &GoMetrClient{"rambler.ru", 20}
	time.Sleep(6 * time.Second)
	checker.Stop(cancel)
	wg.Wait()
}
