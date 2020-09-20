package main

import "github.com/ishangupta-ds/monitor/pkg/data_collector/kube_state_metrics"

func main() {
	kube_state_metrics.GetMetrics()
}
