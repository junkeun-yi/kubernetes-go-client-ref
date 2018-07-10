package main

import (
	
	"github.com/junkeun-yi/kubernetestest/pkg/state"
	"github.com/junkeun-yi/kubernetestest/pkg/utils"
	"github.com/junkeun-yi/kubernetestest/pkg/test"
	"github.com/op/go-logging"
)


// Runs the controller and starts the server
func main() {

	kubeconfig := state.GetKubeConfig();

	utils.InitLogging()

	// Initialise a controller
	var control = test.Controller{
		MetClient: state.GetMetricsClientOutOfCluster(kubeconfig),
		Client: state.GetClientOutOfCluster(kubeconfig),
		Logger: logging.MustGetLogger("control"),
	}

	control.Run()

}