package main

import "math/rand"

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
	healthCheck := g.getHealth()
	if healthCheck.status == PassStatus {
		return true
	} else {
		return false
	}
}

func (g *GoMetrClient) getHealth() HealthCheck {
	n := 1 + rand.Intn(99-1+1)
	var status string
	if n%2 == 0 {
		status = PassStatus
	} else {
		status = FailStatus
	}
	return HealthCheck{g.URL, status}
}
