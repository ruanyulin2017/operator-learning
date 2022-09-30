package v1alpha1

import (
	"test-deployment/api/types/v1alpha1"

	scheme "k8s.io/apimachinery/pkg/runtime/schema"
	goclientScheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

type TestDeploymentV1Alpha1Interface interface {
	TestDeployments(namespace string) TestDeploymentInterface
}

type TestDeploymentV1Alpha1Client struct {
	restClient rest.Interface
}

func NewForConfig(c *rest.Config) (*TestDeploymentV1Alpha1Client, error) {
	config := *c
	config.ContentConfig.GroupVersion = &scheme.GroupVersion{Group: v1alpha1.GroupName, Version: v1alpha1.GroupVersion}
	config.APIPath = "/apis"
	config.NegotiatedSerializer = goclientScheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()

	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &TestDeploymentV1Alpha1Client{restClient: client}, nil
}

func (c TestDeploymentV1Alpha1Client) TestDeployments(namespace string) TestDeploymentInterface {
	return &testDeploymentClient{
		restClient: c.restClient,
		namespace:  namespace,
	}
}
