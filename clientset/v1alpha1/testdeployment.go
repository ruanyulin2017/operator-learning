package v1alpha1

import (
	"context"
	"test-deployment/api/types/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type TestDeploymentInterface interface {
	List(opts metav1.ListOptions) (*v1alpha1.TestDeploymentList, error)
	Get(name string, opts metav1.GetOptions) (*v1alpha1.TestDeployment, error)
	Create(*v1alpha1.TestDeployment) (*v1alpha1.TestDeployment, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
}

type testDeploymentClient struct {
	restClient rest.Interface
	namespace  string
}

func (c *testDeploymentClient) List(opts metav1.ListOptions) (*v1alpha1.TestDeploymentList, error) {
	result := v1alpha1.TestDeploymentList{}
	err := c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("testdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)
	return &result, err
}

func (c *testDeploymentClient) Get(name string, opts metav1.GetOptions) (*v1alpha1.TestDeployment, error) {
	result := v1alpha1.TestDeployment{}
	err := c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("testdeployments").
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(&result)
	return &result, err
}

func (c *testDeploymentClient) Create(testDeployment *v1alpha1.TestDeployment) (*v1alpha1.TestDeployment, error) {
	result := v1alpha1.TestDeployment{}
	err := c.restClient.
		Post().
		Namespace(c.namespace).
		Resource("testdeployments").
		Body(testDeployment).
		Do(context.TODO()).
		Into(&result)
	return &result, err
}

func (c *testDeploymentClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.
		Get().
		Namespace(c.namespace).
		Resource("testdeployments").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(context.TODO())
}
