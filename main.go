package main

import "fmt"

func main() {
	checker := Checker{}
	checker.Add(&GoMetrClient{URL: "yandex.ru", TimeOut: 100})
	checker.Add(&GoMetrClient{URL: "mail.ru", TimeOut: 1000})
	checker.Add(&GoMetrClient{URL: "google.com", TimeOut: 500})
	fmt.Println(checker)
	checker.Check()
	/*for _, health := range GenerateCheck() {
		if health.status == PassStatus {
			fmt.Println(health.ServiceId)
		}
	}*/
}
