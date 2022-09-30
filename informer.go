package main

import (
	"test-deployment/api/types/v1alpha1"
	clientV1alpha1 "test-deployment/clientset/v1alpha1"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

func WatchResources(clientSet clientV1alpha1.TestDeploymentV1Alpha1Interface) cache.Store {
	testDeploymentStore, testDeploymentController := cache.NewInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return clientSet.TestDeployments("default").List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return clientSet.TestDeployments("default").Watch(options)
			},
		},
		&v1alpha1.TestDeployment{},
		1*time.Minute,
		cache.ResourceEventHandlerFuncs{},
	)
	go testDeploymentController.Run(wait.NeverStop)
	return testDeploymentStore
}
