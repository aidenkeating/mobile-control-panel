package operations

import (
	"net/http"

	openservicebroker "github.com/feedhenry/mobile-control-panel/server/pkg/mobile/broker"
)

// Unbind is an implementation of the service broker unbind api
func (b *BrokerOperations) Unbind(instanceID, bindingID string) *openservicebroker.Response {
	// in principle, unbind should alter state somewhere (e.g. invalidating the credentials
	// associated with this binding, for brokers that provide unique credentials for each binding)
	return &openservicebroker.Response{Code: http.StatusOK, Body: &openservicebroker.UnbindResponse{}, Err: nil}
}
