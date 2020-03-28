package main

import (
	"flag"
	"pubkey/pkg/operator"
)

func main() {
	config := operator.GetDefaultConfig()

	kubeconfig := flag.String("kubeconfig", "~/.kube/config", "kubeconfig file")
	debugLevel := flag.String("debug", "INFO", "Debug level")
	port := flag.String("port", "8080", "Webserver port")
	namespace := flag.String("namespace", "default", "Kubernetes namespace")

	flag.Parse()

	config.KubeConfig = *kubeconfig
	config.DebugLevel = *debugLevel
	config.WebServerPort = *port
	config.Namespace = *namespace

	operator, err := operator.New(config)
	if err != nil {
		panic(err)
	}

	operator.Init()
	operator.Run()
}
