// Code is generated by ucloud-model, DO NOT EDIT IT.

package udts

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UDTSClient is the client of UDTS
type UDTSClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UDTSClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UDTSClient {
	meta := ucloud.ClientMeta{Product: "UDTS"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &UDTSClient{
		client,
	}
}
