package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	ApiRequests         func(ApiRequestsLabels) prometheus.Counter `name:"api_requests" help:"API requests"`
	ApiRequestsDuration func(ApiRequestsLabels) prometheus.Summary `name:"api_requests_duration" help:"API requests duration" max_age:"20m" objectives:"0.50,0.95,0.99"`
	WebRequests         func(WebRequestsLabels) prometheus.Counter `name:"web_requests" help:"Web requests"`
	WebRequestsDuration func(WebRequestsLabels) prometheus.Summary `name:"web_requests_duration" help:"Web requests duration" max_age:"20m" objectives:"0.50,0.95,0.99"`
}

type ApiRequestsLabels struct {
	Code   int    `label:"code"`
	Method string `label:"method"`
	URL    string `label:"url"`
}

type WebRequestsLabels struct {
	Code   int    `label:"code"`
	Method string `label:"method"`
	URL    string `label:"url"`
}

func New() *Metrics {
	return &Metrics{}
}
