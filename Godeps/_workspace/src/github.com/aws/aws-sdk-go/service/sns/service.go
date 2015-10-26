// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package sns

import (
	"github.com/convox/rack/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws"
	"github.com/convox/rack/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/convox/rack/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/request"
	"github.com/convox/rack/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/service"
	"github.com/convox/rack/Godeps/_workspace/src/github.com/aws/aws-sdk-go/aws/service/serviceinfo"
	"github.com/convox/rack/Godeps/_workspace/src/github.com/aws/aws-sdk-go/private/protocol/query"
	"github.com/convox/rack/Godeps/_workspace/src/github.com/aws/aws-sdk-go/private/signer/v4"
)

// Amazon Simple Notification Service (Amazon SNS) is a web service that enables
// you to build distributed web-enabled applications. Applications can use Amazon
// SNS to easily push real-time notification messages to interested subscribers
// over multiple delivery protocols. For more information about this product
// see http://aws.amazon.com/sns (http://aws.amazon.com/sns/). For detailed
// information about Amazon SNS features and their associated API calls, see
// the Amazon SNS Developer Guide (http://docs.aws.amazon.com/sns/latest/dg/).
//
// We also provide SDKs that enable you to access Amazon SNS from your preferred
// programming language. The SDKs contain functionality that automatically takes
// care of tasks such as: cryptographically signing your service requests, retrying
// requests, and handling error responses. For a list of available SDKs, go
// to Tools for Amazon Web Services (http://aws.amazon.com/tools/).
type SNS struct {
	*service.Service
}

// Used for custom service initialization logic
var initService func(*service.Service)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// New returns a new SNS client.
func New(config *aws.Config) *SNS {
	service := &service.Service{
		ServiceInfo: serviceinfo.ServiceInfo{
			Config:      defaults.DefaultConfig.Merge(config),
			ServiceName: "sns",
			APIVersion:  "2010-03-31",
		},
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(query.Build)
	service.Handlers.Unmarshal.PushBack(query.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(query.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(query.UnmarshalError)

	// Run custom service initialization if present
	if initService != nil {
		initService(service)
	}

	return &SNS{service}
}

// newRequest creates a new request for a SNS operation and runs any
// custom request initialization.
func (c *SNS) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}