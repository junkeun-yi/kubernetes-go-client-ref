package test

import (
	"time"
	// "fmt"
	// "sync/atomic"
	"math/rand"

	// v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "github.com/netsys/triggers2/kubehandler/pkg/utils"
	// "github.com/netsys/triggers2/kubehandler/pkg/state"
	"k8s.io/client-go/kubernetes"
	"github.com/op/go-logging"
)

// The controller object holds the state, timeline, queue of incoming events,
// and list of triggers to follow.
type Controller struct {
	Client			*kubernetes.Clientset
	Logger			*logging.Logger
}

func (c *Controller) Run() {

	c.testPodRedeploy()

	// c.timedEnforce(time.Second * 10)
}

func (c *Controller) timedEnforce(t time.Duration) {
	ticker := time.NewTicker(t)

	for {
		select {
			case <- ticker.C:
				c.randomDelete()
		}
	}
}

func (c *Controller) testPodRedeploy() {
	pods, err := c.Client.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		c.Logger.Debug(err)
	}

	for _, p := range pods.Items {
		result, getErr := c.Client.CoreV1().Pods("default").Get(p.Name, metav1.GetOptions{})
		if getErr != nil {
			c.Logger.Debug(getErr)
		}
		c.Logger.Infof("%v", result.Name)

		c.Logger.Infof("%v", result.Spec.NodeName)

		c.Logger.Infof("%T", result.Spec.NodeName)

		c.Logger.Infof("%T", result)

		result.Spec.NodeName = "ip-172-20-39-172.us-west-1.compute.internal"


		_, updateErr := c.Client.CoreV1().Pods("default").Update(result)

		if updateErr != nil {
			c.Logger.Debug(updateErr)
		}

		return
	}



}


func (c *Controller) randomDelete() {
	/**
	pods, err := c.Client.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		c.Logger.Debug(err)
	}

	//configPods := make([]*v1.Pod, len(pods.Items))
	podNames := make(string, len(pods.Items))

	for _, pod := range pods.Items {
		if pod.Namespace == "default" {
			podNames = append(podNames, pod.Name)
			//configPods = append(configPods, &pod)
		}
	}
	
	randIndex = rand.
	*/

	podDeleteOption := metav1.NewDeleteOptions(0)
	
	pods, err := c.Client.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		c.Logger.Debug(err)
	}

	for _, pod := range pods.Items {
		temp := rand.Int31n(10)
		if temp < 3 {
			name := pod.Name
			if pod.Namespace == "default" {
				err = c.Client.CoreV1().Pods(pod.Namespace).Delete(name, podDeleteOption)
			}
			if err != nil {
				c.Logger.Debug(err)
				panic(err)
			}
			c.Logger.Infof("deleted pod: %v", name)
			return
		}
	}

	defer c.randomDelete()
	return
	
	/**
	podDeleteOption := metav1.NewDeleteOptions(0)

	err := c.Client.CoreV1().Pods("frontend").Delete("frontend-685d7ff496-8q8n9", podDeleteOption)
	if err != nil {
		c.Logger.Debug(err)
	}
	*/
}
