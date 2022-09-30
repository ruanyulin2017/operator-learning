package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"time"

	// "k8s.io/client-go/applyconfigurations/apiserverinternal/v1alpha1"
	"test-deployment/api/types/v1alpha1"
	clientV1alpha1 "test-deployment/clientset/v1alpha1"

	// "k8s.io/client-go/kubernetes/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientScheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var config *rest.Config
	var err error
	if home := homedir.HomeDir(); home != "" {
		kubeconfig := flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "")
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			panic(err)
		}
	} else {
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err)
		}
	}

	v1alpha1.AddToScheme(clientScheme.Scheme)
	clientSet, err := clientV1alpha1.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	testDeploymentList, err := clientSet.TestDeployments("default").List(v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("test deployment list: %v\n", testDeploymentList)

	store := WatchResources(clientSet)
	// for {
	testDeploymentStore := store.List()
	fmt.Printf("test deployment in store number: %v\n", len(testDeploymentStore))
	fmt.Printf("test deployment in store: %v\n", testDeploymentStore)
	time.Sleep(3 * time.Second)
	// }
}
