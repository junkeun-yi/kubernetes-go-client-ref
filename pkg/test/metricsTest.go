package test

import (
	// "fmt"
	// "sync/atomic"

	// v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "github.com/netsys/triggers2/kubehandler/pkg/utils"
	// "github.com/netsys/triggers2/kubehandler/pkg/state"
)

// prints the metrics for pods
func (c *Controller) printMetrics() {
	metrics, err := c.MetClient.MetricsV1beta1().PodMetricses("").List(metav1.ListOptions{});
	if (err != nil) {
		c.Logger.Debugf("%v", err)
	}
	for i, metric := range metrics.Items {
		c.Logger.Infof("========================================");
		c.Logger.Infof("Pod %v\n", i);
		c.Logger.Infof("%v", metric);
		c.Logger.Infof("========================================\n\n");
	}
}
