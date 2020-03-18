package sakuracloud

import (
	"fmt"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
)

type sakuraCloudProvider struct {
	manager         *interface{}
	resourceLimiter *cloudprovider.ResourceLimiter
}

func BuildSakuraCloudProvider(manager *interface{}, resourceLimiter *cloudprovider.ResourceLimiter) (cloudprovider.CloudProvider, error) {
	mcp := &sakuraCloudProvider{
		manager: &manager,
		resourceLimiter: resourceLimiter
	}

	return mcp, nil
}
