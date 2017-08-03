package operations

import (
	clientset "github.com/feedhenry/mobile-control-panel/server/pkg/client/clientset_generated/internalclientset"
)

// BrokerOperations provides the implementation of the service broker
// catalog/provision/bind/unbind/unprovision/last_operation apis.
// It users a resource client to access the ServiceInstance resources
// that are used to store instance/binding state information.
type BrokerOperations struct {
	Client clientset.Interface
}
