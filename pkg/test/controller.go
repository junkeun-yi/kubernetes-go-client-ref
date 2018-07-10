package test

import (
	"k8s.io/client-go/kubernetes"
	"github.com/op/go-logging"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

type Controller struct {
	Client			*kubernetes.Clientset
	MetClient		*metrics.Clientset
	Logger			*logging.Logger
}

func (c *Controller) Run() {

	c.printMetrics();

	// c.printNodes();
	// c.printPods();
	// c.printDeployments();

	// c.failAllPods()
	// c.testPodRedeploy()
	// c.timedEnforce(time.Second * 10)
}

