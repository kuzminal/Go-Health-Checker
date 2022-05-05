package main

import (
	"time"
)

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

type GoMetrClient struct {
	URL     string
	TimeOut int
}

type HealthCheck struct {
	ServiceId string
	status    string
}

func (g *GoMetrClient) GetMetrics() string {
	return ""
}

func (g *GoMetrClient) Ping() string {
	return g.URL + " работает"
}

func (g *GoMetrClient) GetID() string {
	return g.URL
}

func (g *GoMetrClient) Health() bool {
	to := time.After(time.Duration(g.TimeOut) * time.Second)
	ch := make(chan HealthCheck)
	var health HealthCheck
	go func() {
		ch <- g.getHealth()
	}()
	select {
	case health = <-ch:
		if health.status == PassStatus {
			return true
		} else {
			return false
		}
	case <-to:
		health.status = FailStatus
		return false
	}
}

func (g *GoMetrClient) getHealth() HealthCheck {
	/*n := 1 + rand.Intn(99-1+1)
	var status string
	if n%2 == 0 {
		status = PassStatus
	} else {
		status = FailStatus
	}*/
	time.Sleep(3 * time.Second)
	return HealthCheck{g.URL, PassStatus}
}
